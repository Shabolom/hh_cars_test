package model

type Car struct {
	Mark    string `json:"mark"`
	OwnerID string `json:"ownerID"`
	Model   string `json:"model"`
	RegNum  string `json:"regMum"`
	Year    int    `json:"year"`
}
