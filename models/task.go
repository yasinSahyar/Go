package models

import "time"

type Task struct {
    ID        string    `json:"id" bson:"_id,omitempty"`
    Title     string    `json:"title" bson:"title"`
    Completed bool      `json:"completed" bson:"completed"`
    CreatedAt time.Time `json:"created_at" bson:"created_at"`
}
