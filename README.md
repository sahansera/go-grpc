# gRPC on Go

This project is a small starting point for building a Go client and server that communicate with gRPC. The accompanying [gRPC article series](https://sahansera.dev/introduction-to-grpc/) explains the contract, generated code, server, and client.

## Install 🏗

Install the Protocol Buffers compiler using the instructions for your platform. On macOS with Homebrew:

```bash
brew install protobuf
```

Install the Go plugins for the Protocol Buffers compiler. The [official gRPC-Go quick start](https://grpc.io/docs/languages/go/quickstart/#prerequisites) lists the same prerequisites.

```bash
make install
```

Generate the client and server bindings, then verify both modules:

```bash
make gen
make test
```

Run the server:

```bash
make server
```

In another terminal, run the client:

```bash
make client
```

The sample uses plaintext transport only for local development. Configure TLS and authentication before exposing a gRPC service outside your trusted development network.

## Invoking RPCs 🚀

```bash
# Note: since we are not using TLS all the calls are with -plaintext flag
grpcurl -plaintext localhost:8080 list
grpcurl -plaintext localhost:8080 Inventory.GetBookList
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

## Common issues and FAQ ❓

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
    // ...
    "editor.formatOnSave": true
    // ...
}
```
