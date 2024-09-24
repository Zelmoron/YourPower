package service

import (
	"log"
	"net/http"
	"text/template"
)

func SignUp(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("templates/html/sign_up.html")

	if err != nil {
		log.Println("Ошибка обработки html")
		return
	}
	tmpl.ExecuteTemplate(w, "SignUp", nil)

}
