package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"glassgreenhouse.io/plants-service/infrastructure/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"
)

const (
	port    = ":50051"
	webPort = ":8080"
)

const (
	reset  = "\033[0m"
	purple = "\033[35m"
	banner = `planet-service`
)

type Plant struct {
	proto.UnimplementedPlantServer
}

type HealthCheck struct {
	proto.UnimplementedPlantServer
}

func (g *Plant) Hello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloResponse, error) {
	log.Printf("Received: %v", req.GetName())
	return &proto.HelloResponse{Greeting: "Hello " + req.Name}, nil
}

func (s *HealthCheck) Check(ctx context.Context, in *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	log.Printf("Received: %v", in.GetService())
	return &grpc_health_v1.HealthCheckResponse{Status: grpc_health_v1.HealthCheckResponse_SERVING}, nil
}

func (s *HealthCheck) Watch(in *grpc_health_v1.HealthCheckRequest, _ grpc_health_v1.Health_WatchServer) error {
	return status.Error(codes.Unimplemented, "unimplemented")
}

func main() {
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	service := grpc.NewServer()

	proto.RegisterPlantServer(service, &Plant{})

	grpc_health_v1.RegisterHealthServer(service, &HealthCheck{})

	log.Print(purple + banner + reset)
	log.Printf("High performance, minimalist Go for %v", lis.Addr())

	go func() {
		if err := service.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", lis.Addr())
		}
	}()

	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.DialContext(context.Background(), "0.0.0.0"+port, grpc.WithBlock(), grpc.WithInsecure())

	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()

	if err := proto.RegisterPlantHandler(context.Background(), gwmux, conn); err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    webPort,
		Handler: gwmux,
	}

	log.Println("serving gRPC-Gateway on http://0.0.0.0" + webPort)

	if err := gwServer.ListenAndServe(); err != nil {
		log.Fatalf("failed to serve: http://0.0.0.0%v", webPort)
	}
}
