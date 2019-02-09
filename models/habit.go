package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

// Habit is something you are trying to make a regular occurrence
type Habit struct {
	ID          uuid.UUID `json:"id" db:"id"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	Timeframe   int       `json:"timeframe" db:"timeframe"`

	UserID uuid.UUID `json:"user" db:"user_id"`
	User   User      `belongs_to:"user"`
}

// String is not required by pop and may be deleted
func (h Habit) String() string {
	jh, _ := json.Marshal(h)
	return string(jh)
}

// Habits is not required by pop and may be deleted
type Habits []Habit

// String is not required by pop and may be deleted
func (h Habits) String() string {
	jh, _ := json.Marshal(h)
	return string(jh)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (h *Habit) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: h.Name, Name: "Name"},
		&validators.StringIsPresent{Field: h.Description, Name: "Description"},
		&validators.IntIsPresent{Field: h.Timeframe, Name: "Timeframe"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (h *Habit) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (h *Habit) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
