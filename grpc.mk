gen: deps gen-server

deps:
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

gen-server: .pre-gen-server .gen-server
.pre-gen-server:
	mkdir -p pkg/api

.gen-server:
	protoc --go_out=./pkg --go-grpc_out=./pkg \
	-I $(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis \
	--proto_path=. \
	./api/grpc/*.proto