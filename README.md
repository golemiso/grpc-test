# Install
```
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get -u github.com/golang/protobuf/protoc-gen-go
```


# Generate And Run

```
protoc -I helloworld/ -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis helloworld/helloworld.proto --go_out=plugins=grpc:helloworld --grpc-gateway_out=logtostderr=true:helloworld

go run greeter_server/main.go
```