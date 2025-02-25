package dataloader

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/vikstrous/dataloadgen"
	"github.com/vshakitskiy/reddit_comments/internal/model"
	"github.com/vshakitskiy/reddit_comments/internal/repository"
)

type dataloader string

const dataloaderKey = dataloader("dataloader")

type Loader struct {
	UserLoader   *dataloadgen.Loader[string, *model.User]
	ParentLoader *dataloadgen.Loader[string, *model.Comment]
}

func NewLoader(repo *repository.Repository) *Loader {
	return &Loader{
		UserLoader:   NewUserLoader(repo),
		ParentLoader: NewParentLoader(repo),
	}
}

func Middleware(repo *repository.Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		loader := NewLoader(repo)
		ctx := context.WithValue(
			c.Request.Context(),
			dataloaderKey,
			loader,
		)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func LoaderFromCtx(ctx context.Context) *Loader {
	return ctx.Value(dataloaderKey).(*Loader)
}
