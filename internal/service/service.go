package service

import (
	"L0/internal/model"
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
		fmt.Println(order)
	}
}

func ShowInfo(http.ResponseWriter, *http.Request) {
	fmt.Println("info")
}
