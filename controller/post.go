package controller

import (
	"net/http"
	"post/common/logger"
	"post/errors"
	"post/models"
	"post/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PostController struct {
	errors.GinError
	service service.PostServiceInterface
}

func NewPostController(service service.PostServiceInterface) *PostController {
	return &PostController{service: service}
}

func (c *PostController) CreatePost(ginContext *gin.Context) {

	var post models.Post
	if err := ginContext.Bind(&post); err != nil {
		logger.LogError(err)
		ginContext.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id,err := c.service.CreatePost(&post);
	if  err != nil {
		logger.LogError(err)
		ginContext.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ginContext.JSON(http.StatusCreated, gin.H{"id": id})

}

func (c *PostController) ViewPost(ginContext *gin.Context) {

	postIDString := ginContext.Params.ByName("id")
	postID, err := strconv.Atoi(postIDString)
	if err != nil {
		logger.LogError(err)
		ginContext.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	post, err := c.service.ViewPost(postID)
	if err != nil {
		logger.LogError( err)
		ginContext.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ginContext.JSON(http.StatusCreated, gin.H{"post": post})

}

func (c *PostController) UpdatePost(ginContext *gin.Context) {

	var post models.Post
	if err := ginContext.Bind(&post); err != nil {
		logger.LogError( err)
		ginContext.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	postIDString := ginContext.Params.ByName("id")
	postID, err := strconv.Atoi(postIDString)
	if err != nil {
		logger.LogError(err)
		ginContext.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	post.ID = postID
	if err := c.service.UpdatePost(&post); err != nil {
		logger.LogError( err)
		ginContext.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ginContext.JSON(http.StatusCreated, gin.H{"post": post})

}
