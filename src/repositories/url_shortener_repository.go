package repositories

import (
	"../database"
	"../models"
	"../objects"
	"time"
)

type UrlShortRepository struct {
}

type UrlShortRepositoryInterface interface {
	CreateShortcode(requestData objects.URLRequestShortRequest)
	GetByShortcode(shortcode string) models.UrlShortCode
	UpdateShortcodeData(urlShortCode models.UrlShortCode)
}

func (repository *UrlShortRepository) CreateShortcode(requestData objects.URLRequestShortRequest) {
	now := time.Now()
	urlShortCode := models.UrlShortCode{
		Uri:        requestData.Url,
		Shortcode:  requestData.ShortCode,
		PublishAt:  now,
		LastUsedAt: now,
	}
	db := database.GetConnection()
	db.Create(&urlShortCode)
}

func (repository *UrlShortRepository) GetByShortcode(shortcode string) models.UrlShortCode {
	urlShortCode := models.UrlShortCode{}
	db := database.GetConnection()
	db.Where(&models.UrlShortCode{Shortcode: shortcode}).First(&urlShortCode)
	return urlShortCode
}

func (repository *UrlShortRepository) UpdateShortcodeData(urlShortCode models.UrlShortCode) {
	urlShortCodeModel := models.UrlShortCode{}
	db := database.GetConnection()
	db.Model(&urlShortCodeModel).Where("shortcode = ?", urlShortCode.Shortcode).Update(urlShortCode)
}
