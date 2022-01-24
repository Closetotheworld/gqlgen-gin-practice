package graph

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/closetotheworld/gqlgen-gin-practice/api/v1/graph/generated"
	"github.com/closetotheworld/gqlgen-gin-practice/internal/api/v1/graph/stub"
	"github.com/gin-gonic/gin"
)

func PlaygroundHandler(queryPath string) gin.HandlerFunc {
	h := playground.Handler("GQL Playground", queryPath)
	return func(gc *gin.Context) {
		h.ServeHTTP(gc.Writer, gc.Request)
	}
}

func QueryHandler() gin.HandlerFunc {
	r := &stub.Stub{}

	c := generated.Config{Resolvers: r}
	s := generated.NewExecutableSchema(c)
	h := handler.New(s)

	h.SetQueryCache(lru.New(1000))

	h.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(100),
	})

	return func(gc *gin.Context) {
		h.ServeHTTP(gc.Writer, gc.Request)
	}
}
