# gRPC on Go

This projects you a start point for building a client and a server that communicates with gRPC. I also published a [blog article series](https://sahansera.dev/introduction-to-grpc/) if you are interested.

## Install 🏗

Protobufs

```bash
brew instal protobuf
```

Go plugins for the protobuf compiler or refer [here](https://grpc.io/docs/languages/go/quickstart/#prerequisites)

```bash
make install
```

## Invoking RPCs 🚀

```bash
# Note: since we are not using TLS all the calls are with -plaintext flag
grpcurl -plaintext localhost:8080 list # introspect the service
grpcurl -plaintext localhost:8080 Inventory.GetBookList # to get a list of books
```


## Useful commands 📡

Generate Go stubs

```bash
make gen
```

Clean stubs

```bash
make clean
```

## Issues and FAQ

1. VS Code complains about imports?

Known issue with the Proto3 VS Code plugin. Add the following to your `settings.json` file

```
"protoc": {
        "path": "/path/to/protoc",
        "compile_on_save": false,
        "options": [
            "--proto_path=protos/v3",
            "--proto_path=protos/v2",
            "--proto_path=${workspaceRoot}/proto",
            "--proto_path=${env.GOPATH}/src",
            "--java_out=gen/java"
        ]
    }
```

Run `which protoc` to find your proto compiler path.

2. Auto formatter doesn't work?

```bash
brew install clang-format
```

Add the following to your settings.json

```json
{
    ...
    "editor.formatOnSave": true
    ...
}
```
