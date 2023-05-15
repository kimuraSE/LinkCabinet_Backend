package repository

import (
	"LinkCabinet_Backend/internal/api/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IUserRepository interface {
	GetUserByEmail(user *model.User ,email string) error
	CreateUser(user *model.User) error
	DestroyUser(userId uint) error
	UpdateUserName(user *model.User,userId uint) error
	UpdateUserEmail(user *model.User,userId uint) error
	
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

func (ur *userRepository) UpdateUserName(user *model.User,userId uint) error {
	result:=ur.db.Model(user).Clauses(clause.Returning{}).Where("id=?",userId).Update("name",user.Name)
	if result.Error != nil{
		return result.Error
	}
	if result.RowsAffected < 1{
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (ur *userRepository) UpdateUserEmail(user *model.User,userId uint) error {
	result:=ur.db.Model(user).Clauses(clause.Returning{}).Where("id=?",userId).Update("email",user.Email)
	if result.Error != nil{
		return result.Error
	}
	if result.RowsAffected < 1{
		return gorm.ErrRecordNotFound
	}
	return nil
}