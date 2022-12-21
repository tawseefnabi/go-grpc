# go-grpc


install 
  - go get -u gtihub.com/golang/protobuf/protoc-gen-go

generate the Go specific gRPC code using the protoc tool:
```
protoc --go_out=plugins=grpc:chat chat.proto
```

###  run the commands for server
```
go run server.go
```
```
go run client.go
```