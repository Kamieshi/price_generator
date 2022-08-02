// Package model Models  Price Generator
package model

import (
	"github.com/google/uuid"
)

// Company some company
type Company struct {
	ID   uuid.UUID `json:"ID,omitempty"`
	Name string    `json:"Name,omitempty"`
}
