package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/phanvanpeter/my-portfolio/internal/models"
	"github.com/phanvanpeter/my-portfolio/internal/repository"
	"time"
)

func NewRepo(db *sql.DB) *postgresRepo {
	return &postgresRepo{
		DB: db,
	}
}

type postgresRepo struct {
	DB *sql.DB
}


// GetTasks loads all the tasks saved in the database
func (p *postgresRepo) GetTasks(done repository.TaskDone) ([]models.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()

	var tasks []models.Task
	var query string

	switch done {
	case repository.All:
		query = `
			select id, task, done, created_at, updated_at
			from tasks
			order by updated_at`
	case repository.Done:
		query = `
			select id, task, done, created_at, updated_at
			from tasks
			where done = true
			order by updated_at`
	case repository.NotDone:
		query = `
			select id, task, done, created_at, updated_at
			from tasks
			where done = false
			order by updated_at`
	default:
		return tasks, errors.New("invalid TaskDone")
	}

	rows, err := p.DB.QueryContext(ctx, query)
	if err != nil {
		return tasks, errors.New(fmt.Sprintf("Unable to fetch the tasks from the database: %s", err))
	}

	for rows.Next() {
		var t models.Task
		err = rows.Scan(
			&t.ID,
			&t.Task,
			&t.Done,
			&t.CreatedAt,
			&t.UpdatedAt,
		)
		if err != nil {
			return tasks, errors.New(fmt.Sprintf("Unable to scan the loaded tasks: %s", err))
		}
		tasks = append(tasks, t)
	}

	if err = rows.Err(); err != nil {
		return tasks, errors.New(fmt.Sprintf("Error scanning the tasks: %s", err))
	}

	return tasks, nil
}

// AddTask adds a new task to the database
func (p *postgresRepo) AddTask(task string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()

	query := `
		insert into tasks 
		(task, created_at, updated_at) 
		values ($1, $2, $3)`

	_, err := p.DB.ExecContext(ctx, query,
		task,
		time.Now(),
		time.Now(),
	)

	if err != nil {
		return errors.New(fmt.Sprintf("Error inserting the task: %s", err))
	}

	return nil
}

// CompleteTask marks the given task as done
func (p *postgresRepo) CompleteTask(id int, completed bool) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()

	query := `
		update tasks
		set done = $1
		where id = $2`

	_, err := p.DB.ExecContext(ctx, query, completed, id)

	if err != nil {
		return errors.New(fmt.Sprintf("Error updating the task with id %d: %s", id, err))
	}

	return nil
}

// DeleteTask deletes a task with the given id
func (p *postgresRepo) DeleteTask(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()

	query := `
		delete from tasks 
		where id = $1`

	_, err := p.DB.ExecContext(ctx, query, id)

	if err != nil {
		return errors.New(fmt.Sprintf("Error deleting the task with id %d: %s", id, err))
	}

	return nil
}
