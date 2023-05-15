package repository

import (
	"LinkCabinet_Backend/internal/api/model"

	"gorm.io/gorm"
)

type IUserRepository interface {
	GetUserByEmail(user *model.User ,email string) error
	CreateUser(user *model.User) error
	DestroyUser(userId uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db}
}

func (ur *userRepository) GetUserByEmail(user *model.User ,email string) error {
	if err:=ur.db.Where("email = ?", email).First(user).Error;err!=nil{
		return err
	}
	return nil
}

func (ur *userRepository) CreateUser(user *model.User) error {
	if err:=ur.db.Create(user).Error;err!=nil{
		return err
	}
	return nil
}

func (ur *userRepository) DestroyUser(userId uint) error {
	if err:=ur.db.Where("id=?",userId).Delete(&model.User{}).Error;err!=nil{
		return err
	}
	return nil
}