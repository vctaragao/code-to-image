package entity

type Layout struct {
	Body   []byte
	Header []byte
	Style  []byte
}

func NewLayout(body, header, style []byte) *Layout {
	return &Layout{
		Body:   body,
		Header: header,
		Style:  style,
	}
}
