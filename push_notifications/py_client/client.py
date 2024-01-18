import grpc

import notif_pb2
import notif_pb2_grpc
from datetime import datetime

def get_facts(stub,count):
    """Sending request to the server, and receive the response, then prints each response"""
    req = notif_pb2.MeowReq(count=count)
    res = stub.StreamMeow(req)
    count = 1
    for r in res:
        dt = datetime.now().strftime("%Y-%m-%d %H:%M:%S") # timestamp
        print(f"[{dt}] Meow Fact {count}: {r.fact}")
        count += 1

def currency_rate(stub, from_cur, to_cur, count):
    """Sending request to the server, and receive the response, then prints each response"""
    req = notif_pb2.CurReq(
        from_cur = from_cur,
        to_cur = to_cur,
        count = count
    )
    responses = stub.StreamCurrency(req)
    for response in responses:
        print(f"[{response.date}] Base: {response.base}")
        for currency, rate in response.rates.items():
            print(f"- exchange rate for {currency}: {rate}")

def topic_1(stub):
    # validate variable count
    try:
        count = int(input("Enter the number of iteration: "))
    except ValueError:
        print("Invalid argument. Please enter a valid integer")
        return

    get_facts(stub, count) # function call for request and response
    

def topic_2(stub):
    codes = ["AUD", "BGN", "BRL", "CAD", "CHF", "CNY", "CZK", "DKK", "EUR", "GBP",
             "HKD", "HUF", "IDR", "ILS", "INR", "ISK", "JPY", "KRW", "MXN", "MYR",
             "NOK", "NZD", "PHP", "PLN", "RON", "SEK", "SGD", "THB", "TRY", "USD", "ZAR"]
    print("Available currencies: ")
    print(codes)

    from_cur = input("Enter base currency code: ")
    # validate the input currency
    if from_cur.upper() not in codes:
        print("Invalid currency code")
        return

    to_cur = input("Enter destined currency code (seperate with comma for multi-input): ")
    to_cur_list = [code.strip().upper() for code in to_cur.split(",")] # turns into a list to be iterable
    invalid_cur = [code for code in to_cur_list if code not in codes] # check invalid codes
    if invalid_cur:
        print(f"Invalid 'to' currency codes: {', '.join(invalid_cur)}")
        return
    to_cur_str = ",".join(to_cur_list) # turns it back into a string with no whitespaces

    # validate variable count 
    try:
        count = int(input("Enter the number of iteration: "))
    except ValueError:
        print("Invalid argument. Please enter a valid integer")
        return
    if count < 1:
        print("Invalid argument. Iteration request should be bigger than 0")
        return

    currency_rate(stub, from_cur,to_cur_str,count) # function call for request and response


if __name__=="__main__":
    channel = grpc.insecure_channel('localhost:50051')
    stub = notif_pb2_grpc.NotifStub(channel)

    topic = input("\nTopics:\n1. Daily Meow Fact\n2. Currency Exchange Rate\nPick a topic for your push notif [1/2]: ")

    if topic =="1":
        topic_1(stub)
    elif topic == "2":
        topic_2(stub)
    else:
        print("Invalid topic")
