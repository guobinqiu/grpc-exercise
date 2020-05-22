import os
import sys

sys.path.append(os.path.join(os.path.dirname(__file__), "../.."))

import grpc
import pb_python.user_pb2_grpc

if __name__ == '__main__':
    with grpc.insecure_channel('localhost:9999') as channel:
        stub = pb_python.user_pb2_grpc.UserStub(channel)
        response = stub.GetUserInfo(pb_python.user_pb2.UserRequest(name='Guobin'))
    print("Received user name: %s" % response.name)
