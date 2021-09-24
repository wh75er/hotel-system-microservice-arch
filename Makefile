hotel-service:
	go run cmd/hotel-service/main.go

gen-proto-hotel:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    internal/pkg/delivery/grpc/hotel-service/proto/scheme.proto

client:
	go run cmd/client/main.go

