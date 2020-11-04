package main

import (
	"net/url"
	"strings"
)

type RunConfig struct {
	ApiKey     string
	ApiSecret  string
	BaseUrl    string
	UseTestNet bool
}

func NewRunConfig(apiKey string, apiSecret string, hostUsed *string) (*RunConfig, error) {
	config := &RunConfig{
		ApiKey:     apiKey,
		ApiSecret:  apiSecret,
		UseTestNet: false,
		BaseUrl:    *hostUsed,
	}
	if hostUsed == nil || *hostUsed == "" {
		config.BaseUrl = "https://api.gateio.ws/api/v4"
	}
	if !strings.HasPrefix(config.BaseUrl, "http") {
		config.BaseUrl = "https://" + config.BaseUrl
	}
	if !strings.HasSuffix(config.BaseUrl, "/api/v4") {
		config.BaseUrl += "/api/v4"
	}
	parsedUrl, err := url.Parse(config.BaseUrl)
	if err != nil {
		return nil, err
	}
	if parsedUrl.Host == "fx-api-testnet.gateio.ws" {
		config.UseTestNet = true
	}
	return config, nil
}
