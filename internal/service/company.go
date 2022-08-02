package service

import (
	"fmt"

	"github.com/Kamieshi/price_generator/internal/model"
	"github.com/google/uuid"
)

// NewCompany Constructor
func NewCompany(name string) *model.Company {
	return &model.Company{
		ID:   uuid.New(),
		Name: name,
	}
}

func GetCountCompany(countCompany int) []*model.Company {
	companies := make([]*model.Company, 0, countCompany)
	for i := 0; i < countCompany; i++ {
		companies = append(companies, NewCompany(fmt.Sprintf("Company %d", i+1)))
	}
	return companies
}
