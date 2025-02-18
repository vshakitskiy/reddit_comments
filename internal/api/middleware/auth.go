package middleware

import (
	"context"
	"net/http"

	"github.com/vshakitskiy/reddit_comments/internal/service"
	"github.com/vshakitskiy/reddit_comments/pkg/jwt"
)

type authKey string

const ctxKey = authKey("auth")

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
		token = token[len(bearer):]

		val, err := jwt.Validate(token)
		if err != nil || !val.Valid {
			http.Error(w, "unauthorized", http.StatusForbidden)
		}

		customClaim, _ := val.Claims.(*jwt.JwtCustomClaims)
		ctx := context.WithValue(r.Context(), ctxKey, customClaim)

		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func ExtractJwtClaims(
	ctx context.Context,
) *service.JwtCustomClaims {
	raw, _ := ctx.Value(ctxKey).(*service.JwtCustomClaims)
	return raw
}
