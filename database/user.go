package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

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

func (u UserEntity) String() string {
	return fmt.Sprintf("User{%s}", u.Email)
}

type UserRepo struct {
	Client *mongo.Client
}

func (repo *UserRepo) Save(ctx context.Context,  entity *UserEntity) (*UserEntity, error) {
	coll := repo.Client.Database(DBName).Collection(UserColl)
	_, err := coll.InsertOne(ctx, entity)
	if err != nil {
		log.Printf("UserRepo - Save - Error when create account %s, %s", entity, err)
		return nil, err
	}
	entity.SecrectPwd = ""
	return entity, nil
}

func (repo *UserRepo) GetUserByEmail(ctx context.Context, email string) (*UserEntity, error) {
	coll := repo.Client.Database(DBName).Collection(UserColl)
	result := coll.FindOne(ctx, bson.M{"email": email})
	if err := result.Err(); err != nil {
		log.Printf("UserRepo - getUSerByEmail - Error when find account %s, %s", email, err)
		return nil, err
	}
	user := UserEntity{}
	err := result.Decode(&user)
	if err != nil {
		log.Printf("UserRepo - getUSerByEmail - Error when find account %s, %s", email, err)
		return nil, err
	}
	return &user, nil
}