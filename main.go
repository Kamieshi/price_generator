package main

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	log "github.com/sirupsen/logrus"

	"priceGenerator/internal/conf"
	"priceGenerator/internal/models"
	"priceGenerator/internal/repository"
)

func initValue(ctx context.Context, repCompany *repository.Company) {
	companies, err := repCompany.GetAll(ctx)
	if err != nil {
		log.WithError(err).Fatal()
	}
	if len(companies) == 0 {
		if err = repCompany.Insert(ctx, models.NewCompany(
			"Company 1",
			100000,
			50000,
			150,
			110000,
			55000,
			150,
			500,
		)); err != nil {
			log.WithError(err).Fatal()
		}
		if err = repCompany.Insert(ctx, models.NewCompany(
			"Company 2",
			100000,
			50000,
			150,
			110000,
			55000,
			150,
			500,
		)); err != nil {
			log.WithError(err).Fatal()
		}
	}
}

func main() {
	ctx := context.Background()
	config, err := conf.NewConfiguration()
	if err != nil {
		log.WithError(err).Fatal()
	}
	pool, err := pgxpool.Connect(ctx, config.PostgresConnString)
	if err != nil {
		log.WithError(err).Fatal()
	}
	repCompany := repository.NewRepositoryCompany(pool)
	initValue(ctx, repCompany)

	companies, err := repCompany.GetAll(ctx)
	if err != nil {
		log.WithError(err).Fatal()
	}
	genereator1 := models.NewGenerator(companies[0], models.NewCourse(companies[0]))
	for i := 0; ; i++ {
		genereator1.GenerateCourse()
		if genereator1.LastCourse.Bid >= genereator1.LastCourse.Ask {
			log.Error(genereator1.LastCourse, i)
			break
		}
	}
}
