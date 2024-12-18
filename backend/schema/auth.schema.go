package schema

import (
	"errors"
	"regexp"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthSchema struct {
	ID          primitive.ObjectID   `bson:"_id,omitempty" json:"_id"`
	FullName    string               `bson:"fullName" json:"fullName"`
	Email       string               `bson:"email" json:"email"`
	Password    string               `bson:"password" json:"password"`
	RoleId      primitive.ObjectID   `bson:"roleId" json:"roleId"`
	Permissions []primitive.ObjectID `bson:"permissions" json:"permissions"`
	SpaceIds    []primitive.ObjectID `bson:"spaceIds" json:"spaceIds"`
 	CreatedAt   time.Time            `bson:"createdAt" json:"createdAt"`
	UpdatedAt   time.Time            `bson:"updatedAt" json:"updatedAt"`
}

func (a *AuthSchema) Validate() error {
	if a.FullName == "" {
		return errors.New("missing required field: fullName")
	}
	if a.Email == "" {
		return errors.New("missing required field: email")
	}
	if !isValidEmail(a.Email) {
		return errors.New("invalid email format")
	}
	if a.Password == "" {
		return errors.New("missing required field: password")
	}
	if len(a.Password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}
	if a.RoleId.IsZero() {
		return errors.New("missing required field: roleId")
	}
	if len(a.Permissions) == 0 {
		a.Permissions = []primitive.ObjectID{}
	}
	if len(a.SpaceIds) == 0 {
		a.SpaceIds = []primitive.ObjectID{}
	}
	return nil
}

func isValidEmail(email string) bool {
	// Simple regex for email validation
	const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}

func CreateAuthDoc(fullName, email, password string) *AuthSchema {
	currentTime := time.Now()
	return &AuthSchema{
		ID:        primitive.NewObjectID(),
		FullName:  fullName,
		Email:     email,
		Password:  password,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}
}
