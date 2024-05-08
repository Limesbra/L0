package service

import (
	"L0/internal/cache"
	"L0/internal/model"
	"L0/internal/nats"
	"L0/internal/table"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
)

// HandleHTTPRequests настраивает маршруты HTTP для различных обработчиков
func HandleHTTPRequests() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/upload", Upload)
	http.HandleFunc("/orders", ShowInfo)
}

// handleIndex обрабатывает главную страницу, отображая шаблон из файла index.html
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

// Upload обрабатывает загрузку файла JSON, парсит его и выполняет действия с полученными данными
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

// ShowInfo отображает информацию о заказе на веб-странице, используя данные из кэша
func ShowInfo(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	order, ok := (*cache.PtrCache)[id]

	if !ok {
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Информация по заказу: %s\n Информация по заказу\n", id)
	t := table.MakeOrderTable(order)
	fmt.Fprint(w, (*t).Render())
	fmt.Fprintf(w, "\n Информация по доставке\n")
	t = table.MakeDeliveryTable(order)
	fmt.Fprint(w, (*t).Render())
	fmt.Fprintf(w, "\n Информация по платежу\n")
	t = table.MakePaymentTable(order)
	fmt.Fprint(w, (*t).Render())
	fmt.Fprintf(w, "\n Информация по позициям\n")
	t = table.MakeItemsTable(order)
	fmt.Fprint(w, (*t).Render())
}
