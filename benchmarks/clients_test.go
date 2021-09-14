package benchmarks

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	pb "stats"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/encoding/gzip"
)

const (
	address = "localhost:5000"
)

func makeGrpcGzipConn() (*grpc.ClientConn, error) {
	return grpc.Dial(address, grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(1024*1024*8), grpc.UseCompressor(gzip.Name)), grpc.WithInsecure(), grpc.WithBlock())
}

func makeGrpcConn() (*grpc.ClientConn, error) {
	return grpc.Dial(address, grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(1024*1024*8), grpc.UseCompressor(gzip.Name)), grpc.WithInsecure(), grpc.WithBlock())
}

func MakeGrpcCall(withGzip bool, callback func(ctx context.Context, client pb.StatsClient)) {
	var Conn *grpc.ClientConn
	var Err error

	if withGzip {
		conn, err := makeGrpcGzipConn()
		Conn = conn
		Err = err
	} else {
		conn, err := makeGrpcConn()
		Conn = conn
		Err = err
	}

	if Err != nil {
		log.Fatalf("Unable to connect: %v", Err)
	}
	defer Conn.Close()

	c := pb.NewStatsClient(Conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	callback(ctx, c)
}

func MakeRestRequest(path string) {
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

func BenchmarkStocks_Grpc_20MB(b *testing.B) {
	for n := 0; n < b.N; n++ {
		MakeGrpcCall(false, func(ctx context.Context, client pb.StatsClient) {
			_, err := client.GetAllStocks(ctx, &pb.Empty{})
			if err != nil {
				log.Fatalf("Could not get stocks: %v", err)
			}
		})
	}
}

func BenchmarkStocks_Grpc_20MB_gzip(b *testing.B) {
	for n := 0; n < b.N; n++ {
		MakeGrpcCall(true, func(ctx context.Context, client pb.StatsClient) {
			_, err := client.GetAllStocks(ctx, &pb.Empty{})
			if err != nil {
				log.Fatalf("Could not get stocks: %v", err)
			}
		})
	}
}

func BenchmarkStocks_Rest_20MB(b *testing.B) {
	for n := 0; n < b.N; n++ {
		MakeRestRequest("stocks")
	}
}

func BenchmarkStocks_Rest_20MB_gzip(b *testing.B) {
	for n := 0; n < b.N; n++ {
		MakeRestRequest("stocksGzip")
	}
}

func BenchmarkCovidCases_Grpc_345KB_gzip(b *testing.B) {
	for n := 0; n < b.N; n++ {
		MakeGrpcCall(true, func(ctx context.Context, client pb.StatsClient) {
			_, err := client.GetAllCovidCases(ctx, &pb.Empty{})
			if err != nil {
				log.Fatalf("Could not get stocks: %v", err)
			}
		})
	}
}

func BenchmarkCovidCases_Grpc_345KB(b *testing.B) {
	for n := 0; n < b.N; n++ {
		MakeGrpcCall(false, func(ctx context.Context, client pb.StatsClient) {
			_, err := client.GetAllCovidCases(ctx, &pb.Empty{})
			if err != nil {
				log.Fatalf("Could not get stocks: %v", err)
			}
		})
	}
}

func BenchmarkCovidCases_Rest_345KB(b *testing.B) {
	for n := 0; n < b.N; n++ {
		MakeRestRequest("covidCases")
	}
}

func BenchmarkCovidCases_Rest_345KB_gzip(b *testing.B) {
	for n := 0; n < b.N; n++ {
		MakeRestRequest("covidCasesGzip")
	}
}
