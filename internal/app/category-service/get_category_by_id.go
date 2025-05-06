package categoryservice

import (
	"context"
	"errors"

	"github.com/Iowel/category-service/internal/service/category"
	category_service "github.com/Iowel/category-service/pkg/api/category-service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implimentation) GetCategoryById(ctx context.Context, req *category_service.GetCategoryByIdRequest) (*category_service.GetCategoryByIdResponse, error) {

	cat, err := i.categoryService.GetCategoryByID(ctx, req.GetId())
	if err != nil {
		if errors.Is(err, category.ErrNoCategory) {
			return nil, status.Errorf(codes.NotFound, "category not found")
		}
		return nil, status.Errorf(codes.Internal, "internal error: %v", err)
	}

	return makeGetCategoryByIdResponse(cat), nil 
}


func makeGetCategoryByIdResponse(cat *category.Category) *category_service.GetCategoryByIdResponse {
	pbCat := &category_service.Category {
		Id: cat.ID,
		Name: cat.Name,
	}
	return &category_service.GetCategoryByIdResponse{
		Category: pbCat,
	}
}