Go as server side
---

Go as server side, other programming languages as client side.

### Protobuf Compiler (Interface)

| OS | Command                          |
|----------|-----------------------------------|
| macOS    | `brew install protobuf`           |
| CentOS   | `yum install -y protobuf-compiler` |
| Ubuntu   | `apt install -y protobuf-compiler` |

### Protobuf Compiler Plugin (Implementation of Interface)

| Lang | Command |
|---|---|
| Go | `go install google.golang.org/protobuf/cmd/protoc-gen-go` 和 `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc` |
| Python | `python -m pip install grpcio-tools` |
|	Ruby | `gem install grpc-tools` |

### Code Generator

| Lang | Command |
|---|---|
| Go | `protoc --proto_path=./proto --go_out=plugins=grpc:. ./proto/user.proto` |
| Python | `python -m grpc_tools.protoc --proto_path=./proto --python_out=./pb_python --grpc_python_out=./pb_python ./proto/user.proto` |
|	Ruby | `grpc_tools_ruby_protoc --proto_path=./proto --ruby_out=./pb_ruby --grpc_out=./pb_ruby ./proto/user.proto` |

### SDK 

| Lang | Command |
|---|---|
| Go | `go get google.golang.org/grpc` |
| Python | `python -m pip install grpcio` |
|	Ruby | `gem install grpc` |

### Run

- Server side:

  `go run server/main.go`

- Client side:

| Lang | Command |
|---|---|
| Go | `go run client/golang/client.go` |
| Python | `python client/python/client.py` |
|	Ruby | `ruby client/ruby/client.rb` |
