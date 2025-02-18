package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/mux"
	"github.com/vektah/gqlparser/v2/ast"
	"github.com/vshakitskiy/reddit_comments/internal/api/middleware"
	"github.com/vshakitskiy/reddit_comments/internal/directive"
	"github.com/vshakitskiy/reddit_comments/internal/graph"
	"github.com/vshakitskiy/reddit_comments/internal/repository"
	"github.com/vshakitskiy/reddit_comments/internal/resolver"
	"github.com/vshakitskiy/reddit_comments/internal/service"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	repo := repository.NewRepository("inmemory")
	service := service.NewService(repo)
	resolver := resolver.NewResolver(service)

	r := mux.NewRouter()
	r.Use(middleware.AuthMiddleware)

	config := graph.Config{
		Resolvers: resolver,
		Directives: graph.DirectiveRoot{
			Auth: directive.Auth,
		},
	}

	srv := handler.New(graph.NewExecutableSchema(config))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	r.Handle("/", playground.Handler(
		"GraphQL playground",
		"/query",
	))
	r.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
