package table

import (
	"L0/internal/model"

	"github.com/jedib0t/go-pretty/table"
)

// func MakeTables(order model.Order) []*table.Writer {
// 	var t []*table.Writer
// 	t = append(t, makeOrderTable(model.Order{}))
// 	t = append(t, makeDeliveryTable(model.Order{}))
// 	t = append(t, makePaymentTable(model.Order{}))
// 	t = append(t, makeItemsTable(model.Order{}))
// 	return t

// }

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
	t.Render()
	return &t
}

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
	t.Render()
	return &t
}

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
	t.Render()
	return &t
}

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
	t.Render()
	return &t
}
