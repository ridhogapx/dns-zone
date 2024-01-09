package config

import "os"

type Cloudflare = string

var Token = os.Getenv("CLOUDFLARE_TOKEN")
var GlobalKey = os.Getenv("GLOBAL_KEY")
var Authorization Cloudflare = "Bearer " + Token
var BaseURL Cloudflare = "https://api.cloudflare.com/client/v4"

type CloudflareConfig struct {
	Authorization Cloudflare
	URL           Cloudflare
	Token         Cloudflare
}
