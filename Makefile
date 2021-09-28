hotel-service:
	go run cmd/hotel-service/main.go

auth-service:
	go run cmd/auth-service/main.go

hotel-client:
	go run cmd/hotel-client/main.go

gen-proto-hotel:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    internal/pkg/delivery/grpc/hotel-service/proto/scheme.proto

gen-proto-auth:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    internal/pkg/delivery/grpc/auth-service/proto/scheme.proto

gen-proto-loyalty:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    internal/pkg/delivery/grpc/loyalty-service/proto/scheme.proto

fmt:
	go fmt ./internal/... && go fmt ./cmd/...

