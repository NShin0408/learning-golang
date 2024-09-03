package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"learning-golang/controllers"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("views/*")
	router.GET("/", func(c *gin.Context) {
		controllers.GetPosts(c, db)
	})
	router.GET("/post/:id", func(c *gin.Context) {
		controllers.GetPost(c, db)
	})
	return router
}
