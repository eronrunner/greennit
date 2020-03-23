package service

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"time"

	"crypto/sha256"

	jwtGo "github.com/dgrijalva/jwt-go"
	"github.com/greennit/database"
	appErr "github.com/greennit/error"
	"github.com/greennit/util"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UserService - Handle USER's logics
type UserService struct {
	Repository *database.UserRepo
}

// Register - Register an user
func (s *UserService) Register(nickname string, pwd, birth, email string) (*database.UserEntity, error) {
	ctx := context.TODO()
	// Validate
	if len(email) <= 0 && len(pwd) <= 0 {
		return nil, fmt.Errorf("User;%s", appErr.ErrBadCredential)
	}
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

	// hasher := sha256.New()
	// hasher.Write([]byte(pwd))
	// sha := hex.EncodeToString(hasher.Sum(nil))

	entity := database.UserEntity{}
	entity.ID = primitive.NewObjectID()
	entity.Nickname = nickname
	entity.Birth = birth
	entity.Email = email
	entity.CreatedAt = time.Now()
	entity.UpdatedAt = time.Now()
	secret := sha256.Sum256([]byte(pwd))
	entity.SecretPwd = base64.URLEncoding.EncodeToString(secret[:])

	user, saveErr := s.Repository.Save(ctx, &entity)
	if saveErr != nil {
		log.Printf("UserService;Register;Error when register %s", email)
		return nil, saveErr
	}
	return user, nil
}

// Login - User login by email and secret pwd
func (s *UserService) Login(ctx context.Context, email string, pwd string) (*string, error) {
	// Validate
	if len(email) <= 0 || len(pwd) <= 0 {
		log.Printf("UserService;Login; %s", appErr.ErrBadCredential.Error())
		return nil, fmt.Errorf("User;%s", appErr.ErrBadCredential)
	}
	// Check exist
	user, checkErr := s.Repository.GetUserByEmail(ctx, email)
	if checkErr != nil {
		log.Printf("UserService;Login;Error when check out %s", email)
		return nil, checkErr
	}
	if user != nil {
		issueTime := time.Now()
		expirationTime := issueTime.Add(1 * time.Minute)
		checked, err := s.Repository.CheckUserPwd(ctx, email, pwd)
		if err != nil {
			log.Printf("UserService;Login;Error when check PWD %s", email)
			return nil, err
		}
		fmt.Println(checked)
		if checked {
			claims := util.BaseClaims{
				Email: user.Email,
				StandardClaims: jwtGo.StandardClaims{
					IssuedAt:  issueTime.Unix(),
					ExpiresAt: expirationTime.Unix(),
				},
			}
			token := jwtGo.NewWithClaims(jwtGo.SigningMethodHS256, claims)
			tokenString, err := token.SignedString(util.SecretKey)
			if err != nil {
				return nil, err
			}
			accessToken := fmt.Sprintf("%s %s", util.BasicKey, tokenString)
			return &accessToken, nil
		} else {
			log.Printf("UserService;Login;User %s has the incorrect pwd", email)
			return nil, fmt.Errorf("User;%s", appErr.ErrBadCredential)
		}
	}
	log.Printf("UserService;Login;Account %s is not exist", email)
	return nil, fmt.Errorf("User;%s", appErr.ErrObjectNotExist)
}
