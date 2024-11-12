package schema

import (
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ConversationSchema struct {
	ID        primitive.ObjectID   `bson:"_id,omitempty" json:"_id"`
	MemberIds []primitive.ObjectID `bson:"memberIds" json:"memberIds"`
	MessageIds []primitive.ObjectID `bson:"messageIds" json:"messageIds"`
	CreatedAt time.Time            `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time            `bson:"updatedAt" json:"updatedAt"`
}

func (r *ConversationSchema) Validate() error {
	if len(r.MemberIds) < 2 {
		return errors.New("missing required field: members must be 2 or more")
	}
	return nil
}

func CreateConversationDoc(members []primitive.ObjectID) *ConversationSchema {
	currentTime := time.Now()
	return &ConversationSchema{
		ID:        primitive.NewObjectID(),
		MemberIds: members,
		MessageIds: []primitive.ObjectID{},
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}
}