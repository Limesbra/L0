package server

import (
	"L0/internal/service"
	"net/http"
)

// RunServer запускает HTTP-сервер, обрабатывая HTTP-запросы с помощью функции HandleHTTPRequests из пакета service
func RunServer() {

	service.HandleHTTPRequests()
	http.ListenAndServe(":8080", nil)
}

//http://localhost:8080
