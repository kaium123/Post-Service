package route

import (
	"net/http"
	"post/common/logger"
	"post/controller"
	"post/db"
	"post/repository"
	"post/service"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Setup() *gin.Engine {
	gin.SetMode(viper.GetString("GIN_MODE"))

	r := gin.New()
	setupCors(r)

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	api := r.Group("/api")

	db := db.InitDB()

	raventClient := logger.NewRavenClient()
	logger := logger.NewLogger(raventClient)
	repo := repository.NewPostRepository(db, logger)
	service := service.NewPostService(repo)
	postController := controller.NewPostController(service)

	post := api.Group("/post")

	post.POST("/create", postController.CreatePost)
	post.GET("/view/:id", postController.ViewPost)
	post.POST("/update/:id", postController.UpdatePost)
	return r
}

func setupCors(r *gin.Engine) {
	allowConf := viper.GetString("CORS_ALLOW_ORIGINS")
	if allowConf == "" {
		r.Use(cors.Default())
		return
	}
	allowSites := strings.Split(allowConf, ",")
	config := cors.DefaultConfig()
	config.AllowOrigins = allowSites
}
