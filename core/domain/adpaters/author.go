package adapters

type AuthorGetByIdResponse struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Email string  `json:"email"`
	Books []Books `json:"books"`
}

type Books struct {
}
