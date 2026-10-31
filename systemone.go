package systemone

import (
	"context"
	"os"

	"github.com/MarcelArt/constructor"
	"resty.dev/v3"
)

type SystemOne struct {
	Model   string
	BaseURL string
	APIKey  string
	client  *resty.Client
}

type Option = constructor.OptionFunc[SystemOne]

func New(opts ...Option) *SystemOne {
	s1 := constructor.New(&SystemOne{
		Model:   os.Getenv("SYSTEMONE_MODEL"),
		BaseURL: os.Getenv("SYSTEMONE_BASE_URL"),
		APIKey:  os.Getenv("SYSTEMONE_API_KEY"),
	}, opts...)

	s1.client = resty.New().SetAuthToken(s1.APIKey)

	return s1
}

func WithModel(model string) Option {
	return func(s1 *SystemOne) {
		s1.Model = model
	}
}

func WithBaseURL(baseURL string) Option {
	return func(s1 *SystemOne) {
		s1.BaseURL = baseURL
	}
}

func WithAPIKey(apiKey string) Option {
	return func(s1 *SystemOne) {
		s1.APIKey = apiKey
	}
}

func (s1 *SystemOne) Post(ctx context.Context, body Body) (*Response, *ResponseError) {
	res, err := s1.client.R().
		SetBody(body).
		SetResult(&Response{}).
		SetResultError(&ResponseError{}).
		WithContext(ctx).
		SetDebug(true).
		Post(s1.BaseURL)

	if err != nil {
		return nil, res.ResultError().(*ResponseError)
	}

	return res.Result().(*Response), nil
}

func (s1 *SystemOne) Close() error {
	return s1.client.Close()
}
