package directive

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"github.com/vshakitskiy/reddit_comments/internal/api/middleware"
)

func Auth(
	ctx context.Context,
	obj interface{},
	next graphql.Resolver,
) (interface{}, error) {
	tokenData := middleware.ExtractJwtClaims(ctx)

	if tokenData == nil {
		return nil, &gqlerror.Error{
			Message: "unauthorized",
		}
	}

	return next(ctx)
}
