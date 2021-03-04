package main

import (
	"github.com/naag/f1-api/pkg/server"
	"github.com/naag/f1-plugin-dummy/cmd/f1-plugin-dummy/scenarios"
)

func main() {
	server.NewServer().
		WithScenario("dummyAdmission", scenarios.AdmissionScenario).
		WithScenario("dummySubmission", scenarios.SubmissionScenario).
		WithScenario("sqsScenario", scenarios.SqsScenario).
		Serve()
}
