package filerepo

import (
	"bufio"
	"github.com/phanvanpeter/my-portfolio/models"
	"os"
)

const fileName = "./repository/filerepo/tasks.txt"

type FileRepo struct {
}

// GetTasks loads all the tasks saved in the file
func (f *FileRepo) GetTasks() ([]models.Task, error) {
	var tasks []models.Task

	file, err := os.Open(fileName)
	if err != nil {
		return tasks, err
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	sc.Split(bufio.ScanLines)

	id := 1
	for sc.Scan() {
		t := models.Task{
			ID: id,
			Task: sc.Text(),
		}

		tasks = append(tasks, t)
		id++
	}
	if err = sc.Err(); err != nil {
		return tasks, err
	}

	return tasks, nil
}
