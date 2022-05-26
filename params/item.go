package params

type CreateItem struct {
	ItemID      int    `json:"item_id"`
	ItemCode    string `json:"item_code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantitiy"`
	OrderID     uint   `json:"order_id"`
}
