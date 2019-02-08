package objects

import "time"

type URLRequestShortResponse struct {
	ShortCode string `json:"shortcode" binding:"required"`
}

type URLRequestShortRequest struct {
	URLRequestShortResponse
	Url string `json:"url" binding:"required"`
}

type ShortedUrlStatResponse struct {
	StartDate time.Time `json:"startDate"`
	LastSeenDate time.Time `json:"lastSeenDate"`
	RedirectCount int `json:"redirectCount"`
}
