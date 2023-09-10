package controller

import (
	"net/http"
	"post/common/logger"
	"post/errors"
	"post/models"
	"post/service"

	"github.com/gin-gonic/gin"
)

type ReactController struct {
	errors.GinError
	service service.ReactServiceInterface
}

func NewReactController(service service.ReactServiceInterface) *ReactController {
	return &ReactController{service: service}
}

func (c *ReactController) Like(ginContext *gin.Context) {

	userID := int(ginContext.GetInt64("user_id"))
	logger.LogInfo(userID)
	var React models.React
	if err := ginContext.Bind(&React); err != nil {
		logger.LogError(err)
		ginContext.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	React.ReactedUserID = uint(userID)
	if err := c.service.CreateReact(&React); err != nil {
		logger.LogError(err)
		ginContext.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ginContext.JSON(http.StatusCreated, gin.H{"React": React})

}

func (c *ReactController) Unlike(ginContext *gin.Context) {

	userID := int(ginContext.GetInt64("user_id"))
	var React models.React
	if err := ginContext.Bind(&React); err != nil {
		logger.LogError(err)
		ginContext.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	React.ReactedUserID = uint(userID)
	if err := c.service.Unlike(&React); err != nil {
		logger.LogError(err)
		ginContext.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ginContext.JSON(http.StatusCreated, gin.H{"React": React})

}
