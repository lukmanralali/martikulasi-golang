package services
import (
	"../helpers"
	"../objects"
	"../repositories"
	"../models"
	"time"
)

type UrlShortService struct {
	urlShortRepository repositories.UrlShortRepository
}

func  (service *UrlShortService) MakeShortUrl(requestData objects.URLRequestShortRequest) (objects.URLRequestShortResponse){
	shortCode := helpers.BuildRandomString(6)
	if requestData.ShortCode != ""{
		shortcodeDB := service.urlShortRepository.GetByShortcode(requestData.ShortCode)
		if "" != shortcodeDB.Shortcode {
			shortCode = "Already in Used"
		}
	}
	result := objects.URLRequestShortResponse{}
	result.ShortCode = shortCode
	
	requestData.ShortCode = shortCode
	if shortCode != "Already in Used" {
		service.urlShortRepository.CreateShortcode(requestData)
	}
	return result
}

func (service *UrlShortService) GetUrlShortUrl(shortCode string) (models.UrlShortCode){
	shortCodeDB := service.urlShortRepository.GetByShortcode(shortCode)
	shortCodeDB.LastUsedAt = time.Now()
	shortCodeDB.Counter++
	service.urlShortRepository.UpdateShortcodeData(shortCodeDB)
	return shortCodeDB
}

func (service *UrlShortService) GetUrlShortUrlStat(shortCode string) (objects.ShortedUrlStatResponse){
	shortcodeDB := service.urlShortRepository.GetByShortcode(shortCode)
	if "" == shortcodeDB.Shortcode {
		return objects.ShortedUrlStatResponse{}
	}
	result := objects.ShortedUrlStatResponse{}
	result.StartDate = shortcodeDB.PublishAt
	result.LastSeenDate = shortcodeDB.LastUsedAt
	result.RedirectCount = shortcodeDB.Counter
	return result
}