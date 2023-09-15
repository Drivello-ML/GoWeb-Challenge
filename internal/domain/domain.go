package domain

type Task struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
}
