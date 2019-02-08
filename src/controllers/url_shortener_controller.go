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
	// router.GET(":shortcode", getUrlShortcode)

	group := router.Group("shorten")
	group.POST("", shortUrl)
	group.GET(":shortcode", getUrlShortcode)
	group.GET(":shortcode/stats", getUrlShortcodeStat)

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
	if result.ShortCode == "Already in Used" {
		context.JSON(http.StatusConflict, result)
		return
	}
	context.JSON(http.StatusOK, result)
}

func getUrlShortcode(context *gin.Context) {
	reqData := context.Param("shortcode")
	fmt.Println(reqData)
	if "" == reqData {
		context.JSON(http.StatusBadRequest, nil)
		return
	}
	fmt.Println("Validation Success!!")
	result := services.GetUrlShortUrl(reqData)
	context.Header("Location", result.Uri)
	context.JSON(http.StatusFound, result)
}

func getUrlShortcodeStat(context *gin.Context) {
	reqData := context.Param("shortcode")
	fmt.Println(reqData)
	if "" == reqData {
		context.JSON(http.StatusBadRequest, nil)
		return
	}
	fmt.Println("Validation Success!!")
	result := services.GetUrlShortUrlStat(reqData)
	
	context.JSON(http.StatusFound, result)
}