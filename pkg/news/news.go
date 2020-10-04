package news

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Tak1za/go-news/config"
)

type Result struct {
	Articles []ResultArticles `json:"articles"`
}

type ResultArticles struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Url     string `json:"url"`
}

func GetTopNews(config *config.Config) []ResultArticles {
	resp, err := http.Get(config.UrlConfig.TopHeadlines + "?country=in&apiKey=" + config.ApiKey)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	var result Result

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Fatalln(err)
	}

	return result.Articles[:5]
}
