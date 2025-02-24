package dataloader

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vikstrous/dataloadgen"
	"github.com/vshakitskiy/reddit_comments/internal/model"
	"github.com/vshakitskiy/reddit_comments/internal/repository"
)

type dataloader string

const dataloaderKey = dataloader("dataloader")

type Loader struct {
	UserLoader *dataloadgen.Loader[string, *model.User]
}

func NewLoader(repo *repository.Repository) *Loader {
	return &Loader{
		UserLoader: dataloadgen.NewLoader(
			repo.GetUsersByIDs,
			dataloadgen.WithWait(time.Millisecond),
		),
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