package handler

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/vektah/gqlparser/v2/ast"
	"github.com/vshakitskiy/reddit_comments/internal/graph"
	"github.com/vshakitskiy/reddit_comments/internal/resolver"
)

func ApplyHandlers(
	r *gin.Engine,
	resolver *resolver.Resolver,
) {
	r.POST("/query", GraphQLHandler(resolver))
	r.GET("/playground", PlaygroundHandler())
}

func GraphQLHandler(resolver *resolver.Resolver) gin.HandlerFunc {
	h := handler.New(graph.NewExecutableSchema(
		graph.Config{
			Resolvers: resolver,
		},
	))

	h.AddTransport(transport.Options{})
	h.AddTransport(transport.GET{})
	h.AddTransport(transport.POST{})

	h.SetQueryCache(
		lru.New[*ast.QueryDocument](1000),
	)

	h.Use(extension.Introspection{})
	h.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})
	// h.Use(extension.FixedComplexityLimit(1))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func PlaygroundHandler() gin.HandlerFunc {
	h := playground.Handler(
		"GraphQL playground",
		"/query",
	)

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
