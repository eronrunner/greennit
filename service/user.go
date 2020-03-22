package service

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/greennit/database"
	appErr "github.com/greennit/error"
)

type UserService struct {
	Repository *database.UserRepo
}

func (s *UserService) Register(nickname, pwd, birth, email string) (*database.UserEntity, error) {
	ctx := context.TODO()
	// Check exist
	user, checkErr := s.Repository.GetUserByEmail(ctx, email)
	if checkErr != nil {
		log.Printf("UserService;Register;Error when check out %s", email)
		return nil, checkErr
	}
	if user != nil {
		log.Printf("UserService;Register;Account %s is exist", email)
		return nil, fmt.Errorf("User;%s", appErr.ErrObjectExist)
	}
	entity := database.UserEntity{}
	entity.ID = primitive.NewObjectID()
	entity.Nickname = nickname
	entity.Birth = birth
	entity.Email = email
	entity.CreatedAt = time.Now()

	user, saveErr := s.Repository.Save(ctx, &entity)
	if saveErr != nil {
		log.Printf("UserService;Register;Error when register %s", email)
		return nil, saveErr
	}
	return user, nil
}