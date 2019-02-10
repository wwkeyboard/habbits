package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/pkg/errors"
	"github.com/wwkeyboard/habbits/models"
)

// ChecksResource is the resource for the Check model
type ChecksResource struct {
	buffalo.Resource
}

// Create adds a Check to the DB. This function is mapped to the
// path POST /habit/{id}/checks
// TODO: pull habit id from URL
// TODO: sort out checks-per-day logic
func (v ChecksResource) Create(c buffalo.Context) error {
	// Allocate an empty Check
	check := &models.Check{}

	// Bind check to the html form elements
	if err := c.Bind(check); err != nil {
		return errors.WithStack(err)
	}

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Validate the data from the html form
	verrs, err := tx.ValidateAndCreate(check)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the new.html template that the user can
		// correct the input.
		return c.Render(422, r.Auto(c, check))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", "Check was created successfully")

	// and redirect to the checks index page
	return c.Render(201, r.Auto(c, check))
}

// Destroy deletes a Check from the DB. This function is mapped
// to the path DELETE /checks/{check_id}
func (v ChecksResource) Destroy(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Check
	check := &models.Check{}

	// To find the Check the parameter check_id is used.
	if err := tx.Find(check, c.Param("check_id")); err != nil {
		return c.Error(404, err)
	}

	if err := tx.Destroy(check); err != nil {
		return errors.WithStack(err)
	}

	// If there are no errors set a flash message
	c.Flash().Add("success", "Check was destroyed successfully")

	// Redirect to the checks index page
	return c.Render(200, r.Auto(c, check))
}
