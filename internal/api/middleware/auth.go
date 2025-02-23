package middleware

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/vshakitskiy/reddit_comments/pkg/jwt"
)

type auth string

const authKey = auth("auth")

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.Next()
			return
		}

		bearer := "Bearer "
		if len(token) < len(bearer) {
			c.Next()
			return
		}

		token = token[len(bearer):]
		val, err := jwt.Validate(token)
		if err != nil || !val.Valid {
			c.Next()
			return
		}

		customClaim, _ := val.Claims.(*jwt.JwtCustomClaims)
		ctx := context.WithValue(c.Request.Context(), authKey, customClaim)

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func ExtractJwtClaims(
	ctx context.Context,
) *jwt.JwtCustomClaims {
	raw, _ := ctx.Value(authKey).(*jwt.JwtCustomClaims)
	return raw
}

func ExtractUserID(
	ctx context.Context,
) string {
	return ExtractJwtClaims(ctx).ID
}
