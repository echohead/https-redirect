package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func responseFor(method string, url string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	r, _ := http.NewRequest(method, url, nil)
	redirectHandler(w, r)
	return w
}

func shouldRedirect(t *testing.T, method string, url string, newUrl string) {
	w := responseFor(method, url)
	if w.Code != http.StatusMovedPermanently {
		t.Errorf("%v %v should give a 301 response.", method, url)
	}
	if w.HeaderMap["Location"][0] != newUrl {
		t.Errorf("%v should redirect to %v", url, newUrl)
	}
}

func TestRedirects(t *testing.T) {
	shouldRedirect(t, "GET", "foo.com", "https://foo.com")
	shouldRedirect(t, "POST", "foo.com", "https://foo.com")
	shouldRedirect(t, "GET", "http://foo.com/blah", "https://foo.com/blah")
	shouldRedirect(t, "GET", "foo.com/bar?a=b", "https://foo.com/bar")
}

