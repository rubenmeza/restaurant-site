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
	c.Validation.Required(account.FristName)
	c.Validation.Required(account.LastName)
	c.Validation.Required(account.Address1)
	c.Validation.Required(account.City)
	c.Validation.Required(account.State)
	c.Validation.Required(account.ZipCode)
	c.Validation.Length(account.ZipCode, 5)

	fmt.Printf("Has error: %v\n", c.Validation.HasErrors())

	fmt.Printf("Account info: %v\n", account)

	return c.RenderTemplate("accounts/create.html")
}
