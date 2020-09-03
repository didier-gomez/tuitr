package bd

import (
	"context"
	"time"

	"github.com/didier-gomez/tuitr/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
func UpdateRegister(u models.User, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()
	db := MongoCN.Database("tuitr")
	col := db.Collection("users")
	register := make(map[string] interface{})
	if len(u.Name) >0 {
		register["name"] = u.Name
	}
	if len(u.Name) >0 {
		register["name"] = u.Name
	}
	if len(u.Surname) >0 {
		register["surname"] = u.Surname
	}
	if len(u.Avatar) >0 {
		register["avatar"] = u.Avatar
	}
	if len(u.Banner) >0 {
		register["banner"] = u.Banner
	}
	if len(u.Bio) >0 {
		register["bio"] = u.Bio
	}
	if len(u.Location) >0 {
		register["location"] = u.Location
	}
	updateStr := bson.M{
		"$set": register,
	}

	objId, _ := primitive.ObjectIDFromHex(ID)
	where := bson.M{
		"_id": bson.M{"$eq": objId},
	}

	_, err := col.UpdateOne(ctx, where, updateStr)
	if err != nil {
		return false, err
	}
	return true, nil
}