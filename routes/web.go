package routes

import (
	"DouyinParser/app/http/controllers"
	"github.com/gin-gonic/gin"
)

// Register used for register router.
func Register(r *gin.Engine) {
	v := r.Group("api/v1")
	userGroup := v.Group("videos")
	{
		controller := controllers.VideosController{}
		userGroup.GET("parser", controller.ParserURL)
	}
}
