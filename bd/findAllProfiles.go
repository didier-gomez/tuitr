package bd

import (
	"context"
	"time"
	"log"
	"github.com/didier-gomez/tuitr/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)
func FindAllProfiles(ID string, page int64, search string, type_ string) ([]*models.User, bool) {
	var PageSize int64 = int64(20)
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()
	db := MongoCN.Database("tuitr")
	col := db.Collection("users")

	var results []*models.User

	findOptions := options.Find()
	findOptions.SetLimit(PageSize)
	findOptions.SetSkip((page-1) * PageSize)

	whereData := bson.M {
		"name": bson.M{"$regex": `(?i)`+search},
	}

	cursor, err := col.Find(ctx, whereData, findOptions)

	if err != nil {
		log.Print("error: " + err.Error())
		return results, false
	}

	var found, incl bool

	for cursor.Next(context.TODO()) {
		var s models.User
		err := cursor.Decode(&s)
		if err != nil {
			log.Print(err.Error())
			return results, false
		}

		var r models.Relation
		r.UserID = ID
		r.RelatedUserID = s.ID.Hex()
		incl = false
		found, err = CheckRelation(r)
		if type_ == "new" && !found {
			incl = true
		}
		
		if (type_ == "follow" && found == true) {
			incl = true
		}

		if r.RelatedUserID == ID {
			incl = false
		}

		if incl {
			s.Password = ""
			s.Bio =""
			s.Location = ""
			s.Email = ""
			s.Banner = ""
			results = append(results, &s)
		}
	}
	err = cursor.Err()
	if err != nil {
		log.Print("error: " +err.Error())
	}
	cursor.Close(ctx)
	return results, true

}