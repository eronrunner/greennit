package service

import (
	"context"
	"log"
	"time"

	"github.com/greennit/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService struct {
	Repository *database.UserRepo
}

func (s *UserService) Register(nickname, pwd, birth, email string) (*database.UserEntity, error) {
	ctx := context.TODO()
	// Check exist
	user, checkErr := s.Repository.GetUserByEmail(ctx, email)
	if checkErr != nil {
		log.Printf("UserService - Register - Error when check out %s, %s", email, checkErr)
		return nil, checkErr
	}
	if user != nil {
		log.Printf("UserService - Register - Account %s is exist", email)
		return nil, checkErr
	}
	entity := database.UserEntity{}
	entity.ID = primitive.NewObjectID()
	entity.Nickname = nickname
	entity.Birth = birth
	entity.Email = email
	entity.CreatedAt = time.Now()

	user, saveErr := s.Repository.Save(ctx, &entity)
	if saveErr != nil {
		log.Printf("UserService - Register - Error when register %s, %s", email, saveErr)
		return nil, saveErr
	}
	return user, nil
}