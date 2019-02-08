package repositories

import (
	"../models"
	"../database"
	"../objects"
	"time"
)

func CreateShortcode(requestData objects.URLRequestShortRequest) {
	now := time.Now()
	urlShortCode := models.UrlShortCode{
		Uri: requestData.Url,
    	Shortcode: requestData.ShortCode,
    	PublishAt: now,
    	LastUsedAt: now,
	}
	db := database.GetConnection()
	db.Create(&urlShortCode)
}

func GetByShortcode(shortcode string) (models.UrlShortCode) {
	urlShortCode := models.UrlShortCode{}
	db := database.GetConnection()
	db.Where(&models.UrlShortCode{Shortcode: shortcode}).First(&urlShortCode)
	return urlShortCode
}

func UpdateShortcodeData(urlShortCode models.UrlShortCode) {
	urlShortCodeModel := models.UrlShortCode{}
	db := database.GetConnection()
	db.Model(&urlShortCodeModel).Where("shortcode = ?", urlShortCode.Shortcode).Update(urlShortCode)
}