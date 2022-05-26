package params

import "time"

type Response struct {
	Status         int         `json:"status"`
	Message        string      `json:"message,omitempty"`
	Error          string      `json:"error,omitempty"`
	AdditionalInfo interface{} `json:"additional_info,omitempty"`
	Payload        interface{} `json:"payload,omitempty"`
}

type AllResponseData struct {
	OrderID      int            `json:"order_id"`
	OrderedAt    time.Time      `json:"ordered_at"`
	CustomerName string         `json:"customer_name"`
	Items        []ItemResponse `json:"items"`
}

type ItemResponse struct {
	ItemID      int    `json:"item_id"`
	ItemCode    string `json:"item_code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	OrderID     int    `json:"order_id"`
}
