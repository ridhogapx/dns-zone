package config

import "os"

type Cloudflare = string

var Token = os.Getenv("CLOUDFLARE_TOKEN")
var Authorization Cloudflare = "Bearer " + Token
var BaseURL Cloudflare = "https://api.cloudflare.com/client/v4"

type CloudflareConfig struct {
	Authorization Cloudflare
	URL           Cloudflare
}
