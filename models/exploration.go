package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

/*Exploration exploration

swagger:model Exploration
*/
type Exploration struct {
	/* created at
	 */
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	/* Serialization of the exploration query configuration.
	 */
	Data interface{} `json:"data,omitempty"`

	/* link
	 */
	Link *Link `json:"link,omitempty"`

	/* Exploration name given by user.
	 */
	Name string `json:"name,omitempty"`

	/* Latest time the exploration was updated.
	 */
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`
}

// Validate validates this exploration
func (m *Exploration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLink(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Exploration) validateLink(formats strfmt.Registry) error {

	if swag.IsZero(m.Link) { // not required
		return nil
	}

	if m.Link != nil {

		if err := m.Link.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}
