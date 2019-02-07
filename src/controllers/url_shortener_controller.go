package controllers

import (
	"../constants"
	"../services"
	"../helpers"
	"../objects"
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)

type UrlShrotenerController struct {
}

func UrlShrotenerControllerHandler(router *gin.Engine) {
	group := router.Group("shorten")
	// test url
	// group.GET("/1", handler.GetById)
	// real url
	group.POST("", shortUrl)

}

func shortUrl(context *gin.Context) {
	reqData := objects.URLRequestShortRequest{}
	if err := context.ShouldBindJSON(&reqData); err != nil {
		fmt.Println("not valid data")
		helpers.HTTPResponseError2(context, err, constants.RequestParameterInvalid)
		return
	}
	fmt.Println("Validation Success!!")
	result := services.MakeShortUrl(reqData)
	context.JSON(http.StatusOK, result)
}