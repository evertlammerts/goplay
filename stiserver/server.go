// Package main implements a simple gRPC server that allows you to offload the
// computations involved in estimating definite integrals on single term
// polynomials.
//
// Look at the integrator/integrator.proto definition for the specification of
// the interface.
package main

import (
	"log"
	"math"
	"net"

	"google.golang.org/grpc"

	"golang.org/x/net/context"

	"github.com/evertlammerts/goplay/integral"
	pb "github.com/evertlammerts/goplay/integrator"
)

// The port to start the server on
const port = ":50505"

// A type to implement the server
type integratorServer struct{}

// Polynomial returns a function based on the constants given in the
// STPolynomial instance
func (s *integratorServer) polynomial(p *pb.STPolynomial) func(float64) float64 {
	return func(x float64) (result float64) {
		for i, constant := range p.Constants {
			result += float64(constant) * math.Pow(x, float64(i))
		}
		return
	}
}

// Estimate the integral of the given single term polynomial over the given
// interval, using the given amount of strips
func (s *integratorServer) SingleTermPolynomialIntegral(ctx context.Context, q *pb.Query) (*pb.IntegralEstimate, error) {
	return &pb.IntegralEstimate{
		Result: integral.Trapezoid(q.Def.NumStrips, q.Def.Start, q.Def.End, s.polynomial(q.Poly)),
	}, nil
}

func main() {
	// Create the TCP port listener
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// Create the server
	s := grpc.NewServer()
	pb.RegisterIntegratorServer(s, &integratorServer{})
	// Start listening
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
