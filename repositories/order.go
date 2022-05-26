package repositories

import (
	"hacktiv8-assignment2/models"

	"github.com/jinzhu/gorm"
)

type OrderRepo interface {
	CreateOrder(order *models.Order) (*models.Order, error)
	GetOrderByID(orderId int) (*models.Order, error)
	GetAllOrdersWithItems() (*[]models.Order, error)
	GetOrderByIDWithItems(orderId int) (*[]models.Order, error)
	UpdateOrderByID(orderId int, order *models.Order) (*models.Order, error)
	DeleteOrder(orderId int) error
}

type orderRepo struct {
	db *gorm.DB
}

func NewOrderRepo(db *gorm.DB) OrderRepo {
	return &orderRepo{db}
}

func (o *orderRepo) CreateOrder(order *models.Order) (*models.Order, error) {
	return order, o.db.Create(order).Error
}

func (o *orderRepo) GetOrderByIDWithItems(orderId int) (*[]models.Order, error) {
	var order []models.Order
	err := o.db.Preload("Items").Where("id=?", orderId).Find(&order).Error
	return &order, err
}

func (o *orderRepo) GetAllOrdersWithItems() (*[]models.Order, error) {
	var order []models.Order
	err := o.db.Preload("Items").Find(&order).Error
	return &order, err
}

func (o *orderRepo) GetOrderByID(orderId int) (*models.Order, error) {
	var order models.Order

	err := o.db.Preload("Items").First(&order, "id=?", orderId).Error
	return &order, err
}

func (o *orderRepo) UpdateOrderByID(orderId int, updateOrder *models.Order) (*models.Order, error) {
	var order models.Order

	err := o.db.Model(&order).Where("id=?", orderId).Updates(updateOrder).Error
	return &order, err
}

func (o *orderRepo) DeleteOrder(orderId int) error {
	var order models.Order

	err := o.db.Where("id=?", orderId).Delete(&order).Error
	return err
}
