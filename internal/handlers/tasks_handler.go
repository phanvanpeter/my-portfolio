package handlers

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/phanvanpeter/my-portfolio/internal/config"
	"github.com/phanvanpeter/my-portfolio/internal/forms"
	"github.com/phanvanpeter/my-portfolio/internal/render"
	"github.com/phanvanpeter/my-portfolio/internal/repository"
	"log"
	"net/http"
	"strconv"
)

// Tasks renders a task HTML page
func Tasks(w http.ResponseWriter, r *http.Request) {
	file := "tasks.page"

	strMap, formErrors := taskGetSessions(r.Context())
	data := taskGetData()

	render.Template(w, r, file, &config.TemplateData{
		StringMap: strMap,
		FormErrors: formErrors,
		Data: data,
	})

	app.Session.Remove(r.Context(), "newTask")
	app.Session.Remove(r.Context(), "task")
	app.Session.Remove(r.Context(), "formErrors")
}

// taskGetSessions gets sessions, particularly "task" session and put them in the map of strings
func taskGetSessions(c context.Context) (map[string]string, forms.FormErrors) {
	strMap := map[string]string{}
	errMap := forms.FormErrors{}

	if app.Session.Exists(c, "formErrors") {
		errMap = app.Session.Get(c, "formErrors").(forms.FormErrors)
	}

	if app.Session.Exists(c, "task") {
		task := app.Session.Get(c, "task").(string)
		strMap["task"] = task
	}
	return strMap, errMap
}

// taskGetData loads the tasks from the file and returns them in the map
func taskGetData() map[string]interface{} {
	tasks, err := db.GetTasks(repository.NotDone)
	if err != nil {
		log.Fatal("Error loading tasks:", err)
	}

	completedTasks, err := db.GetTasks(repository.Done)
	if err != nil {
		log.Fatal("Error loading tasks:", err)
	}

	data := make(map[string]interface{})
	data["tasks"] = tasks
	data["completedTasks"] = completedTasks
	return data
}

// PostTask adds a new task the to list of tasks
func PostTask(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal("The task could not be processed:", err)
	}

	form := forms.New(r.PostForm)
	form.MinLength("task", 3)
	form.Required("task")

	task := r.Form.Get("task")

	if !form.IsValid() {
		formErrors := form.Errors.GetAll()

		app.Session.Put(r.Context(), "task", task)
		app.Session.Put(r.Context(), "formErrors", formErrors)
		http.Redirect(w, r, "/tasks", http.StatusSeeOther)
		return
	}

	app.Session.Put(r.Context(), "newTask", task)

	err = db.AddTask(task)
	if err != nil {
		log.Fatalf("Error storing the task: %s", err)
	}

	app.Session.Put(r.Context(), "task", task)

	http.Redirect(w, r, "/tasks", http.StatusSeeOther)
}

// CompleteTask marks the task as done
func CompleteTask(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatalf("Error parsing a form: %s", err)
	}

	taskID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Fatalf("Invalid task ID, expected number, got %s", chi.URLParam(r, "id"))
	}

	completed, err := strconv.ParseBool(r.PostForm.Get("completed"))
	if err != nil {
		log.Fatalf("Invalid 'completed', expected bool, got %s", r.PostForm.Get("completed"))
	}

	err = db.CompleteTask(taskID, completed)
	if err != nil {
		log.Fatalf("Error completing the task %d: %s", taskID, err)
	}

	http.Redirect(w, r, "/tasks", http.StatusSeeOther)
}

// EditTask edits the task (only it's text, not any other attribute).
func EditTask(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatalf("Error parsing a form: %s", err)
	}

	taskID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Fatalf("Invalid task ID, expected number, got %s", chi.URLParam(r, "id"))
	}

	task := r.PostForm.Get("task")

	err = db.EditTask(taskID, task)
	if err != nil {
		log.Fatalf("Error editing the task %d: %s", taskID, err)
	}

	http.Redirect(w, r, "/tasks", http.StatusSeeOther)
}

// DeleteTask deletes task from the file
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatalf("Error parsing a form: %s", err)
	}

	taskID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Fatalf("Invalid task ID, expected number, got %s", chi.URLParam(r, "id"))
	}

	err = db.DeleteTask(taskID)
	if err != nil {
		log.Fatalf("Error deleting the task %d: %s", taskID, err)
	}

	http.Redirect(w, r, "/tasks", http.StatusSeeOther)
}
