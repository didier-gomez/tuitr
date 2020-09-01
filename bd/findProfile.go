package bd

import (
	"context"
	"log"
	"time"
	"github.com/didier-gomez/tuitr/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
func FindProfile(ID string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()
	db := MongoCN.Database("tuitr")
	col := db.Collection("users")
	
	var profile models.User

	objId, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M {
		"_id": objId,
	}
	err := col.FindOne(ctx, condition).Decode(&profile)
	profile.Password =""
	if err != nil {
		log.Fatal("register not found" +err.Error())
		return profile, err
	}
	return profile, nil
}