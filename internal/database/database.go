package database

import (
	"L0/internal/model"
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/lib/pq"
)

type Database struct {
	db *sql.DB
}

func (database *Database) Connect() {
	db, err := sql.Open("postgres", "user=postgres dbname=postgres sslmode=disable")
	if err != nil {
		fmt.Println(err)
		return
	}
	database.db = db
}

func (db *Database) AddOrder(order model.Order) error {
	_, err := db.db.Exec("INSERT INTO orders VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)",
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
		order.Oof_shard,
	)
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = db.addDeliveryInfo(order)
	if err != nil {
		return err
	}
	err = db.addPaymentInfo(order)
	if err != nil {
		return err
	}
	err = db.addItemsInfo(order)
	if err != nil {
		return err
	}
	return nil
}

func (db *Database) addDeliveryInfo(order model.Order) error {
	_, err := db.db.Exec("INSERT INTO delivery VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
		order.Order_uid,
		order.Delivery.Name,
		order.Delivery.Phone,
		order.Delivery.Zip,
		order.Delivery.City,
		order.Delivery.Address,
		order.Delivery.Region,
		order.Delivery.Email,
	)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (db *Database) addPaymentInfo(order model.Order) error {
	_, err := db.db.Exec("INSERT INTO payment VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)",
		order.Payment.Transaction,
		order.Payment.Request_id,
		order.Payment.Currency,
		order.Payment.Provider,
		order.Payment.Amount,
		order.Payment.Payment_dt,
		order.Payment.Bank,
		order.Payment.Delivery_cost,
		order.Payment.Goods_total,
		order.Payment.Custom_fee,
	)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (db *Database) addItemsInfo(order model.Order) error {
	for _, item := range order.Items {
		_, err := db.db.Exec("Insert into items VALUES ($1, $2, $3, $4, $5)",
			item.Chrt_id,
			item.Track_number,
			item.Price,
			item.Rid,
			item.Name,
			item.Sale,
			item.Size,
			item.Total_price,
			item.Nm_id,
			item.Brand,
			item.Status,
		)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
	return nil
}

func (db *Database) GetAllOrders() map[string]model.Order {
	var orders []model.Order
	type item_info struct {
		Chrt_id      int    `json:"chrt_id" db:"chrt_id"`
		Track_number string `json:"track_number" db:"track_number"`
		Price        int    `json:"price" db:"price"`
		Rid          string `json:"rid" db:"rid"`
		Name         string `json:"name" db:"name"`
		Sale         int    `json:"sale" db:"sale"`
		Size         string `json:"size" db:"size"`
		Total_price  int    `json:"total_price" db:"total_price"`
		Nm_id        int    `json:"nm_id" db:"nm_id"`
		Brand        string `json:"brand" db:"brand"`
		Status       int    `json:"status" db:"status"`
	}
	var item item_info
	rows, err := db.db.Query(`SELECT
			o.order_uid, o.track_number, o.entry, o.locale, o.internal_signature, o.customer_id, 
			o.delivery_service, o.shardkey, o.sm_id, o.date_created, o.oof_shard,
			d.name, d.phone, d.zip, d.city, d.address, d.region, d.email,
			p.transaction, p.request_id, p.currency, p.provider, p.amount, p.payment_dt, p.bank, p.delivery_cost, p.goods_total, p.custom_fee,
			oi.chrt_id,
			i.track_number, i.price, i.rid, i.name, i.sale, i.size, i.total_price, i.nm_id, i.brand, i.status
	 		FROM orders o 
			LEFT JOIN delivery d ON o.order_uid = d.order_uid 
			LEFT JOIN payment p ON o.order_uid = p.transaction
			LEFT Join orderitems oi ON o.order_uid = oi.order_uid
			LEFT JOIN items i ON oi.chrt_id = i.chrt_id`)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer rows.Close()
	for rows.Next() {
		var order model.Order
		if err := rows.Scan(
			&order.Order_uid,
			&order.Track_number,
			&order.Entry,
			&order.Locale,
			&order.Internal_signature,
			&order.Customer_id,
			&order.Delivery_service,
			&order.Shardkey,
			&order.Sm_id,
			&order.Date_created,
			&order.Oof_shard,
			&order.Delivery.Name,
			&order.Delivery.Phone,
			&order.Delivery.Zip,
			&order.Delivery.City,
			&order.Delivery.Address,
			&order.Delivery.Region,
			&order.Delivery.Email,
			&order.Payment.Transaction,
			&order.Payment.Request_id,
			&order.Payment.Currency,
			&order.Payment.Provider,
			&order.Payment.Amount,
			&order.Payment.Payment_dt,
			&order.Payment.Bank,
			&order.Payment.Delivery_cost,
			&order.Payment.Goods_total,
			&order.Payment.Custom_fee,
			&item.Chrt_id,
			&item.Track_number,
			&item.Price,
			&item.Rid,
			&item.Name,
			&item.Sale,
			&item.Size,
			&item.Total_price,
			&item.Nm_id,
			&item.Brand,
			&item.Status,
		); err != nil {
			fmt.Println(err)
			return nil
		}
		order.Items = append(order.Items, item)
		orders = append(orders, order)
	}
	orderMap := make(map[string]model.Order)
	for _, order := range orders {
		orderMap[order.Order_uid] = order
	}
	return orderMap
}

// func DbSubscribe(db *Database) {
// 	var srvNats nats.Service
// 	err := srvNats.Connect("consumer_db")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	err = srvNats.Subscribe("upload_consumer", db)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// }

func (database *Database) Consume(data []byte) error {
	var order model.Order
	err := json.Unmarshal(data, &order)
	if err != nil {
		return err
	}

	database.Connect()

	err = database.AddOrder(order)
	if err != nil {
		return err
	}
	return nil
}
