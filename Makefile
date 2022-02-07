gen:
	protoc --proto_path=proto proto/*.proto --go_out=.
clean:
	rm -f *.pb.go
run:
	go run main.go