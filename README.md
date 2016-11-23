# Go Play [![Build Status](https://travis-ci.org/evertlammerts/goplay.svg?branch=master)](https://travis-ci.org/evertlammerts/goplay) [![Coverage Status](https://coveralls.io/repos/github/evertlammerts/goplay/badge.svg?branch=master)](https://coveralls.io/github/evertlammerts/goplay?branch=master)

Playing around with Go.

## `integral.Trapezoid` [![GoDoc](https://godoc.org/github.com/evertlammerts/goplay/integral?status.svg)](https://godoc.org/github.com/evertlammerts/goplay/integral)

```
$ go get github.com/evertlammerts/goplay/integral
```

This package provides an implementation of a parallel trapezoid function for estimating definite integrals of single variable functions. It will try to keep all available cores busy. In the paper you can read about the tradeoff between the overhead of creating a goroutine versus parallel computation. In practice this means you should make sure that this overhead is neglectable compared to the amount of work the goroutine is handling. As a rough heuristic, setting numStrips >= 1e8 should show almost linear scalability.

```
// Estimate Pi with 1e7 trapezoids in the interval [0,1] on the f(x)
pi := integral.Trapezoid(1e7, 0, 1, func(x float64) float64 { return 4.0 / (1.0 + x*x) })
```

## Optimal BST

TODO

