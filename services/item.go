package services

import (
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

func (i *ItemService) GetItemsByOrderID(orderId int) (*[]models.Item, *params.Response) {
	response, err := i.itemRepo.GetItemsByOrderID(orderId)
	if err != nil {
		return response, &params.Response{
			Status:         http.StatusNotFound,
			Error:          "Error - Item Not Found",
			AdditionalInfo: err.Error(),
		}
	}

	return response, &params.Response{
		Status:  http.StatusOK,
		Message: "Success",
		Payload: response,
	}
}

func (i *ItemService) UpdateItemByID(itemModel *[]models.Item, request params.CreateOrder) *params.Response {

	items := request.Items

	for _, v := range *itemModel {
		for _, itemRequest := range items {
			if v.ID == uint(itemRequest.ItemID) {
				updateItem := models.Item{
					ItemCode:    itemRequest.ItemCode,
					Description: itemRequest.Description,
					Quantity:    itemRequest.Quantity,
				}

				_, err := i.itemRepo.UpdateItemByID(itemRequest.ItemID, &updateItem)

				if err != nil {
					return &params.Response{
						Status:         400,
						Error:          "BAD REQUEST",
						AdditionalInfo: err,
					}
				}
			}
		}
	}

	return &params.Response{
		Status:  200,
		Message: "Success - Update Order & Items",
	}
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
