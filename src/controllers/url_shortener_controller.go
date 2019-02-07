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

func UrlShrotenerControllerHandler(router *gin.Engine) {
	group := router.Group("shorten")
	// test url
	// group.GET("/1", handler.GetById)
	// real url
	group.POST("", shortUrl)

}

func shortUrl(context *gin.Context) {
	reqData := objects.URLRequestShortRequest{}
	
	// validate payload
	if err := context.ShouldBindJSON(&reqData); err != nil {
		fmt.Println("not valid data")
		helpers.HTTPResponseError2(context, err, constants.RequestParameterInvalid)
		return
	}
	
	// validate url should be valid
	if !helpers.UrlValidator(reqData.Url) {
		context.JSON(http.StatusUnprocessableEntity, nil)
		return
	}

	// validate shortcode if match with our regex requirement
	if !helpers.ValidatorShortCode(reqData.ShortCode) && reqData.ShortCode != "" {
		context.JSON(http.StatusUnprocessableEntity, nil)
		return
	}
	fmt.Println("Validation Success!!")
	
	result := services.MakeShortUrl(reqData)
	context.JSON(http.StatusOK, result)
}