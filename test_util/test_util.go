package test_util

import (
	"net/http"
	"strings"
)

// A mock http.ResponseWriter
type MockResponseWriter struct {
	header http.Header
	builder strings.Builder
	status int
}

func (m *MockResponseWriter) GetWrittenText() string {
	return m.builder.String()
}

func (m *MockResponseWriter) Header() http.Header {
	return m.header
}

func (m *MockResponseWriter) Write(p []byte) (int, error) {
	return m.builder.Write(p)
}

func (m *MockResponseWriter) WriteHeader(statusCode int) {
	m.status = statusCode
}

// Create a new mock http.ResponseWriter
func NewMockResponseWriter() MockResponseWriter {
	return MockResponseWriter{
		header:  make(http.Header),
		builder: strings.Builder{},
		status:  0,
	}
}



