package handlers

import (
	"fmt"
	"github.com/phanvanpeter/my-portfolio/internal/render"
	"github.com/phanvanpeter/my-portfolio/models"
	"github.com/phanvanpeter/my-portfolio/repository/filerepo"
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

	tasks, err := filerepo.GetTasks()
	if err != nil {
		log.Fatal("Error loading tasks:", err)
	}

	data := make(map[string]interface{})
	data["tasks"] = tasks

	render.Template(w, r, file, &models.TemplateData{
		StringMap: strMap,
		Data: data,
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