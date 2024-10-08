package models

import "time"

type Task struct {
	Id        string    `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	Completed string    `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Tasks struct {
	Tasks []Task
}

type ListTasksParams struct {
	Completed bool   `json:"completed"`
	Order     string `json:"order"`
	SortBy    string `json:"sort_by"`
}
