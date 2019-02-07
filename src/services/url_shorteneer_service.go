package services
import (
	"../helpers"
	"../objects"
)

type UrlShortenerService struct {
	// stringUtil helpers.StringUtil
}

func UrlShortenerServiceHandler() (UrlShortenerService) {
	return UrlShortenerService{}
}

func (service *UrlShortenerService) MakeShortUrl(requestData objects.URLRequestShortRequest) (objects.URLRequestShortResponse, error){
	// validUrl := helpers.UrlValidator(requestData.RealUrl)
	// if false == validUrl {
	// 	return objects.URLRequestShortRequest{}, error
	// }
	shortCode := helpers.BuildRandomString(6)
	result := objects.URLRequestShortResponse{}
	result.ShortedUrl = shortCode
	return result, nil
}