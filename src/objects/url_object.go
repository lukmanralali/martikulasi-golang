package objects

type URLRequestShortResponse struct {
	ShortCode string `json:"shortcode" binding:"required"`
}

type URLRequestShortRequest struct {
	URLRequestShortResponse
	Url string `json:"url" binding:"required"`
}
