package service

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func Index(w http.ResponseWriter, r *http.Request) {
	flag, id := CheckToken(r)
	if flag {
		tmpl, err := template.ParseFiles("templates/html/index.html")

		if err != nil {
			log.Println("Ошибка обработки html")
			return
		}
		tmpl.ExecuteTemplate(w, "Index", nil)
		fmt.Println(id)
	} else {
		http.Redirect(w, r, "http://127.0.0.1:8080/", http.StatusSeeOther)

	}

}
