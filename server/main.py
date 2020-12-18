import os
import sys
curPath = os.path.abspath(os.path.dirname(__file__))
rootPath = os.path.split(curPath)[0]
sys.path.append(os.path.split(rootPath)[0])

from service_center.grpc_service import EtcdHandleServ
from grpc_proto import zarten_pb2_grpc
from service_center.logic_service import ZartenServ
from service_center.utils import get_arguments_parser

import grpc
from concurrent import futures
import time
import threading
import signal


def main(service_ip, service_port, etcd_ip, etcd_port, etcd_prefix):
    print('***service is starting...')
    grpc_server = grpc.server(futures.ThreadPoolExecutor(max_workers=500))
    zarten_pb2_grpc.add_ZartenServicer_to_server(ZartenServ(), grpc_server)
    grpc_server.add_insecure_port(f'{service_ip}:{service_port}')
    grpc_server.start()

    etcd_handle_serv = EtcdHandleServ(service_port=service_port, etcd_ip=etcd_ip, etcd_port=etcd_port, etcd_prefix=etcd_prefix)
    etcd_handle_serv.register_service()

    event = threading.Event()
    def signal_handler(*args):
        etcd_handle_serv.logout_service()
        event.set()
    signal.signal(signal.SIGTERM, signal_handler)

    print("***serveice started")
    try:
        while True:
            time.sleep(60 * 60 * 24)
    except KeyboardInterrupt:
        etcd_handle_serv.logout_service()
        grpc_server.stop(0)

# python main.py --service_port 65510 --etcd_ip xxxx --etcd_prefix /zarten
if __name__ == '__main__':
    parser = get_arguments_parser()
    args = parser.parse_args()
    main(args.service_ip, args.service_port, args.etcd_ip, args.etcd_port, args.etcd_prefix)