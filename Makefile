hotel-service:
	go run cmd/hotel-service/main.go

auth-service:
	go run cmd/auth-service/main.go

loyalty-service:
	go run cmd/loyalty-service/main.go

payment-service:
	go run cmd/payment-service/main.go

reservation-service:
	go run cmd/reservation-service/main.go

gateway-service:
	go run cmd/gateway-service/main.go

stat-service:
	go run cmd/stat-service/main.go

hotel-client:
	go run cmd/hotel-client/main.go

gen-proto-common:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
	internal/pkg/delivery/grpc/commonProto/common.proto

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

gen-proto-payment:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    internal/pkg/delivery/grpc/payment-service/proto/scheme.proto

gen-proto-reservation:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    internal/pkg/delivery/grpc/reservation-service/proto/scheme.proto

gen-proto-gateway:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    internal/pkg/delivery/grpc/gateway-service/proto/scheme.proto

gen-proto-stat:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    internal/pkg/delivery/grpc/stat-service/proto/scheme.proto

gen-proto-frontend:
	protoc --js_out=import_style=commonjs:./frontend/src/proto \
	--grpc-web_out=import_style=commonjs,mode=grpcwebtext:./frontend/src/proto \
    internal/pkg/delivery/grpc/gateway-service/proto/scheme.proto
	protoc --js_out=import_style=commonjs:./frontend/src/proto \
	--grpc-web_out=import_style=commonjs,mode=grpcwebtext:./frontend/src/proto \
    internal/pkg/delivery/grpc/reservation-service/proto/scheme.proto
	protoc --js_out=import_style=commonjs:./frontend/src/proto \
	--grpc-web_out=import_style=commonjs,mode=grpcwebtext:./frontend/src/proto \
    internal/pkg/delivery/grpc/payment-service/proto/scheme.proto
	protoc --js_out=import_style=commonjs:./frontend/src/proto \
	--grpc-web_out=import_style=commonjs,mode=grpcwebtext:./frontend/src/proto \
    internal/pkg/delivery/grpc/loyalty-service/proto/scheme.proto
	protoc --js_out=import_style=commonjs:./frontend/src/proto \
	--grpc-web_out=import_style=commonjs,mode=grpcwebtext:./frontend/src/proto \
    internal/pkg/delivery/grpc/auth-service/proto/scheme.proto
	protoc --js_out=import_style=commonjs:./frontend/src/proto \
	--grpc-web_out=import_style=commonjs,mode=grpcwebtext:./frontend/src/proto \
    internal/pkg/delivery/grpc/hotel-service/proto/scheme.proto
	protoc --js_out=import_style=commonjs:./frontend/src/proto \
	--grpc-web_out=import_style=commonjs,mode=grpcwebtext:./frontend/src/proto \
	internal/pkg/delivery/grpc/commonProto/common.proto

fmt:
	go fmt ./internal/... && go fmt ./cmd/...

