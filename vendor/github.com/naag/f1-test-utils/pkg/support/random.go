package support

import (
	"github.com/naag/f1-test-utils/pkg/support/rand"
)

func RandomReference() string {
	b := make([]byte, 17)
	charset := "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
