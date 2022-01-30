package entities

type Upload struct {
	Id   *int    `json:"id"`
	Path *string `json:"path"`
	Name *string `json:"name"`
	Type *string `json:"type"`
}
