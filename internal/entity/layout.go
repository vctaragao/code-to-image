package entity

type Layout struct {
	Body   string
	Header string
	Style  string
}

func NewLayout(body, header, style string) *Layout {
	return &Layout{
		Body:   body,
		Header: header,
		Style:  style,
	}
}
