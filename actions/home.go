package actions

import (
	"mask/models"

	"github.com/gobuffalo/buffalo"
)

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	orderCount, err := models.DB.Count(models.Orders{})
	if err != nil {
		return err
	}
	captchaCount, err := models.DB.Where("answer = null").Count(models.Captchas{})
	if err != nil {
		return err
	}
	c.Set("orderCount", orderCount)
	c.Set("captchaCount", captchaCount)
	return c.Render(200, r.HTML("index.html"))
}

// func loadUsers() (models.Users, error) {
// 	users := models.Users{}
// 	err := models.DB.All(&users)
// 	return users, err
// }
