package handler

import (
	"fmt"
	"net/http"
	"slices"
	"strconv"
	"strings"

	"html/template"
	"htmxProj/models"

	"github.com/angelofallars/htmx-go"
)

type Handler struct {
	Tasks []models.Tasks
}

func New(t []models.Tasks) *Handler {
	return &Handler{Tasks: t}
}

func (h *Handler) IndexPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("html/index.html"))
	data := map[string][]models.Tasks{
		"Tasks": h.Tasks,
	}

	err := tmpl.Execute(w, data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		// fmt.Print("Error is :", err)
	}
}

func (h *Handler) AddTask(w http.ResponseWriter, r *http.Request) {
	if !htmx.IsHTMX(r) {
		fmt.Print("Not a htmx request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tsk := r.PostFormValue("task")
	if strings.TrimSpace(tsk) == "" {
		w.WriteHeader(http.StatusBadRequest)
		htmx.NewResponse().Write(w)
		return
	}

	desc := r.PostFormValue("desc")
	if strings.TrimSpace(desc) == "" {
		desc = "N/A"
	}

	totalTasks := len(h.Tasks)

	t := models.Tasks{Task: tsk, ID: totalTasks + 1, Desc: desc}

	h.Tasks = append(h.Tasks, t)

	// w.WriteHeader(http.StatusOK)
	err := htmx.NewResponse().Redirect("/").StatusCode(http.StatusOK).Write(w)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
}

func (h *Handler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	if !htmx.IsHTMX(r) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id := r.PathValue("id")

	fmt.Println("ID : ", id)

	val, err := strconv.Atoi(id)
	if err != nil {
		fmt.Printf("Error in atoi: %v", err)
		htmx.NewResponse().StatusCode(http.StatusBadRequest)
	}

	if val > len(h.Tasks)+1 {
		fmt.Print("Delete Out of Index")
		htmx.NewResponse().StatusCode(http.StatusBadRequest)
	}

	idx := 0
	for i := range h.Tasks {
		if h.Tasks[i].ID == val {
			idx = i
		}
	}

	h.Tasks = slices.Delete(h.Tasks, idx, idx+1)

	htmx.NewResponse().StatusCode(http.StatusAccepted).Redirect("/").Refresh(true)
}

func (h *Handler) Done(w http.ResponseWriter, r *http.Request) {
	if !htmx.IsHTMX(r) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id := r.PathValue("id")

	fmt.Println("ID : ", id)

	val, err := strconv.Atoi(id)
	if err != nil {
		fmt.Printf("Error in atoi: %v", err)
		htmx.NewResponse().StatusCode(http.StatusBadRequest)
	}

	for i := range h.Tasks {
		if h.Tasks[i].ID == val {
			h.Tasks[i].IsDone = true
		}
	}

	_ = htmx.NewResponse().StatusCode(http.StatusOK).Redirect("/").Refresh(true).Write(w)
}
