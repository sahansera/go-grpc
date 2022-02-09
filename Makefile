gen:
	protoc --proto_path=proto proto/*.proto --go_out=. --go-grpc_out=.

clean:
	rm -f server/*.pb.go

server:
	go run server/main.go

install:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
	brew install protobuf
	brew install clang-format
	export PATH=$PATH:$(go env GOPATH)/bin

test:
	rm -rf tmp && mkdir tmp
	go test -cover -race serializer/*.go
