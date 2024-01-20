LOCAL_BIN:=$(CURDIR)/bin
LOCAL_MIGRATION_DIR=./migrations
LOCAL_MIGRATION_DSN="host=localhost port=54321 dbname=note-service user=note-service-user password=note-service-password sslmode=disable"

.PHONY: install-goose
.install-goose:
	go install github.com/pressly/goose/v3/cmd/goose@latest

.PHONY: local-migration-status
local-migration-status:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} status -v

.PHONY: local-migration-up
local-migration-up: 
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} up -v

.PHONY: local-migration-down
local-migration-down: 
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} down -v

install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@latest
	GOBIN=$(LOCAL_BIN) go install github.com/envoyproxy/protoc-gen-validate@v0.10.1
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.15.2
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.15.2
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger@v1.16.0

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

generate:
	mkdir -p pkg/note_v1
	protoc	--proto_path vendor.protogen --proto_path api/note_v1 \
			--go_out=pkg/note_v1 --go_opt=paths=source_relative \
			--plugin=protoc-gen-go=bin/protoc-gen-go \
			--go-grpc_out=pkg/note_v1 --go-grpc_opt=paths=source_relative \
			--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
			--grpc-gateway_out=pkg/note_v1 --grpc-gateway_opt=paths=source_relative \
			--plugin=protoc-gen-grpc-gateway=bin/protoc-gen-grpc-gateway \
			--swagger_out=allow_merge=true,merge_file_name=api:pkg/note_v1 \
			--plugin=protoc-gen-swagger=bin/protoc-gen-swagger \
			--validate_out lang=go:pkg/note_v1 --validate_opt=paths=source_relative \
			--plugin=protoc-gen-validate=bin/protoc-gen-validate \
			api/note_v1/note.proto


PHONY: vendor-proto
vendor-proto: .vendor-proto

PHONY: .vendor-proto
.vendor-proto:
			@if [ ! -d vendor.protogen/google ]; then \
				git clone https://github.com/googleapis/googleapis vendor.protogen/googleapis &&\
				mkdir -p vendor.protogen/google/ &&\
				mv vendor.protogen/googleapis/google/api vendor.protogen/google &&\
				rm -rf vendor.protogen/googleapis ;\
			fi
			@if [ ! -d vendor.protogen/github.com/envoyproxy ]; then \
				mkdir -p vendor.protogen/validate &&\
				git clone https://github.com/envoyproxy/protoc-gen-validate vendor.protogen/protoc-gen-validate &&\
				mv vendor.protogen/protoc-gen-validate/validate/*.proto vendor.protogen/validate &&\
				rm -rf vendor.protogen/protoc-gen-validate ;\
			fi
			@if [ ! -d vendor.protogen/google/protobuf ]; then \
				git clone https://github.com/protocolbuffers/protobuf vendor.protogen/protobuf &&\
				mkdir -p vendor.protogen/google/protobuf &&\
				mv vendor.protogen/protobuf/src/google/protobuf/*.proto vendor.protogen/google/protobuf &&\
				rm -rf vendor.protogen/protobuf ;\
			fi
