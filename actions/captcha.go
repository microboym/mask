package actions

import (
	"encoding/base64"
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
	imageBase64 := base64.StdEncoding.EncodeToString(captchas[0].Image)
	c.Set("image", "data:image/jpg;base64,"+imageBase64)

	// imageSHABase64 := base64.StdEncoding.EncodeToString(captchas[0].Image)
	// c.Cookies().Set("image_sha", imageSHABase64, time.Minute*5)
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
