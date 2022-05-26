package repositories

import (
	"hacktiv8-assignment2/models"

	"github.com/jinzhu/gorm"
)

type ItemRepo interface {
	CreateItem(item *models.Item) (*models.Item, error)
	GetItemsByOrderID(orderId int) (*[]models.Item, error)
	UpdateItemByID(id int, item *models.Item) (*models.Item, error)
	DeleteItem(orderId int) error
}

type itemRepo struct {
	db *gorm.DB
}

func NewItemRepo(db *gorm.DB) ItemRepo {
	return &itemRepo{db}
}

func (i *itemRepo) CreateItem(item *models.Item) (*models.Item, error) {
	return item, i.db.Create(item).Error
}

func (i *itemRepo) GetItemsByOrderID(orderId int) (*[]models.Item, error) {
	var items []models.Item

	err := i.db.Where("order_id=?", orderId).Find(&items).Error
	return &items, err
}

func (i *itemRepo) UpdateItemByID(id int, updateItem *models.Item) (*models.Item, error) {
	var item models.Item

	err := i.db.Model(&item).Where("id=?", id).Updates(updateItem).Error
	return &item, err
}

func (i *itemRepo) DeleteItem(orderId int) error {
	var item models.Item

	err := i.db.Where("order_id=?", orderId).Delete(&item).Error
	return err
}
