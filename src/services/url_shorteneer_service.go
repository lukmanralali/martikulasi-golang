package services
import "../helpers"

type UrlShortenerService struct {
	stringUtil helpers.StringUtil
}

func UrlShortenerServiceHandler() (UrlShortenerService) {
	return UrlShortenerService{}
}

func (service *UrlShortenerService) makeShortUrl(requestData object.URLRequestShortRequest) (object.URLRequestShortResponse, error){
	validUrl := stringUtil.urlValidator(requestData.RealUrl)
	if nil == validUrl {
		return objects.URLRequestShortRequest{}, err
	}
	shortCode := stringUtil.BuildRandomString(6)
	result := object.URLRequestShortResponse{}
	result.ShortedUrl = shortCode
	return result, nil
}