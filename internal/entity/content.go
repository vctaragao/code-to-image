package entity

type Content struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

func NewContent(name, code string) *Content {
	return &Content{
		Name: name,
		Code: code,
	}
}
