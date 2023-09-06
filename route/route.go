// Application
//
// Application description
//
//	Schemes: http
//	Host: localhost:8080
//	BasePath: /api
//	Version: 0.0.1
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
package route

import (
	"auth/common/logger"
	"auth/controller"
	"auth/db"
	"auth/repository"
	"auth/service"
	"net/http"
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
	repo := repository.NewUserRepository(db, logger)
	service := service.NewUserService(repo)
	userController := controller.NewUserController(service)

	user := api.Group("/user")

	user.GET("/:sign-in", userController.SignIn)
	user.POST("/register", userController.Register)
	//user.POST("/update/:post_id", postController.UpdatePost)
	// post.GET("/list", postController.AllPost)
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
