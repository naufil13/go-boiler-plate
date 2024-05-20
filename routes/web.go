package routes

import (
	"go-first-project/controllers"
	"go-first-project/middleware"

	"github.com/gin-gonic/gin"
)

func Webroutes() {
	r := gin.Default()

	/*Set Configs*/
	r.MaxMultipartMemory = 8 << 20 // 8 MiB

	/* Post Routes*/
	r.GET("/posts", controllers.PostsIndex)
	r.GET("/posts/:id", controllers.PostsShow)

	r.POST("/posts", controllers.PostsCreate)
	r.PUT("/posts/:id", controllers.PostsUpdate)
	r.DELETE("/posts/:id", controllers.PostsDelete)

	/* User Routes*/
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	/* Static Serving files*/
	r.LoadHTMLGlob("templates/*")

	r.POST("/upload", middleware.RequireAuth, controllers.MultiUpload)
	r.POST("/upload-aws", controllers.MultiUploadS3)

	/* Email Sending routes*/
	r.GET("/sendmail", controllers.SendMailSimple)

	r.Run() // listen and serve on 0.0.0.0:8080
}
