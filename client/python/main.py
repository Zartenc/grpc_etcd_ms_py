import os
import sys
curPath = os.path.abspath(os.path.dirname(__file__))
rootPath = os.path.split(curPath)[0]
sys.path.append(os.path.split(rootPath)[0])

from client_center.grpc_client import EtcdHandleClient
from grpc_proto import zarten_pb2_grpc, zarten_pb2
import grpc

def main():
    etcd_client = EtcdHandleClient(etcd_ip='xxxx', etcd_port=2379, etcd_prefix='/zarten')
    endpoint = etcd_client.get_grpc_serv_ip()
    print('endpoint:', endpoint)

    with grpc.insecure_channel(endpoint) as channel:
        stub = zarten_pb2_grpc.ZartenStub(channel)
        response = stub.GetInfo(zarten_pb2.ZartenRequest(zhihu_name='Zarten123'))
        print(f'receive response: {response}')


if __name__ == '__main__':
    main()
