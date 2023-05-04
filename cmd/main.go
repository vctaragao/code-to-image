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
	TEMPLATES_FOLDER = "/template/"
)

type content struct{
	Name string
	Code string
}

type args struct {
	TemplateId string
	OutputId string
	TextFile string
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
	a := &args{}
	for i, arg := range os.Args[1:] {
		switch (i){
		case 0:
			fmt.Println(arg)
			a.TemplateId = arg
		case 1:
			a.OutputId = arg
		case 2:
			a.TextFile = arg
		}
	}

	fmt.Println(a)

	data, err := os.ReadFile(a.TextFile);
	if err != nil{
		fmt.Println("err on reading context file: ", err)
		return
	}
	fmt.Println(string(data))
	var cont *content
	err = json.Unmarshal(data, &cont)

	if err != nil{
		fmt.Println("error unmarshaling the content: ", err)
		return
	}

	_, err = buildFromTemplate(a, cont)
	if err != nil{
		fmt.Println("error on building from template: ", err)
		return
	}

	// uri := createPngFromHtml(buffer.String())
	// downloadAndSaveFile(uri)
}

func buildFromTemplate(a *args, cont *content) (string, error) {
	outputFile, err := os.Create("result/"+a.OutputId+".html")
	if err != nil {
		fmt.Println("error on creating output file: ", err)
		return "", err
	}
	defer outputFile.Close()

	currentFolder, err := os.Getwd()
	if err != nil{
		fmt.Println("error on getting current folder: ", err)
		return "", err
	}

	t := template.Must(template.New(a.TemplateId).ParseFiles(currentFolder+TEMPLATES_FOLDER+a.TemplateId))
	err = t.Execute(outputFile, cont)

	if err != nil {
		fmt.Println("error on executing the template: ", err)
		return "", err
	}

	return a.OutputId, nil
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
