package adapters

type LibraryGetByIdResponse struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Authors any    `json:"authors"`
}

type AuthorsResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type UpdateLibraryParams struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}
