package nosql

import (
	"context"

	"github.com/Reljod/Reljod-Portfolio-Backend/app/logger"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IDatabaseRepository interface {
	GetInfo(ctx context.Context, filters interface{}) (Info, error)
}

type MongoDbRepository struct {
	NoSqlDb *NoSqlDatabase
}

func (r *MongoDbRepository) GetInfo(ctx context.Context, filters interface{}) (Info, error) {
	var result Info
	err := r.NoSqlDb.Db.Collection("Portfolio Collection").FindOne(ctx, filters).Decode(&result)
	if err != nil {
		logger.Logger.Fatal("Collection: ", err)
	}

	logger.Logger.Debug("Result: ", result)
	return result, err
}

type Info struct {
	ID        primitive.ObjectID `bson:"_id"`
	Name      Name_t             `bson:"name"`
	Birthdate int                `bson:"birthdate"`
	Career    string             `bson:"career"`
}

type Name_t struct {
	Firstname  string `bson:"firstname,omitempty"`
	MiddleName string `bson:"middlename,omitempty"`
	LastName   string `bson:"lastname,omitempty"`
	NickName   string `bson:"nickname,omitempty"`
}
