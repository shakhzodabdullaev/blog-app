package handlers

import (
	"blog/app/api/docs"
	"blog/app/api/handlers"
	"blog/app/config"

	_ "blog/app/api/docs"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @description This is a sample article demo.
// @termsOfService https://udevs.io
func SetUpAPI(r *gin.Engine, h handlers.Handler, cfg config.Config) {
	docs.SwaggerInfo.Title = cfg.App
	docs.SwaggerInfo.Version = cfg.Version
	docs.SwaggerInfo.Host = cfg.ServiceHost + cfg.HTTPPort
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r.POST("/articles", h.CreateArticle)
	r.GET("/articles", h.GetArticleList)
	// r.GET("/articles/:id", GetByIDHandler)
	// r.PUT("/articles/:id", UpdateHandler)
	// r.DELETE("/articles/:id", DeleteHandler)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
