package gocors

import (
	"net/http"
	"strings"
)

//cors 跨域
type cors struct {
	nextHandler  http.Handler
	allowOrigin  string
	allowMethods string
	allowHeaders string
}

func NewCors(handler http.Handler, opts ...Option) http.Handler {
	o := &options{
		origin:  "*",
		methods: []string{"*"},
		headers: []string{"*"},
	}
	for _, opt := range opts {
		opt.apply(o)
	}
	return &cors{
		nextHandler:  handler,
		allowOrigin:  o.origin,
		allowMethods: strings.Join(o.methods, ", "),
		allowHeaders: strings.Join(o.headers, ", "),
	}
}

func (c *cors) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", c.allowOrigin)
	w.Header().Set("Access-Control-Allow-Methods", c.allowMethods)
	w.Header().Set("Access-Control-Allow-Headers", c.allowHeaders)
	if r.Method == http.MethodOptions && r.Header.Get("Access-Control-Request-Method") != "" {
		w.WriteHeader(http.StatusOK)
		return
	}
	c.nextHandler.ServeHTTP(w, r)
}
