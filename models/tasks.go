package models

import "time"

type Task struct {
	ID        int
	Task      string
	Done      bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
