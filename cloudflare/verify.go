package cloudflare

import (
	"dns-zone/config"
	"dns-zone/entity"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type RequestCloudflareImpl struct {
	Config config.CloudflareConfig
}

func (c *RequestCloudflareImpl) Verify() {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/user/tokens/verify", c.Config.URL), nil)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", c.Config.Authorization)

	if err != nil {
		fmt.Println("Failed to create instance of http", err)
		return
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println("Failed making http request ", err)
		return
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		fmt.Println("Failed to read body response ", err)
		return
	}

	var data entity.CloudflareEntity

	err = json.Unmarshal(body, &data)

	if err != nil {
		fmt.Println("Failed to unmarshal data ", err)
		return
	}

	fmt.Println("Message:", data.Messages[0].Message)
	fmt.Println("Success:", data.Success)
	fmt.Println("ID:", data.Result.ID)
}
