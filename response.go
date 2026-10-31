package systemone

type Response struct {
	Model   string         `json:"model"`
	Answers map[string]any `json:"answers"`
	Usage   Usage          `json:"usage"`
}

type envelope struct {
	Result  Response `json:"result"`
	Success bool     `json:"success"`
}

type Usage struct {
	InputTokens  int `json:"input_tokens"`
	OutputTokens int `json:"output_tokens"`
}

type ResponseError struct {
	Detail any `json:"detail"`
}
