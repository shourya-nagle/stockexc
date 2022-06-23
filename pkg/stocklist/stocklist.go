package stocklist

import (
	"context"
	"time"

	"golang.org/x/sync/errgroup"
)

var Stocks = []string{
	"Swiss Life AG", "Spotify", "SolarCity",
}

type StockLister struct {
	stockticker time.Duration
}

func (s *StockLister) start(ctx context.Context) error {
	errorgroup, errorcontext := errgroup.WithContext(ctx)
	for _, stock := range Stocks {
		errorgroup.Go(func() error {
			return s.runstockworker(stock, errorcontext)
		})
	}
	return nil
}
func (a *StockLister) runstockworker(stock string, ctx context.Context) error {
	for {
		continue
	}

}
func NewStockLister(t time.Duration) *StockLister {
	return &StockLister{
		stockticker: t,
	}
}
