package service

import (
	"context"
	"errors"
	"io"

	"github.com/gabrielmq/grpc/internal/database"
	"github.com/gabrielmq/grpc/internal/pb"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDB database.Category
}

func NewCategoryService(db database.Category) *CategoryService {
	return &CategoryService{CategoryDB: db}
}

func (c *CategoryService) CreateCategory(ctx context.Context, in *pb.CreateCategoyRequest) (*pb.CategoryResponse, error) {
	category, err := c.CategoryDB.Create(in.Name, in.Description)
	if err != nil {
		return nil, err
	}

	return &pb.CategoryResponse{
		Category: &pb.Category{
			Id:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		},
	}, nil
}

func (c *CategoryService) ListCategories(ctx context.Context, blank *pb.Blank) (*pb.CategoryList, error) {
	categories, err := c.CategoryDB.FindAll()
	if err != nil {
		return nil, err
	}

	var categoryList []*pb.Category
	for _, category := range categories {
		categoryList = append(categoryList, &pb.Category{
			Id:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		})
	}
	return &pb.CategoryList{Categories: categoryList}, nil
}

func (c *CategoryService) GetCategory(ctx context.Context, in *pb.GetCategoryRequest) (*pb.Category, error) {
	category, err := c.CategoryDB.FindByID(in.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}, nil
}

func (c *CategoryService) CreateCategoryStream(stream pb.CategoryService_CreateCategoryStreamServer) error {
	categories := &pb.CategoryList{}
	for {
		// vai recebendo
		category, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			return stream.SendAndClose(categories)
		}

		if err != nil {
			return err
		}

		categoryResult, err := c.CategoryDB.Create(category.Name, category.Description)
		if err != nil {
			return err
		}

		// no final retorna tudo de uma vez
		categories.Categories = append(categories.Categories, &pb.Category{
			Id:          categoryResult.ID,
			Name:        categoryResult.Name,
			Description: categoryResult.Description,
		})
	}
}

func (c *CategoryService) CreateCategoryStreamBidirectional(stream pb.CategoryService_CreateCategoryStreamBidirectionalServer) error {
	for {
		// vai recebendo
		category, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			return nil
		}

		if err != nil {
			return err
		}

		categoryResult, err := c.CategoryDB.Create(category.Name, category.Description)
		if err != nil {
			return err
		}

		// e devolvendo o resultado
		err = stream.Send(&pb.Category{
			Id:          categoryResult.ID,
			Name:        categoryResult.Name,
			Description: categoryResult.Description,
		})
		if err != nil {
			return err
		}
	}
}
