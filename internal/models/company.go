// Package models Models  Price Generator
package models

import "github.com/google/uuid"

// Company some company
type Company struct {
	ID   uuid.UUID `json:"ID,omitempty"`
	Name string    `json:"Name,omitempty"`
}

// NewCompany Constructor
func NewCompany(name string) *Company {
	return &Company{
		ID:   uuid.New(),
		Name: name,
	}
}
