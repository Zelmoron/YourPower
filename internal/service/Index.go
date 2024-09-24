package service

import (
	"log"
	"net/http"
	"text/template"
)

func Index(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/html/index.html")

	if err != nil {
		log.Println("Ошибка обработки html")
		return
	}
	tmpl.ExecuteTemplate(w, "Index", nil)
}
