package support

import (
	"strconv"

	"github.com/form3tech-oss/f1/pkg/f1/testing"
	"github.com/naag/f1-test-utils/pkg/config"
)

type combinedScenarios struct {
	run      []testing.RunFn
	teardown []testing.TeardownFn
}

func CombineScenarios(scenarios ...testing.SetupFn) testing.SetupFn {
	return func(t *testing.T) (testing.RunFn, testing.TeardownFn) {

		combo := &combinedScenarios{
			run:      []testing.RunFn{},
			teardown: []testing.TeardownFn{},
		}

		cleanOnStart := config.CleanOnStart
		Cleanup(t)

		// stop functions from cleaning each others queues.
		config.CleanOnStart = false
		defer func() { config.CleanOnStart = cleanOnStart }()

		for _, s := range scenarios {
			run, teardown := s(t)
			combo.run = append(combo.run, run)
			combo.teardown = append(combo.teardown, teardown)
		}
		return combo.Run, combo.Teardown
	}
}

func (s *combinedScenarios) Run(t *testing.T) {
	user, err := strconv.Atoi(t.VirtualUser)
	t.Require.Nil(err)
	s.run[user%len(s.run)](t)
}

func (s *combinedScenarios) Teardown(t *testing.T) {
	for _, fn := range s.teardown {
		fn(t)
	}
}
