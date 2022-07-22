package models

import (
	"time"
)

// Course object current course
type Course struct {
	Comp *Company  `json:"comp,omitempty"`
	Bid  int64     `json:"bid,omitempty"`
	Ask  int64     `json:"ask,omitempty"`
	Time time.Time `json:"time"`
}
