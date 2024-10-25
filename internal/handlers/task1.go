package handlers

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"myapp/internal/database"
	"myapp/internal/models"
	"myapp/internal/rabbitmq"
)

func Task1Handler(w http.ResponseWriter, r *http.Request) {
	input1 := r.URL.Query().Get("input1")
	input2 := r.URL.Query().Get("input2")

	db, err := database.Connect()
	if err != nil {
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	query := "SELECT col1, col2, col3, col4, col5 FROM table1 WHERE col1 = :1 AND col2 = :2"
	rows, err := db.Query(query, input1, input2)
	if err != nil {
		http.Error(w, "Query error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var data []models.Data
	for rows.Next() {
		var d models.Data
		if err := rows.Scan(&d.Col1, &d.Col2, &d.Col3, &d.Col4, &d.Col5); err != nil {
			log.Println(err)
			continue
		}
		data = append(data, d)
	}

	tmpl, err := template.ParseFiles("web/templates/layout.html", "web/templates/task1.html")
	if err != nil {
		http.Error(w, "Template parsing error", http.StatusInternalServerError)
		return
	}

	tmpl.ExecuteTemplate(w, "layout", map[string]interface{}{
		"Title": "Task 1",
		"Data":  data,
	})
}

func Task1EventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var event struct {
		Action   string `json:"action"`
		ID       string `json:"id"`
		NewValue string `json:"newValue,omitempty"`
	}

	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Create a message for RabbitMQ
	message := map[string]interface{}{
		"action":   event.Action,
		"id":       event.ID,
		"newValue": event.NewValue,
	}

	// Send the message to RabbitMQ
	if err := rabbitmq.PublishMessage("task1_events", message); err != nil {
		log.Printf("Error publishing message to RabbitMQ: %v", err)
		http.Error(w, "Error processing event", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}
