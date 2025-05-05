package categoryservice

import (
	category_service "github.com/Iowel/category-service/pkg/category_service"
)

type Server struct {
	category_service.UnimplementedCategoryServiceServer
}
