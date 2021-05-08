package nosql

import (
	"context"

	"github.com/Reljod/Reljod-Portfolio-Backend/app/environ"
	"github.com/Reljod/Reljod-Portfolio-Backend/app/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type NoSqlDatabase struct {
	Db      *mongo.Database
	Session *mongo.Client
}

func NewNoSqlDatabase() *NoSqlDatabase {
	client, err := mongo.NewClient(options.Client().ApplyURI(environ.DatabaseUrl))
	if err != nil {
		logger.Logger.Fatal("Cannot connect to MongoDB Server.")
	}
	client.Connect(context.TODO())

	db := client.Database("PortfolioDB")
	if db == nil {
		logger.Logger.Fatal("Cannot connect to Database.")
	}

	var sqlDb *NoSqlDatabase = &NoSqlDatabase{Db: db, Session: client}
	return sqlDb
}
