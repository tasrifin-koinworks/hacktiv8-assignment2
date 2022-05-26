package services

import (
	"fmt"
	"hacktiv8-assignment2/models"
	"hacktiv8-assignment2/params"
	"hacktiv8-assignment2/repositories"
	"net/http"
)

type ItemService struct {
	itemRepo repositories.ItemRepo
}

func NewItemService(repo repositories.ItemRepo) *ItemService {
	return &ItemService{
		itemRepo: repo,
	}
}

var createdItems []params.ItemResponse

func (i *ItemService) CreateItem(responseOrder params.Response, request params.CreateOrder) *params.Response {
	createdItems = nil

	orderData := responseOrder.Payload
	order, err := orderData.(*models.Order)

	if !err {
		return &params.Response{
			Status:         400,
			Error:          "BAD REQUEST",
			AdditionalInfo: err,
		}
	}

	items := request.Items

	for _, item := range items {
		itemModel := models.Item{
			ItemCode:    item.ItemCode,
			Description: item.Description,
			Quantity:    item.Quantity,
			OrderId:     uint(order.ID),
		}

		fmt.Println(&itemModel)
		itemData, err := i.itemRepo.CreateItem(&itemModel)

		if err != nil {
			return &params.Response{
				Status:         400,
				Error:          "BAD REQUEST",
				AdditionalInfo: err,
			}
		}

		createdItems = append(createdItems, params.ItemResponse{ItemID: int(itemData.ID),
			ItemCode:    itemData.ItemCode,
			Description: itemData.Description,
			Quantity:    itemData.Quantity,
			OrderID:     order.ID})
	}

	return &params.Response{
		Status:  200,
		Message: "Success - Create Order & Items",
		AdditionalInfo: params.AllResponseData{
			OrderID:      order.ID,
			OrderedAt:    order.CreatedAt,
			CustomerName: order.CustomerName,
			Items:        createdItems,
		},
	}
}

func (i *ItemService) GetItemsByOrderID(orderId int) *params.Response {
	response, err := i.itemRepo.GetItemsByOrderID(orderId)
	fmt.Println(err)
	if err != nil {
		return &params.Response{
			Status:         http.StatusNotFound,
			Error:          "Error - Item Not Found",
			AdditionalInfo: err.Error(),
		}
	}

	return &params.Response{
		Status:  http.StatusOK,
		Message: "Success",
		Payload: response,
	}
}

func (i *ItemService) UpdateItemByID(orderId int, request params.AllResponseData) *params.Response {
	// _, err := i.ItemService.GetOrderByID(orderId)
	// fmt.Println(err)
	// if err != nil {
	return &params.Response{
		Status: http.StatusNotFound,
		Error:  "Error - Item Not Found",
	}
	// }
}

func (i *ItemService) DeleteItems(orderId int) *params.Response {
	err := i.itemRepo.DeleteItem(orderId)

	if err != nil {
		return &params.Response{
			Status:         http.StatusBadRequest,
			Error:          "BAD REQUEST",
			AdditionalInfo: err,
		}
	}

	return &params.Response{
		Status:  http.StatusOK,
		Message: "Success - Delete Items",
	}
}
