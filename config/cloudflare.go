package config

import "os"

var Token = os.Getenv("CLOUDFLARE_TOKEN")
