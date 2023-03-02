package jsonapi

type Error struct {
	Status int    `json:"status"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

func NewError(status int, title, detail string) Error {
	return Error{
		Status: status,
		Title:  title,
		Detail: detail,
	}
}
