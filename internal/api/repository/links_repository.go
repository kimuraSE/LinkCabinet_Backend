package repository

import (
	"LinkCabinet_Backend/internal/api/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ILinksRepository interface {
	AllGetLinks(links *[]model.Link,userId uint) error
	GetLinkById(link *model.Link,userId uint ,linkId uint) error
	CreateLink(link *model.Link) error
	UpdateLink(link *model.Link,userId uint, linkId uint) error
	DeleteLink(userId uint, linkId uint) error
}


type linksRepository struct {
	db *gorm.DB
}

func NewLinksRepository(db *gorm.DB) ILinksRepository {
	return &linksRepository{db}
}

func (lr *linksRepository) AllGetLinks(links *[]model.Link,userId uint) error {
	if err := lr.db.Table("users").Select("links.id , links.title , links.url").Joins("JOIN links ON users.id = links.user_id").Where("links.user_id = ?",userId).Find(links).Error; err != nil {
		return err
	}
	return nil
}

func (lr *linksRepository) GetLinkById(link *model.Link,userId uint ,linkId uint) error {
	if err := lr.db.Table("users").Select("links.id , links.title , links.url").Joins("JOIN links ON users.id = links.user_id").Where("links.user_id = ? AND links.id=?",userId,linkId).First(link).Error; err != nil {
		return err
	}
	return nil
}

func (lr *linksRepository) CreateLink(link *model.Link) error {
	if err := lr.db.Create(link).Error; err != nil {
		return err
	}
	return nil
}

func (lr *linksRepository) UpdateLink(link *model.Link,userId uint, linkId uint) error {
	if err := lr.db.Model(link).Clauses(clause.Returning{}).Where("id=? AND user_id=?",linkId,userId).Updates(
		map[string]interface{}{
		"title": link.Title,
		"url": link.Url,
		},
	).Error; err != nil {
		return err
	}
	return nil
}

func (lr *linksRepository) DeleteLink(userId uint, linkId uint) error {
	if err := lr.db.Where("user_id = ? AND id = ?",userId,linkId).Delete(&model.Link{}).Error; err != nil {
		return err
	}
	return nil
}

