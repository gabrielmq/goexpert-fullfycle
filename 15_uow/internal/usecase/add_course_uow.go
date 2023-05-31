package usecase

import (
	"context"

	"github.com/gabrielmq/uow/internal/entity"
	"github.com/gabrielmq/uow/internal/repository"
	"github.com/gabrielmq/uow/pkg/uow"
)

type InputUseCaseUow struct {
	CategoryName     string
	CourseName       string
	CourseCategoryID int
}

type AddCourseUseCaseUow struct {
	Uow uow.UowInterface
}

func NewAddCourseUseCaseUow(uow uow.UowInterface) *AddCourseUseCaseUow {
	return &AddCourseUseCaseUow{Uow: uow}
}

func (a *AddCourseUseCaseUow) Execute(ctx context.Context, input InputUseCaseUow) error {
	return a.Uow.Do(ctx, func(uow *uow.Uow) error {
		category := entity.Category{
			Name: input.CategoryName,
		}

		categoryRepository := a.getCategoryRepository(ctx)
		if err := categoryRepository.Insert(ctx, category); err != nil {
			return err
		}

		course := entity.Course{
			Name:       input.CourseName,
			CategoryID: input.CourseCategoryID,
		}

		courseRepository := a.getCourseRepository(ctx)
		if err := courseRepository.Insert(ctx, course); err != nil {
			return err
		}
		return nil
	})
}

func (a *AddCourseUseCaseUow) getCategoryRepository(ctx context.Context) repository.CategoryRepositoryInterface {
	repo, err := a.Uow.GetRepository(ctx, "CategoryRepository")
	if err != nil {
		panic(err)
	}
	return repo.(repository.CategoryRepositoryInterface)
}

func (a *AddCourseUseCaseUow) getCourseRepository(ctx context.Context) repository.CourseRepositoryInterface {
	repo, err := a.Uow.GetRepository(ctx, "CourseRepository")
	if err != nil {
		panic(err)
	}
	return repo.(repository.CourseRepositoryInterface)
}
