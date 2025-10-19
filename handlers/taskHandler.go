package handlers

import (
    "context"
    "encoding/json"
    "net/http"
    "strconv"
    "time"

    "github.com/go-chi/chi/v5"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "yasin-todo-api/models"
)

type TaskHandler struct {
    Collection *mongo.Collection
}

func (h *TaskHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    filter := bson.M{}
    completedQuery := r.URL.Query().Get("completed")
    if completedQuery != "" {
        completed, _ := strconv.ParseBool(completedQuery)
        filter["completed"] = completed
    }

    cur, err := h.Collection.Find(ctx, filter)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer cur.Close(ctx)

    tasks := []models.Task{}
    for cur.Next(ctx) {
        var t models.Task
        cur.Decode(&t)
        tasks = append(tasks, t)
    }

    json.NewEncoder(w).Encode(tasks)
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    var t models.Task
    json.NewDecoder(r.Body).Decode(&t)
    t.CreatedAt = time.Now()

    res, err := h.Collection.InsertOne(ctx, t)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    t.ID = res.InsertedID.(string)
    json.NewEncoder(w).Encode(t)
}

func (h *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
    id := chi.URLParam(r, "id")
    var t models.Task
    json.NewDecoder(r.Body).Decode(&t)

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    update := bson.M{
        "$set": bson.M{
            "title":     t.Title,
            "completed": t.Completed,
        },
    }

    _, err := h.Collection.UpdateByID(ctx, id, update)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(t)
}

func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
    id := chi.URLParam(r, "id")
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    _, err := h.Collection.DeleteOne(ctx, bson.M{"_id": id})
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}
