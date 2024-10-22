package main

import (
	"log"
	"net/http"

	"github.com/oantoniosilva/go-crud-in-memory/internal/handlers"
	"github.com/oantoniosilva/go-crud-in-memory/internal/repositories"
)

func main() {
	taskRepository := repositories.NewTaskInMemory()
	tasksHandler := handlers.NewTask(taskRepository)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /tasks/{uuid}", tasksHandler.Show)
	mux.HandleFunc("POST /tasks", tasksHandler.Register)
	mux.HandleFunc("PUT /tasks/{uuid}", tasksHandler.Update)
	mux.HandleFunc("DELETE /tasks/{uuid}", tasksHandler.Destroy)

	log.Println("Server is running on port: 8082")

	if err := http.ListenAndServe(":8082", mux); err != nil {
		log.Fatal(err)
	}
}
