module github.com/Iowel/category-service

go 1.24.1

require (
	github.com/Iowel/category-service/pkg/api/category-service v0.0.0-20250505163423-3eb98ed7471d
	google.golang.org/grpc v1.72.0
)

require (
	github.com/pkg/errors v0.9.1
	golang.org/x/net v0.37.0 // indirect
	golang.org/x/sys v0.31.0 // indirect
	golang.org/x/text v0.23.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250425173222-7b384671a197 // indirect
	google.golang.org/protobuf v1.36.6 // indirect
)

replace github.com/Iowel/category-service/pkg/api/category-service => ./pkg/api/category-service
