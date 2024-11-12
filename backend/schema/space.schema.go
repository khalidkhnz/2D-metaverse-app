package schema

import (
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SpaceSchema struct {
	ID           primitive.ObjectID   `bson:"_id,omitempty" json:"_id"`
	Name         string               `bson:"name" json:"name"`
	Description  string               `bson:"description" json:"description"`
	Logo         string               `bson:"logo" json:"logo"`
	CoverArt     string               `bson:"coverArt" json:"coverArt"`
	CreatorId    primitive.ObjectID   `bson:"creatorId" json:"creatorId"`
	MemberIds    []primitive.ObjectID `bson:"memberIds" json:"memberIds"`
	AdminIds     []primitive.ObjectID `bson:"adminIds" json:"adminIds"`
	ModeratorIds []primitive.ObjectID `bson:"moderatorIds" json:"moderatorIds"`
	ElderIds     []primitive.ObjectID `bson:"elderIds" json:"elderIds"`
	ChannelIds   []primitive.ObjectID `bson:"channelIds" json:"channelIds"`
	CreatedAt    time.Time            `bson:"createdAt" json:"createdAt"`
	UpdatedAt    time.Time            `bson:"updatedAt" json:"updatedAt"`
}

func (r *SpaceSchema) Validate() error {
	if r.Name == "" {
		return errors.New("missing required field: name")
	}
	if r.CreatorId.IsZero() {
		return errors.New("missing required field: creatorId")
	}
	if r.AdminIds == nil {
		r.AdminIds = []primitive.ObjectID{}
	}
	if r.ChannelIds == nil {
		r.ChannelIds = []primitive.ObjectID{}
	}
	if r.ElderIds == nil {
		r.ElderIds = []primitive.ObjectID{}
	}
	if r.MemberIds == nil {
		r.MemberIds = []primitive.ObjectID{}
	}
	if r.ModeratorIds == nil {
		r.ModeratorIds = []primitive.ObjectID{}
	}
	
	return nil
}

func CreateSpaceDoc(name, description string, creatorId primitive.ObjectID, members, admins, channels []primitive.ObjectID) *SpaceSchema {
	currentTime := time.Now()
	return &SpaceSchema{
		ID:          primitive.NewObjectID(),
		Name:        name,
		Description: description,
		CreatorId:   creatorId,
		MemberIds:   members,
		AdminIds:    admins,
		ChannelIds:  channels,
		CreatedAt:   currentTime,
		UpdatedAt:   currentTime,
	}
}