package pg

import (
	"math"

	"github.com/vshakitskiy/reddit_comments/internal/model"
	"github.com/vshakitskiy/reddit_comments/pkg/pagination"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CreateDB() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=12345 dbname=reddit port=5555 sslmode=disable"
	dialector := postgres.Open(dsn)
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = AutoMigrate(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.User{},
		&model.Post{},
		&model.Comment{},
	)
}

type paginateCallback func(db *gorm.DB) *gorm.DB



func Paginate(
	value interface{},
	pagination *pagination.Pagination,
	db *gorm.DB,
	totalRows int64,
) paginateCallback {
	pagination.TotalRows = totalRows

	totalPages := int(math.Ceil(
		float64(totalRows) / float64(pagination.GetLimit()),
	))
	pagination.TotalPages = totalPages

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(int(pagination.GetOffset())).Limit(
			int(pagination.GetLimit()),
		)
	}
}
