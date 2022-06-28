package handlers

import (
	"context"
	"github.com/phanvanpeter/my-portfolio/config"
	"github.com/phanvanpeter/my-portfolio/internal/render"
	"log"
	"net/http"
)

// Tasks renders a task HTML page
func Tasks(w http.ResponseWriter, r *http.Request) {
	file := "tasks.page"

	strMap := taskGetSessions(r.Context())
	data := taskGetData()

	render.Template(w, r, file, &config.TemplateData{
		StringMap: strMap,
		Data: data,
	})

	app.Session.Remove(r.Context(), "task")
	app.Session.Remove(r.Context(), "bad_task")
}

// taskGetSessions gets sessions, particularly "task" session and put them in the map of strings
func taskGetSessions(c context.Context) map[string]string {
	strMap := map[string]string{}

	if app.Session.Exists(c, "bad_task") {
		task := app.Session.Get(c, "bad_task").(string)
		strMap["badTask"] = task
	}

	if app.Session.Exists(c, "task") {
		task := app.Session.Get(c, "task").(string)
		strMap["task"] = task
	}
	return strMap
}

// taskGetData loads the tasks from the file and returns them in the map
func taskGetData() map[string]interface{} {
	tasks, err := db.GetTasks()
	if err != nil {
		log.Fatal("Error loading tasks:", err)
	}

	data := make(map[string]interface{})
	data["tasks"] = tasks
	return data
}

// PostTasks adds a new task the to list of tasks
func PostTasks(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal("The task could not be processed:", err)
	}

	values := r.PostForm
	task := values.Get("task")

	if task == "bad task" {
		http.Error(w, "Task cannot be 'bad task'", http.StatusSeeOther)
		app.Session.Put(r.Context(), "bad_task", task)
		http.Redirect(w, r, "/tasks", http.StatusSeeOther)
		return
	}

	app.Session.Put(r.Context(), "task", task)

	http.Redirect(w, r, "/tasks", http.StatusSeeOther)
}