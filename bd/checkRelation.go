package bd
import (
	"context"
	"time"
	"log"
	"github.com/didier-gomez/tuitr/models"
	"go.mongodb.org/mongo-driver/bson"
)
func CheckRelation(t models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()
	db := MongoCN.Database("tuitr")
	col := db.Collection("relations")

	whereData := bson.M {
		"userid": t.UserID,
		"relateduserid": t.RelatedUserID,
	}
	var result models.Relation
	err := col.FindOne(ctx, whereData).Decode(&result)
	if err != nil {
		log.Fatal("db error " +err.Error())
		return false, err
	}
	return true,nil
}