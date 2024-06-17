package main

import (
	"bytes"
	"net/http"
)

type nopReader struct{}

func (nopReader) Read(p []byte) (n int, err error) {
	return 0, nil
}

type responseWriter struct {
	header http.Header
	buf    bytes.Buffer
	status int
}

func newResponseWriter() *responseWriter {
	return &responseWriter{header: http.Header{}}
}

func (r *responseWriter) Header() http.Header {
	return r.header
}

func (r *responseWriter) Write(bytes []byte) (int, error) {
	return r.buf.Write(bytes)
}

func (r *responseWriter) WriteHeader(statusCode int) {
	r.status = statusCode
}
