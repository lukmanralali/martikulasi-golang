package controller

import (
	"../constants"
	"../services"
	"../helpers"
	"../objects"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

type UrlShrotenerController struct {
	urlShortenerService services.UrlShortenerService
	errorHelper helpers.ErrorHelper
}

func UrlShrotenerControllerHandler(router *gin.Engine) {
	handler := &UrlShrotenerController{
		urlShortenerService: services.UrlShortenerServiceHandler(),
		errorHelper: helpers.ErrorHelperHandler(),
	}
	group := router.Group("shorten")
	// test url
	// group.GET("/1", handler.GetById)
	// real url
	group.POST("", handler.ShortUrl)

}

func (handler *UrlShrotenerController) ShortUrl(context *gin.Context) {

	if err := context.ShouldBindWith(&objects.URLRequestShortRequest, binding.Query); err != nil {
		handler.errorHelper.HTTPResponseError(context, err, constants.RequestParameterInvalid)
	}

	var reqData URLRequestShortRequest
    context.BindJSON(&reqData)

	result, err := handler.urlShortenerService.MakeShortUrl(reqData.RealUrl)
	context.JSON(http.StatusOK, result)

}