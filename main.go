package main

import (
	"fmt"
	stdlog "log"
	"net"
	"os"
	"os/signal"
	"syscall"

	// "net/http"

	// "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"glassgreenhouse.io/plants-service/infrastructure/proto"
	"glassgreenhouse.io/plants-service/infrastructure/transports"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	// "google.golang.org/grpc/credentials/insecure"
)

const (
	port    = ":50051"
	webPort = ":8080"
	banner = `starting plant-service`
)

func main() {	
	logger := log.NewJSONLogger(log.NewSyncWriter(os.Stdout))
	logger = level.NewFilter(logger, level.AllowInfo())
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)
	
	stdlog.SetOutput(log.NewStdlibAdapter(logger))

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGALRM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	lis, err := net.Listen("tcp", port)

	if err != nil {
		level.Error(logger).Log("failed to listen: %v", err)
		os.Exit(1)
	}

	stdlog.Print(banner)

	plantsServer := transports.NewPlantsServer(logger)
	healthCheckServer := transports.NewHealthCheck(logger)

	go func() {
		service := grpc.NewServer()
		proto.RegisterPlantsServer(service, plantsServer)
		grpc_health_v1.RegisterHealthServer(service, healthCheckServer)
		stdlog.Printf("High performance, minimalist Go for %v", lis.Addr())
		level.Info(logger).Log("msg", "Server started successfully 🚀")
		if err := service.Serve(lis); err != nil {
			level.Error(logger).Log("failed to serve: %v", lis.Addr())
		}
	}()

	level.Error(logger).Log("exit", <-errs)

	// go func() {
	// if err := service.Serve(lis); err != nil {
	// 	log.Fatalf("failed to serve: %v", lis.Addr())
	// }
	// }()

	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	// conn, err := grpc.DialContext(
	// 	context.Background(),
	// 	"0.0.0.0"+port,
	// 	grpc.WithBlock(),
	// 	grpc.WithTransportCredentials(insecure.NewCredentials()),
	// )

	// if err != nil {
	// 	log.Fatalln("Failed to dial server:", err)
	// }

	// gwmux := runtime.NewServeMux()

	// if err := proto.RegisterPlantHandler(context.Background(), gwmux, conn); err != nil {
	// 	log.Fatalln("Failed to register gateway:", err)
	// }

	// gwServer := &http.Server{
	// 	Addr:    webPort,
	// 	Handler: gwmux,
	// }

	// log.Println("serving gRPC-Gateway on http://0.0.0.0" + webPort)

	// if err := gwServer.ListenAndServe(); err != nil {
	// 	log.Fatalf("failed to serve: http://0.0.0.0%v", webPort)
	// }
}
