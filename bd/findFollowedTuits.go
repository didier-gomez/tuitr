package bd

import (
	"context"
	"time"
	"log"
	"github.com/didier-gomez/tuitr/models"
	"go.mongodb.org/mongo-driver/bson"
)
func FindFollowedTuits(ID string, page int) ([]models.TuitFollowed, bool) {
	var PageSize = 20
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()
	db := MongoCN.Database("tuitr")
	col := db.Collection("relations")
	skip := (page-1) * PageSize

	conditions := make([]bson.M, 0)
	conditions = append(conditions, bson.M {"$match": bson.M{"userid": ID}})
	conditions = append(conditions, bson.M{
		"$lookup": bson.M{
			"from":         "tuits",
			"localField":   "relateduserid",
			"foreignField": "userid",
			"as":           "tuit",
		}})
	conditions = append(conditions, bson.M {"$unwind": "$tuit"})
	conditions = append(conditions, bson.M {"$sort": bson.M{"tuit.date": -1}})
	conditions = append(conditions, bson.M {"$skip": skip})
	conditions = append(conditions, bson.M {"$limit": PageSize})

	cursor, err := col.Aggregate(ctx, conditions)

	var results [] models.TuitFollowed
	if err != nil {
		log.Print(err.Error())
		return results, false
	}
	err = cursor.All(ctx, &results)
	if err != nil {
		log.Print(err.Error())
		return results, false
	}
	return results, true

}