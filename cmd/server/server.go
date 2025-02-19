package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vshakitskiy/reddit_comments/internal/api/handler"
	"github.com/vshakitskiy/reddit_comments/internal/api/middleware"
)

func main() {
	r := gin.Default()
	r.Use(middleware.GinCtxToCtxMiddleware())
	handler.ApplyHandlers(r)

	r.Run()
}
