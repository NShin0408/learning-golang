package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"learning-golang/controllers"
)

// SetupV1Routes はv1 APIグループのルートを設定します。
func SetupV1Routes(r *gin.RouterGroup, db *gorm.DB) {
	eg := r.Group("/")
	{
		// GetPosts ハンドラーをラップして直接使用
		eg.GET("/", func(c *gin.Context) {
			controllers.GetPosts(c, db)
		})

		// GetPost ハンドラーをラップして直接使用
		eg.GET("/post/:id", func(c *gin.Context) {
			controllers.GetPost(c, db)
		})
	}
}
