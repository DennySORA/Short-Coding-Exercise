package common

type ShortParameter struct {
	Url     string `json:"url"`
	Expires string `json:"expireAt"`
}

type ShortResponse struct {
	ID       string `json:"id"`
	ShortUrl string `json:"shortUrl"`
}
