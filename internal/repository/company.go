// Package repository Work with repository
package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"

	"priceGenerator/internal/models"
)

type Company struct {
	pull *pgxpool.Pool
}

func NewRepositoryCompany(conn *pgxpool.Pool) *Company {
	return &Company{
		pull: conn,
	}
}

// Insert Company to db
func (c *Company) Insert(ctx context.Context, comp *models.Company) error {
	querySQL := "INSERT INTO Company(id,name, max_ask, min_ask,max_ask_different,max_bid,min_bid,max_bid_different,max_ask_bid_diff) values ($1,$2,$3,$4,$5,$6,$7,$8,$9)"
	comp.ID = uuid.New()
	comTag, err := c.pull.Exec(ctx, querySQL,
		comp.ID,
		comp.Name,
		comp.MinAsk,
		comp.MaxAsk,
		comp.MaxAskDifferent,
		comp.MaxBid,
		comp.MinBid,
		comp.MaxBidDifferent,
		comp.MaxAskBidDiff,
	)
	if err != nil {
		comp.ID = uuid.UUID{}
		return fmt.Errorf("repository Company/Create:%v", err)
	}
	if !comTag.Insert() {
		comp.ID = uuid.UUID{}
		return fmt.Errorf("repository Company/Create:%v", err)
	}
	return nil
}

func (c *Company) Get(ctx context.Context, id uuid.UUID) (*models.Company, error) {
	querySQL := "SELECT  id,name, max_ask, min_ask,max_ask_different,max_bid,min_bid,max_bid_different FROM Company WHERE id=$1;"
	var comp models.Company
	err := c.pull.QueryRow(ctx, querySQL, id).Scan(
		&comp.ID,
		&comp.Name,
		&comp.MinAsk,
		&comp.MaxAsk,
		&comp.MaxAskDifferent,
		&comp.MaxBid,
		&comp.MinBid,
		&comp.MaxBidDifferent,
	)
	if err != nil {
		return nil, fmt.Errorf("repository Company/Get:%v", err)
	}
	return &comp, nil
}
func (c *Company) GetAll(ctx context.Context) ([]*models.Company, error) {
	querySQL := "SELECT  id,name, max_ask, min_ask,max_ask_different,max_bid,min_bid,max_bid_different FROM Company"
	rows, err := c.pull.Query(ctx, querySQL)
	if err != nil {
		return nil, fmt.Errorf("repository Company/GetAll:%v", err)
	}
	defer rows.Close()
	companies := make([]*models.Company, 0, len(rows.FieldDescriptions()))

	for rows.Next() {
		comp := models.Company{}
		err = rows.Scan(
			&comp.ID,
			&comp.Name,
			&comp.MinAsk,
			&comp.MaxAsk,
			&comp.MaxAskDifferent,
			&comp.MaxBid,
			&comp.MinBid,
			&comp.MaxBidDifferent,
		)
		if err != nil {
			logrus.WithError(err).Error()
			continue
		}
		companies = append(companies, &comp)
	}
	return companies, nil
}

func (c *Company) Delete(ctx context.Context, id uuid.UUID) error {
	querySQL := "DELETE FROM Company WHERE id=$1"
	comTag, err := c.pull.Exec(ctx, querySQL, id)
	if err != nil {
		return fmt.Errorf("repository Company/Delete:%v", err)
	}
	if !comTag.Delete() {
		logrus.Infof("repository Company/Delete commandTagDelete:%v", comTag.Delete())
	}
	return nil
}
