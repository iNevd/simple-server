package simple_server

import (
	"context"
	"io"
)

type Request struct {
	ctx    context.Context
	Header Header
	Body   io.Reader
}
