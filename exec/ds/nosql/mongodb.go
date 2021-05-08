package main

import (
	"context"
	"time"

	"github.com/Reljod/Reljod-Portfolio-Backend/ds/nosql"
	"go.mongodb.org/mongo-driver/bson"

	log "github.com/Reljod/Reljod-Portfolio-Backend/app/logger"
)

var logger = log.Logger

func main() {
	db := nosql.NewNoSqlDatabase()
	var repo nosql.IDatabaseRepository = &nosql.MongoDbRepository{NoSqlDb: db}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	info, err := repo.GetInfo(ctx, bson.M{})
	checkErr(err)

	logger.Info(info.Name)
}

func checkErr(err error) {
	if err != nil {
		logger.Fatal(err)
	}
}
