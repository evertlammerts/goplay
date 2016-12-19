# Go Play [![Build Status](https://travis-ci.org/evertlammerts/goplay.svg?branch=amsterdam)](https://travis-ci.org/evertlammerts/goplay?branch=amsterdam) [![Coverage Status](https://coveralls.io/repos/github/evertlammerts/goplay/badge.svg?branch=amsterdam)](https://coveralls.io/github/evertlammerts/goplay?branch=amsterdam)

Playing around with Go, showoff version for the municipality of Amsterdam.

This version allows you to offload the computations needed for estimating a definite integral of any single term polynomial to some server.

* Computing the polynomial (multi-core) is done in the `integral` package
* The protobuf files are in `/integrator`
* The server implementation can be found in `stiserver` and the client in `sticlient`

Since this is a branch this doesn't play nice with `go get`. To run this you'll have to clone the repo and get the dependencies, then build the client and the server:

```
git clone https://github.com/evertlammerts/goplay.git
go get google.golang.org/grpc
go get golang.org/x/net/context
go get github.com/evertlammerts/goplay/integral
cd goplay/stiserver
go build
cd ../sticlient
go build
cd ..
```

Now, if all is well, you can run the server:

```
./stiserver/stiserver
```

And the client:

```
./sticlient/client -help
```

## `integral.Trapezoid` [![GoDoc](https://godoc.org/github.com/evertlammerts/goplay/integral?status.svg)](https://godoc.org/github.com/evertlammerts/goplay/integral)

```
$ go get github.com/evertlammerts/goplay/integral
```

This package provides an implementation of a parallel trapezoid function for estimating definite integrals of single variable functions. It will try to keep all available cores busy. In the paper you can read about the tradeoff between the overhead of creating a goroutine versus parallel computation. In practice this means you should make sure that this overhead is neglectable compared to the amount of work the goroutine is handling. As a rough heuristic, setting numStrips >= 1e8 should show almost linear scalability.

```
// Estimate Pi with 1e7 trapezoids in the interval [0,1] on the f(x)
pi := integral.Trapezoid(1e7, 0, 1, func(x float64) float64 { return 4.0 / (1.0 + x*x) })
```


