# Quick Start Guide for gRPC PoC with Go and Kubernetes

# What's in this README

- Quick Start Guide
- Detailed Explanation

# Quick Start Guide
This gRPC Proof of Concept (PoC) project demonstrates a simple client-server communication using gRPC. The server provides a `SayHello` RPC method that takes a `HelloRequest` containing a name and responds with a `HelloReply` containing a greeting message. The client sends multiple requests to the server and logs the responses.

### Project Structure
- **Server**:
  - Implements the `Greeter` gRPC service.
  - Defined in greeter.proto.
  - Server code is in main.go.
  - Dockerfile for building the server image is in Dockerfile.

- **Client**:
  - Sends requests to the `Greeter` service.
  - Defined in greeter.proto.
  - Client code is in main.go.
  - Dockerfile for building the client image is in Dockerfile.

- **Kubernetes**:
  - Deployment and service YAML files for both client and server are in the k8s folder.

---

### Steps to Get It Up and Running

#### 1. **Generate gRPC Code**
Ensure you have `protoc` and the Go plugins installed. Run the following commands to generate the gRPC code:
```bash
protoc --go_out=. --go-grpc_out=. greeter.proto
```
Run this in both the server and client directories to generate the `pb` files.

---

#### 2. **Build the Go Applications**
Navigate to the respective directories and build the binaries:
```bash
# Server
cd server
go build -o server .

# Client
cd ../client
go build -o client .
```

---

#### 3. **Build Docker Images**
Use the provided Dockerfiles to build the Docker images for the server and client:
```bash
# Server
cd server
docker build -t grpc-server:latest .

# Client
cd ../client
docker build -t grpc-client:latest .
```

---

#### 4. **Deploy to Kubernetes**
Apply the Kubernetes manifests to deploy the server and client:
```bash
# Deploy the server
kubectl apply -f k8s/server-deployment.yaml
kubectl apply -f k8s/server-service.yaml

# Deploy the client
kubectl apply -f k8s/client-deployment.yaml
```

---

#### 5. **Verify the Deployment**
Check the status of the pods and services:
```bash
kubectl get pods
kubectl get services
```

---

#### 6. **Logs and Debugging**
To view the logs of the client or server:
```bash
# Server logs
kubectl logs <server-pod-name>

# Client logs
kubectl logs <client-pod-name>
```

---

### Summary of Commands

#### Go Commands
```bash
go build -o server .  # Build server binary
go build -o client .  # Build client binary
```

#### Docker Commands
```bash
docker build -t grpc-server:latest .  # Build server image
docker build -t grpc-client:latest .  # Build client image
```

#### Kubernetes Commands
```bash
kubectl apply -f k8s/server-deployment.yaml  # Deploy server
kubectl apply -f k8s/server-service.yaml     # Expose server
kubectl apply -f k8s/client-deployment.yaml  # Deploy client
kubectl get pods                             # Check pod status
kubectl logs <pod-name>                      # View logs
```

# Details Explanation

# gRPC Proof of Concept (PoC) with Go and Kubernetes

This repository contains a Proof of Concept (PoC) for a gRPC client-server application written in Go, containerized with Docker, and deployed to a Kubernetes cluster (Docker Desktop). The server implements a `Greeter` service that responds to `SayHello` requests with a greeting message. The client acts as a background worker, sending a `SayHello` request every 10 seconds indefinitely. This README documents the steps to set up, build, and deploy the PoC.

## Project Structure

```
grpc-poc/
├── client/
│   ├── main.go
│   ├── Dockerfile
│   ├── go.mod
│   ├── go.sum
│   ├── greeter.proto
│   └── pb/
│       ├── greeter.pb.go
│       └── greeter_grpc.pb.go
├── server/
│   ├── main.go
│   ├── Dockerfile
│   ├── go.mod
│   ├── go.sum
│   ├── greeter.proto
│   └── pb/
│       ├── greeter.pb.go
│       └── greeter_grpc.pb.go
├── k8s/
│   ├── client-deployment.yaml
│   ├── server-deployment.yaml
│   └── server-service.yaml
├── README.md
```

## Prerequisites

Before starting, ensure the following tools are installed on your system (Windows instructions provided; adjust for other OS):

1. **Go** (v1.21 or later):
   - Download and install from https://go.dev/dl/.
   - Verify: `go version`

2. **Docker Desktop** (with Kubernetes enabled):
   - Download and install from https://www.docker.com/products/docker-desktop/.
   - Enable Kubernetes: Settings > Kubernetes > Enable Kubernetes.
   - Verify: `docker --version`, `kubectl version --client`

3. **kubectl**:
   - Installed with Docker Desktop or download from https://kubernetes.io/docs/tasks/tools/install-kubectl-windows/.
   - Verify: `kubectl version --client`

4. **Protocol Buffers Compiler (`protoc`)**:
   - Download the latest release (e.g., `protoc-27.3-win64.zip`) from https://github.com/protocolbuffers/protobuf/releases.
   - Extract to `C:\protoc` and add `C:\protoc\bin` to your system PATH:
     - Open "Edit the system environment variables" > Environment Variables > Edit `Path` > Add `C:\protoc\bin`.
   - Verify: `protoc --version`

## Step 1: Install Dependencies

### Install Go gRPC and Protobuf Plugins

The `protoc` compiler requires Go-specific plugins to generate gRPC code.

1. Install `protoc-gen-go` and `protoc-gen-go-grpc`:
   ```bash
   go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
   go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
   ```

2. Add Go binary directory to PATH:
   - The plugins are installed in `%USERPROFILE%\go\bin` (e.g., `C:\Users\timot\go\bin`).
   - Add this directory to your system PATH via "Edit the system environment variables."
   - Verify:
     ```bash
     protoc-gen-go --version
     protoc-gen-go-grpc --version
     ```

## Step 2: Create Project Resources

### Protobuf Definition (`greeter.proto`)

The `greeter.proto` file defines the `Greeter` service and is identical for both client and server. Place it in `client/` and `server/` directories.

**File**: `client/greeter.proto`, `server/greeter.proto`
```protobuf
syntax = "proto3";

option go_package = "./pb";

package greeter;

service Greeter {
  rpc SayHello (HelloRequest) returns (HelloReply) {}
}

message HelloRequest {
  string name = 1;
}

message HelloReply {
  string message = 1;
}
```

### Generate Protobuf Code

Generate Go code from `greeter.proto` for both client and server.

1. **Client**:
   ```bash
   cd client
   protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative greeter.proto
   ```
   This creates `client/pb/greeter.pb.go` and `client/pb/greeter_grpc.pb.go`.

2. **Server**:
   ```bash
   cd server
   protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative greeter.proto
   ```
   This creates `server/pb/greeter.pb.go` and `server/pb/greeter_grpc.pb.go`.

### Server Code (`server/main.go`)

The server implements the `Greeter` service, listening on port 50051.

**File**: `server/main.go`
```go
package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "server/pb"
)

type greeterServer struct {
	pb.UnimplementedGreeterServer
}

func (s *greeterServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	message := fmt.Sprintf("Hello, %s!", req.GetName())
	return &pb.HelloReply{Message: message}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &greeterServer{})
	log.Println("Server is running on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
```

### Client Code (`client/main.go`)

The client acts as a background worker, sending a `SayHello` request every 10 seconds indefinitely using `grpc.NewClient` (to avoid the deprecated `grpc.Dial`).

**File**: `client/main.go`
```go
package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "client/pb"
)

func main() {
	conn, err := grpc.NewClient("server-service:50051", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewGreeterClient(conn)

	requestNum := 1

	for {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		r, err := c.SayHello(ctx, &pb.HelloRequest{Name: "World"})
		if err != nil {
			log.Printf("Request %d failed: %v", requestNum, err)
		} else {
			log.Printf("Request %d - Greeting: %s", requestNum, r.GetMessage())
		}
		cancel()

		requestNum++

		time.Sleep(10 * time.Second)
	}
}
```

### Initialize Go Modules

Set up Go modules for both client and server to manage dependencies.

1. **Client**:
   ```bash
   cd client
   go mod init client
   go mod tidy
   ```

2. **Server**:
   ```bash
   cd server
   go mod init server
   go mod tidy
   ```

This creates `go.mod` and `go.sum` files in each directory, including dependencies like `google.golang.org/grpc` and `google.golang.org/protobuf`.

### Dockerfiles

#### Server Dockerfile

**File**: `server/Dockerfile`
```dockerfile
FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy
COPY . .
RUN go build -o server .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/server .

EXPOSE 50051

CMD ["./server"]
```

#### Client Dockerfile

**File**: `client/Dockerfile`
```dockerfile
FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy
COPY . .
RUN go build -o client .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/client .

CMD ["./client"]
```

### Kubernetes YAMLs

#### Server Deployment (`k8s/server-deployment.yaml`)

Deploys the server with a readiness probe to ensure port 50051 is open.

**File**: `k8s/server-deployment.yaml`
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: grpc-server
  namespace: default
spec:
  replicas: 1
  selector:
    matchLabels:
      app: grpc-server
  template:
    metadata:
      labels:
        app: grpc-server
    spec:
      containers:
      - name: grpc-server
        image: grpc-server:latest
        imagePullPolicy: Never
        ports:
        - containerPort: 50051
        readinessProbe:
          tcpSocket:
            port: 50051
          initialDelaySeconds: 5
          periodSeconds: 10
        resources:
          limits:
            cpu: "500m"
            memory: "512Mi"
          requests:
            cpu: "200m"
            memory: "256Mi"
```

#### Client Deployment (`k8s/client-deployment.yaml`)

Deploys the client with a readiness probe to check the `client` process.

**File**: `k8s/client-deployment.yaml`
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: grpc-client
  namespace: default
spec:
  replicas: 1
  selector:
    matchLabels:
      app: grpc-client
  template:
    metadata:
      labels:
        app: grpc-client
    spec:
      containers:
      - name: grpc-client
        image: grpc-client:latest
        imagePullPolicy: Never
        readinessProbe:
          exec:
            command:
            - /bin/sh
            - -c
            - "ps aux | grep client | grep -v grep"
          initialDelaySeconds: 5
          periodSeconds: 10
        resources:
          limits:
            cpu: "500m"
            memory: "512Mi"
          requests:
            cpu: "200m"
            memory: "256Mi"
```

#### Server Service (`k8s/server-service.yaml`)

Exposes the server within the cluster at `server-service:50051`.

**File**: `k8s/server-service.yaml`
```yaml
apiVersion: v1
kind: Service
metadata:
  name: server-service
  namespace: default
spec:
  selector:
    app: grpc-server
  ports:
  - protocol: TCP
    port: 50051
    targetPort: 50051
  type: ClusterIP
```

## Step 3: Build the Application

1. **Build Server Docker Image**:
   ```bash
   cd server
   docker build -t grpc-server:latest .
   ```

2. **Build Client Docker Image**:
   ```bash
   cd client
   docker build -t grpc-client:latest .
   ```

3. **Verify Images**:
   ```bash
   docker images | findstr grpc
   ```
   Expected output:
   ```
   grpc-client      latest    <image-id>     <timestamp>    <size>
   grpc-server      latest    <image-id>     <timestamp>    <size>
   ```

## Step 4: Deploy to Kubernetes

1. **Apply Kubernetes Manifests**:
   ```bash
   kubectl apply -f k8s/server-deployment.yaml
   kubectl apply -f k8s/server-service.yaml
   kubectl apply -f k8s/client-deployment.yaml
   ```

2. **Verify Pods**:
   ```bash
   kubectl get pods
   ```
   Expected output:
   ```
   NAME                             READY   STATUS    RESTARTS   AGE
   grpc-client-<hash>               1/1     Running   0          <age>
   grpc-server-<hash>               1/1     Running   0          <age>
   ```

3. **Verify Service**:
   ```bash
   kubectl get svc
   ```
   Expected output:
   ```
   NAME             TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)     AGE
   server-service   ClusterIP   <cluster-ip>    <none>        50051/TCP   <age>
   ```

4. **Check Client Logs**:
   ```bash
   kubectl logs -l app=grpc-client --follow
   ```
   Expected output (new log every ~10 seconds):
   ```
   Request 1 - Greeting: Hello, World!
   Request 2 - Greeting: Hello, World!
   Request 3 - Greeting: Hello, World!
   ...
   ```

5. **Check Server Logs**:
   ```bash
   kubectl logs -l app=grpc-server
   ```
   Expected output:
   ```
   Server is running on port 50051...
   ```

## Troubleshooting

### `protoc` Errors
- **Error**: `'protoc' is not recognized`:
  - Ensure `C:\protoc\bin` is in your PATH.
  - Reinstall `protoc` from https://github.com/protocolbuffers/protobuf/releases.
- **Error**: `'protoc-gen-go' is not recognized`:
  - Reinstall plugins: `go install google.golang.org/protobuf/cmd/protoc-gen-go@latest` and `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest`.
  - Add `C:\Users\timot\go\bin` to PATH.

### Go Build Errors
- **Error**: `package client/pb is not in std`:
  - Ensure `client/pb/` contains `greeter.pb.go` and `greeter_grpc.pb.go`.
  - Run `protoc` to regenerate files.
  - Verify `go.mod` exists: `go mod init client`, `go mod tidy`.

### Docker Build Errors
- Ensure `pb/`, `go.mod`, and `go.sum` are in `client/` and `server/` directories.
- Test locally: `go build` in each directory.

### Kubernetes `ImagePullBackOff`
- **Error**: `Failed to pull image "grpc-client:latest": pull access denied`:
  - Confirm images exist: `docker images | findstr grpc`.
  - Ensure `imagePullPolicy: Never` is set in `k8s/client-deployment.yaml` and `k8s/server-deployment.yaml` for Docker Desktop.
  - Alternatively, push images to Docker Hub:
    ```bash
    docker tag grpc-client:latest <your-username>/grpc-client:latest
    docker tag grpc-server:latest <your-username>/grpc-server:latest
    docker login
    docker push <your-username>/grpc-client:latest
    docker push <your-username>/grpc-server:latest
    ```
    Update YAMLs with `<your-username>/grpc-client:latest` and `<your-username>/grpc-server:latest`.

### Client Connection Errors
- **Error**: `Request <n> failed: rpc error: code = Unavailable desc = connection refused`:
  - Verify server pod: `kubectl get pods -l app=grpc-server`.
  - Check service: `kubectl describe svc server-service`.
  - Test connectivity:
    ```bash
    kubectl exec -it <client-pod-name> -- sh
    apk add busybox-extras
    telnet server-service 50051
    ```

## Notes
- **Client Behavior**: The client runs indefinitely, sending a request every 10 seconds. It uses `grpc.NewClient` to avoid the deprecated `grpc.Dial`.
- **Insecure Connection**: The client uses `insecure.NewCredentials()` for simplicity. For production, enable TLS.
- **Kubernetes Cluster**: This PoC uses Docker Desktop’s Kubernetes. For remote clusters, push images to a registry and remove `imagePullPolicy: Never`.
- **Readiness Probes**: The server uses a TCP probe (port 50051), and the client uses an `exec` probe to check the `client` process.

## Conclusion

This PoC demonstrates a gRPC client-server application in Go, containerized with Docker, and deployed to Kubernetes. The client sends periodic requests to the server, which responds with greetings. The setup includes robust error handling, readiness probes, and resource limits, making it a solid foundation for further development.

For issues or questions, refer to the troubleshooting section or contact the repository maintainer.