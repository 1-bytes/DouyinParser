package bootstrap

import (
	"DouyinParser/app/http/middlewares"
	"DouyinParser/routes"
	"github.com/gin-gonic/gin"
	"net/http"
)

// SetupRoute used for init Router and middleware.
func SetupRoute() http.Handler {
	router := gin.New()
	middlewares.Register(router)
	routes.Register(router)
	return router
}
