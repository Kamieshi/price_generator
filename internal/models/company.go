// Package models Models  Price Generator
package models

import "github.com/google/uuid"

// Company some company
type Company struct {
	ID   uuid.UUID `json:"id,omitempty"`
	Name string    `json:"name,omitempty"`
}

// NewCompany Constructor
func NewCompany(name string) *Company {
	return &Company{
		ID:   uuid.New(),
		Name: name,
	}
}
