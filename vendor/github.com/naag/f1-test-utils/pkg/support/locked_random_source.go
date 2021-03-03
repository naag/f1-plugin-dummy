package support

import (
	"sync"

	"math/rand"
)

// NewLockedSource provides a thread safe rand source
//
// rand.NewSource is not thread safe, only the global rand source is
// This lifts Go's private implementation used for the global lock source
func NewLockedSource(seed int64) *LockedSource {
	return &LockedSource{
		src: rand.NewSource(seed).(rand.Source64),
	}
}

type LockedSource struct {
	lk  sync.Mutex
	src rand.Source64
}

func (r *LockedSource) Int63() (n int64) {
	r.lk.Lock()
	n = r.src.Int63()
	r.lk.Unlock()
	return
}

func (r *LockedSource) Uint64() (n uint64) {
	r.lk.Lock()
	n = r.src.Uint64()
	r.lk.Unlock()
	return
}

func (r *LockedSource) Seed(seed int64) {
	r.lk.Lock()
	r.src.Seed(seed)
	r.lk.Unlock()
}
