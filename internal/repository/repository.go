package repository

import (
	"github.com/phanvanpeter/my-portfolio/internal/models"
)

// DBRepository is an interface for allowing various database systems to connect
type DBRepository interface {
	GetTasks(done TaskDone) ([]models.Task, error)
	AddTask(task string) error
	CompleteTask(id int, completed bool) error
	EditTask(id int, task string) error
	DeleteTask(id int) error
}

type TaskDone uint

const Done TaskDone = 1
const NotDone TaskDone = 2
const All TaskDone = 3
