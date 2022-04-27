package apiModel

type Blog struct {
	PublishedDate string `json:"published_date"`
	Title         string `json:"title"`
	Content       string `json:"content"`
	Public        bool   `json:"public"`
	Groups        string `json:"groups"`
}
