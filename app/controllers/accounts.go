package controllers

import (
	"fmt"
	"restaurant-site/app/models"

	"github.com/revel/revel"
)

type Accounts struct {
	*revel.Controller
}

func (c Accounts) Create() revel.Result {
	return c.Render()
}

func (c Accounts) CreatePost() revel.Result {
	var account models.Account
	c.Params.Bind(&account, "account")
	fmt.Printf("Account info: %v\n", account)

	return c.RenderTemplate("accounts/create.html")
}
