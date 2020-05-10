package actions

import (
	"mask/models"

	"github.com/gobuffalo/buffalo"
)

// OrdersDisplay default implementation.
func OrdersDisplay(c buffalo.Context) error {
	orders := models.Orders{}
	models.DB.All(orders)
	print(orders)
	c.Set("orders", orders)
	return c.Render(200, r.HTML("orders/display.html"))
}

// OrdersImage default implementation.
func OrdersImage(c buffalo.Context) error {
	return c.Render(200, r.HTML("orders/image.html"))
}
