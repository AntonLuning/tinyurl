package models

type CreateShortenURLResponse struct {
	Original string `json:"original"`
	Shorten  string `json:"shorten"`
}

type GetOriginalURLResponse struct {
	Shorten  string `json:"shorten"`
	Original string `json:"original"`
}
