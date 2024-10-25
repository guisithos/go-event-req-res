package handlers

import (
	"html/template"
	"log"
	"net/http"

	"myapp/internal/database"
	"myapp/internal/models"
)

func Task2Handler(w http.ResponseWriter, r *http.Request) {
	input1 := r.URL.Query().Get("input1")

	db, err := database.Connect()
	if err != nil {
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	query := "SELECT col1, col2, col3, col4, col5, col6, col7, col8 FROM table2 WHERE col1 = :1"
	rows, err := db.Query(query, input1)
	if err != nil {
		http.Error(w, "Query error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var data []models.Data
	for rows.Next() {
		var d models.Data
		if err := rows.Scan(&d.Col1, &d.Col2, &d.Col3, &d.Col4, &d.Col5, &d.Col6, &d.Col7, &d.Col8); err != nil {
			log.Println(err)
			continue
		}
		data = append(data, d)
	}

	tmpl, err := template.ParseFiles("web/templates/task2.html")
	if err != nil {
		http.Error(w, "Template parsing error", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, data)
}
