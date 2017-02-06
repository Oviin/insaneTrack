package models

import (
	"time"
)

type TaskScope struct {
	Epic       Epic   `json:"epic"`
	ID         string `json:"id"`
	Decription string `json:"decription"`
}

type Task struct {
	ID uint `json:"id"`

	Status   Status   `json:"id"`
	Priority Priority `json:"priority"`

	TaskScope *TaskScope `json:"taskscope"`

	Header       string `json:"header"`
	Body         string `json:"body"`
	AssignedUser *User  `json:"user"`

	IsTrash bool `json:"istrash"`

	Estimate          *time.Time `json:"estimate"`
	RemainingEstimate *time.Time `json:"remainingestimate"`
}

type Status struct {
	ID          uint   `json:"id"`
	Description string `json:"decription"`
}

type Priority struct {
	ID          uint   `json:"id"`
	Description string `json:"decription"`
}
