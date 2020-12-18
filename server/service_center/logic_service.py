from grpc_proto import zarten_pb2_grpc, zarten_pb2

class ZartenServ(zarten_pb2_grpc.ZartenServicer):
    def GetInfo(self, request, context):
        zhihu_name = request.zhihu_name
        homepage = 'https://www.zhihu.com/people/zarten/posts'
        return zarten_pb2.ZartenResponse(name=zhihu_name, homepage=homepage)