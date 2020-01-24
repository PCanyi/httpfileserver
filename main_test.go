package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func BenchmarkServer(b *testing.B) {
	ts := httptest.NewServer(New("/", ".").Handle())
	defer ts.Close()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		resp, _ := http.Get(ts.URL + "/README.md")
		resp.Body.Close()
	}
}

func TestServer(t *testing.T) {
	ts := httptest.NewServer(New("/", ".").Handle())
	defer ts.Close()

	for i := 0; i < 2; i++ {
		res, err := http.Get(ts.URL + "/README.md")
		if err != nil {
			t.Errorf("%s", err)
		}
		greeting, _ := ioutil.ReadAll(res.Body)
		res.Body.Close()

		fmt.Printf("%s", greeting)
	}

}
