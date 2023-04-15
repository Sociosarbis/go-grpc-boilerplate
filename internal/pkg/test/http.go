package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
)

type Request struct {
	t       *testing.T
	Method  string
	Path    string
	Params  map[string]any
	Headers map[string]string
	Data    []byte
}

func NewRequest(t *testing.T, path string) *Request {
	t.Helper()
	return &Request{
		t:       t,
		Method:  "",
		Path:    path,
		Params:  make(map[string]any, 0),
		Headers: make(map[string]string, 0),
		Data:    make([]byte, 0),
	}
}

func (req *Request) Get() *Request {
	req.t.Helper()
	req.Method = http.MethodGet
	return req
}

func (req *Request) Post() *Request {
	req.t.Helper()
	req.Method = http.MethodPost
	return req
}

func (req *Request) Header(key string, value string) *Request {
	req.t.Helper()
	req.Headers[http.CanonicalHeaderKey(key)] = value
	return req
}

func (req *Request) JSON(data map[string]any) *Request {
	res, err := json.Marshal(data)
	require.NoError(req.t, err)
	req.Data = res
	return req
}

func (req *Request) Build() *http.Request {
	urlObj, err := url.Parse(req.Path)
	require.NoError(req.t, err)
	queryObj := urlObj.Query()
	for k, v := range req.Params {
		queryObj.Add(k, fmt.Sprintf("%s", v))
	}
	urlObj.RawQuery = queryObj.Encode()
	httpReq := httptest.NewRequest(req.Method, urlObj.String(), http.NoBody)
	if len(req.Data) != 0 {
		httpReq.Body = io.NopCloser(bytes.NewReader(req.Data))
	}
	for k, v := range req.Headers {
		httpReq.Header.Add(k, v)
	}
	return httpReq
}

func ReadJSONResponse[T any](t *testing.T, res *http.Response, out *T) *T {
	t.Helper()

	data, err := io.ReadAll(res.Body)

	require.NoError(t, err)

	err = json.Unmarshal(data, out)

	require.NoError(t, err)

	return out
}
