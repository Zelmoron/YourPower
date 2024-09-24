package app

import (
	"myApp/internal/service"
	"net/http"

	"github.com/gorilla/mux"
)

// запуск обработчика хендлера, запускается через main
func Handlers() {
	rtr := mux.NewRouter()
	//подключил свои стили и js
	http.Handle("/templates/",
		http.StripPrefix("/templates", http.FileServer(http.Dir("./templates/"))))

	rtr.HandleFunc("/", service.SignIn)       // обработчик страницы авторизации
	rtr.HandleFunc("/SignUp", service.SignUp) // обработчик страницы регистрации

	http.Handle("/", rtr) // все обраоботчики через роутер

}
