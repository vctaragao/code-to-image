package entity

type Draft struct {
	Id string
}

func NewDraft(id string) *Draft {
	return &Draft{
		Id: id,
	}
}
