package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"

	"golang.org/x/net/context"

	pb "github.com/evertlammerts/goplay/integrator"
	"google.golang.org/grpc"
)

const address = "localhost:50505"

type polyConstants []int32

func (i *polyConstants) String() string {
	return fmt.Sprint(*i)
}

func (i *polyConstants) Set(value string) error {
	if len(*i) > 0 {
		return errors.New("constants flag already set")
	}
	for _, c := range strings.Split(value, ",") {
		constant, err := strconv.Atoi(c)
		if err != nil {
			return err
		}
		*i = append(*i, int32(constant))
	}
	return nil
}

var (
	server        = flag.String("server", "localhost:50505", "Host and port of the server")
	numStrips     = flag.Int64("strips", 1e6, "the amount of strips to use for the integral estimate")
	intervalStart = flag.Float64("start", 0, "start of the interval")
	intervalEnd   = flag.Float64("end", 1, "end of the interval")
	constants     polyConstants
)

func main() {
	// Parse CLI arguments
	flag.Var(&constants, "constants", "comma-seperated list of constants to use for the polynomial (default 1,0)")
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewIntegratorClient(conn)
	// Create the constants array for the polynomial
	if len(constants) == 0 {
		constants = polyConstants{1, 0}
	}
	polynomial := pb.STPolynomial{Constants: constants}
	// Create the integral config
	conf := pb.IntegralConfig{NumStrips: *numStrips, Start: *intervalStart, End: *intervalEnd}
	// Create the query
	query := pb.Query{Poly: &polynomial, Conf: &conf}
	integral, err := c.SingleTermPolynomialIntegral(context.Background(), &query)
	if err != nil {
		log.Fatalf("could not compute integral: %v", err)
	}
	log.Printf("Result: %v", integral.Result)
}
