package routes

import (
    "github.com/go-chi/chi/v5"
    "yasin-todo-api/handlers"
)

func TaskRoutes(h *handlers.TaskHandler) *chi.Mux {
    r := chi.NewRouter()

    r.Get("/tasks", h.GetTasks)
    r.Post("/tasks", h.CreateTask)
    r.Put("/tasks/{id}", h.UpdateTask)
    r.Delete("/tasks/{id}", h.DeleteTask)

    return r
}
