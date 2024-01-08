package entity

type MessageResp struct {
	Code    int32       `json:"code"`
	Message string      `json:"message"`
	Type    interface{} `json:"type"`
}

type ResultResp struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

type CloudflareEntity struct {
	Result   ResultResp    `json:"result"`
	Success  bool          `json:"success"`
	Errors   []interface{} `json:"errors"`
	Messages []MessageResp `json:"messages"`
}
