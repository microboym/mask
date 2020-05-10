package actions

import (
	"mask/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/nulls"
)

// CaptchaDisplay default implementation.
func CaptchaDisplay(c buffalo.Context) error {
	captchas := models.Captchas{}
	if err := models.DB.All(&captchas); err != nil {
		return err
	}
	c.Set("captcha", captchas[0])
	return c.Render(200, r.HTML("captcha/display.html"))
}

// CaptchaUploadAnswer default implementation.
func CaptchaUploadAnswer(c buffalo.Context) error {
	answer := c.Params().Get("Answer")
	imageSHA := c.Params().Get("ImageSHA")

	// imageSHA, err := base64.StdEncoding.DecodeString(imageSHABase64)
	// if err != nil {
	// 	return err
	// }

	captcha := models.Captcha{}
	if err := models.DB.Find(captcha, imageSHA); err != nil {
		return err
	}
	captcha.Answer = nulls.NewString(answer)
	_, err := models.DB.ValidateAndSave(captcha)
	return err
}
