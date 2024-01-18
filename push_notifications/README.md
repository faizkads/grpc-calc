# gRPC-based Push Notification
The task is to develop a server-side streaming for simplified push notification. The server is developed in Golang and the client is in Python.

The topics chosen are:
- Meow Facts: [Github Repo](https://github.com/wh-iterabb-it/meowfacts)
- Currency Exchange: [Web Documentation](https://www.frankfurter.app/docs/)

The Meow Facts will display random facts about cats, the user-input needed for the client side will only be the number of iteration

the Currency Exchange will display currency exchange rate between 2 currencies at a given date. For example exchange rate from USD to GBP at 2024-01-18. The user input will be base currency (from), destination currency (to), and days of iteration.

## Project Initiation
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
4. Run the server
```bash
go run server.go
```
5. Run the client
```bash
python client.py
```

## Python gRPC Setup *
Just as a notes for myself to remember 
1. Set python version for the environment using `pyenv`
```bash 
pyenv local 3.10.5
```
2. Set virtual environemnt and activate
```bash
virtualenv venv
source venv/Scripts/activate
```
3. Install libraries needed
```bash
pip install grpcio grpc_tools
```