package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"time"

	"google.golang.org/grpc"
	pb "github.com/faizkads/push-notifications/pb"
)

type notifServer struct{
	pb.UnimplementedNotifServer
}

func fetchMeow(count int) ([]string, error) {
	meowurl := fmt.Sprintf("https://meowfacts.herokuapp.com/?count=%d", count)
	resp, err := http.Get(meowurl)
	if err != nil {
		return nil, fmt.Errorf("fetch error: %v", err)
	}
	defer resp.Body.Close()

	var data map[string][]string
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response error: %v", err)
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, fmt.Errorf("unmarshal error: %v", err)
	}

	return data["data"],nil
}

func (s *notifServer) StreamMeow(req *pb.MeowReq, stream pb.Notif_StreamMeowServer) error{
	if req.Count < 1 {
		return fmt.Errorf("request must include at least one time")
	}
	log.Printf("Received request for: %d iteration", req.Count)
	count := req.Count
	facts, err := fetchMeow(int(count))
	if err != nil {
		return err
	}
	i := 1
	for _,fact := range facts {
		// log.Printf("%v",fact)
		log.Printf("Meow Fact %d Sent",i)
		stream.Send(&pb.MeowRes{Fact: fact})
		i += 1
		time.Sleep(1*time.Second)
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterNotifServer(grpcServer, &notifServer{})
	log.Printf("Starting notification server on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}