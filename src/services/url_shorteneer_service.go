package services
import (
	"../helpers"
	"../objects"
)

func MakeShortUrl(requestData objects.URLRequestShortRequest) (objects.URLRequestShortResponse){
	shortCode := helpers.BuildRandomString(6)
	result := objects.URLRequestShortResponse{}
	result.ShortCode = shortCode
	return result
}