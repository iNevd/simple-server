package simple_server

import (
	"io"
)

const (
	HeaderPayloadLenUnlimited = -1
)

type Header interface {
	Write(io.Writer) (int, error)
	Read(io.Reader) (int, error)
	PayLoadLen() int
	NewInstance() Header
}

type DefaultHeader struct{}

func (h *DefaultHeader) Write(i io.Writer) (int, error) {
	return 0, nil
}

func (h *DefaultHeader) Read(i io.Reader) (int, error) {
	return 0, nil
}

func (h *DefaultHeader) PayLoadLen() int {
	return HeaderPayloadLenUnlimited
}

func (h *DefaultHeader) NewInstance() Header {
	return &DefaultHeader{}
}
