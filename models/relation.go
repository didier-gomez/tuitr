package models

type Relation struct {
	UserID string 				`bson:"userid" json:"userid"`
	RelatedUserID string	`bson:"relateduserid" json:"relateduserid"`
}