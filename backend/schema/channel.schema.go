package schema

import (
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ChannelSchema struct {
	ID             primitive.ObjectID   `bson:"_id,omitempty" json:"_id"`
	Name           string               `bson:"name" json:"name"`
	Description    string               `bson:"description" json:"description"`
	SpaceId        primitive.ObjectID   `bson:"spaceId" json:"spaceId"`
	Logo           string               `bson:"logo" json:"logo"`
	CreatorId      primitive.ObjectID   `bson:"creatorId" json:"creatorId"`
	MemberIds      []primitive.ObjectID `bson:"memberIds" json:"memberIds"`
	ConversationId primitive.ObjectID   `bson:"conversationId" json:"conversationId"`
	CreatedAt      time.Time            `bson:"createdAt" json:"createdAt"`
	UpdatedAt      time.Time            `bson:"updatedAt" json:"updatedAt"`
}

func (r *ChannelSchema) Validate() error {
	if r.Name == "" {
		return errors.New("missing required field: name")
	}
	if r.Description == "" {
		return errors.New("missing required field: description")
	}
	if r.CreatorId.IsZero() {
		return errors.New("missing required field: creatorId")
	}
	if r.SpaceId.IsZero() {
		return errors.New("missing required field: spaceId")
	}
	if r.ConversationId.IsZero() {
		return errors.New("missing required field: conversationId")
	}
	return nil
}

func CreateChannelDoc(name, description string, spaceId primitive.ObjectID, logo string, creatorId, conversationId primitive.ObjectID, members []primitive.ObjectID) *ChannelSchema {
	currentTime := time.Now()
	return &ChannelSchema{
		ID:             primitive.NewObjectID(),
		Name:           name,
		Description:    description,
		SpaceId:        spaceId,
		Logo:           logo,
		CreatorId:      creatorId,
		MemberIds:      members,
		ConversationId: conversationId,
		CreatedAt:      currentTime,
		UpdatedAt:      currentTime,
	}
}
