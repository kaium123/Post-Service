package route

import (
	"net/http"
	"post/common/logger"
	"post/controller"
	"post/db"
	"post/middlewares"
	"post/pb"
	"post/redis"
	"post/repository"
	"post/service"

	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
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
	conn, err := grpc.Dial(viper.GetString("ATTACHMENTURL"), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	gRPCCLient := pb.NewAttachmentServiceClient(conn)

	raventClient := logger.NewRavenClient()
	logger := logger.NewLogger(raventClient)
	postRepo := repository.NewPostRepository(db, logger)
	redisConn := redis.NewRedisDb()
	redisRepo := repository.NewRedisRepository(redisConn, logger)
	reactRepo := repository.NewReactRepository(db, logger)
	reactService := service.NewReactService(reactRepo)
	reactController := controller.NewReactController(reactService)

	react := api.Group("/react").Use(middlewares.Auth())

	react.POST("/like", reactController.Like)
	react.POST("/unlike", reactController.Unlike)

	commentRepo := repository.NewCommentRepository(db, logger)
	commentService := service.NewCommentService(commentRepo)
	commentController := controller.NewCommentController(commentService)

	comment := api.Group("/comment").Use(middlewares.Auth())

	comment.POST("/create", commentController.CreateComment)
	comment.GET("/view/:id", commentController.ViewComment)
	comment.POST("/update/:id", commentController.UpdateComment)
	comment.GET("/list/:post_id", commentController.AllComment)
	comment.DELETE("/delete/:id", commentController.Delete)

	postService := service.NewPostService(gRPCCLient, postRepo, commentRepo, reactRepo, redisRepo)
	postController := controller.NewPostController(postService)

	post := api.Group("/post").Use(middlewares.Auth())
	post.POST("/create", postController.CreatePost)
	post.GET("/view/:id", postController.ViewPost)
	post.POST("/update/:id", postController.UpdatePost)
	post.GET("/list", postController.AllPost)

	// shareRepo := repository.NewShareRepository(db, logger)
	// shareService := service.NewShareService(shareRepo)
	// shareController := controller.NewShareController(shareService)

	// share := api.Group("/share")

	// share.POST("/create", shareController.CreateShare)
	// share.GET("/view/:id", shareController.ViewShare)
	// share.POST("/update/:id", shareController.UpdateShare)
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
