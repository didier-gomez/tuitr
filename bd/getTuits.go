package bd
import (
	"context"
	"time"
	"log"
	"github.com/didier-gomez/tuitr/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)
var PageSize int64 = 20
func GetTuits(ID string, page int64) ([]*models.Tuits, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()
	db := MongoCN.Database("tuitr")
	col := db.Collection("tuits")
	var result []*models.Tuits
	whereData := bson.M {
		"userid": ID,
	}
	opts := options.Find()
	opts.SetLimit(PageSize)
	opts.SetSort(bson.D{{Key: "date", Value: -1}}) // order by date desc
	opts.SetSkip((page-1) * PageSize)
	cursor, err := col.Find(ctx, whereData)
	if err != nil {
		log.Fatal(err.Error())
		return result, false
	}
	for cursor.Next(context.TODO()) {
		var register models.Tuits
		err := cursor.Decode(&register)
		if(err != nil) {
			return result, false
		}
		result = append(result, &register)
	}
	return result, true
}