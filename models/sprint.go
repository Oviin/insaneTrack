package models

import (
	"time"
)

type SprintModel struct {
	Description string     `json:"decription"`
	Teams       []*Team    `json:"teams"`
	StartTime   *time.Time `json:"start"`
	EndTime     *time.Time `json:"end"`
}
