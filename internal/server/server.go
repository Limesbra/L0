package server

import (
	"L0/internal/service"
	"net/http"
)

func RunServer() {

	service.HandleHTTPRequests()
	http.ListenAndServe(":8080", nil)
}

//http://localhost:8080
