package services
import (
	"../helpers"
	"../objects"
	"../repositories"
	"../models"
	"time"
)

func MakeShortUrl(requestData objects.URLRequestShortRequest) (objects.URLRequestShortResponse){
	shortCode := helpers.BuildRandomString(6)
	if requestData.ShortCode != ""{
		shortcodeDB := repositories.GetByShortcode(requestData.ShortCode)
		if "" != shortcodeDB.Shortcode {
			shortCode = "Already in Used"
		}
	}
	result := objects.URLRequestShortResponse{}
	result.ShortCode = shortCode
	
	requestData.ShortCode = shortCode
	if shortCode != "Already in Used" {
		repositories.CreateShortcode(requestData)
	}
	return result
}

func GetUrlShortUrl(shortCode string) (models.UrlShortCode){
	shortCodeDB := repositories.GetByShortcode(shortCode)
	shortCodeDB.LastUsedAt = time.Now()
	shortCodeDB.Counter++
	repositories.UpdateShortcodeData(shortCodeDB)
	return shortCodeDB
}

func GetUrlShortUrlStat(shortCode string) (objects.ShortedUrlStatResponse){
	shortcodeDB := repositories.GetByShortcode(shortCode)
	if "" == shortcodeDB.Shortcode {
		return objects.ShortedUrlStatResponse{}
	}
	result := objects.ShortedUrlStatResponse{}
	result.StartDate = shortcodeDB.PublishAt
	result.LastSeenDate = shortcodeDB.LastUsedAt
	result.RedirectCount = shortcodeDB.Counter
	return result
}