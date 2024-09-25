package service

import (
	"log"
	"net/http"
	"text/template"
)

func SignIn(w http.ResponseWriter, r *http.Request) {
	flag, _ := CheckToken(r)
	if flag {

		http.Redirect(w, r, "http://127.0.0.1:8080/index", http.StatusSeeOther)
	} else {
		tmpl, err := template.ParseFiles("templates/html/autorization.html")

		if err != nil {
			log.Println("Ошибка обработки html")
			return
		}
		tmpl.ExecuteTemplate(w, "SignIn", nil)
	}

}
