# Project Initiation
1. Initialize the Go Module
```bash 
go mod init github.com/faizkads/push-notifications
```
2. Compile the `.proto` file to Go and gRPC for the server
```bash
protoc --go_out=. --go-grpc_out=. --proto_path=./proto ./proto/notif.proto
```
3. Compile the `.proto` file to Python and gRPC for the client
```bash
python -m grpc_tools.protoc --proto_path=../proto --python_out=. --grpc_python_out=. ../proto/notif.proto
```
