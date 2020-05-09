package models

import (
	"encoding/json"

	"github.com/gobuffalo/nulls"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
)

// Captcha is used by pop to map your .model.Name.Proper.Pluralize.Underscore database table to your go code.
type Captcha struct {
	ID     []byte       `json:"image_sha" db:"img_sha"`
	Image  []byte       `db:"img"`
	Answer nulls.String `json:"answer" db:"answer"`
}

// String is not required by pop and may be deleted
func (c Captcha) String() string {
	jc, _ := json.Marshal(c)
	return string(jc)
}

// Captchas is not required by pop and may be deleted
type Captchas []Captcha

// String is not required by pop and may be deleted
func (c Captchas) String() string {
	jc, _ := json.Marshal(c)
	return string(jc)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (c *Captcha) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (c *Captcha) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (c *Captcha) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
