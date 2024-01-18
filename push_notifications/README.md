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

starting up in this project:
```bash
pip install grpcio grpc_tools
```
or generally:
```bash
pip install -r requirements.txt
```
4. Generate the proto files

## Sample Case
### Topic Selection
```
python client.py

Topics:
1. Daily Meow Fact
2. Currency Exchange Rate
Pick a topic for your push notif [1/2]:
```
### Topic 1: Meow Facts
- Client
```
python client.py

Topics:
1. Daily Meow Fact
2. Currency Exchange Rate
Pick a topic for your push notif [1/2]: 1
Enter the number of iteration: 4
[2024-01-18 15:46:01] Meow Fact 1: Cats bury their feces to cover their trails from predators.
[2024-01-18 15:46:06] Meow Fact 2: A cat has been mayor of Talkeetna, Alaska, for 15 years. His name is Stubbs.
[2024-01-18 15:46:11] Meow Fact 3: The worlds richest cat is worth $13 million after his human passed away and left her fortune to him.     
[2024-01-18 15:46:16] Meow Fact 4: Almost 10% of a cat's bones are in its tail, and the tail is used to maintain balance.
```

- Server
```
go run server.go
2024/01/18 15:10:32 Starting notification server on port 50051
2024/01/18 15:45:59 Received request for Meow Facts with [4] iteration
2024/01/18 15:46:01 Meow Fact 1 Sent
2024/01/18 15:46:06 Meow Fact 2 Sent
2024/01/18 15:46:11 Meow Fact 3 Sent
2024/01/18 15:46:16 Meow Fact 4 Sent
```

### Topic 2: Currency Exchange
#### Single Dest
- Client
```
Topics:
1. Daily Meow Fact
2. Currency Exchange Rate
Pick a topic for your push notif [1/2]: 2
Available currencies:
['AUD', 'BGN', 'BRL', 'CAD', 'CHF', 'CNY', 'CZK', 'DKK', 'EUR', 'GBP', 'HKD', 'HUF', 'IDR', 'ILS', 'INR', 'ISK', 'JPY', 'KRW', 'MXN', 'MYR', 'NOK', 'NZD', 'PHP', 'PLN', 'RON', 'SEK', 'SGD', 'THB', 'TRY', 'USD', 'ZAR']
Enter base currency code: usd
Enter destined currency code (seperate with comma for multi-input): gbp
Enter the number of iteration: 3
[2024-01-18] Base: USD
- exchange rate for GBP: 0.7626699805259705
[2024-01-19] Base: USD
- exchange rate for GBP: 0.7745299935340881
[2024-01-20] Base: USD
- exchange rate for GBP: 0.7762799859046936
```
- Server
```
2024/01/18 15:29:48 Received exchange rate from [usd] to [GBP] for: [3] iteration
2024/01/18 15:29:48 Exchange rate from usd to GBP on 2024-01-18 is Sent
2024/01/18 15:29:54 Exchange rate from usd to GBP on 2024-01-19 is Sent
2024/01/18 15:29:59 Exchange rate from usd to GBP on 2024-01-20 is Sent
```

#### Multi Dest
- Client
```
Topics:
1. Daily Meow Fact
2. Currency Exchange Rate
Pick a topic for your push notif [1/2]: 2
Available currencies:
['AUD', 'BGN', 'BRL', 'CAD', 'CHF', 'CNY', 'CZK', 'DKK', 'EUR', 'GBP', 'HKD', 'HUF', 'IDR', 'ILS', 'INR', 'ISK', 'JPY', 'KRW', 'MXN', 'MYR', 'NOK', 'NZD', 'PHP', 'PLN', 'RON', 'SEK', 'SGD', 'THB', 'TRY', 'USD', 'ZAR']
Enter base currency code: usd
Enter destined currency code (seperate with comma for multi-input): gbp,sgd,eur
Enter the number of iteration: 3
[2024-01-18] Base: USD
- exchange rate for EUR: 0.8884900212287903
- exchange rate for GBP: 0.7626699805259705
- exchange rate for SGD: 1.3207999467849731
[2024-01-19] Base: USD
- exchange rate for EUR: 0.8911100029945374
- exchange rate for GBP: 0.7745299935340881
- exchange rate for SGD: 1.3255000114440918
[2024-01-20] Base: USD
- exchange rate for EUR: 0.8931000232696533
- exchange rate for GBP: 0.7762799859046936
- exchange rate for SGD: 1.3223999738693237
```
- Server
```
Received exchange rate from [usd] to [GBP,SGD,EUR] for: [3] iteration
2024/01/18 15:28:07 Exchange rate from usd to GBP,SGD,EUR on 2024-01-18 is Sent
2024/01/18 15:28:12 Exchange rate from usd to GBP,SGD,EUR on 2024-01-19 is Sent
2024/01/18 15:28:17 Exchange rate from usd to GBP,SGD,EUR on 2024-01-20 is Sent
```