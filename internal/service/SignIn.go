package service

import (
	"log"
	"net/http"
	"text/template"
)

func SignIn(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("templates/html/autorization.html")

	if err != nil {
		log.Println("Ошибка обработки html")
		return
	}
	tmpl.ExecuteTemplate(w, "SignIn", nil)

}
