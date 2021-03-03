package support

import (
	"fmt"

	"github.com/naag/f1-test-utils/pkg/support/rand"
)

func GenerateTransactionReferenceNumber() string {
	return fmt.Sprintf("%018d", rand.Intn(999999999999999999))
}
