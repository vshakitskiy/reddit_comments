package middleware

import (
	"context"
	"net/http"

	"github.com/vshakitskiy/reddit_comments/pkg/jwt"
)

type auth string

const authKey = auth("auth")

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(
		w http.ResponseWriter,
		r *http.Request,
	) {
		token := r.Header.Get("Authorization")
		if token == "" {
			next.ServeHTTP(w, r)
			return
		}

		bearer := "Bearer "
		if len(token) < len(bearer) {
			next.ServeHTTP(w, r)
			return
		}

		token = token[len(bearer):]
		val, err := jwt.Validate(token)
		if err != nil || !val.Valid {
			next.ServeHTTP(w, r)
			return
		}

		customClaim, _ := val.Claims.(*jwt.JwtCustomClaims)
		ctx := context.WithValue(r.Context(), authKey, customClaim)

		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func ExtractJwtClaims(
	ctx context.Context,
) *jwt.JwtCustomClaims {
	raw, _ := ctx.Value(authKey).(*jwt.JwtCustomClaims)
	return raw
}
