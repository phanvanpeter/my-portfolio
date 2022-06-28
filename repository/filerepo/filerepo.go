package filerepo

import (
	"bufio"
	"github.com/phanvanpeter/my-portfolio/models"
	"log"
	"os"
	"strings"
)

const fileName = "./repository/filerepo/tasks.txt"

type FileRepo struct {
}

const readFlag = os.O_APPEND|os.O_CREATE|os.O_RDONLY
const writeFlag = os.O_APPEND|os.O_CREATE|os.O_RDWR
const editFlag = os.O_CREATE|os.O_RDWR

func openFile(flag int) (*os.File, error) {
	return os.OpenFile(fileName, flag, 0644)
}

// GetTasks loads all the tasks saved in the file
func (f *FileRepo) GetTasks() ([]models.Task, error) {
	var tasks []models.Task

	file, err := openFile(readFlag)
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

func (f *FileRepo) AddTask(task string) error {
	file, err := openFile(writeFlag)
	if err != nil {
		return err
	}
	defer file.Close()

	task = strings.Trim(task, " ")
	_, err = file.WriteString(task + "\n")
	if err != nil {
		return err
	}
	return nil
}

func (f *FileRepo) DeleteTask(id int) error {
	tasks, err := f.GetTasks()
	if err != nil {
		return err
	}

	file, err := openFile(editFlag)
	if err != nil {
		return err
	}
	defer file.Close()

	// deleting the content of the file first
	if err := os.Truncate(fileName, 0); err != nil {
		log.Printf("Failed to truncate: %v", err)
	}

	for _, t := range tasks {
		if t.ID == id {
			continue
		}

		task := strings.Trim(t.Task, " ")
		_, err = file.WriteString(task + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}