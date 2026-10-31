// Package systemone is a small client for SystemOne-style classification
// models such as Cloudflare Workers AI's clef-flash.
//
// You describe a piece of text (the state) and a set of named questions.
// The model answers each question with a probability, a choice, or a score.
//
// Configuration is read from the environment by [New]:
//
//	SYSTEMONE_MODEL     model name sent in the request body
//	SYSTEMONE_BASE_URL  full URL of the run endpoint
//	SYSTEMONE_API_KEY   bearer token
//
// Basic usage:
//
//	s1 := systemone.New()
//	defer s1.Close()
//
//	res, rerr := s1.Post(ctx, systemone.Body{
//		Model: s1.Model,
//		State: "Checkout has been failing for every customer.",
//		Questions: systemone.Questions{
//			"urgent": systemone.Noul("Is this urgent?", nil),
//		},
//	})
package systemone

import (
	"context"
	"os"

	"github.com/MarcelArt/constructor"
	"resty.dev/v3"
)

// SystemOne is the API client. Create one with [New] and release it with
// [SystemOne.Close].
type SystemOne struct {
	// Model is the model name, defaulting to $SYSTEMONE_MODEL.
	Model string
	// BaseURL is the endpoint requests are posted to, defaulting to
	// $SYSTEMONE_BASE_URL.
	BaseURL string
	// APIKey is sent as a bearer token, defaulting to $SYSTEMONE_API_KEY.
	APIKey string
	client *resty.Client
}

// Option customises a [SystemOne] client in [New].
type Option = constructor.OptionFunc[SystemOne]

// New returns a client configured from the SYSTEMONE_MODEL,
// SYSTEMONE_BASE_URL and SYSTEMONE_API_KEY environment variables. Any opts
// are applied afterwards and override the environment.
func New(opts ...Option) *SystemOne {
	s1 := constructor.New(&SystemOne{
		Model:   os.Getenv("SYSTEMONE_MODEL"),
		BaseURL: os.Getenv("SYSTEMONE_BASE_URL"),
		APIKey:  os.Getenv("SYSTEMONE_API_KEY"),
	}, opts...)

	s1.client = resty.New().SetAuthToken(s1.APIKey)

	return s1
}

// WithModel sets the model name.
func WithModel(model string) Option {
	return func(s1 *SystemOne) {
		s1.Model = model
	}
}

// WithBaseURL sets the endpoint requests are posted to.
func WithBaseURL(baseURL string) Option {
	return func(s1 *SystemOne) {
		s1.BaseURL = baseURL
	}
}

// WithAPIKey sets the bearer token.
func WithAPIKey(apiKey string) Option {
	return func(s1 *SystemOne) {
		s1.APIKey = apiKey
	}
}

// Post sends body to BaseURL and returns the decoded answers.
//
// The endpoint is expected to wrap its payload in a Cloudflare-style
// {"result": ...} envelope. On a transport error it returns a nil
// [Response] and the decoded [ResponseError].
func (s1 *SystemOne) Post(ctx context.Context, body Body) (*Response, *ResponseError) {
	res, err := s1.client.R().
		SetBody(body).
		SetResult(&envelope{}).
		SetResultError(&ResponseError{}).
		WithContext(ctx).
		// SetDebug(true).
		Post(s1.BaseURL)

	if err != nil {
		return nil, res.ResultError().(*ResponseError)
	}

	return &res.Result().(*envelope).Result, nil
}

// Close releases the underlying HTTP client.
func (s1 *SystemOne) Close() error {
	return s1.client.Close()
}
