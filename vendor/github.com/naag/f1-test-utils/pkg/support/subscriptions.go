package support

import (
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/service/sqs"

	"github.com/form3tech-oss/f1/pkg/f1/testing"
	"github.com/go-openapi/strfmt"
	"github.com/hako/durafmt"
)

func WaitForAll(t *testing.T, channels map[strfmt.UUID]chan *sqs.Message, requiredContent string) {

	timeout := time.Duration(30+len(channels)) * time.Second

	t.Log.Infof("Waiting for up to %s for all resources to notify %s", durafmt.Parse(timeout), requiredContent)

	completionChan := make(chan bool, 500)
	defer close(completionChan)
	stopChan := make(chan bool)
	defer close(stopChan)

	for id, ch := range channels {
		go func(id string, ch chan *sqs.Message) {
			for {
				select {
				case <-stopChan:
					return
				case message := <-ch:
					if strings.Contains(*message.Body, requiredContent) {
						t.Log.Debugf("Received %s for %s", requiredContent, id)
						completionChan <- true
						return
					}
				}
			}
		}(id.String(), ch)
	}

	timeoutChan := time.After(timeout)
	completed := 0
	for {
		select {
		case <-timeoutChan:
			t.Require.Failf("timeout", "timed out waiting for %d items to receive notifications for %s with %d/%d", len(channels), requiredContent, completed, len(channels))
		case <-stopChan:
			return
		case <-completionChan:
			completed++
			t.Log.Infof("%d of %d notifications received for %s", completed, len(channels), requiredContent)
			if completed == len(channels) {
				return
			}
		}
	}
}
