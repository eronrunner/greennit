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
	ID        primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Nickname  string             `json:"nickname" json:"nickname,omitempty"`
	SecretPwd string             `json:"secret_pwd" json:"secretpwd,omitempty"`
	Birth     string             `json:"birth" json:"birth,omitempty"`
	Email     string             `json:"email" json:"email,omitempty"`
	Picture   string             `json:"picture"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at,omitempty"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at,omitempty"`
}

func (u UserEntity) String() string {
	return fmt.Sprintf("User <%s>", u.Email)
}

// UserRepo - Handle USER's CRUD
type UserRepo struct {
	Client *mongo.Client
}

// Save - Create an user
func (repo *UserRepo) Save(ctx context.Context, entity *UserEntity) (*UserEntity, error) {
	coll := repo.Client.Database(DBName).Collection(UserColl)
	_, err := coll.InsertOne(ctx, entity)
	if err != nil {
		log.Printf("UserRepo;Save;Error when create account %s, %s", entity, err)
		return nil, err
	}
	entity.SecretPwd = ""
	return entity, nil
}

/*GetUserByEmail - Get user by email address,
If got ErrNoDocuments return nil, nil
If get other Error return nill, Error
If got User return *UserEntity, nil.
*/
func (repo *UserRepo) GetUserByEmail(ctx context.Context, email string) (*UserEntity, error) {
	coll := repo.Client.Database(DBName).Collection(UserColl)
	result := coll.FindOne(ctx, bson.M{"email": email})
	if err := result.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		log.Printf("UserRepo;GetUserByEmail;Error when find user %s, %s", email, err)
		return nil, err
	}
	user := UserEntity{}
	err := result.Decode(&user)
	if err != nil {
		log.Printf("UserRepo;GetUserByEmail;Error when extract user %s, %s", email, err)
		return nil, err
	}
	user.SecretPwd = ""
	return &user, nil
}

// CheckUserPwd - return true if PWD matches to USER in DB
func (repo *UserRepo) CheckUserPwd(ctx context.Context, email, pwd string) (bool, error) {
	coll := repo.Client.Database(DBName).Collection(UserColl)
	result := coll.FindOne(ctx, bson.M{"email": email, "secretpwd": pwd})
	if err := result.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		log.Printf("UserRepo;CheckUserPwd;Error when find user %s, %s", email, err)
		return false, err
	}
	user := UserEntity{}
	err := result.Decode(&user)
	if err != nil {
		log.Printf("UserRepo;CheckUserPwd;Error when extract user %s, %s", email, err)
		return false, err
	}
	return true, nil
}
