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

type CommentController struct {
	errors.GinError
	service service.CommentServiceInterface
}

func NewCommentController(service service.CommentServiceInterface) *CommentController {
	return &CommentController{service: service}
}

func (c *CommentController) CreateComment(ginContext *gin.Context) {

	var Comment models.Comment
	if err := ginContext.Bind(&Comment); err != nil {
		logger.LogError(err)
		ginContext.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.CreateComment(&Comment); err != nil {
		logger.LogError(err)
		ginContext.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ginContext.JSON(http.StatusCreated, gin.H{"Comment": Comment})

}

func (c *CommentController) ViewComment(ginContext *gin.Context) {

	CommentIDString := ginContext.Params.ByName("id")
	CommentID, err := strconv.Atoi(CommentIDString)
	if err != nil {
		logger.LogError(err)
		ginContext.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	Comment, err := c.service.ViewComment(CommentID)
	if err != nil {
		logger.LogError(err)
		ginContext.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ginContext.JSON(http.StatusCreated, gin.H{"Comment": Comment})

}

func (c *CommentController) UpdateComment(ginContext *gin.Context) {

	var Comment models.Comment
	if err := ginContext.Bind(&Comment); err != nil {
		logger.LogError(err)
		ginContext.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	CommentIDString := ginContext.Params.ByName("id")
	CommentID, err := strconv.Atoi(CommentIDString)
	if err != nil {
		logger.LogError(err)
		ginContext.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	Comment.ID = CommentID
	if err := c.service.UpdateComment(&Comment); err != nil {
		logger.LogError(err)
		ginContext.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ginContext.JSON(http.StatusCreated, gin.H{"Comment": Comment})

}

func (c *CommentController) AllComment(ginContext *gin.Context) {

	postIDString := ginContext.Params.ByName("post_id")
	postID, err := strconv.Atoi(postIDString)
	if err != nil {
		logger.LogError(err)
		ginContext.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	Comment, err := c.service.AllComment(postID)
	if err != nil {
		logger.LogError(err)
		ginContext.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ginContext.JSON(http.StatusCreated, gin.H{"Comment": Comment})

}

func (c *CommentController) Delete(ginContext *gin.Context) {

	CommentIDString := ginContext.Params.ByName("id")
	CommentID, err := strconv.Atoi(CommentIDString)
	if err != nil {
		logger.LogError(err)
		ginContext.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	err = c.service.Delete(CommentID)
	if err != nil {
		logger.LogError(err)
		ginContext.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ginContext.JSON(http.StatusCreated, gin.H{"Comment": "successfully deleted"})

}
