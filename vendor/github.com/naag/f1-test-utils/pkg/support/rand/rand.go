package rand

import (
	crand "crypto/rand"
	"math/big"
)

func Intn(max int) int {
	return int(Int63n(int64(max)))
}

func Int63n(max int64) int64 {
	m := big.NewInt(max)
	r, err := crand.Int(crand.Reader, m)
	if err != nil {
		panic(err)
	}
	return r.Int64()
}
