import thriftpy

price_thrift = thriftpy.load("service.thrift",module_name="price_thrift")

from thriftpy.rpc import make_client

client = make_client(price_thrift.PriceService, '127.0.0.1', 9090)
# secure
#client = make_client(price_thrift.PriceService, host='127.0.0.1',port=9090, cafile='keys/server.crt')
p =client.price("oneforge")
print(p)
