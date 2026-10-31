package systemone

import "fmt"

// Response is the decoded result of a successful [SystemOne.Post].
type Response struct {
	// Model is the model that produced the answers.
	Model string `json:"model"`
	// Answers holds one entry per question key. Each is a map with a
	// "type" field plus "noul", "choice" or "score" and "probabilities",
	// depending on the question type.
	Answers map[string]any `json:"answers"`
	// Usage reports token usage for the request.
	Usage Usage `json:"usage"`
}

type envelope struct {
	Result  Response `json:"result"`
	Success bool     `json:"success"`
}

// Usage reports token consumption for a request.
type Usage struct {
	InputTokens  int `json:"input_tokens"`
	OutputTokens int `json:"output_tokens"`
}

// ResponseError is the error body returned by the API.
type ResponseError struct {
	Detail any `json:"detail"`
}

// Error implements the error interface.
func (e *ResponseError) Error() string {
	return fmt.Sprint(e.Detail)
}
