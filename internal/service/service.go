package service

import (
	"L0/internal/cache"
	"L0/internal/model"
	"L0/internal/nats"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
)

func HandleHTTPRequests() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/upload", Upload)
	http.HandleFunc("/show", Upload)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	// Открываем файл с шаблоном
	tmpl, err := template.ParseFiles("../web/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
	// Отображаем шаблон
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}
}

func Upload(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var order model.Order
		file, _, err := r.FormFile("jsonfile")
		if err != nil {
			fmt.Println("this", err)
			return
		}
		filybytes, _ := io.ReadAll(file)
		err = json.Unmarshal(filybytes, &order)
		if err != nil {
			fmt.Println("There", err)
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)

		var srvNats nats.Service
		srvNats.Connect("produce_upload")
		srvNats.Publish("upload_consumer", filybytes)
		srvNats.Close()
	}
}

func ShowInfo(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	fmt.Println("info")

	order, ok := (*cache.PtrCache)[id]

	if !ok {
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	}

	fileJSON, err := json.MarshalIndent(order, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(fileJSON)
}
