package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRoot(t *testing.T) {
	expectHeader := "application/json"
	expectURL := "/test"
	r, _ := http.NewRequest("GET", expectURL, nil)
	w := httptest.NewRecorder()
	root(w, r)
	returnedHeader := w.Header().Get("Content-Type")
	if returnedHeader != expectHeader {
		t.Errorf("App return wrong header %v. Not '%v'", returnedHeader, expectHeader)
	}

	m := make(map[string]string)
	err := json.Unmarshal(w.Body.Bytes(), &m)
	if err != nil {
		t.Errorf("Json don't valid")
	}
	if m["url"] != expectURL {
		t.Errorf("Wron data in returned json. Expect: %v., Got: %v", expectURL, m["url"])
	}
}
