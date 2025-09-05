package handlers

import (
	"fmt"
	functions "go-web/internal/functions"
	models "go-web/internal/models"
	"html/template"
	"net/http"
)

func Length(w http.ResponseWriter, r *http.Request) {

	var pgData models.PgData

	tmpl, err := template.ParseFiles("templates/length.html")
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	if r.Method == "POST" {
		strVal := r.FormValue("value")
		from := r.FormValue("from")
		to := r.FormValue("to")

		strResult := functions.Calculation("length", strVal, from, to)
		pgData = models.PgData{
			InputValue: strVal,
			Result:     strResult,
			From:       from,
			To:         to,
		}
	}

	tmpl.Execute(w, pgData)
}
