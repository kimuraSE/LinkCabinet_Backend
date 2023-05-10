package usecase

import (
	"LinkCabinet_Backend/internal/api/model"
	"LinkCabinet_Backend/internal/api/repository"
	"LinkCabinet_Backend/internal/api/validator"
)

type ILinksUsecase interface {
	AllGetLinks(userId uint) ([]model.LinkResponse, error)
	GetLinksByUserID(userID uint,linkId uint) (model.LinkResponse, error)
	CreateLink(link model.Link) (model.LinkResponse, error)
	UpdateLink(link model.Link,userId uint,linkId uint) (model.LinkResponse, error)
	DeleteLink(userId uint,linkId uint) error
}

type linksUsecase struct {
	lr repository.ILinksRepository
	lv validator.ILinksValidator
}

func NewLinksUsecase(lr repository.ILinksRepository,lv validator.ILinksValidator) ILinksUsecase {
	return &linksUsecase{lr,lv}
}

func (lu *linksUsecase) AllGetLinks(userId uint) ([]model.LinkResponse, error) {
	links := []model.Link{}
	if err := lu.lr.AllGetLinks(&links,userId); err != nil {
		return nil, err
	}
	resLinks := []model.LinkResponse{}
	for _, link := range links {
		v := model.LinkResponse{
			ID:          link.ID,
			Title: 	 link.Title,
			Url: 	 link.Url,
		}
		resLinks = append(resLinks, v)
	}
	return resLinks, nil
}

func (lu *linksUsecase) GetLinksByUserID(userId uint,linkId uint) (model.LinkResponse, error) {
	link := model.Link{}
	if err := lu.lr.GetLinkById(&link,userId,linkId); err != nil {
		return model.LinkResponse{}, err
	}
	resLink := model.LinkResponse{
		ID:          link.ID,
		Title: 	 link.Title,
		Url: 	 link.Url,
	}
	return resLink, nil
}

func (lu *linksUsecase) CreateLink(link model.Link) (model.LinkResponse, error) {
	
	if err := lu.lv.LinksValidator(link); err != nil {
		return model.LinkResponse{}, err
	}

	
	if err := lu.lr.CreateLink(&link); err != nil {
		return model.LinkResponse{}, err
	}
	resLink := model.LinkResponse{
		ID:          link.ID,
		Title: 	 link.Title,
		Url: 	 link.Url,
	}
	return resLink, nil
}

func (lu *linksUsecase) UpdateLink(link model.Link,userId uint,linkId uint) (model.LinkResponse, error) {
	
	if err:= lu.lv.LinksValidator(link); err != nil {
		return model.LinkResponse{}, err
	}
	

	if err := lu.lr.UpdateLink(&link,userId,linkId); err != nil {
		return model.LinkResponse{}, err
	}
	resLink := model.LinkResponse{
		ID:          link.ID,
		Title: 	 link.Title,
		Url: 	 link.Url,
	}
	return resLink, nil
}

func (lu *linksUsecase) DeleteLink(userId uint,linkId uint) error {
	if err := lu.lr.DeleteLink(userId,linkId); err != nil {
		return err
	}
	return nil
}