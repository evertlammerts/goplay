package bst

import (
	"math"
	"math/rand"
	"testing"
	"time"
)

func TestSequential(t *testing.T) {
	prob := prob(2e2)
	root, cost := Sequential(prob)
	t.Logf("root=%d cost=%f \n", root, cost)
}

func prob(n int) []float64 {
	var (
		prob = make([]float64, n)
		sum  float64
		min  float64 = math.MaxFloat64
	)
	// Seed the PRNG
	rand.Seed(time.Now().UnixNano())
	// Generate float64's on a normal distribution (mean=0, stdev=1)
	for i := 0; i < n; i++ {
		prob[i] = rand.NormFloat64()
		if prob[i] < min {
			min = prob[i]
		}
	}
	// Normalize to positive numbers
	for i := 0; i < n; i++ {
		prob[i] -= min
		sum += prob[i]
	}
	// Turn into probabilities that (appr) sum up to 1
	for i := 0; i < n; i++ {
		prob[i] /= sum
	}
	return prob
}
