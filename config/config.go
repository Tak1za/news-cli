package config

import "os"

type Config struct {
	UrlConfig *UrlOptions
	ApiKey    string
}

func defaultConfig() *Config {
	return &Config{
		UrlConfig: &UrlOptions{
			TopHeadlines: "https://newsapi.org/v2/top-headlines",
			Everything:   "https://newsapi.org/v2/everything",
		},
		ApiKey: "",
	}
}

func New() *Config {
	config := defaultConfig()

	if apiKey, ok := os.LookupEnv("NEWS_API_KEY"); ok {
		config.ApiKey = apiKey
	}

	return config
}
