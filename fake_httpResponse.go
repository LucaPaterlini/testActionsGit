package helpers

/*
Copyright (c) 2019, Sportable Technologies. All rights reserved.
author: Luca Paterlini <lucaSportable>
date: 2019-19-07
module: testing.go
last modified by:  Luca Paterlini <lucaSportable>
last modified time: 2019-19-07
description:
*/

import (
	"bytes"
	"net/http"
	"testing"
)

// FakeResponseWriter implements a struct that implements the interface
// http.ResponseWriter and its used for testing purpose.
type FakeResponseWriter struct {
	status  int
	headers http.Header
	body    []byte
}

// NewFakeResponseWriter instantiate and return a FakeResponseWriter.
func NewFakeResponseWriter() *FakeResponseWriter {
	return &FakeResponseWriter{
		headers: make(http.Header),
	}
}

// WriteHeader implements the WriteHeader of http.ResponseWriter.
func (r *FakeResponseWriter) WriteStatusCode(status int) {
	r.status = status
}

// WriteHeader implements the WriteHeader of http.ResponseWriter.
func (r *FakeResponseWriter) ReadStatusCode(status int) {
	r.status = status
}

// WriteHeader implements the WriteHeader of http.ResponseWriter.
func (r *FakeResponseWriter) WriteHeader(status int) {
	r.status = status
}

// ReadHeader is not part of the interface readwrite but is useful
// to retrieve what has been written inside the header during testing.
func (r *FakeResponseWriter) Header() http.Header {
	return r.headers
}

// Write implements the Write of http.ResponseWriter.
func (r *FakeResponseWriter) Write(body []byte) (int, error) {
	r.body = body
	return len(body), nil
}

// ReadBody is not part of the interface readwrite but is useful
// to retrieve what has been written inside the body during testing.
func (r *FakeResponseWriter) Read() []byte {
	return r.body
}

// InitHTTPRW generate a fake responsewriter and request having as input
// a byte array with the content of the body of the request.
func InitHTTPRW(t *testing.T, jsonBody []byte) (w http.ResponseWriter, r *http.Request) {
	w = NewFakeResponseWriter()
	r, err := http.NewRequest("POST", "/testpath", bytes.NewReader(jsonBody))
	if err != nil {
		t.Fatal(err)
	}
	return
}
