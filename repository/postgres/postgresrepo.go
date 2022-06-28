package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/phanvanpeter/my-portfolio/models"
)

func NewRepo(db *sql.DB) *postgresRepo {
	return &postgresRepo{
		DB: db,
	}
}

type postgresRepo struct {
	DB *sql.DB
}

var fakeTasks = []models.Task{
	{1, "First task"},
	{2, "Second task"},
	{3, "Third task"},
	{4, "Fourth task"},
}
var counter = 5

// GetTasks loads all the tasks saved in the file
func (p *postgresRepo) GetTasks() ([]models.Task, error) {
	return fakeTasks, nil
}

// AddTask adds a new task to the database
func (p *postgresRepo) AddTask(task string) error {
	t := models.Task{counter, task}
	counter++
	fakeTasks = append(fakeTasks, t)
	return nil
}

// DeleteTask deletes a task with the given id
func (p *postgresRepo) DeleteTask(id int) error {
	// finds the index of the task to be deleted
	var index = -1
	for i, t := range fakeTasks {
		if t.ID == id {
			index = i
			break
		}
	}
	if index == -1 {
		return errors.New(fmt.Sprintf("There is no task with the id %d", id))
	}

	// deleting the task
	fakeTasks = append(fakeTasks[:index], fakeTasks[index + 1:]...)

	return nil
}
