package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

const (
	API_USER = ""
	API_KEY  = ""
)

type response struct {
	Url string `json:"url"`
}

type request struct {
	method      string
	uri         string
	queryParams map[string]string
	body        []byte
}

type CreatePngFromHtml struct {
}

func Execute(html string) string {
	reqBody, err := json.Marshal(map[string]string{
		"html": html,
	})

	if err != nil {
		log.Fatalf("unable to marshal data: %s", err.Error())
	}

	req := &request{
		method: "POST",
		uri:    "https://hcti.io/v1",
		body:   reqBody,
	}

	resp := makeRequest(req)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("unable to read response body: %s", err.Error())
	}

	var r response
	json.Unmarshal(body, &r)

	fmt.Println("url:", r.Url)

	return r.Url
}

func downloadAndSaveFile(uri string) {
	req := &request{
		method: "GET",
		uri:    uri,
		queryParams: map[string]string{
			"dl": "1",
		},
	}

	resp := makeRequest(req)
	defer resp.Body.Close()

	file, err := os.Create("code.png")
	if err != nil {
		log.Fatalf("error creating flie: %s", err.Error())
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		log.Fatalf("error coping response body: %s", err.Error())
	}

	fmt.Println("Image saved to code.png")
}

func makeRequest(r *request) *http.Response {
	uri := addQueryParams(r.uri, r.queryParams)
	req, err := http.NewRequest(r.method, uri, bytes.NewReader(r.body))
	if err != nil {
		log.Fatalf("unable to create new request: %s", err.Error())
	}

	req.SetBasicAuth(API_USER, API_KEY)
	client := &http.Client{Timeout: time.Second * 10}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("request was unsuccessful: %s", err.Error())
	}

	return resp
}

func addQueryParams(uri string, queryParams map[string]string) string {
	if queryParams == nil {
		return uri
	}

	values := url.Values{}
	for key, param := range queryParams {
		values.Add(key, param)
	}
	uri += values.Encode()

	return uri
}
