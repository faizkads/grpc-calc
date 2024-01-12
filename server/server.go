package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "github.com/faizkads/grpc-calc/proto"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedCalculatorServer
}

// Function for addition log and operation
func (s *server) Add(ctx context.Context, in *pb.CalcReq) (*pb.AddRes, error) {
	log.Printf("Received: %.2f + %.2f", in.N1, in.N2)
	result := in.N1 + in.N2
	return &pb.AddRes{R: result}, nil
}

// Function for substraction log and operation
func (s *server) Sub(ctx context.Context, in *pb.CalcReq) (*pb.SubRes, error) {
	log.Printf("Received: %.2f - %.2f", in.N1, in.N2)
	result := in.N1 - in.N2
	return &pb.SubRes{R: result}, nil
}

// Function for multiplication log and operation
func (s *server) Mul(ctx context.Context, in *pb.CalcReq) (*pb.MulRes, error) {
	log.Printf("Received: %.2f * %.2f", in.N1, in.N2)
	result := in.N1 * in.N2
	return &pb.MulRes{R: result}, nil
}

// Function for division log and operation
func (s *server) Div(ctx context.Context, in *pb.CalcReq) (*pb.DivRes, error) {
	log.Printf("Received: %.2f / %.2f", in.N1, in.N2)
	if in.N2 == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "Cannot divide by zero")
	}
	result := in.N1 / in.N2
	return &pb.DivRes{R: result}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %s", err)
	}
	s := grpc.NewServer()
	pb.RegisterCalculatorServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}