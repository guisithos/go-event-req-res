package handlers

import (
	"html/template"
	"log"
	"net/http"

	"myapp/internal/database"
	"myapp/internal/models"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")

	db, err := database.Connect()
	if err != nil {
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	query := "SELECT id, name, start_date, end_date, type FROM services WHERE start_date >= :1 AND end_date <= :2"
	rows, err := db.Query(query, startDate, endDate)
	if err != nil {
		http.Error(w, "Query error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var data []models.Data
	for rows.Next() {
		var d models.Data
		if err := rows.Scan(&d.ID, &d.Name, &d.StartDate, &d.EndDate, &d.Type); err != nil {
			log.Println(err)
			continue
		}
		data = append(data, d)
	}

	tmpl, err := template.ParseFiles("web/templates/results.html")
	if err != nil {
		http.Error(w, "Template parsing error", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, data)
}
