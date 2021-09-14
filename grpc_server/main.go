package main

import (
	"context"
	"log"
	"net"

	pb "stats"

	"google.golang.org/grpc"
	_ "google.golang.org/grpc/encoding/gzip"
)

type statsServer struct {
	pb.UnimplementedStatsServer
}

func (s *statsServer) GetAllCovidCases(ctx context.Context, in *pb.Empty) (*pb.CovidCaseStatEntryList, error) {
	return &pb.CovidCaseStatEntryList{Entries: pb.CovidCases}, nil
}

func (s *statsServer) GetAllCovidCasesStream(empty *pb.Empty, stream pb.Stats_GetAllCovidCasesStreamServer) error {
	for _, entry := range pb.CovidCases {
		if err := stream.Send(entry); err != nil {
			return err
		}
	}

	return nil
}

func (s *statsServer) GetAllStocks(ctx context.Context, in *pb.Empty) (*pb.StockList, error) {
	return &pb.StockList{Stocks: pb.Stocks}, nil
}

func (s *statsServer) GetAllStocksStream(empty *pb.Empty, stream pb.Stats_GetAllStocksStreamServer) error {
	for _, stock := range pb.Stocks {
		if err := stream.Send(stock); err != nil {
			return err
		}
	}

	return nil
}

func (s *statsServer) GetAllVaccineEntries(ctx context.Context, in *pb.Empty) (*pb.VaccineEntryList, error) {
	return &pb.VaccineEntryList{Entries: pb.VaccineEntries}, nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:5000")
	if err != nil {
		log.Fatalf("Failed to start listen: %v", err)
	}

	server := grpc.NewServer(grpc.MaxRecvMsgSize(1024*1024*8), grpc.MaxSendMsgSize(1024*1024*20))
	pb.RegisterStatsServer(server, &statsServer{})
	log.Printf("Server listening at %v", lis.Addr())

	server.Serve(lis)
}
