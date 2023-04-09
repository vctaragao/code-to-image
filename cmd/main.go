package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"text/template"
	"time"
)

const (
	API_USER = ""
	API_KEY  = ""
)

type data struct {
	Text string
}

type response struct {
	Url string `json:"url"`
}

type request struct {
	method      string
	uri         string
	queryParams map[string]string
	body        []byte
}

func main() {
	var d *data
	for i, arg := range os.Args[1:] {
		if i == 0 {
			d.Text = arg
		}
	}

	buffer := buildFromTemplate(d)
	uri := createPngFromHtml(buffer.String())
	downloadAndSaveFile(uri)
}

func buildFromTemplate(d *data) *bytes.Buffer {
	buffer := bytes.NewBuffer([]byte{})
	t := template.Must(template.New("test.html").ParseFiles("/home/victor.aragao/personal/code-to-image/template/test.html"))
	err := t.Execute(buffer, d)
	if err != nil {
		fmt.Println("error on executing the template", err)
	}

	return buffer
}

func createPngFromHtml(html string) string {
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
	uri = uri + "?" + values.Encode()

	return uri
}
