package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/vshakitskiy/reddit_comments/internal/api/handler"
	"github.com/vshakitskiy/reddit_comments/internal/api/middleware"
	"github.com/vshakitskiy/reddit_comments/internal/dataloader"
	"github.com/vshakitskiy/reddit_comments/internal/pubsub"
	"github.com/vshakitskiy/reddit_comments/internal/repository"
	"github.com/vshakitskiy/reddit_comments/internal/repository/pg"
	"github.com/vshakitskiy/reddit_comments/internal/resolver"
	"github.com/vshakitskiy/reddit_comments/internal/service"
)

func main() {
	db, err := pg.CreateDB()
	if err != nil {
		os.Exit(1)
	}
	commentPubSub := pubsub.NewCommentPubSub()
	repo := repository.NewRepository(db)
	svc := service.NewService(repo, commentPubSub)
	resolver := resolver.NewResolver(svc)

	r := gin.Default()
	r.Use(middleware.AuthMiddleware())
	r.Use(dataloader.Middleware(repo))
	r.Use(middleware.GinCtxToCtxMiddleware())
	handler.ApplyHandlers(r, resolver)

	r.Run()
}
