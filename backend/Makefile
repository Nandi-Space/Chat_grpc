gen:
	protoc --proto_path=cmd/api/proto cmd/api/proto/*.proto  --go_out=plugins=grpc:cmd/api/

clean:
	rm cmd/api/pb/*.go

run:
	go run cmd/api/main.go
test:
	go test -v -cover ./...

.PHONY: gen clean run test
