package params

type CreateOrder struct {
	OrderID      int    `json:"order_id"`
	CustomerName string `json:"customer_name"`
	Items        []CreateItem
}
