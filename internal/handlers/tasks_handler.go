package handlers

import (
	"fmt"
	"github.com/phanvanpeter/my-portfolio/internal/render"
	"github.com/phanvanpeter/my-portfolio/models"
	"log"
	"net/http"
)

// Tasks renders a task HTML page
func Tasks(w http.ResponseWriter, r *http.Request) {
	file := "tasks.page"

	strMap := map[string]string{}

	if app.Session.Exists(r.Context(), "task") {
		task := app.Session.Get(r.Context(), "task").(string)
		msg := fmt.Sprintf("Your task has been processed: %s", task)
		strMap["newTask"] = msg
	}

	render.Template(w, r, file, &models.TemplateData{
		StringMap: strMap,
	})

	app.Session.Remove(r.Context(), "task")
}

// PostTasks adds a new task the to list of tasks
func PostTasks(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal("The task could not be processed:", err)
	}

	values := r.PostForm
	task := values.Get("task")
	app.Session.Put(r.Context(), "task", task)

	http.Redirect(w, r, "/tasks", http.StatusSeeOther)
}