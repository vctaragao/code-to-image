package create

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"text/template"
)

const (
	TEMPLATES_FOLDER = "template"
	LAYOUTS_FOLDER   = "layout"
)

type content struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type config struct {
	LayoutId string `json:"layout_id"`
}

type layout struct {
	Body   string
	Header string
	Style  string
}

var currentFolder string

func Execute(dto *InputDto) error {
	cont, err := getContentFromFile(dto.TextFile)
	if err != nil {
		return fmt.Errorf("error on getting content of text file: %v", err)
	}

	currentFolder, err = os.Getwd()
	if err != nil {
		fmt.Println("error on getting current folder: ", err)
		return err
	}

	_, err = buildFromTemplate(dto, cont)
	if err != nil {
		fmt.Println("error on building from template: ", err)
		return err
	}

	// uri := createPngFromHtml(buffer.String())
	// downloadAndSaveFile(uri)

	return nil
}

func getContentFromFile(file string) (*content, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return &content{}, fmt.Errorf("err on reading context file: %v", err)
	}

	var cont *content
	if err := json.Unmarshal(data, &cont); err != nil {
		fmt.Println("error unmarshaling the content: ", err)
		return &content{}, err
	}

	return cont, nil
}

func buildFromTemplate(dto *InputDto, cont *content) (string, error) {
	body, err := fillTemplateBody(dto.TemplateId, cont)
	if err != nil {
		fmt.Println("error on filling template body: ", err)
		return "", err
	}

	post, err := fillLayout(dto.TemplateId, body)
	if err != nil {
		fmt.Println("error on filling layout: ", err)
		return "", err
	}

	if err = os.WriteFile("result/"+dto.OutputId, post, 0644); err != nil {
		fmt.Println("error on creating/writing to output file: ", err)
		return "", err
	}

	return dto.OutputId, nil
}

func fillTemplateBody(templateFolder string, cont *content) ([]byte, error) {
	buffer := bytes.NewBuffer([]byte{})
	templateBodyPath := getTemplatePath(templateFolder, "body.html")

	if err := template.Must(template.New("body.html").ParseFiles(templateBodyPath)).Execute(buffer, cont); err != nil {
		fmt.Println("error on filling template body: ", err)
		return []byte{}, err
	}

	return buffer.Bytes(), nil
}

func fillLayout(templateFolder string, body []byte) ([]byte, error) {
	conf, err := getTemplateConfig(templateFolder, currentFolder)
	if err != nil {
		fmt.Println("error on getting template config: ", err)
		return []byte{}, err
	}

	header, err := getFileFromTemplate(templateFolder, "header.html")
	if err != nil {
		fmt.Println("error on getting header.html: ", err)
		return []byte{}, err
	}

	style, err := getFileFromTemplate(templateFolder, "style.css")
	if err != nil {
		fmt.Println("error on getting style.css: ", err)
		return []byte{}, err
	}

	lay := &layout{
		Body:   string(body),
		Header: string(header),
		Style:  string(style),
	}

	layoutPath := getLayoutPath(conf.LayoutId)

	buffer := bytes.NewBuffer([]byte{})
	if err = template.Must(template.New(conf.LayoutId).ParseFiles(layoutPath)).Execute(buffer, lay); err != nil {
		fmt.Println("error on filling layout: ", err)
		return []byte{}, err
	}

	return buffer.Bytes(), nil
}

func getTemplateConfig(templateFolder string, currentFolder string) (*config, error) {
	data, err := getFileFromTemplate(templateFolder, "config.json")
	if err != nil {
		fmt.Println("error on getting getting template config: ", err)
		return &config{}, err
	}

	var conf *config
	json.Unmarshal(data, &conf)

	return conf, err
}

func getTemplatePath(templateFolder, file string) string {
	return fmt.Sprintf("%s\\%s\\%s\\%s", currentFolder, TEMPLATES_FOLDER, templateFolder, file)
}

func getLayoutPath(layout string) string {
	return fmt.Sprintf("%s\\%s\\%s\\%s", currentFolder, TEMPLATES_FOLDER, LAYOUTS_FOLDER, layout)
}

func getFileFromTemplate(templateFolder, file string) ([]byte, error) {
	filePath := getTemplatePath(templateFolder, file)

	f, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("error on getting %s/%s: %v", templateFolder, file, err)
		return []byte{}, err
	}

	return f, nil
}
