package main

import (
	v1 "github.com/closetotheworld/gqlgen-gin-practice/internal/api/v1/graph"
	"github.com/closetotheworld/gqlgen-gin-practice/internal/models"
	database "github.com/closetotheworld/gqlgen-gin-practice/internal/pkg/db/mysql"
	"github.com/closetotheworld/gqlgen-gin-practice/internal/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

const (
	v1QueryPath      = "/v1/query"
	v1PlaygroundPath = "/v1"
)

func main() {
	database.InitDB()
	models.Migrate()

	r := gin.Default()

	r.Use(
		gin.Logger(),
		cors.New(cors.Config{
			AllowOrigins:     []string{"https://lunit-care-mtp.s.lunit.io", "https://lunit-care-mtp.d.lunit.io"},
			AllowMethods:     []string{"GET", "OPTIONS", "POST"},
			AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			AllowOriginFunc: func(origin string) bool {
				return strings.HasSuffix(origin, "lunit.io")
			},
			MaxAge: 1 * time.Hour,
		}),
		// jwt middleware 추가
		utils.GinContextToContextMiddleware(),
	)

	v1QueryHandler := v1.QueryHandler()
	r.GET(v1QueryPath, v1QueryHandler)
	r.OPTIONS(v1QueryPath, v1QueryHandler)
	r.POST(v1QueryPath, v1QueryHandler)
	r.GET(v1PlaygroundPath, v1.PlaygroundHandler(v1QueryPath))
	r.Run()
}
