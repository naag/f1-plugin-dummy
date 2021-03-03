package scenarios

import (
	"errors"
	"time"

	f1 "github.com/naag/f1-api/pkg/api"
	"github.com/naag/f1-test-utils/pkg/support"
)

func AdmissionScenario(t *f1.T) (f1.RunFn, f1.TeardownFn) {
	runFunc := func(t *f1.T) {
		// assert.Fail(t, "I'm failing")
		time.Sleep(50 * time.Millisecond)
	}

	teardownFunc := func(t *f1.T) {
	}

	return runFunc, teardownFunc
}

func SubmissionScenario(t *f1.T) (f1.RunFn, f1.TeardownFn) {
	runFunc := func(t *f1.T) {
		// assert.Fail(t, "I'm failing")
		time.Sleep(10 * time.Millisecond)
	}

	teardownFunc := func(t *f1.T) {
	}

	return runFunc, teardownFunc
}

func SqsScenario(t *f1.T) (f1.RunFn, f1.TeardownFn) {
	runFunc := func(t *f1.T) {
		r := support.RandomReference()
		panic(errors.New("foobar"))
		t.Log.Infof("got random reference %s", r)
	}

	teardownFunc := func(t *f1.T) {
	}

	return runFunc, teardownFunc
}
