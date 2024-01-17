# gRPC for Calculator
A simple calculator gRPC to count operations of: addition, subtraction, multiplication, and division

The client is only limited to present 2 arguments in the type of and in the range of float32 data

The variables will only be displayed in 2 decimal places

## Protobuf Protoc Initiation

1. Download protobuf protoc file via [Github Repo](https://github.com/protocolbuffers/protobuf/releases)

2. Add path of `protoc.exe` to system environment variables

3. Install protocol buffer compiler for Go code from `.proto` files using:
```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go
```

4. Install protocol buffer compiler for Go code for gRPC services using:
```bash
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

## Project Initiation
1. Initialize the Go Module
```bash 
go mod init github.com/faizkads/grpc-calc
```

2. Compile the `.proto` file for Go and gRPC
```bash
protoc --go_out=. --go-grpc_out=. calc/proto/calc.proto
```

3. Run the server
```bash
go run server/server.go
```

4. Run the client with 2 arguments in float32, for example:
```bash
go run client/client.go 10 2.5
```

## Test Cases
- Check results
```bash
go run client/client.go 17 5
go run client/client.go 17 5.7
go run client/client.go 1.7 5.7
```

- Check error by invalid number of arguments
```bash
go run client/client.go 7
go run client/client.go 7 8 2
```

- Check error by invalid argument data type
```bash
go run client/client.go 10.2 lala
```

- Check error by invalid argument for division by zero
```bash
go run client/client.go 3 0
```