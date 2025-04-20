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