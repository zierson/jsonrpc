package jsonrpc

import (
	"github.com/asaskevich/govalidator"
	"github.com/pkg/errors"
)

type (
	Client struct {
		endpoints []string
	}

	Request struct {
		Version string      `json:"jsonrpc"`
		Method  string      `json:"method"`
		Params  interface{} `json:"params"`
		Id      uint64      `json:"id"`
	}
)

func NewClient() *Client {
	return &Client{}
}

func (c *Client) AddEndpoint(url string) error {
	if !govalidator.IsURL(url) {
		return errors.WithMessage(errors.New("invalid http url"), "AddEndpoint")
	}

	c.endpoints = append(c.endpoints, url)

	return nil
}

func (c *Client) Call(req *Request) (*Response, error) {
	resp := &Response{}

	return resp, nil
}
