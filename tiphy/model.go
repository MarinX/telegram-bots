package main

type GiphyResponse struct {
	Data []GiphyData `json:"data"`
}

type GiphyData struct {
	Url string `json:"url"`
}
