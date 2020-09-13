package internal

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
)

// IsUp checks whether a website is accessible, if it is not, returns an error
func IsUp(url string) error {
	_, err := http.Get(url)
	if err != nil {
		return err
	}
	return nil
}

// StartTestServer starts a test server to serve a website page. Should only be used when testing
func StartTestServer(filePath string) *httptest.Server {

	mux := http.NewServeMux()

	job, err := ioutil.ReadFile(filePath)

	if err != nil {
		log.Fatal(err)
	}

	mux.HandleFunc("/job", func(w http.ResponseWriter, r *http.Request) {
		w.Write(job)
	})

	return httptest.NewServer(mux)
}
