package usecase

import (
	"LinkCabinet_Backend/internal/api/model"
	"LinkCabinet_Backend/internal/api/repository"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	bycrypt "golang.org/x/crypto/bcrypt"
)

type IUserUsecase interface {
	Login(user model.User) (string, error)
	Register(user model.User) (string, error)
}

type userUsecase struct {
	ur repository.IUserRepository
}

func NewUserUsecase(ur repository.IUserRepository) IUserUsecase {
	return &userUsecase{ur}
}

func (uu *userUsecase) Login(user model.User) (string, error) {

	storeUser := model.User{}

	if err := uu.ur.GetUserByEmail(&storeUser, user.Email); err != nil {
		return "", err
	}

	err := bycrypt.CompareHashAndPassword([]byte(user.Password), []byte(user.Password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":    time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (uu *userUsecase) Register(user model.User) (string, error) {

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




