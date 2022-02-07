# go-grpc

## Install

Protobufs
```bash
brew instal protobuf
```

Go plugins for the protobuf compiler or refer [here](https://grpc.io/docs/languages/go/quickstart/#prerequisites)

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
```

## Compiling Protocs
```bash
protoc --proto_path=proto proto/*.proto --go_out=.
```