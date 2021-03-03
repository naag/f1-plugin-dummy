// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"

	"github.com/form3tech-oss/go-form3/v2/pkg/client"
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RoleAttributes role attributes
// swagger:model RoleAttributes
type RoleAttributes struct {

	// Name of the role
	Name string `json:"name,omitempty"`

	// Unique resource ID of the parent Role
	// Format: uuid
	ParentRoleID *strfmt.UUID `json:"parent_role_id,omitempty"`
}

func RoleAttributesWithDefaults(defaults client.Defaults) *RoleAttributes {
	return &RoleAttributes{

		Name: defaults.GetString("RoleAttributes", "name"),

		ParentRoleID: defaults.GetStrfmtUUIDPtr("RoleAttributes", "parent_role_id"),
	}
}

func (m *RoleAttributes) WithName(name string) *RoleAttributes {

	m.Name = name

	return m
}

func (m *RoleAttributes) WithParentRoleID(parentRoleID strfmt.UUID) *RoleAttributes {

	m.ParentRoleID = &parentRoleID

	return m
}

func (m *RoleAttributes) WithoutParentRoleID() *RoleAttributes {
	m.ParentRoleID = nil
	return m
}

// Validate validates this role attributes
func (m *RoleAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateParentRoleID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RoleAttributes) validateParentRoleID(formats strfmt.Registry) error {

	if swag.IsZero(m.ParentRoleID) { // not required
		return nil
	}

	if err := validate.FormatOf("parent_role_id", "body", "uuid", m.ParentRoleID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RoleAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RoleAttributes) UnmarshalBinary(b []byte) error {
	var res RoleAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *RoleAttributes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
