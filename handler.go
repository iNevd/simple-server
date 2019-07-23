package simple_server

import "fmt"

type serverHandler struct {
	srv *Server
}

func (sh serverHandler) Serve(rw ResponseWriter, req *Request) {
	handler := sh.srv.Handler
	if handler == nil {
		panic("simple_server: invalid handler")
	}
	handler.Serve(rw, req)
}

type Handler interface {
	Serve(ResponseWriter, *Request)
}

type HandlerFunc func(ResponseWriter, *Request)

// Serve calls f(w, r).
func (f HandlerFunc) Serve(w ResponseWriter, r *Request) {
	f(w, r)
}

// HelloWorld replies to the request with handler not found error.
func HelloWorld(w ResponseWriter, r *Request) {
	fmt.Fprintf(w, "simple_server: hello world!")
}

// HelloWorldHandler returns a simple request handler
// that replies to each request with a ``simple_server: hello world!'' reply.
func HelloWorldHandler() Handler { return HandlerFunc(HelloWorld) }
