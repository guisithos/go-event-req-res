package main

import (
	"log"
	"net/http"

	"myapp/internal/handlers"
	"myapp/internal/rabbitmq"
)

func main() {
	if err := rabbitmq.Init(); err != nil {
		log.Fatalf("Failed to initialize RabbitMQ: %v", err)
	}
	defer rabbitmq.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/templates/index.html")
	})
	http.HandleFunc("/task1", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/templates/task1.html")
	})
	http.HandleFunc("/task1/results", handlers.Task1Handler)
	http.HandleFunc("/task1/event", handlers.Task1EventHandler)
	http.HandleFunc("/task2", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/templates/task2.html")
	})
	http.HandleFunc("/task2/results", handlers.Task2Handler)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
