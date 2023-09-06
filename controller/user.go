package controller

import (
	"auth/common/logger"
	"auth/common/utils"
	"auth/errors"
	"auth/models"
	"auth/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	errors.GinError
	service service.UserServiceInterface
}

func NewUserController(service service.UserServiceInterface) *UserController {
	return &UserController{service: service}
}

func (c *UserController) Register(ginContext *gin.Context) {
	var user models.User
	if err := ginContext.Bind(&user); err != nil {
		logger.LogError("failed to query estimate ", err)
		ginContext.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": utils.Trans("invalidRequestParams", nil)})
		return
	}

	if err := c.service.Register(user); err!= nil {
        logger.LogError("failed to query estimate ", err)
        ginContext.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": utils.Trans("invalidRequestParams", nil)})
        return
    }


}

func (c *UserController) SignIn(ginContext *gin.Context) {
	var signInInfo models.SignInData
	if err := ginContext.Bind(&signInInfo); err != nil {
		logger.LogError("failed to query estimate ", err)
		ginContext.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": utils.Trans("invalidRequestParams", nil)})
		return
	}

	if err := c.service.SignIn(signInInfo); err!= nil {
        logger.LogError("failed to query estimate ", err)
        ginContext.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": utils.Trans("invalidRequestParams", nil)})
        return
    }


}
