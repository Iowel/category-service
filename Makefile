gen:
	@protoc -I. -I./googleapis \
       --go_out=pkg/ --go_opt=paths=source_relative \
       --go-grpc_out=pkg/ --go-grpc_opt=paths=source_relative \
       api/category-service/category-service.proto


grpcCurl:
	@/bin/sh example/get-category-by-id.sh



deps:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.27.1
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1.0
	go installgithub.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.5.0
	go installgithub.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.5.0

