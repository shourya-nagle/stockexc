package engine

import (
	"encoding/json"
)

type Order struct {
	Name      string  `json:"Name"`
	Amount    int     `json:"Amount"`
	Price     float64 `json:"Price"`
	ID        string  `json:"ID"`
	Intent    string  `json:"Intent"`
	Timestamp int64   `json:"Timestamp"`
}

func (order *Order) FromJSON(msg []byte) error {
	return json.Unmarshal(msg, order)
}

func (order *Order) ToJSON() []byte {
	str, _ := json.Marshal(order)
	return str
}
