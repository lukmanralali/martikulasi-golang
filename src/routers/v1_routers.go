package routers

import (
	"../controllers"
	"github.com/gin-gonic/gin"
)

type Router struct {
	mainController controllers.UrlShrotenerController
}

func (routerReceiver *Router) RouterHandler(router *gin.Engine) {
	// router.GET(":shortcode", getUrlShortcode)
	group := router.Group("shorten")
	group.POST("", routerReceiver.mainController.ShortUrl)
	group.GET(":shortcode", routerReceiver.mainController.GetUrlShortcode)
	group.GET(":shortcode/stats", routerReceiver.mainController.GetUrlShortcodeStat)

}
