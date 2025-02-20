package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vshakitskiy/reddit_comments/internal/api/handler"
	"github.com/vshakitskiy/reddit_comments/internal/api/middleware"
	"github.com/vshakitskiy/reddit_comments/internal/repository"
	"github.com/vshakitskiy/reddit_comments/internal/repository/pg"
	"github.com/vshakitskiy/reddit_comments/internal/resolver"
	"github.com/vshakitskiy/reddit_comments/internal/service"
)

func main() {
	db, err := pg.CreateDB()
	if err != nil {
		panic(err)
	}
	repo := repository.NewRepository(db)
	svc := service.NewService(repo)
	resolver := resolver.NewResolver(svc)

	r := gin.Default()
	r.Use(middleware.GinCtxToCtxMiddleware())
	handler.ApplyHandlers(r, resolver)

	r.Run()
}
