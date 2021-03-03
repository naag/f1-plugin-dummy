package support

import (
	"fmt"
	"os"

	"github.com/stretchr/testify/require"

	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/go-openapi/swag"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/form3tech/go-logger/log"
)

func GetAwsConfig(t require.TestingT) aws.Config {
	containerCredentialsUri, hasContainerCredentialsUri := os.LookupEnv("AWS_CONTAINER_CREDENTIALS_RELATIVE_URI")
	if hasContainerCredentialsUri {
		log.Info(fmt.Sprintf("Using credentials from IAM role at: %s", containerCredentialsUri))
	} else if os.Getenv("IS_RUNNING_IN_K8S") == "true" {
		log.Info("Using instance profile AWS auth")
	} else {
		_, hasAccessKey := os.LookupEnv("AWS_ACCESS_KEY_ID")
		require.True(t, hasAccessKey, "must be running in a valid f3 auth shell to create sqs queue. AWS_ACCESS_KEY_ID missing")
		log.Info("Using credentials from F3 shell.")
	}

	region, found := os.LookupEnv("AWS_REGION")
	if !found {
		region = "eu-west-1"
	}

	if os.Getenv("AWS_ENDPOINT") != "" {
		return aws.Config{
			Region:                  aws.String(region),
			DisableComputeChecksums: aws.Bool(true),
			DisableParamValidation:  aws.Bool(true),
			Endpoint:                swag.String(os.Getenv("AWS_ENDPOINT")),
			Credentials:             credentials.NewStaticCredentials("foo", "bar", ""),
		}
	} else {
		return aws.Config{
			Region:                  aws.String(region),
			DisableComputeChecksums: aws.Bool(true),
			DisableParamValidation:  aws.Bool(true),
		}
	}
}
