package usecase

import (
	"LinkCabinet_Backend/internal/api/model"
	"LinkCabinet_Backend/internal/api/repository"
	"LinkCabinet_Backend/internal/api/validator"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	bycrypt "golang.org/x/crypto/bcrypt"
)

type IUserUsecase interface {
	Login(user model.User) (string, error)
	Register(user model.User) (string, error)
	DeleteUser(userId uint) error
	UpdateUserName(user model.User, userId uint) error
	UpdateUserEmail(user model.User, userId uint) error
	UpdateUserPassword(user model.User, userId uint) error
}

type userUsecase struct {
	ur repository.IUserRepository
	uv validator.IUserValidator
}

func NewUserUsecase(ur repository.IUserRepository,uv validator.IUserValidator) IUserUsecase {
	return &userUsecase{ur,uv}
}

func (uu *userUsecase) Login(user model.User) (string, error) {

	storeUser := model.User{}

	if err := uu.ur.GetUserByEmail(&storeUser, user.Email); err != nil {
		return "", err
	}

	err := bycrypt.CompareHashAndPassword([]byte(storeUser.Password), []byte(user.Password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": storeUser.ID,
		"exp":    time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (uu *userUsecase) Register(user model.User) (string, error) {

	if err:=uu.uv.UserValidate(user);err!=nil{
		return "",err
	}

	hash,err := bycrypt.GenerateFromPassword([]byte(user.Password),10)
	if err!=nil{
		return "",err
	}

	newUser := model.User{
		Name: user.Name,
		Email: user.Email,
		Password: string(hash),
	}

	if err:=uu.ur.CreateUser(&newUser);err!=nil{
		return "",err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": newUser.ID,
		"exp":    time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (uu *userUsecase) DeleteUser(userId uint) error {
	if err:=uu.ur.DestroyUser(userId);err!=nil{
		return err
	}
	return nil
}

func (uu *userUsecase) UpdateUserName(user model.User, userId uint) error {
	if err:=uu.ur.UpdateUserName(&user,userId);err!=nil{
		return err
	}
	return nil
}

func (uu *userUsecase) UpdateUserEmail(user model.User, userId uint) error {
	if err:=uu.ur.UpdateUserEmail(&user,userId);err!=nil{
		return err
	}
	return nil
}

func (uu *userUsecase) UpdateUserPassword(user model.User, userId uint) error {
	
	hash,err := bycrypt.GenerateFromPassword([]byte(user.Password),10)
	if err!=nil{
		return err
	}

	newUser := model.User{
		Password: string(hash),
	}

	if err:=uu.ur.UpdateUserPassword(&newUser,userId);err!=nil{
		return err
	}
	return nil
}

