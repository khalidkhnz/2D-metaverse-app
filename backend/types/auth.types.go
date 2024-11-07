package types

import "github.com/khalidkhnz/2D-metaverse-app/backend/schema"

type LoginBody struct {
	Email    string `bson:"email" json:"email"`
	Password string `bson:"password" json:"password"`
}

type FullProfile struct {
	Auth    *schema.AuthSchema		`bson:"auth" json:"auth"`
	Profile *schema.ProfileSchema	`bson:"profile" json:"profile"`
}