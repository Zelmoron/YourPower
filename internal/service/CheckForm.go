package service

import (
	"net/http"
)

func CheckForm(w http.ResponseWriter, r *http.Request) {

	nameUser := r.FormValue("username")
	if nameUser == "admin" {
		CreateToken(nameUser, w, r)
		http.Redirect(w, r, "http://127.0.0.1:8080/index", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "http://127.0.0.1:8080/", http.StatusSeeOther)
	}

}
