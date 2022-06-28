package repository

import "github.com/phanvanpeter/my-portfolio/models"

// DBRepository is an interface for allowing various database systems to connect
type DBRepository interface {
	GetTasks() ([]models.Task, error)
}
