package main

const n int = 1e6

var (
	cost [n + 1][n + 1]float64
	root [n + 1][n + 1]int
	prob [n]float64
)

func mst(i, j int) {
	var bestCost float64 = 1e9
	var bestRoot int = -1
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

func main() {

}
