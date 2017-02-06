package models

import (
	"time"
)

type TaskScope struct {
	Epic       Epic
	ID         string
	Decription string
}

type Task struct {
	ID uint

	Status   Status
	Priority Priority

	TaskScope *TaskScope

	Header       string
	Body         string
	assignedUser *User

	trash bool

	Estimate          *time.Time
	RemainingEstimate *time.Time
}

type Status struct {
	ID          uint
	Description string
}

type Priority struct {
	ID          uint
	Description string
}
