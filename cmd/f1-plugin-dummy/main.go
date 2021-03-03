package main

import (
	"github.com/naag/f1-api/pkg/plugin"
	"github.com/naag/f1-plugin-dummy/cmd/f1-plugin-dummy/scenarios"
)

func main() {
	plugin.NewServer().
		WithScenario("dummyAdmission", scenarios.AdmissionScenario).
		WithScenario("dummySubmission", scenarios.SubmissionScenario).
		Serve()
}
