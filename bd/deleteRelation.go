package bd

import (
	"context"
	"time"
	"github.com/didier-gomez/tuitr/models"

)
func DeleteRelation(t models.Relation) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("tuitr")
	col := db.Collection("relations")
	_, err := col.DeleteOne(ctx, t)
	return err
}