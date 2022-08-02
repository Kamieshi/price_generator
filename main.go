package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Kamieshi/price_generator/internal/config"
	"github.com/Kamieshi/price_generator/internal/service"
	rds "github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

func main() {
	conf, err := config.GetConfig()
	if err != nil {
		logrus.WithError(err).Fatal("Parse config from OS ENV")
		return
	}
	ctx := context.Background()
	companies := service.GetCountCompany(conf.CountCompany)
	generators := make([]*service.Generator, 0, len(companies))
	for _, c := range companies {
		generators = append(generators, service.NewGenerator(c))
	}
	client := rds.NewClient(&rds.Options{
		Addr:     fmt.Sprintf("%s:%s", conf.RedisHost, conf.RedisPort),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	for {
		for _, g := range generators {
			t_N := time.Now()
			g.GenerateAddCourse()
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
			log.Info(strCmd.Val(), "  ", time.Since(t_N))
		}
		log.Info(generators[0].LastCourse.Ask, " > ", generators[0].LastCourse.Bid)
		log.Info("UPDATE COMPLETE")
		time.Sleep(time.Duration(1000/conf.CountUpdatePerSecond) * time.Millisecond)
	}
}
