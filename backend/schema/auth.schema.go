package schema

import (
	"errors"
	"regexp"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthSchema struct {
	ID          primitive.ObjectID   `bson:"_id,omitempty" json:"_id"`
	FullName    string               `bson:"fullName" json:"fullName"`
	Email       string               `bson:"email" json:"email"`
	Password    string               `bson:"password" json:"password"`
	RoleId      primitive.ObjectID   `bson:"roleId" json:"roleId"`
	Permissions []primitive.ObjectID `bson:"permissions" json:"permissions"`
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
	if a.Permissions == nil {
		return errors.New("missing field permissions Array")
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
	return &AuthSchema{
		ID:        primitive.NewObjectID(),
		FullName:  fullName,
		Email:     email,
		Password:  password,
	}
}
