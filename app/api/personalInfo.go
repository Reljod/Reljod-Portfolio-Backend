package api

import (
	"context"
	"net/http"
	"time"

	"github.com/Reljod/Reljod-Portfolio-Backend/ds/nosql"
	"go.mongodb.org/mongo-driver/bson"
)

// About godoc
// @Summary Fetches Information About me.
// @Description Fetches Information About me.
// @Produce  json
// @Success 200 {object} nosql.Info
// @Router /about [get]
func Get(w http.ResponseWriter, r *http.Request) {
	Log.Info("Get info.")

	info, err := getInfo()
	CheckError(err)

	WriteJsonToResponse(w, info)
}

// About godoc
// @Summary Fetches my name.
// @Description Fetches my name.
// @Produce  json
// @Success 200 {object} nosql.Name_t
// @Router /about/name [get]
func GetName(w http.ResponseWriter, r *http.Request) {
	Log.Info("Get name info.")

	info, err := getInfo()
	CheckError(err)

	WriteJsonToResponse(w, info.Name)
}

// About godoc
// @Summary Fetches my birthday.
// @Description Fetches my birthday.
// @Produce  json
// @Success 200 {object} int
// @Router /about/birthday [get]
func GetBirthDate(w http.ResponseWriter, r *http.Request) {
	Log.Info("Get birthday info.")

	info, err := getInfo()
	CheckError(err)

	WriteJsonToResponse(w, info.Birthdate)
}

func getInfo() (nosql.Info, error) {
	db := nosql.NewNoSqlDatabase()
	var repo nosql.IDatabaseRepository = &nosql.MongoDbRepository{NoSqlDb: db}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	info, err := repo.GetInfo(ctx, bson.M{})
	CheckError(err)

	Log.Debug(info)

	return info, nil
}
