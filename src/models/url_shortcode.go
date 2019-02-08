package models

import "github.com/jinzhu/gorm"
import "time"

type UrlShortCode struct {
	gorm.Model
    Uri string `gorm:"size:255"`
    Shortcode string `gorm:"unique;unique_index;not null;size:6"`
    Counter int `gorm:"default:0;not null"`
    PublishAt time.Time
    LastUsedAt time.Time
}

func (UrlShortCode) TableName() string {
	return "rl_url_shortcodes"
}
