package actions

import (
	"mask/models"

	"github.com/gobuffalo/buffalo"
)

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	users, err := loadUsers()
	if err != nil {
		return err
	}
	c.Set("users", users)
	return c.Render(200, r.HTML("index.html"))
}

func loadUsers() (models.Users, error) {
	users := models.Users{}
	err := models.DB.All(&users)
	return users, err
}
