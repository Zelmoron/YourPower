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

	rtr.HandleFunc("/", service.SignIn).Methods("GET")              // обработчик страницы авторизации
	rtr.HandleFunc("/signUp", service.SignUp).Methods("GET")        // обработчик страницы регистрации
	rtr.HandleFunc("/index", service.Index).Methods("GET")          // обработчик главной страницы
	rtr.HandleFunc("/checkform", service.CheckForm).Methods("POST") // обработчик для проверки пароля и логина страницы

	http.Handle("/", rtr) // все обраоботчики через роутер

}
