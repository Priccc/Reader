package router

import (
	"github.com/gin-gonic/gin"
)

func initStatic(router *gin.Engine) {

	// router.Static("/css", "./assets/static/css")
	// router.Static("/images", "./assets/static/images")
	// router.Static("/js", "./assets/static/js")
	router.Static("/static", "./assets/static")
}
