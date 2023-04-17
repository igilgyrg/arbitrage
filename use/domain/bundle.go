package domain

import (
	"fmt"
	"time"
)

type Bundle struct {
	Id                   int       `json:"id"`
	Symbol               string    `json:"symbol"`
	ExchangeFrom         string    `json:"exchange_from"`
	PriceFrom            float64   `json:"price_from"`
	ExchangeTo           string    `json:"exchange_to"`
	PriceTo              float64   `json:"price_to"`
	PercentageDifference float64   `json:"percentage_difference"`
	UpdatedAt            time.Time `json:"updated_at"`
}

func (b *Bundle) String() string {
	return fmt.Sprintf("SYMBOL: %s\nFROM-TO: %s-%s\nPRICES: %f-%f\nPERCENT: %f", b.Symbol, b.ExchangeFrom, b.ExchangeTo, b.PriceFrom, b.PriceTo, b.PercentageDifference)
}
