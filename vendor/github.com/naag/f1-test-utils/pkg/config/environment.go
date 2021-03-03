package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/form3tech-oss/f1/pkg/f1/testing"
	"github.com/go-openapi/strfmt"
)

// prefer the specific int size version GetIntOrDefault
func GetNumberOrDefault(t *testing.T, name string, defaultValue int) int {
	stringValue := GetStringOrDefault(t, name, fmt.Sprintf("%d", defaultValue))

	intValue, err := strconv.ParseInt(stringValue, 10, 32)
	t.Require.Nil(err)
	return int(intValue)
}

func GetOptionalNumber(t *testing.T, name string) *int {
	stringValue, err := GetString(t, name)
	if err != nil {
		return nil
	}

	intValue, err := strconv.ParseInt(stringValue, 10, 32)
	t.Require.Nil(err)
	return Int(int(intValue))
}

func GetIntOrDefault(t *testing.T, name string, defaultValue int64) int64 {
	stringValue := GetStringOrDefault(t, name, fmt.Sprintf("%d", defaultValue))

	intValue, err := strconv.ParseInt(stringValue, 10, 64)
	t.Require.Nil(err)
	return intValue
}

func GetFloatOrDefault(t *testing.T, name string, defaultValue float64) float64 {
	stringValue := GetStringOrDefault(t, name, fmt.Sprintf("%f", defaultValue))

	intValue, err := strconv.ParseFloat(stringValue, 64)
	t.Require.Nil(err)
	return intValue
}

func GetStringOrDefault(t *testing.T, name string, defaultValue string) string {
	value, err := GetString(t, name)
	if err != nil {
		return defaultValue
	}

	return value
}

func GetDurationOrDefault(t *testing.T, name string, defaultValue *time.Duration) *time.Duration {
	value, err := GetString(t, name)
	if err != nil {
		return defaultValue
	}

	duration, err := time.ParseDuration(value)
	if err != nil {
		t.Require.NoError(err)
	}

	return &duration
}

func GetStringSliceOrDefault(t *testing.T, name string, defaultValue []string) []string {
	if value, err := GetString(t, name); err == nil {
		return strings.Split(value, ",")
	}
	return defaultValue
}

func GetUUIDsOrDefault(t *testing.T, name string, defaultValue []strfmt.UUID) []strfmt.UUID {
	if value, err := GetString(t, name); err == nil {
		split := strings.Split(value, ",")
		var uuids []strfmt.UUID
		for _, s := range split {
			uuids = append(uuids, strfmt.UUID(s))
		}
		return uuids
	}
	return defaultValue
}

func GetDateOrDefault(t *testing.T, name string, defaultValue time.Time) time.Time {
	if value, err := GetString(t, name); err == nil {
		date, err := time.Parse(time.RFC3339, value)
		t.Require.Nil(err)
		return date
	}
	return defaultValue
}

func GetBoolOrDefault(t *testing.T, name string, defaultValue bool) bool {
	if value, err := GetString(t, name); err == nil {
		b, err := strconv.ParseBool(value)
		t.Require.Nil(err)
		return b
	}
	return defaultValue
}

func GetString(t *testing.T, name string) (string, error) {
	k6EnvName := fmt.Sprintf("K6_TESTS_%s", name)
	if t != nil {
		value, ok := t.Environment[k6EnvName]
		if ok && value != "" {
			return value, nil
		}
	}

	if val := os.Getenv(name); val != "" {
		return val, nil
	}

	return "", fmt.Errorf("could not find env var %s or %s", k6EnvName, name)
}

func MustGetString(t *testing.T, name string) string {
	value, err := GetString(t, name)
	if err != nil {
		t.Errorf("missing environment variable %s", name)
		t.FailNow()
	}
	return value
}

// Int returns a pointer to the int value passed in.
func Int(v int) *int {
	return &v
}

var (
	// Tests will have to wait for the final poll to compete during teardown, so this is lower than the default 20s.
	SqsLongPollingWaitTime = int64(5)
)
