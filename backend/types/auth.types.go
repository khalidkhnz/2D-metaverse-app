package types

type LoginBody struct {
	Email    string `bson:"email" json:"email"`
	Password string `bson:"password" json:"password"`
}
