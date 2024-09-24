package main

import (
	"fmt"
	"log"
	"myApp/internal/app"
	"myApp/internal/config"
	"net/http"
)

func main() {
	port := config.Port()
	log.Printf("Server on %s is starting", port)
	app.Handlers()
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
