package bd

import (
	"context"
	"time"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
func DeleteTuit(ID string, UserID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("tuitr")
	col := db.Collection("tuits")
	objID,_ := primitive.ObjectIDFromHex(ID)
	whereData := bson.M {
		"_id": objID,
		"userid": UserID,
	}
	_, err := col.DeleteOne(ctx, whereData)
	return err
}