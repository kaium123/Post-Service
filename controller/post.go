package controller

import (
	"net/http"
	"post/common/logger"
	"post/errors"
	"post/models"
	"post/service"
	"strconv"
	"strings"

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
	userID := int(ginContext.GetInt64("user_id"))


	var post models.Post
	if err := ginContext.Bind(&post); err != nil {
		logger.LogError(err)
		ginContext.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	post.UserID=userID

	id, err := c.service.CreatePost(&post)
	if err != nil {
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
		logger.LogError(err)
		ginContext.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ginContext.JSON(http.StatusCreated, gin.H{"post": post})

}

func (c *PostController) UpdatePost(ginContext *gin.Context) {

	var post models.Post
	if err := ginContext.Bind(&post); err != nil {
		logger.LogError(err)
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
		logger.LogError(err)
		ginContext.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ginContext.JSON(http.StatusCreated, gin.H{"post": post})

}

func (c *PostController) AllPost(ginContext *gin.Context) {

	cookie, err := ginContext.Cookie("access_token")
	authorizationHeader := ginContext.Request.Header.Get("Authorization")
	fields := strings.Fields(authorizationHeader)
	accessToken := ""
	if len(fields) != 0 && fields[0] == "Bearer" {
		accessToken = fields[0] + " " + fields[1]
	} else if err == nil {
		accessToken = "Bearer " + cookie
	}
	userID := int(ginContext.GetInt64("user_id"))
	keyword:=ginContext.Query("keyword")
	requestParam:=&models.RequestParams{
		Keyword: keyword,
	}

	posts, err := c.service.AllPosts(userID,accessToken,*requestParam)
	if err != nil {
		logger.LogError(err)
		ginContext.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ginContext.JSON(http.StatusCreated, gin.H{"posts": posts})

}
