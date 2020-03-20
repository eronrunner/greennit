package database

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User - manipulate with USER mongo
type User interface {
	ListAll() ([]*User, error)
	GetById(id primitive.ObjectID) (*User, error)
	Save(nickname, pwd, birth, email, picture string) (*User, error)
	Delete(id primitive.ObjectID) (*User, error)
}

// UserEntity - USER domain model
type UserEntity struct {
	ID         primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Nickname   string             `json:"nickname" json:"nickname,omitempty"`
	SecrectPwd string             `json:"secrect_pwd" json:"SecrectPwd,omitempty"`
	Birth      string             `json:"birth" json:"SecrectPwd,omitempty"`
	Email      string             `json:"email" json:"email,omitempty"`
	Picture    string             `json:"picture"`
	CreatedAt  time.Time          `bson:"created_at" json:"created_at,omitempty"`
	UpdatedAt  time.Time          `bson:"updated_at"`
}
