package models

// Course object current course
type Course struct {
	Comp *Company `json:"Company,omitempty"`
	Bid  int64    `json:"Bid,omitempty"`
	Ask  int64    `json:"Ask,omitempty"`
	Time string   `json:"Time"`
}
