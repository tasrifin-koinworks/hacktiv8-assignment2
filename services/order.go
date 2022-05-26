package services

import (
	"hacktiv8-assignment2/models"
	"hacktiv8-assignment2/params"
	"hacktiv8-assignment2/repositories"
	"net/http"
)

type OrderService struct {
	orderRepo repositories.OrderRepo
}

func NewOrderService(repo repositories.OrderRepo) *OrderService {
	return &OrderService{
		orderRepo: repo,
	}
}

func (o *OrderService) CreateOrder(request params.CreateOrder) *params.Response {

	if request.CustomerName == "" {
		return &params.Response{
			Status: 500,
			Error:  "DATA EMPTY",
			AdditionalInfo: map[string]string{
				"message": "Customer Name can't be null",
			},
		}
	}

	orderModel := models.Order{
		CustomerName: request.CustomerName,
	}

	storeOrder, err := o.orderRepo.CreateOrder(&orderModel)

	if err != nil {
		return &params.Response{
			Status:         400,
			Error:          "BAD REQUEST",
			AdditionalInfo: err,
		}
	}

	return &params.Response{
		Status:  200,
		Message: "CREATE SUKSES",
		Payload: storeOrder,
	}
}

func (o *OrderService) GetOrderByIDWithItems(orderId int) *params.Response {
	response, err := o.orderRepo.GetOrderByIDWithItems(orderId)

	if err != nil {
		return &params.Response{
			Status:  400,
			Error:   "BAD REQUEST",
			Payload: err,
		}
	}

	return &params.Response{
		Status:  200,
		Message: "Success - Order with Items",
		Payload: response,
	}

}

func (o *OrderService) GetAllOrdersWithItems() *params.Response {
	response, err := o.orderRepo.GetAllOrdersWithItems()

	if err != nil {
		return &params.Response{
			Status:  400,
			Error:   "BAD REQUEST",
			Payload: err,
		}
	}

	return &params.Response{
		Status:  200,
		Message: "Success - Get All Orders with Items",
		Payload: response,
	}

}

func (o *OrderService) GetOrderByID(orderId int) *params.Response {
	response, err := o.orderRepo.GetOrderByID(orderId)
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

func (o *OrderService) UpdateOrderByID(orderId int, request params.CreateOrder) *params.Response {
	_, err := o.orderRepo.GetOrderByID(orderId)
	if err != nil {
		return &params.Response{
			Status:         http.StatusNotFound,
			Error:          "Error - Item Not Found",
			AdditionalInfo: err.Error(),
		}
	}

	if request.CustomerName == "" {
		return &params.Response{
			Status: 500,
			Error:  "DATA EMPTY",
			AdditionalInfo: map[string]string{
				"message": "Customer Name can't be null",
			},
		}
	}

	orderModel := models.Order{
		CustomerName: request.CustomerName,
	}

	storeOrder, err := o.orderRepo.UpdateOrderByID(orderId, &orderModel)

	if err != nil {
		return &params.Response{
			Status:         400,
			Error:          "BAD REQUEST",
			AdditionalInfo: err.Error(),
		}
	}

	return &params.Response{
		Status:  200,
		Message: "Success - Update Order",
		Payload: storeOrder,
	}
}

func (o *OrderService) DeleteOrder(orderId int) *params.Response {
	err := o.orderRepo.DeleteOrder(orderId)

	if err != nil {
		return &params.Response{
			Status:         http.StatusBadRequest,
			Error:          "BAD REQUEST",
			AdditionalInfo: err,
		}
	}

	return &params.Response{
		Status:  http.StatusOK,
		Message: "Success - Delete Order & Items",
	}
}
