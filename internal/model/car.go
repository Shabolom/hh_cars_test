package model

type Car struct {
	Mark    string `json:"mark"`
	OwnerID string `json:"owner"`
	Model   string `json:"model"`
	RegNum  string `json:"regNum"`
	Year    int    `json:"year"`
}
