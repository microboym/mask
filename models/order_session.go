package models

import (
	"encoding/json"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
)

// OrderSession is used by pop to map your .model.Name.Proper.Pluralize.Underscore database table to your go code.
type OrderSession struct {
	Phone string `json:"phone" db:"phone"`
	// CreatedAt time.Time `json:"created_at" db:"created_at"`
	// UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (o OrderSession) String() string {
	jo, _ := json.Marshal(o)
	return string(jo)
}

// OrderSessions is not required by pop and may be deleted
type OrderSessions []OrderSession

// String is not required by pop and may be deleted
func (o OrderSessions) String() string {
	jo, _ := json.Marshal(o)
	return string(jo)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (o *OrderSession) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (o *OrderSession) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (o *OrderSession) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
