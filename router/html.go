package router
import (
	"github.com/gin-gonic/gin"
	"project/controllers"
)

func initHtml(router *gin.Engine)  {
	router.LoadHTMLGlob("assets/*.html")

	router.GET("/index", controllers.IndexController)
}
