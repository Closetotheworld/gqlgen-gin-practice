package graph

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/closetotheworld/gqlgen-gin-practice/api/v1/graph/generated"
	"github.com/closetotheworld/gqlgen-gin-practice/internal/api/v1/graph/stub"
	"github.com/closetotheworld/gqlgen-gin-practice/internal/pkg/Todo"
	database "github.com/closetotheworld/gqlgen-gin-practice/internal/pkg/db/mysql"
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

	todoService := Todo.InitTodoService(database.Db)
	r.QueryResolver.Todo = todoService.Todo
	r.QueryResolver.Todos = todoService.Todos
	r.MutationResolver.CreateTodo = todoService.CreateTodo
	r.MutationResolver.UpdateTodo = todoService.UpdateTodo

	c := generated.Config{Resolvers: r}
	s := generated.NewExecutableSchema(c)
	h := handler.New(s)

	h.AddTransport(transport.Options{})
	h.AddTransport(transport.GET{})
	h.AddTransport(transport.POST{})
	h.AddTransport(transport.MultipartForm{})

	h.SetQueryCache(lru.New(1000))

	h.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(100),
	})

	return func(gc *gin.Context) {
		h.ServeHTTP(gc.Writer, gc.Request)
	}
}
