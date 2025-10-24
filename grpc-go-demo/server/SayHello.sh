#!/bin/bash
grpcurl -plaintext \
-d @ \
-import-path ./proto \
-proto hello.proto 127.0.0.1:40001 hello.HelloGRPC.SayHello <<EOM
{
  "name": "World"
}
EOM


grpcurl -plaintext -d {"name":"World"} -import-path ./proto -proto hello.proto 127.0.0.1:40001 hello.HelloGRPC.SayHello

# 一行执行
# grpcurl -plaintext -d '{"name":"World"}' -import-path ./proto -proto hello.proto 127.0.0.1:40001 hello.HelloGRPC.SayHello