package objects

type URLRequestShortRequest struct {
	Url string `json:"url" binding:"required"`
}

type URLRequestShortResponse struct {
	ShortedUrl string `json:"shortcode"`
}