package objects

type URLRequestShortRequest struct {
	RealUrl  string `json:"url"`
}

type URLRequestShortResponse struct {
	ShortedUrl  string `json:"shortcode"`
}