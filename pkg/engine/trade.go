package engine

import "encoding/json"

type Trade struct {
	Name         string  `json:"name"`
	TakerOrderID string  `json:"taker_order_id"`
	MakerOrderID string  `json:"maker_order_id"`
	Amount       int     `json:"amount"`
	Price        float64 `json:"price"`
	Timestamp    int64   `json:"timestamp"`
}

func (trade *Trade) FromJSON(msg []byte) error {
	return json.Unmarshal(msg, trade)
}

func (trade *Trade) ToJSON() []byte {
	str, _ := json.Marshal(trade)
	return str
}
