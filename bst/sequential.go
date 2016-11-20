package bst

func Sequential(prob []float64) (int, float64) {
	n := len(prob)
	var (
		cost = make([][]float64, n+1)
		root = make([][]int, n+1)
	)
	for i := n; i >= 0; i-- {
		cost[i] = make([]float64, n+1)
		root[i] = make([]int, n+1)
		for j := i; j <= n; j++ {
			mst(i, j, prob, cost, root)
		}
	}
	return root[0][n], cost[0][n]
}

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
