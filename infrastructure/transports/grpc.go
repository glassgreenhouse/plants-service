package transports

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"glassgreenhouse.io/plants-service/infrastructure/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"
)

type plantServer struct {
	logger log.Logger
	proto.UnimplementedPlantsServer
}

type healthCheck struct {
    logger log.Logger
}

// NewGRPCServer initializes a new gRPC server
func NewPlantsServer(logger log.Logger) proto.PlantsServer {
	return &plantServer{
		logger: logger,
	}
}

// NewHealthCheck initializes a new gRPC server
func NewHealthCheck(logger log.Logger) grpc_health_v1.HealthServer {
	return &healthCheck{
		logger: logger,
	}
}

func (server *plantServer) NewPlant(ctx context.Context, req *proto.NewPlantRequest) (*proto.Plant, error) {
	level.Info(server.logger).Log("Received: %v", req.GetName())
	return &proto.Plant{Name: "Hello " + req.Name}, nil
}

func (server *healthCheck) Check(ctx context.Context, in *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	level.Info(server.logger).Log("Received: %v", in.GetService())
	return &grpc_health_v1.HealthCheckResponse{Status: grpc_health_v1.HealthCheckResponse_SERVING}, nil
}

func (server *healthCheck) Watch(in *grpc_health_v1.HealthCheckRequest, _ grpc_health_v1.Health_WatchServer) error {
	return status.Error(codes.Unimplemented, "unimplemented")
}