package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"yasin-todo-api/handlers"
	"yasin-todo-api/routes"

	_ "yasin-todo-api/docs" // Swagger dokümantasyonu import

	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Yasin Todo API
// @version 1.0
// @description Simple REST API for managing tasks.
// @host localhost:9000
// @BasePath /
func main() {
	// MongoDB bağlantısı
	mongoURL := os.Getenv("MONGO_URL")
	if mongoURL == "" {
		mongoURL = "mongodb://localhost:27017"
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURL))
	if err != nil {
		log.Fatal("❌ Mongo client oluşturulamadı:", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("❌ Mongo bağlantısı başarısız:", err)
	}

	collection := client.Database("yasin_todo").Collection("tasks")
	handler := &handlers.TaskHandler{Collection: collection}

	// Router
	r := chi.NewRouter()
	r.Mount("/", routes.TaskRoutes(handler))

	// Swagger endpoint
	r.Get("/swagger/*", httpSwagger.WrapHandler)

	log.Println("🚀 Server started at: http://localhost:9000")
	log.Println("📘 Swagger docs available at: http://localhost:9000/swagger/index.html")

	http.ListenAndServe(":9000", r)
}
