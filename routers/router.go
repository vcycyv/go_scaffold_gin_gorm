package routers

import (
	"github.com/gin-gonic/gin"
	"{package_name}/{project_name}/pkg/setting"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})

	r.GET("/{entity_plural}", Get{entity_capital_plural})
	r.GET("/{entity_plural}/:id", Get{entity_capital})
	r.POST("/{entity_plural}", Add{entity_capital})
	r.PUT("/{entity_plural}/:id", Edit{entity_capital})
	r.DELETE("/{entity_plural}/:id", Delete{entity_capital})

	return r
}
