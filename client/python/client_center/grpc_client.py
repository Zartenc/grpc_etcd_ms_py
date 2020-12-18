import etcd3
import json
import random

class EtcdClient(etcd3.Etcd3Client):
    def get_values_by_key(self, key, **kwargs):
        values, _ = self.get(key, **kwargs)
        values_list = []
        if values is not None:
            try:
                values_list = json.loads(values.decode('utf-8'))
                if not isinstance(values_list, list):
                    raise TypeError()
            except:
                raise Exception()

        return values_list

    def put_values_by_key(self, key, values):
        if not isinstance(values, list):
            raise Exception()
        self.put(key, json.dumps(values))

class EtcdHandleClient(EtcdClient):
    _singleton = None
    def __new__(cls, *args, **kwargs):
        if not cls._singleton:
            cls._singleton = super().__new__(cls)
        return cls._singleton

    def __init__(self, etcd_ip, etcd_port, etcd_prefix):
        self.etcd_ip = etcd_ip
        self.etcd_port = etcd_port
        self.etcd_prefix = etcd_prefix
        super().__init__(host=etcd_ip, port=etcd_port)
        self.endpoints_list = self.get_values_by_key(f'{self.etcd_prefix}/grpc')
        self.watched_id = self.add_watch_callback(key=f'{self.etcd_prefix}/grpc', callback=self._update_endpoints)

    def __del__(self):
        self.cancel_watch(self.watched_id)

    def get_grpc_serv_ip(self):
        endpoints_nums = len(self.endpoints_list)
        if endpoints_nums <= 0:
            raise RuntimeError('No grpc services are available.Please notify the administrator to start the grpc service')
        select_id = random.randint(0, len(self.endpoints_list)-1)
        return self.endpoints_list[select_id]

    def _update_endpoints(self, watched_response):
        watched_event = watched_response.events[0]
        try:
            update_endpoint_list = json.loads(watched_event.value)
            if not isinstance(update_endpoint_list, list):
                raise TypeError
        except Exception as e:
            print(e)
            return

        self.endpoints_list = update_endpoint_list





