package models

import "github.com/jinzhu/gorm"

type UrlShortCode struct {
	gorm.Model
    Uri string `gorm:"size:255"`
    Shortcode string `gorm:"unique;unique_index;not null;size:6"`
}

func (UrlShortCode) TableName() string {
	return "rl_url_shortcodes"
}
