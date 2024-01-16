import grpc
import sys
import calc_pb2 as pb
import calc_pb2_grpc as pb_grpc
from grpc._channel import _InactiveRpcError

def arg_parser():
    try:
        n1 = float(sys.argv[1])
        n2 = float(sys.argv[2])
    except ValueError:
        sys.exit("Arguments should be in float32")
    
    return n1,n2

def format_num(num):
    r = str(num).rstrip('0').rstrip('.') if '.' in str(num) else str(num)
    return r

def main():
    if len(sys.argv) != 3:
        sys.exit("Exactly 2 arguments needed")

    n1, n2 = arg_parser()

    with grpc.insecure_channel("localhost:50051") as ch:
        stub = pb_grpc.CalculatorStub(ch)

        req = pb.CalcReq(n1=n1, n2=n2)

        AddRes = stub.Add(req)
        print(f"{format_num(n1)} + {format_num(n2)} = {format_num(AddRes.r)}")

        SubRes = stub.Sub(req)
        print(f"{format_num(n1)} - {format_num(n2)} = {format_num(SubRes.r)}")

        MulRes = stub.Mul(req)
        print(f"{format_num(n1)} * {format_num(n2)} = {format_num(MulRes.r)}")

        try:
            DivRes = stub.Div(req)
            print(f"{format_num(n1)} / {format_num(n2)} = {format_num(DivRes.r)}")
        except _InactiveRpcError as e:
            print(f"Division error: {e.details()}")

if __name__ == "__main__":
    main()