package bd

import (
	"context"
	"time"
	"github.com/didier-gomez/tuitr/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
func CreateTuit(t models.TuitInsert) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()
	db := MongoCN.Database("tuitr")
	col := db.Collection("tuits")
	register := bson.M {
		"userid": t.UserID,
		"message": t.Message,
		"date": t.Date,
	}
	result, err := col.InsertOne(ctx, register)
	if err!= nil {
		return "", false, err
	}
	objId, _ := result.InsertedID.(primitive.ObjectID)
	return objId.String(), true, nil
}