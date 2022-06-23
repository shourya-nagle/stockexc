package engine

import "encoding/json"

type Intent string

const (
	BUY  = "buy"
	SELL = "sell"
)

type Order struct {
	Amount int    `json:"amount"`
	Price  int    `json:"price"`
	ID     string `json:"id"`
	Intent Intent `json:"intent"`
}

func (order *Order) FromJSON(msg []byte) error {
	return json.Unmarshal(msg, order)
}

func (order *Order) ToJSON() []byte {
	str, _ := json.Marshal(order)
	return str
}
