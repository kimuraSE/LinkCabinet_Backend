package validator

import (
	"LinkCabinet_Backend/internal/api/model"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type ILinksValidator interface {
	LinksValidator(links model.Link) error
}

type linksValidator struct {
}

func NewLinksValidator() ILinksValidator {
	return &linksValidator{}
}

func (lv *linksValidator) LinksValidator(links model.Link) error {
	return validation.ValidateStruct(&links,
		validation.Field(&links.Title, validation.Required.Error("title is required"),
			validation.RuneLength(1, 256).Error("limitted 1-256 characters")),
		validation.Field(&links.Url, validation.Required.Error("url is required"),
			validation.RuneLength(1, 2048).Error("limitted 1-2048 characters")),
	)
}
