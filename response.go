package simple_server

import (
	"context"
)

// A response represents the server side of a response.
type response struct {
	conn      *conn
	req       *Request           // request for this response
	cancelCtx context.CancelFunc // when Serve exits
}

func (w *response) Write(p []byte) (n int, err error) {
	n, err = w.conn.bufw.Write(p)
	if err != nil {
		w.conn.rwc.Close()
	}
	return
}

type ResponseWriter interface {
	//WriteHeader(statusCode int)
	Write([]byte) (int, error)
}

func (w *response) finishRequest() {
	w.conn.bufw.Flush()
}
