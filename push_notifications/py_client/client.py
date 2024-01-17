import grpc

import notif_pb2
import notif_pb2_grpc

def get_facts(stub,count):
    req = notif_pb2.MeowReq(count=count)
    res = stub.StreamMeow(req)
    # print(res)
    count = 1
    for r in res:
        print(f"\n[Meow Fact {count}]: {r.fact}")
        count += 1

if __name__=="__main__":
    channel = grpc.insecure_channel('localhost:50051')
    stub = notif_pb2_grpc.NotifStub(channel)

    try:
        count = int(input("Enter the number of iteration: "))
        get_facts(stub, count)
    except ValueError:
        print("Please insert a valid integer.")
