package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	pb "github.com/fakegermano/consumption/proto"
	"github.com/influxdata/influxdb-client-go/v2"
	"google.golang.org/grpc"
)

func getenv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

var (
	port, _ = strconv.Atoi(getenv("DATASERVER_PORT", "50051"))
	start   = getenv("DATABASE_START", "2019-01-01T00:00:00Z")
	stop    = getenv("DATABASE_STOP", "2019-02-01T00:00:00Z")
	bucket  = getenv("DATABASE_TABLE_NAME", "electricity")
	org     = getenv("DATABASE_DB_NAME", "org")
	address = getenv("DATABASE_URL", "http://localhost:8086")
	token   = getenv("DATABASE_TOKEN", "BK6tSeFCSRQ7vRzHejGF9SZ837fp2UMn")
)

type MeterUsageServer struct {
	pb.UnimplementedMeterUsageServiceServer
	repository *MeterUsageRepository
}

type MeterUsageRepository struct {
	influxdb2.Client
}

func NewMeterUsageRepository() *MeterUsageRepository {
	return &MeterUsageRepository{
		influxdb2.NewClient(address, token),
	}
}

func (r *MeterUsageRepository) Close() {
	r.Close()
}

func (r *MeterUsageRepository) Get(limit, offset int32) ([]*pb.MeterUsage, error) {
	query := r.QueryAPI(org)
	query_string := fmt.Sprintf("from(bucket:\"%s\") |> range(start: %s, stop: %s) |> limit(n: %d, offset: %d)",
		bucket, start, stop, limit, offset)
	result, err := query.Query(context.Background(), query_string)
	if err != nil {
		return nil, err
	}
	usages := make([]*pb.MeterUsage, 0)
	for result.Next() {
		usages = append(usages, &pb.MeterUsage{
			Time:  result.Record().Time().Unix(),
			Usage: result.Record().Value().(float64),
		})
	}
	return usages, nil
}

func NewMeterUsageServer(repository *MeterUsageRepository) *MeterUsageServer {
	return &MeterUsageServer{
		repository: repository,
	}
}

func (s *MeterUsageServer) Close() {
	s.repository.Close()
}

func (s *MeterUsageServer) Get(ctx context.Context, in *pb.MeterUsageRequest) (*pb.MeterUsageResponse, error) {
	if in.GetLimit() < 0 || in.GetOffset() < 0 {
		return nil, errors.New("limit and offset cannot be negative")
	}
	usages, err := s.repository.Get(in.GetLimit(), in.GetOffset())
	return &pb.MeterUsageResponse{
		Usages: usages,
	}, err
}

func (s *MeterUsageServer) Run() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterMeterUsageServiceServer(server, s)
	log.Printf("server listening at %v", lis.Addr())
	return server.Serve(lis)

}
func main() {
	meter_usage_server := NewMeterUsageServer(NewMeterUsageRepository())
	defer meter_usage_server.Close()
	if err := meter_usage_server.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
