// Package models Models  Price Generator
package models

import "github.com/google/uuid"

// Company some company
type Company struct {
	ID              uuid.UUID `json:"id" db:"id"`
	Name            string    `json:"name,omitempty" db:"name"`
	MaxAsk          uint32    `json:"max_ask,omitempty" db:"max_ask"`
	MinAsk          uint32    `json:"min_ask,omitempty" db:"min_ask"`
	MaxAskDifferent uint32    `json:"max_ask_different,omitempty" db:"max_ask_different"`
	MaxBid          uint32    `json:"max_bid,omitempty" db:"max_bid"`
	MinBid          uint32    `json:"min_bid,omitempty" db:"min_bid"`
	MaxBidDifferent uint32    `json:"max_bid_different,omitempty" db:"max_bid_different"`
	MaxAskBidDiff   uint32    `json:"max_ask_bid_diff,omitempty" db:"max_ask_bid_diff"`
}

// NewCompany Constructor
func NewCompany(name string, maxAsk, minAsk, maxAskDifferent, maxBid, minBid, maxBidDifferent, maxAskBidDiff uint32) *Company {

	return &Company{
		Name:            name,
		MaxAsk:          maxAsk,
		MinAsk:          minAsk,
		MaxAskDifferent: maxAskDifferent,
		MaxBid:          maxBid,
		MinBid:          minBid,
		MaxBidDifferent: maxBidDifferent,
		MaxAskBidDiff:   maxAskBidDiff,
	}
}
