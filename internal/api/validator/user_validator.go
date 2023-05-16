package validator

import (
	"LinkCabinet_Backend/internal/api/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type IUserValidator interface {
	UserValidate(user model.User) error
}

type userValidator struct {
}

func NewUserValidator() IUserValidator {
	return &userValidator{}
}

func (uv *userValidator) UserValidate(user model.User) error {

	return validation.ValidateStruct(&user,
		validation.Field(
			&user.Name,
			validation.Required.Error("name is required"),
			validation.RuneLength(1,64).Error("limited 1-64 characters"),
		),
		validation.Field(
			&user.Email,
			validation.Required.Error("email is required"),
			validation.RuneLength(1,128).Error("limited 1-128 characters"),
			is.Email.Error("invalid email"),
		),
		validation.Field(
			&user.Password,
			validation.Required.Error("password is required"),
			validation.RuneLength(6,128).Error("limited 6-128 characters"),
		),
	)
	
}
