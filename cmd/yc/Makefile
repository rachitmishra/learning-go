# grpc gateway
# go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
# go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
# go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
# go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

main_package_path = ./
binary_name = todo
GQLGEN_CMD = go run github.com/99designs/gqlgen generate

.PHONY: protogen
protogen: 
	mkdir -p generated/data/proto
	protoc -I ./data/proto \
   --go_out ./generated/data/proto --go_opt paths=source_relative \
   --go-grpc_out ./generated/data/proto --go-grpc_opt paths=source_relative \
   --grpc-gateway_out ./generated/data/proto --grpc-gateway_opt paths=source_relative \
   ./data/proto/todo/*.proto

.PHONY: gqlgen
gqlgen:
	$(GQLGEN_CMD)

.PHONY: sqlgen
sqlgen:
	sqlc generate

.PHONY: run
run:
	gow run ${main_package_path}


.PHONY: tidy
tidy:
	go mod tidy