package integral

import "runtime"

// Trapezoid computes an estimate of a definite integral of a function f(x) in parallel.
//
// The trapezoidal rule works by approximating the region under the graph of
// the function f(x) as a trapezoid and calculating its area. This implementation
// computes the height of rectangles based on the center of an interval, rather than
// the average of the interval's start and end
//
// numStrips is the number of rectangles that will be used. More rectangles provides
// more accurate results. Note that since the operation is parallelized over all
// runtime.NumCPU() cores, you should tune this for optimality. As a rule of thumb,
// use numStrips >= 10e8 for better scalability.
//
// intervalStart is the start of the interval over which we're integrating
//
// intervalEnd is the end of the interval over which we're integrating
//
// f is the function for which we're estimating the integral
func Trapezoid(numStrips int64, intervalStart, intervalEnd float64, f func(x float64) float64) (result float64) {
	np := runtime.NumCPU()
	c := make(chan float64, numStrips)
	h := (intervalEnd - intervalStart) / float64(numStrips) // the width of a rectangle
	// Parallelize the work over all np cores
	for i := 0; i < np; i++ {
		start := int64(i) * numStrips / int64(np)
		end := (int64(i) + 1) * numStrips / int64(np)
		go func() {
			var sum float64 = 0.0
			for j := start; j < end; j++ {
				x := h*(float64(j)+0.5) + intervalStart
				sum += f(x)
			}
			c <- sum * h

		}()
	}
	for i := 0; i < np; i++ {
		result += <-c
	}
	return
}
