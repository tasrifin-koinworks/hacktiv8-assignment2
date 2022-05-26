package controllers

import (
	"hacktiv8-assignment2/params"
	"hacktiv8-assignment2/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	itemService  services.ItemService
	orderService services.OrderService
}

func NewOrderController(service *services.OrderService, service2 *services.ItemService) *OrderController {
	return &OrderController{
		orderService: *service,
		itemService:  *service2,
	}
}

func (o *OrderController) CreateNewOrder(c *gin.Context) {
	var req params.CreateOrder

	err := c.ShouldBind(&req)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
			Status:         http.StatusBadRequest,
			Error:          "BAD REQUEST",
			AdditionalInfo: err,
		})

		return
	}

	if req.Items == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
			Status:         http.StatusBadRequest,
			Error:          "BAD REQUEST",
			AdditionalInfo: "Items Cannot be Null",
		})

		return
	}

	response := o.orderService.CreateOrder(req)

	if response.Status != 200 {
		c.JSON(response.Status, response)
		return
	}

	response2 := o.itemService.CreateItem(*response, req)

	c.JSON(response2.Status, response2)

}

func (o *OrderController) GetAllOrdersWithItems(c *gin.Context) {
	response := o.orderService.GetAllOrdersWithItems()

	c.JSON(response.Status, response)
}

func (o *OrderController) UpdateOrder(c *gin.Context) {
	var req params.CreateOrder
	orderId, err := strconv.Atoi(c.Param("orderId"))

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
			Status:         http.StatusBadRequest,
			Error:          "BAD REQUEST",
			AdditionalInfo: err,
		})

		return
	}

	checkOrderResponse := o.orderService.GetOrderByID(orderId)

	if checkOrderResponse.Status != http.StatusOK {
		c.JSON(checkOrderResponse.Status, checkOrderResponse)
		return
	}

	isErr := c.ShouldBind(&req)

	if isErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
			Status:         http.StatusBadRequest,
			Error:          "BAD REQUEST",
			AdditionalInfo: isErr.Error(),
		})

		return
	}

	updateOrderResponse := o.orderService.UpdateOrderByID(orderId, req)

	if updateOrderResponse.Status != http.StatusOK {
		c.JSON(updateOrderResponse.Status, updateOrderResponse)
		return
	}

	getItemById, _ := o.itemService.GetItemsByOrderID(orderId)

	updateItemResponse := o.itemService.UpdateItemByID(getItemById, req)
	if updateItemResponse.Status != http.StatusOK {
		c.JSON(updateOrderResponse.Status, updateItemResponse)
		return
	}

	getOrderDetail := o.orderService.GetOrderByIDWithItems(orderId)

	c.JSON(updateOrderResponse.Status, getOrderDetail)

}

func (o *OrderController) DeleteOrder(c *gin.Context) {
	orderId, err := strconv.Atoi(c.Param("orderId"))

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
			Status:         http.StatusBadRequest,
			Error:          "BAD REQUEST",
			AdditionalInfo: err,
		})

		return
	}

	checkOrderResponse := o.orderService.GetOrderByID(orderId)

	if checkOrderResponse.Status != http.StatusOK {
		c.JSON(checkOrderResponse.Status, checkOrderResponse)
		return
	}

	deleteItemsResponse := o.itemService.DeleteItems(orderId)

	if deleteItemsResponse.Status != http.StatusOK {
		c.JSON(deleteItemsResponse.Status, deleteItemsResponse)
		return
	}

	deleteOrderResponse := o.orderService.DeleteOrder(orderId)
	c.JSON(deleteOrderResponse.Status, deleteOrderResponse)

}
