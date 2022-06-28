package repository

import "github.com/phanvanpeter/my-portfolio/models"

type DBRepository interface {
	GetTasks() ([]models.Task, error)
}
