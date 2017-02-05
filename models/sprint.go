package models

import (
	"time"
)

type SprintModel struct {
	Description string
	Teams       []*Team
	StartTime   *time.Time
	EndTime     *time.Time
}
