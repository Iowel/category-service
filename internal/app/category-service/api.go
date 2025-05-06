package categoryservice

import (
	"github.com/Iowel/category-service/internal/service/category"
	category_service "github.com/Iowel/category-service/pkg/api/category-service"
)

type Implimentation struct {
	category_service.UnimplementedCategoryServiceServer

	categoryService category.Service
}

func NewCategoryServiceServer(categoryService category.Service) category_service.CategoryServiceServer {
	return &Implimentation{categoryService: categoryService}
}
