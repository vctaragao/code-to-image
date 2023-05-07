package entity

type Template struct {
	Id     string
	Config Config
}

func NewTemplate(id string, conf *Config) *Template {
	return &Template{
		Id:     id,
		Config: *conf,
	}
}
