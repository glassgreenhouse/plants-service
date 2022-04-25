# plants-service
 single plant registration record


Introduction
---

This is a template for doing go-kit development using GitLab. It's based on the
[shipping](https://github.com/go-kit/examples/shipping) Go Kit template.

### Reference links

- [GitLab CI Documentation](https://docs.github.com/ee/ci/)
- [Go Kit Overview](https://github.com/go-kit/kit)
- [Go Kit Toolkit](https://gokit.io)

### Getting started

First thing to do is update `main.go` with your new project path:

```diff
-       proto "github.com/gitlab-org/project-templates/go-kit/proto"
+       proto "github.com/$YOUR_NAMESPACE/$PROJECT_NAME/proto"
```

Note that these are not actual environment variables, but values you should
replace.

### What's contained in this project

```
.
├──application
│  ├── gateways
│  │   ├── port_in.go       # 
│  │   └── port_out.go      # 
│  ├── usecase
│  │   ├── service.go       # contains the service's core business logic
│  │   └── service.go       # contains the protobuf definition of the API
├──domain
│  ├── plants.go            # contains the endpoint definition
│  ├── temperature.go       # contains the protobuf definition of the API
│  └── depth.go             # our gRPC generated code
├──infrastructure
│  ├── endpoints
│  │   └── endpoints.go     # contains the endpoint definition
│  ├── proto                # contains the protobuf definition of the API
│  │   ├── plant.pb.go      # our gRPC generated code
│  │   └── plant.proto      # our protobuf definitions
│  └── transports
│      ├── grpc.go          # contains the gRPC transport
│      └── rest.go          # contains the rest transport
└── main.go                 # is the main definition of the service, handler and client
```

### Dependencies

Install the following

- [go-kit](https://github.com/go-kit)
- [protoc-gen](https://github.com/golang/protobuf/proto)

### Run Service

```shell
go run main.go
```

### Generate by proto

> Install the protocol compiler plugins for Go using the following commands: 


```shell
$ go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.10.0 \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.10.0  \
    google.golang.org/protobuf/cmd/protoc-gen-go@v1.26 \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
```

> Update your PATH so that the protoc compiler can find the plugins:

```shell
$ export PATH="$PATH:$(go env GOPATH)/bin"
```

Before you can use the new service method, you need to recompile the updated `.proto` file.

While still in the root directory, run the following command:

```shell
protoc --proto_path=./infrastructure/proto \
    --go_out=./infrastructure/proto \
    --go_opt=paths=source_relative \
    --go-grpc_out=./infrastructure/proto \
    --go-grpc_opt=paths=source_relative \
    --grpc-gateway_out ./infrastructure/proto \
    --grpc-gateway_opt paths=source_relative \
    infrastructure/proto/plant.proto
```
