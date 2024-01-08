package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type Cloudflare = string

var baseUrl Cloudflare = "https://api.cloudflare.com/client/v4"

func main() {
	verifyURL := baseUrl + "/user/tokens/verify"
	req, err := http.NewRequest(http.MethodGet, verifyURL, nil)

	tokenAuthorization := os.Getenv("CLOUDFLARE_TOKEN")

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", tokenAuthorization))

	if err != nil {
		fmt.Println("Failed to create instance of http request ", err)
		return
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println("Failed making http request ", err)
		return
	}

	body := res.Body
	readBody, err := io.ReadAll(body)

	if err != nil {
		fmt.Println("Failed to read response body ", err)
		return
	}

	fmt.Println("Result:", string(readBody))
}
