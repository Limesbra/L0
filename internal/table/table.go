package table

import (
	"L0/internal/model"

	"github.com/jedib0t/go-pretty/table"
)

// MakeOrderTable создает таблицу для переданного заказа.
// Принимает объект заказа в качестве входных данных и возвращает указатель на объект таблицы.
// Объект таблицы заполняется деталями заказа, такими как order_uid, track_number и т.д.
func MakeOrderTable(order model.Order) *table.Writer {
	t := table.NewWriter()
	t.AppendHeader(table.Row{
		"order_uid",
		"track_number",
		"entry",
		"locale",
		"internal_signature",
		"customer_id",
		"delivery_service",
		"shardkey",
		"sm_id",
		"date_created",
		"oof_shard",
	})
	t.AppendRow([]interface{}{
		order.Order_uid,
		order.Track_number,
		order.Entry,
		order.Locale,
		order.Internal_signature,
		order.Customer_id,
		order.Delivery_service,
		order.Shardkey,
		order.Sm_id,
		order.Date_created,
		order.Oof_shard})
	return &t
}

// MakePaymentTable создает таблицу для платежей по переданному заказу.
// Принимает объект заказа в качестве входных данных и возвращает указатель на объект таблицы.
// Объект таблицы заполняется данными о платеже, такими как transaction, request_id и т.д.
// Функция добавляет строку с данными платежа в таблицу и возвращает указатель на объект таблицы.
func MakePaymentTable(order model.Order) *table.Writer {
	t := table.NewWriter()
	t.AppendHeader(table.Row{
		"transaction",
		"request_id",
		"currency",
		"provider",
		"amount",
		"payment_dt",
		"bank",
		"delivery_cost",
		"goods_total",
		"custom_fee",
	})
	t.AppendRow([]interface{}{
		order.Payment.Transaction,
		order.Payment.Request_id,
		order.Payment.Currency,
		order.Payment.Provider,
		order.Payment.Amount,
		order.Payment.Payment_dt,
		order.Payment.Bank,
		order.Payment.Delivery_cost,
		order.Payment.Goods_total,
		order.Payment.Custom_fee})
	return &t
}

// MakeDeliveryTable создает таблицу для информации о доставке по переданному заказу.
// Принимает объект заказа в качестве входных данных и возвращает указатель на объект таблицы.
// Объект таблицы заполняется данными о доставке, такими как name, phone и т.д.
// Функция добавляет строку с данными о доставке в таблицу и возвращает указатель на объект таблицы.
func MakeDeliveryTable(order model.Order) *table.Writer {
	t := table.NewWriter()
	t.AppendHeader(table.Row{
		"name",
		"phone",
		"zip",
		"city",
		"address",
		"region",
		"email",
	})
	t.AppendRow([]interface{}{
		order.Delivery.Name,
		order.Delivery.Phone,
		order.Delivery.Zip,
		order.Delivery.City,
		order.Delivery.Address,
		order.Delivery.Region,
		order.Delivery.Email,
	})
	return &t
}

// MakeItemsTable создает таблицу для элементов (товаров) по переданному заказу.
// Принимает объект заказа в качестве входных данных и возвращает указатель на объект таблицы.
// Объект таблицы заполняется данными о товарах, такими как chrt_id, track_number и т.д.
// Функция добавляет строки с данными о товарах в таблицу и возвращает указатель на объект таблицы.
func MakeItemsTable(order model.Order) *table.Writer {
	t := table.NewWriter()
	t.AppendHeader(table.Row{
		"chrt_id",
		"track_number",
		"price",
		"rid",
		"name",
		"sale",
		"size",
		"total_price",
		"nm_id",
		"brand",
		"status",
	})
	for i := 0; i < len(order.Items); i++ {
		t.AppendRow(table.Row{
			order.Items[i].Chrt_id,
			order.Items[i].Track_number,
			order.Items[i].Price,
			order.Items[i].Rid,
			order.Items[i].Name,
			order.Items[i].Sale,
			order.Items[i].Size,
			order.Items[i].Total_price,
			order.Items[i].Nm_id,
			order.Items[i].Brand,
			order.Items[i].Status,
		})
	}
	return &t
}
