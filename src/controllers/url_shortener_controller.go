package controller

import (
	"../objects"
	"../services"
)

type UrlShrotenerController struct {
	urlShortenerService services.UrlShortenerService
	errorHelper helpers.ErrorHelper
}

func UrlShrotenerControllerHandler(router *gin.Engine) {

	handler := &V1UserController{
		urlShortenerService: services.UrlShortenerServiceHandler(),
		errorHelper: helpers.ErrorHelperHandler(),
	}
	group := router.Group("shorten")
	// test url
	// group.GET("/1", handler.GetById)
	// real url
	group.POST("", handler.makeShortUrl)

}