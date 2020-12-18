import requests
from argparse import ArgumentParser

def get_outside_ip():
    return requests.get('http://ifconfig.me/ip').text.strip()

def get_arguments_parser():
    parser = ArgumentParser()
    parser.add_argument('--service_ip', type=str, required=False, default='0.0.0.0')
    parser.add_argument('--service_port', type=int, required=True)
    parser.add_argument('--etcd_ip', type=str, required=True)
    parser.add_argument('--etcd_port', type=int, required=False, default=2379)
    parser.add_argument('--etcd_prefix', type=str, required=True)
    return parser
