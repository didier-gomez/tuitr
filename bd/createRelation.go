package bd
import (
	"context"
	"time"
	"github.com/didier-gomez/tuitr/models"
)

func CreateRelation(r models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()
	db := MongoCN.Database("tuitr")
	col := db.Collection("relations")

	_, err := col.InsertOne(ctx, r)
	
	if err != nil {
		return false, err
	}
	return true, nil
}