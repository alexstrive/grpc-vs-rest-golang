package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	pb "stats"
	"time"

	"google.golang.org/grpc"
)

const (
	address = "localhost:5000"
)

func TestRest(path string) {
	var arr []*pb.CovidCaseStatEntry
	resp, err := http.Get(fmt.Sprintf("http://localhost:8080/%v", path))
	if err != nil {
		log.Printf("Couldn't get resource: %v", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Couldn't read content: %v", err)
	}

	json.Unmarshal(body, &arr)
}

func TestGetAllCovidCasesGrpc() {
	conn, err := grpc.Dial(address, grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(1024*1024*8)), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewStatsClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err = c.GetAllCovidCases(ctx, &pb.Empty{})

	if err != nil {
		log.Fatalf("Could not get message: %v", err)
	}
}
