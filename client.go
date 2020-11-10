package jsonrpc

import (
	"github.com/asaskevich/govalidator"
	"github.com/pkg/errors"
)

type (
	Client struct {
		endpoint string
	}

	Request struct {
		Version string      `json:"jsonrpc"`
		Method  string      `json:"method"`
		Params  interface{} `json:"params"`
		Id      uint64      `json:"id"`
	}
)

func NewClient(url string) (*Client, error) {
	if !govalidator.IsURL(url) {
		return nil, errors.WithMessage(errors.New("invalid http url"), "AddEndpoint")
	}

	return &Client{
		endpoint: url,
	}, nil
}

func (c *Client) Call(req *Request) (*Response, error) {
	resp := &Response{}

	return resp, nil
}
