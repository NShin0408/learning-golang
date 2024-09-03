package controllers

import (
	"fmt"
	"gorm.io/gorm"
	"learning-golang/constants"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"learning-golang/models"
)

// ErrorResponse はエラーレスポンスを表す構造体です。
type ErrorResponse struct {
	Error string `json:"error"`
}

// ErrorResponseJSON はBadRequestやNotFoundなどのエラーレスポンスを返します。
func ErrorResponseJSON(c *gin.Context, statusCode int, errorMessage string) {
	c.JSON(statusCode, ErrorResponse{Error: errorMessage})
}

// GetPosts godoc
//
//	@Summary		Get all posts
//	@Description	Get all posts from the db
//	@Tags			posts
//	@Accept			json
//	@Produce		json
//	@Param			NoParams	query		bool	false	"No parameters"
//	@Success		200			{object}	[]models.Post
//	@Failure		500			{object}	httputil.HTTPError
//	@Router			/ [get]
func GetPosts(c *gin.Context, db *gorm.DB) {
	var posts []models.Post
	if err := db.Find(&posts).Error; err != nil {
		fmt.Println("Failed to fetch posts:", err)
		ErrorResponseJSON(c, http.StatusInternalServerError, constants.ErrFetchPosts)
		return
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Home",
		"posts": posts,
	})
}

// GetPost godoc
//
//	@Summary		Get a post by ID
//	@Description	Get a post from the db by its ID
//	@Tags			posts
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Post ID"
//	@Success		200	{object}	models.Post
//	@Failure		400	{object}	string	"Invalid post ID"
//	@Failure		404	{object}	string	"Post not found"
//	@Router			/posts/{id} [get]
func GetPost(c *gin.Context, db *gorm.DB) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ErrorResponseJSON(c, http.StatusBadRequest, constants.ErrInvalidPostID)
		return
	}

	var post models.Post
	if err := db.First(&post, id).Error; err != nil {
		ErrorResponseJSON(c, http.StatusNotFound, constants.ErrPostNotFound)
		return
	}

	c.HTML(http.StatusOK, "post.html", gin.H{
		"title": "Post",
		"post":  post,
	})
}
