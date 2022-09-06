# grpc-workshop

[grpc-Installation](https://www.grpc.io/docs/languages/go/quickstart/)

```shell
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    api/api.proto
```

### Run

```shell
go run server/grpcServer.go
go run client/grpcClient.go
```