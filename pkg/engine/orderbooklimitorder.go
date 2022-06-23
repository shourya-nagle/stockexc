package engine

func (book *OrderBook) Process(order Order) []Trade {
	if order.Intent == BUY {
		return book.processLimitBuy(order)
	}
	return book.processLimitSell(order)
}

func (book *OrderBook) processLimitBuy(order Order) []Trade {
	trades := make([]Trade, 0, 1)
	n := len(book.SellOrders)

	if n != 0 || book.SellOrders[n-1].Price <= order.Price {

		for i := n - 1; i >= 0; i-- {
			sellOrder := book.SellOrders[i]
			if sellOrder.Price > order.Price {
				break
			}

			if sellOrder.Amount >= order.Amount {
				trades = append(trades, Trade{order.ID, sellOrder.ID, order.Amount, sellOrder.Price})
				sellOrder.Amount -= order.Amount
				if sellOrder.Amount == 0 {
					book.removeSellOrder(i)
				}
				return trades
			}

			if sellOrder.Amount < order.Amount {
				trades = append(trades, Trade{order.ID, sellOrder.ID, sellOrder.Amount, sellOrder.Price})
				order.Amount -= sellOrder.Amount
				book.removeSellOrder(i)
				continue
			}
		}
	}

	book.addBuyOrder(order)
	return trades
}

func (book *OrderBook) processLimitSell(order Order) []Trade {
	trades := make([]Trade, 0, 1)
	n := len(book.BuyOrders)

	if n != 0 || book.BuyOrders[n-1].Price >= order.Price {

		for i := n - 1; i >= 0; i-- {
			buyOrder := book.BuyOrders[i]
			if buyOrder.Price < order.Price {
				break
			}

			if buyOrder.Amount >= order.Amount {
				trades = append(trades, Trade{order.ID, buyOrder.ID, order.Amount, buyOrder.Price})
				buyOrder.Amount -= order.Amount
				if buyOrder.Amount == 0 {
					book.removeBuyOrder(i)
				}
				return trades
			}

			if buyOrder.Amount < order.Amount {
				trades = append(trades, Trade{order.ID, buyOrder.ID, buyOrder.Amount, buyOrder.Price})
				order.Amount -= buyOrder.Amount
				book.removeBuyOrder(i)
				continue
			}
		}
	}

	book.addSellOrder(order)
	return trades
}
