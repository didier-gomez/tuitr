package models
import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
type TuitFollowed struct {
	ID primitive.ObjectID		`bson:"_id" json:"_id,omitempty"`
	UserID string						`bson:"userid" json:"userid,omitempty"`
	RelatedUserID string		`bson:"relateduserid" json:"relateduserid,omitempty"`
	Tuit struct {
		Message string		`bson:"message" json:"message,omitempty"`
		Date time.Time		`bson:"date" json:"date,omitempty"`
		ID string		`bson:"_id" json:"_id,omitempty"`
	}
}