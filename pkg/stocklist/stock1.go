package stocklist

import "stockexchange/pkg/engine"

var Book = engine.OrderBook{
	BuyOrders: []engine.Order{
		{Name: "Paytm", Amount: 10, Price: 845.8, ID: "user2.1", Intent: "buy", Timestamp: 1263},
		{Name: "Paytm", Amount: 5, Price: 854.8, ID: "user3.1", Intent: "buy", Timestamp: 1254},
		{Name: "Paytm", Amount: 5, Price: 854.8, ID: "user5.1", Intent: "buy", Timestamp: 1254},
	},
	SellOrders: []engine.Order{
		{Name: "Paytm", Amount: 5, Price: 854.8, ID: "user4.1", Intent: "sell", Timestamp: 1254},
		{Name: "Paytm", Amount: 5, Price: 845.8, ID: "user6.1", Intent: "sell", Timestamp: 1254},
		{Name: "Paytm", Amount: 5, Price: 845.8, ID: "user7.1", Intent: "sell", Timestamp: 1254},
		{Name: "Paytm", Amount: 5, Price: 845.8, ID: "user9.1", Intent: "sell", Timestamp: 1254},
	},
}
var Trades = []engine.Trade{}

func ProcessOrder(order engine.Order) {

	Trades = append(Trades, Book.Process(order)...)
}
