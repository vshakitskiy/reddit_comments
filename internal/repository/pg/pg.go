package pg

import (
	"github.com/vshakitskiy/reddit_comments/internal/model"
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
