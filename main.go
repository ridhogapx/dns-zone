package main

import (
	"dns-zone/cloudflare"
	"dns-zone/config"
)

func main() {
	initConfig := config.CloudflareConfig{
		URL:           config.BaseURL,
		Authorization: config.Authorization,
		Token:         config.Token,
	}

	cloudflareAPI := cloudflare.RequestCloudflareImpl{
		Config: initConfig,
	}

	cloudflareAPI.Verify()

	cloudflareAPI.ListZones()
}
