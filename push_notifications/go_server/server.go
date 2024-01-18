package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"time"

	pb "github.com/faizkads/push-notifications/pb"
	"google.golang.org/grpc"
)

type notifServer struct{
	pb.UnimplementedNotifServer
}

func fetchMeow(count int) ([]string, error) {
	// set the url query corresponding to the 'count' variable
	meowurl := fmt.Sprintf("https://meowfacts.herokuapp.com/?count=%d", count)
	resp, err := http.Get(meowurl) // http get method to fetch response
	if err != nil {
		return nil, fmt.Errorf("fetch error: %v", err)
	}
	defer resp.Body.Close()

	var data map[string][]string
	body, err := io.ReadAll(resp.Body) // read the body of the response
	if err != nil {
		return nil, fmt.Errorf("read response error: %v", err)
	}

	err = json.Unmarshal(body, &data) // unmarshal body dan point ke memory address data
	if err != nil {
		return nil, fmt.Errorf("unmarshal error: %v", err)
	}

	return data["data"],nil
}

func fetchCurrency(url string) (*pb.CurRes, error) {
	resp, err := http.Get(url) // http get method to fetch response
	if err != nil {
		return nil, fmt.Errorf("fetch error: %v", err)
	}
	defer resp.Body.Close()

	var result pb.CurRes
	body, err := io.ReadAll(resp.Body) // read the body of the response
	if err != nil {
		return nil, fmt.Errorf("read response error: %v", err)
	}

	err = json.Unmarshal(body, &result) // unmarshal body dan point ke memory address data
	if err != nil {
		return nil, fmt.Errorf("unmarshal error: %v", err)
	}

	return &result, nil
}

func (s *notifServer) StreamMeow(req *pb.MeowReq, stream pb.Notif_StreamMeowServer) error{
	count := req.Count
	log.Printf("Received request for Meow Facts with [%d] iteration", count)
	facts, err := fetchMeow(int(count))
	if err != nil {
		return err
	}
	i := 1
	for _,fact := range facts {
		log.Printf("Meow Fact %d Sent",i)
		stream.Send(&pb.MeowRes{Fact: fact})
		i += 1
		time.Sleep(5*time.Second) // delay to simulate periodical iteration
	}
	return nil
}

func (s *notifServer) StreamCurrency(req *pb.CurReq, stream pb.Notif_StreamCurrencyServer) error {
	from := req.FromCur
	to := req.ToCur
	count := req.Count

	log.Printf("Received exchange rate from [%s] to [%s] for: [%d] iteration", from, to, count)

	date := time.Now() // the real date
	dateFake := date.AddDate(0,-6,0) // fake date to simulate real-time streaming

	for i := 0; i < int(count); i++ {
		// increment dates
		dateString := date.AddDate(0,0,i).Format("2006-01-02") 
		fakeString := dateFake.AddDate(0,0,i).Format("2006-01-02")

		// modify the url corresponding to the query
		url := fmt.Sprintf("https://api.frankfurter.app/%s?from=%s&to=%s", fakeString, from, to)

		result, err := fetchCurrency(url) 
		if err != nil {
			return err
		}

		response := &pb.CurRes{
			Base:	result.Base,
			Date: 	dateString,
			Rates:	result.Rates,
		}

		err = stream.Send(response) // send the stream response
		if err != nil {
			return err
		}
		log.Printf("Exchange rate from %s to %s on %s is Sent", from, to, dateString)

		time.Sleep(5*time.Second) // delay to simulate periodical iteration
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