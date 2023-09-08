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

type ShareController struct {
	errors.GinError
	service service.ShareServiceInterface
}

func NewShareController(service service.ShareServiceInterface) *ShareController {
	return &ShareController{service: service}
}

func (c *ShareController) CreateShare(ginContext *gin.Context) {

	var Share models.Share
	if err := ginContext.Bind(&Share); err != nil {
		logger.LogError(err)
		ginContext.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.CreateShare(&Share); err != nil {
		logger.LogError(err)
		ginContext.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ginContext.JSON(http.StatusCreated, gin.H{"Share": Share})

}

func (c *ShareController) ViewShare(ginContext *gin.Context) {

	ShareIDString := ginContext.Params.ByName("id")
	ShareID, err := strconv.Atoi(ShareIDString)
	if err != nil {
		logger.LogError(err)
		ginContext.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	Share, err := c.service.ViewShare(ShareID)
	if err != nil {
		logger.LogError(err)
		ginContext.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ginContext.JSON(http.StatusCreated, gin.H{"Share": Share})

}

func (c *ShareController) UpdateShare(ginContext *gin.Context) {

	var Share models.Share
	if err := ginContext.Bind(&Share); err != nil {
		logger.LogError(err)
		ginContext.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ShareIDString := ginContext.Params.ByName("id")
	ShareID, err := strconv.Atoi(ShareIDString)
	if err != nil {
		logger.LogError(err)
		ginContext.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	Share.ID = ShareID
	if err := c.service.UpdateShare(&Share); err != nil {
		logger.LogError(err)
		ginContext.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ginContext.JSON(http.StatusCreated, gin.H{"Share": Share})

}
