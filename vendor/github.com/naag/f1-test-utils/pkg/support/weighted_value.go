package support

import (
	"sync/atomic"

	"math/rand"
)

type WeightedValueChooser interface {
	Next() string
}
type randomWeightedValues struct {
	used   int64
	values []weightedValue
}

type weightedValue struct {
	value         string
	currentWeight float64
	targetWeight  float64
	used          int64
}

func NewRandomWeightedValues(weights map[string]float64) WeightedValueChooser {
	var values []weightedValue
	for value, weight := range weights {
		values = append(values, weightedValue{
			value:         value,
			currentWeight: weight,
			targetWeight:  weight,
			used:          0,
		})
	}
	return &randomWeightedValues{
		used:   0,
		values: values,
	}
}

func (wv *randomWeightedValues) Next() string {
	position := rand.Float64() // #nosec G404 (CWE-338): Use of weak random number generator (math/rand instead of crypto/rand)
	weight := 0.0
	for _, val := range wv.values {
		weight = weight + val.currentWeight
		if position <= weight {
			atomic.AddInt64(&val.used, 1)
			used := atomic.AddInt64(&wv.used, 1)
			if used%100 == 0 {
				wv.adjustWeights()
			}
			return val.value
		}
	}
	return ""
}

func (wv *randomWeightedValues) adjustWeights() {
	used := atomic.LoadInt64(&wv.used)
	for _, value := range wv.values {
		expectedTotal := value.targetWeight * float64(used)
		variance := float64(atomic.LoadInt64(&value.used)) / expectedTotal
		value.currentWeight = value.targetWeight * variance
	}
}

type equalWeightedValues struct {
	position int64
	values   []string
}

func NewEquallyDistributedWeightedValues(weights map[string]int) WeightedValueChooser {
	var values []string
	for value, count := range weights {
		for i := 0; i < count; i++ {
			values = append(values, value)
		}
	}
	rand.Shuffle(len(values), func(i, j int) { values[i], values[j] = values[j], values[i] })

	return &equalWeightedValues{
		position: 0,
		values:   values,
	}
}

func (e *equalWeightedValues) Next() string {
	pos := atomic.AddInt64(&e.position, 1)
	return e.values[pos%int64(len(e.values))]
}
