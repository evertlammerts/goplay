package main

import (
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func mst(i, j int, prob []float64, cost [][]float64, root [][]int) {
	var (
		bestCost float64 = 1e9
		bestRoot int     = -1
	)
	switch {
	case i >= j:
		cost[i][j] = 0.0
		root[i][j] = -1
	case i+1 == j:
		cost[i][j] = prob[i]
		root[i][j] = i + 1
	case i+1 < j:
		psum := 0.0
		for k := i; k < j; k++ {
			psum += prob[k]
		}
		for r := i; r < j; r++ {
			rcost := cost[i][r] + cost[r+1][j]
			if rcost < bestCost {
				bestCost = rcost
				bestRoot = r + 1
			}
		}
		cost[i][j] = bestCost + psum
		root[i][j] = bestRoot
	}
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

func main() {
	n, _ := strconv.Atoi(os.Args[1])
	var (
		cost = make([][]float64, n+1)
		root = make([][]int, n+1)
	)
	prob := prob(n)
	for i := n; i >= 0; i-- {
		cost[i] = make([]float64, n+1)
		root[i] = make([]int, n+1)
		for j := i; j <= n; j++ {
			mst(i, j, prob, cost, root)
		}
	}
}
