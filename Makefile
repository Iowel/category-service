gen:
	@protoc -I. -I./googleapis \
       --go_out=pkg/ --go_opt=paths=source_relative \
       --go-grpc_out=pkg/ --go-grpc_opt=paths=source_relative \
       api/category-service/category-service.proto
