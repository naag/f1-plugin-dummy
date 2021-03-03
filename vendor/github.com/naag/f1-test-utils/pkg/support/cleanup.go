package support

import (
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/form3tech-oss/f1/pkg/f1/testing"
	"github.com/form3tech-oss/go-form3/v2/pkg/generated/client/subscriptions"
	"github.com/go-openapi/swag"
	"github.com/naag/f1-test-utils/pkg/config"
)

func Cleanup(t *testing.T) {
	if config.CleanOnStart {
		CleanupNow(t)
	}
}

func CleanupNow(t *testing.T) {
	stackName := os.Getenv("STACK_NAME")
	if stackName == "" {
		t.Log.Fatalf("Stack name must be present in environment variable `STACK_NAME`")
	}
	t.Log.Info("Checking for subscriptions left by previous runs")
	removeSubscriptions(t, F3ForTest(t).Subscriptions, stackName)
	deleteQueues(t, stackName)
}

func deleteQueues(t *testing.T, stackName string) {
	awsConfig := GetAwsConfig(t)
	sess, err := session.NewSession(&awsConfig)
	if err != nil {
		t.Log.WithError(err).Error("Error getting AWS session to delete queues")
		return
	}
	sqsClient := sqs.New(sess)
	queues, err := sqsClient.ListQueues(&sqs.ListQueuesInput{
		QueueNamePrefix: swag.String(fmt.Sprintf("%s-testing", stackName)),
	})
	if err != nil {
		t.Log.WithError(err).Error("Error listing old queues to delete")
		return
	}
	for _, queue := range queues.QueueUrls {
		t.Log.Infof("Deleting queue %s", *queue)
		_, err := sqsClient.DeleteQueue(&sqs.DeleteQueueInput{
			QueueUrl: queue,
		})
		if err != nil {
			t.Log.WithError(err).Errorf("Unable to delete queue %s", *queue)
		}
	}
}

func removeSubscriptions(t *testing.T, subscriptionsClient *subscriptions.Client, stackName string) {
	subscriptionsList, err := subscriptionsClient.ListSubscriptions().Do()
	if err != nil {
		t.Log.WithError(err).Error("Unable to list existing subscriptions for cleanup")
		return
	}
	for _, s := range subscriptionsList.Data {
		if strings.Contains(*s.Attributes.CallbackURI, fmt.Sprintf("%s-testing", stackName)) {
			t.Log.Infof("Deleting subscription to queue %s", *s.Attributes.CallbackURI)
			_, err := subscriptionsClient.DeleteSubscription().WithID(*s.ID).WithVersion(*s.Version).Do()
			if err != nil {
				t.Log.WithError(err).Errorf("unable to delete subscription to queue %s", *s.Attributes.CallbackURI)
			}
		}
	}
}
