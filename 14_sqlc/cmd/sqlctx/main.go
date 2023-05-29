package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/gabrielmq/sqlc/internal/db"
	_ "github.com/go-sql-driver/mysql"
)

type CourseDB struct {
	*db.Queries
	db *sql.DB
}

type CourseParams struct {
	ID          string
	Name        string
	Price       float64
	Description sql.NullString
}

type CategoryParams struct {
	ID          string
	Name        string
	Description sql.NullString
}

func NewCourseDB(dbConn *sql.DB) *CourseDB {
	return &CourseDB{
		Queries: db.New(dbConn),
		db:      dbConn,
	}
}

func (c *CourseDB) CreateCourseAndCategory(ctx context.Context, argsCategory CategoryParams, argsCourse CourseParams) error {
	return c.callTx(ctx, func(q *db.Queries) error {
		err := q.CreateCategory(ctx, db.CreateCategoryParams{
			ID:          argsCategory.ID,
			Name:        argsCategory.Name,
			Description: argsCategory.Description,
		})
		if err != nil {
			return err
		}

		err = q.CreateCourse(ctx, db.CreateCourseParams{
			ID:          argsCourse.ID,
			Name:        argsCourse.Name,
			Description: argsCourse.Description,
			Price:       argsCourse.Price,
			CategoryID:  argsCategory.ID,
		})
		if err != nil {
			return err
		}
		return nil
	})
}

func (c *CourseDB) callTx(ctx context.Context, fn func(*db.Queries) error) error {
	tx, err := c.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	if err := fn(db.New(tx)); err != nil {
		if errRb := tx.Rollback(); errRb != nil {
			return fmt.Errorf("error on rollback: %v, original error: %w", errRb, err)
		}
		return err
	}
	return tx.Commit()
}

func main() {
	ctx := context.Background()

	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	queries := db.New(dbConn)

	courses, err := queries.ListCourses(ctx)
	if err != nil {
		panic(err)
	}
	for _, course := range courses {
		fmt.Printf("Category: %s, Course ID: %s, Course Name: %s, Course Description: %s, Course Price: %f",
			course.CategoryName, course.ID, course.Name, course.Description.String, course.Price)
	}

	// courseArgs := CourseParams{
	// 	ID:          uuid.New().String(),
	// 	Name:        "Go",
	// 	Description: sql.NullString{String: "Go Course", Valid: true},
	// 	Price:       10.95,
	// }

	// categoryArgs := CategoryParams{
	// 	ID:          uuid.New().String(),
	// 	Name:        "Backend",
	// 	Description: sql.NullString{String: "Backend Course", Valid: true},
	// }

	// courseDB := NewCourseDB(dbConn)
	// err = courseDB.CreateCourseAndCategory(ctx, categoryArgs, courseArgs)
	// if err != nil {
	// 	panic(err)
	// }
}
