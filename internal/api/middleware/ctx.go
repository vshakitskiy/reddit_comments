package middleware

import (
	"context"
	"errors"

	"github.com/gin-gonic/gin"
)

type ginCtx string

const ginCtxKey = ginCtx("GinCtx")

func GinCtxToCtxMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(
			c.Request.Context(),
			ginCtxKey,
			c,
		)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func GinCtxFromCtx(ctx context.Context) (*gin.Context, error) {
	ginCtx := ctx.Value(ginCtxKey)

	if ginCtx == nil {
		return nil, errors.New("gin.Context not found")
	}

	c, ok := ginCtx.(*gin.Context)
	if !ok {
		return nil, errors.New("gin.Context has wrong type")
	}

	return c, nil
}
