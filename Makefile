.PHONY: gen clean server client install test

gen:
	protoc --proto_path=proto proto/*.proto --go_out=server --go-grpc_out=server
	protoc --proto_path=proto proto/*.proto --go_out=client --go-grpc_out=client

clean:
	rm -rf server/pb/
	rm -rf client/pb/

server:
	go -C server run .

client:
	go -C client run .

install:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	@echo 'Add $$(go env GOPATH)/bin to your PATH if protoc cannot find the plugins.'

test:
	go -C server test ./...
	go -C client test ./...
