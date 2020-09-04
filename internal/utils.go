package internal

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
)

func IsUp(url string) error {
	_, err := http.Get(url)
	if err != nil {
		return err
	}
	return nil
}

func StartTestServer(filePath string) *httptest.Server {

	mux := http.NewServeMux()

	job, err := ioutil.ReadFile(filePath)

	if err != nil {
		log.Fatal(err)
	}

	mux.HandleFunc("/job", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write(job)
	})

	return httptest.NewServer(mux)
}
