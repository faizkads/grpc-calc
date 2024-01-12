package main

import (
	"context"
	"flag"
	"log"
	"time"
	"os"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "github.com/faizkads/grpc-calc/proto"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

// Parse input arguments as integers of n1 and n2
func argParser(n1 string, n2 string) (float32, float32) {
	N1, err := strconv.ParseFloat(os.Args[1], 32)
	if err != nil {
		log.Fatalf("Cannot parse arg[1]: arguments should be in float32 and in range")
	}
	N2, err := strconv.ParseFloat(os.Args[2], 32)
	if err != nil {
		log.Fatalf("Cannot parse arg[2]: arguments should be float32 and in range")
	}
	return float32(N1), float32(N2)
}

func main() {
	// Check arguments
	if len(os.Args) != 3 {
		log.Fatalf("2 Arguments Needed")
	}

	// Parse input argument
	n1,n2 := argParser(os.Args[1], os.Args[2])
	
	// define connection credentials
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// initiate client connection
	c := pb.NewCalculatorClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Perform addition of n1 and n2 in the server and return the result
	AddRes, err := c.Add(ctx, &pb.CalcReq{N1:float32(n1), N2:float32(n2)})
	if err != nil {
		log.Fatalf("Addition error: %s", err)
	}
	log.Printf("%.2f + %.2f = %.2f", n1, n2, AddRes.R) // result of division diplayed as float with 2 decimal places

	// Perform substraction of n1 and n2 in the server and return the result
	SubRes, err := c.Sub(ctx, &pb.CalcReq{N1:float32(n1), N2:float32(n2)})
	if err != nil {
		log.Fatalf("Substraction error: %s", err)
	}
	log.Printf("%.2f - %.2f = %.2f", n1, n2, SubRes.R)

	// Perform multiplication of n1 and n2 in the server and return the result
	MulRes, err := c.Mul(ctx, &pb.CalcReq{N1:float32(n1), N2:float32(n2)})
	if err != nil {
		log.Fatalf("Multiplication error: %s", err)
	}
	log.Printf("%.2f * %.2f = %.2f", n1, n2, MulRes.R)

	// Perform division of n1 and n2 in the server and return the result
	DivRes, err := c.Div(ctx, &pb.CalcReq{N1:float32(n1), N2:float32(n2)})
	if err != nil {
		log.Fatalf("Division error: %s", err)
	}
	log.Printf("%.2f / %.2f = %.2f", n1, n2, DivRes.R) 
}