package categoryservice

import (
	"context"

	category_service "github.com/Iowel/category-service/pkg/category-service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implimentation) GetCategoryById(ctx context.Context, req *category_service.GetCategoryByIdRequest) (*category_service.GetCategoryByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCategoryById not implemented")
}
