package main

import (
	"context"
	"encoding/json"
	"time"

	rds "github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"

	"priceGenerator/internal/models"
	"priceGenerator/internal/service"
)

func main() {
	ctx := context.Background()
	companies := []*models.Company{
		models.NewCompany("Company 1"),
		models.NewCompany("Company 2"),
		models.NewCompany("Company 3"),
		models.NewCompany("Company 4"),
		models.NewCompany("Company 5"),
	}
	generators := make([]*service.Generator, 0, len(companies))
	for _, c := range companies {
		generators = append(generators, service.NewGenerator(c))
	}
	client := rds.NewClient(&rds.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	for {
		for _, g := range generators {
			g.GenerateCourse()
			data, err := json.Marshal(g.LastCourse)
			if err != nil {
				log.WithError(err).Error()
				continue
			}
			arg := rds.XAddArgs{
				Stream: "prices",
				MaxLen: 0,
				ID:     "",
				Values: map[string]string{
					"price": string(data),
				},
			}
			strCmd := client.XAdd(ctx, &arg)
			if strCmd.Err() != nil {
				log.WithError(err)
			}
			log.Info(strCmd.Val())

		}
		log.Info(generators[0].LastCourse.Ask, " > ", generators[0].LastCourse.Bid)
		log.Info("UPDATE COMPLETE")
		time.Sleep(100 * time.Millisecond)
	}
}
