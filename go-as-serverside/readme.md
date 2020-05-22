Practice gRPC
---

Golang as server side, other programming languages as client side.

### Install

Components/Languages | Go | Python | Ruby 
--- | --- | --- | ---
gRPC | go get google.golang.org/grpc | python -m pip install grpcio | gem install grpc
gRPC tool | <table><tr><td>Protobuf Compiler</td><td> Protobuf Compiler Plugin</td></tr><tr><td>go get github.com/golang/protobuf/protoc-gen-go</td><td><table><td>mac</td><td>centos</td><td>ubantu</td></tr><tr><td>brew install protobuf</td><td>yum install -y protobuf-compiler</td><td>apt install -y protobuf-compiler</td></table></td></tr></table> | python -m pip install grpcio-tools | gem install grpc-tools

### Generate Stub

Go | Python | Ruby
--- | --- | ---
protoc --proto_path=./proto --go_out=plugins=grpc:. ./proto/user.proto | python -m grpc_tools.protoc --proto_path=./proto --python_out=./pb_python --grpc_python_out=./pb_python ./proto/user.proto | grpc_tools_ruby_protoc --proto_path=./proto --ruby_out=./pb_ruby --grpc_out=./pb_ruby ./proto/user.proto

### Run

Server side:

`go run server/main.go`

Client side:

Go | Python | Ruby
--- | --- | ---
go run client/golang/client.go | python client/python/client.py | ruby client/ruby/client.rb

### Tutorial

> https://grpc.io/docs/tutorials/
