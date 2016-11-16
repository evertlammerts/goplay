# Multi-Core Parallel Programming in Go [![Build Status](https://travis-ci.org/evertlammerts/mcppg.svg?branch=master)](https://travis-ci.org/evertlammerts/mcppg)

Implementation of two parallel algorithms described in "[Multi-Core Parallel Programming in Go](http://www.ualr.edu/pxtang/papers/acc10.pdf)" (Peiyi Tang, 2010).

## Integral.Trapezoid [![GoDoc](https://godoc.org/github.com/evertlammerts/mcppg/integral?status.svg)](https://godoc.org/github.com/evertlammerts/mcppg/integral)

This package provides an implementation of a parallel trapezoid function for estimating definite integrals of single variable functions. It will try to keep all available cores busy. See the paper for implementation details and parameter tuning.

```
// Estimate Pi with 1e7 trapezoids in the interval [0,1] on the f(x)
pi := Trapezoid(1e7, 0, 1, func(x float64) float64 { return 4.0 / (1.0 + x*x) })
```

## Issues

I'd love to hear about them on the github issue tracker.
