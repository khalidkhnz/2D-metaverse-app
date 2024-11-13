package schema

import (
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MessageSchema struct {
	ID              primitive.ObjectID   `bson:"_id,omitempty" json:"_id"`
	From            primitive.ObjectID   `bson:"from" json:"from"`
	Message         string               `bson:"message" json:"message"`
	Attachments     []string             `bson:"attachments" json:"attachments"`
	Mentions        []primitive.ObjectID `bson:"mentions" json:"mentions"`
	ConversationId  primitive.ObjectID   `bson:"conversationId" json:"conversationId"`
	CreatedAt       time.Time            `bson:"createdAt" json:"createdAt"`
	UpdatedAt       time.Time            `bson:"updatedAt" json:"updatedAt"`
}

func (r *MessageSchema) Validate() error {
	if r.Message == "" {
		return errors.New("missing required field: message")
	}
	if r.ConversationId.String() == "" {
		return errors.New("missing required field: conversationId")
	}
	if r.From.IsZero() {
		return errors.New("missing required field: from")
	}
	if r.Mentions == nil {
		r.Mentions = []primitive.ObjectID{}
	}
	if r.Attachments == nil {
		r.Attachments = []string{}
	}
	return nil
}

func CreateMessageDoc(from primitive.ObjectID, message string, attachments []string, mentions []primitive.ObjectID, conversationId primitive.ObjectID) *MessageSchema {
	currentTime := time.Now()
	return &MessageSchema{
		ID:             primitive.NewObjectID(),
		From:           from,
		Message:        message,
		Attachments:    attachments,
		Mentions:       mentions,
		ConversationId: conversationId,
		CreatedAt:      currentTime,
		UpdatedAt:      currentTime,
	}
}