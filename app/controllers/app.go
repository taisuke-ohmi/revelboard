package controllers

import (
	"myapp/app/models"

	"github.com/revel/revel"
)

type App struct {
	ApiV1Controller
}

func (c App) Index() revel.Result {
	comments := []models.Comment{}

	if err := DB.Order("created_at").Find(&comments).Error; err != nil {
		return c.HandleInternalServerError("Record Find Failure")
	}

	g := "Hello World"
	return c.Render(g, comments)
}

func (c App) PostMessage(myName string, body string) revel.Result {
	// validation check
	c.Validation.Required(myName).Message("Your name is required")
	c.Validation.MinSize(myName, 3).Message("Your name is not long enough")
	c.Validation.Required(body).Message("Body is required")

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(App.Index)
	}

	// DB insert
	comment := &models.Comment{Nickname: myName, Body: body}

	if err := DB.Create(comment).Error; err != nil {
		return c.HandleInternalServerError("Record Create Failure")
	}

	return c.Redirect(App.Index)
}
