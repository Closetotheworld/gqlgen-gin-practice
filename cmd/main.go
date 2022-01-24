package main

import (
	v1 "github.com/closetotheworld/gqlgen-gin-practice/internal/api/v1/graph"
	"github.com/closetotheworld/gqlgen-gin-practice/internal/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"os"
	"time"
)

const defaultPort = "8080"

const (
	V1QueryPath      = "/v1/query"
	V1PlaygroundPath = "/v1"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	r := gin.Default()

	r.Use(
		gin.Logger(),
		cors.New(cors.Config{
			AllowOrigins:     []string{"*"},
			AllowMethods:     []string{"GET", "OPTIONS", "POST"},
			AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			MaxAge:           1 * time.Hour,
		}),
		utils.GinContextToContextMiddleware(),
	)
	v1QueryHandler := v1.QueryHandler()
	r.GET(V1QueryPath, v1QueryHandler)
	r.OPTIONS(V1QueryPath, v1QueryHandler)
	r.POST(V1QueryPath, v1QueryHandler)
	r.GET(V1PlaygroundPath, v1.PlaygroundHandler(V1QueryPath))
	r.Run()
}
