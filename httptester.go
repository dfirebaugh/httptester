package httptester

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type TestCase struct {
	Pre  func()
	Post func(res *httptest.ResponseRecorder)
}

type HTTPTest struct {
	Method string
	URL    string
	Body   io.Reader
	Tests  []TestCase
}

func (tc TestCase) Execute(t *testing.T, res *httptest.ResponseRecorder) {
	if tc.Pre != nil {
		tc.Pre()
	}

	if res != nil {
		println(res.Body.String())
		println(res.Code)
	}

	if tc.Post != nil {
		tc.Post(res)
	}
}

func (h HTTPTest) Execute(t *testing.T, handler http.HandlerFunc) {
	println(h.URL)
	req, err := http.NewRequest(h.Method, h.URL, h.Body)
	if err != nil {
		t.Fatal(err)
	}

	res := HTTPRecorder(handler, req)

	for _, tc := range h.Tests {
		tc.Execute(t, res)
	}
}

func HTTPRecorder(next http.HandlerFunc, r *http.Request) *httptest.ResponseRecorder {
	res := httptest.NewRecorder()
	http.HandlerFunc(next).ServeHTTP(res, r)
	return res
}
