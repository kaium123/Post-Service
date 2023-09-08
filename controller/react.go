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

type ReactController struct {
	errors.GinError
	service service.ReactServiceInterface
}

func NewReactController(service service.ReactServiceInterface) *ReactController {
	return &ReactController{service: service}
}

func (c *ReactController) CreateReact(ginContext *gin.Context) {

	var React models.React
	if err := ginContext.Bind(&React); err != nil {
		logger.LogError(err)
		ginContext.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.CreateReact(&React); err != nil {
		logger.LogError(err)
		ginContext.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ginContext.JSON(http.StatusCreated, gin.H{"React": React})

}

func (c *ReactController) ViewReact(ginContext *gin.Context) {

	ReactIDString := ginContext.Params.ByName("id")
	ReactID, err := strconv.Atoi(ReactIDString)
	if err != nil {
		logger.LogError(err)
		ginContext.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	React, err := c.service.ViewReact(ReactID)
	if err != nil {
		logger.LogError(err)
		ginContext.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ginContext.JSON(http.StatusCreated, gin.H{"React": React})

}

func (c *ReactController) UpdateReact(ginContext *gin.Context) {

	var React models.React
	if err := ginContext.Bind(&React); err != nil {
		logger.LogError(err)
		ginContext.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ReactIDString := ginContext.Params.ByName("id")
	ReactID, err := strconv.Atoi(ReactIDString)
	if err != nil {
		logger.LogError(err)
		ginContext.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	React.ID = ReactID
	if err := c.service.UpdateReact(&React); err != nil {
		logger.LogError(err)
		ginContext.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ginContext.JSON(http.StatusCreated, gin.H{"React": React})

}
