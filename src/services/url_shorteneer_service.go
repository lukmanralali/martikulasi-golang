package services
import (
	"../helpers"
	"../objects"
	"fmt"
)

type UrlShortenerService struct {
}

func UrlShortenerServiceHandler() {
}

func MakeShortUrl(requestData objects.URLRequestShortRequest) (objects.URLRequestShortResponse){
	if !helpers.UrlValidator(requestData.Url) {

		fmt.Println("not validUrl")
	}
	// if false == validUrl {
	// 	return objects.URLRequestShortRequest{}, error
	// }
	shortCode := helpers.BuildRandomString(6)
	result := objects.URLRequestShortResponse{}
	result.ShortedUrl = shortCode
	return result
}