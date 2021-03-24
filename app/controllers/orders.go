package controllers

import (
	"fmt"
	"restaurant-site/app/models"

	"github.com/revel/revel"
)

type Orders struct {
	*revel.Controller
}

func (c Orders) Create() revel.Result {
	return c.Render()
}

func (c Orders) GetPayment(orderId int) revel.Result {
	println("The order ID: ", orderId)
	return c.RenderTemplate("orders/payment.html")
}

func (c Orders) ApiCreate() revel.Result {
	order := new(models.Order)
	c.Params.BindJSON(&order)
	fmt.Printf("The order data: %v\n", order)
	return c.RenderText("OK")

}
