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
)

// ContactAccountRelationships contact account relationships
// swagger:model ContactAccountRelationships
type ContactAccountRelationships struct {

	// contact
	Contact *RelationshipsContactRef `json:"contact,omitempty"`

	// party
	Party *RelationshipsPartyRef `json:"party,omitempty"`
}

func ContactAccountRelationshipsWithDefaults(defaults client.Defaults) *ContactAccountRelationships {
	return &ContactAccountRelationships{

		Contact: RelationshipsContactRefWithDefaults(defaults),

		Party: RelationshipsPartyRefWithDefaults(defaults),
	}
}

func (m *ContactAccountRelationships) WithContact(contact RelationshipsContactRef) *ContactAccountRelationships {

	m.Contact = &contact

	return m
}

func (m *ContactAccountRelationships) WithoutContact() *ContactAccountRelationships {
	m.Contact = nil
	return m
}

func (m *ContactAccountRelationships) WithParty(party RelationshipsPartyRef) *ContactAccountRelationships {

	m.Party = &party

	return m
}

func (m *ContactAccountRelationships) WithoutParty() *ContactAccountRelationships {
	m.Party = nil
	return m
}

// Validate validates this contact account relationships
func (m *ContactAccountRelationships) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContact(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParty(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContactAccountRelationships) validateContact(formats strfmt.Registry) error {

	if swag.IsZero(m.Contact) { // not required
		return nil
	}

	if m.Contact != nil {
		if err := m.Contact.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("contact")
			}
			return err
		}
	}

	return nil
}

func (m *ContactAccountRelationships) validateParty(formats strfmt.Registry) error {

	if swag.IsZero(m.Party) { // not required
		return nil
	}

	if m.Party != nil {
		if err := m.Party.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("party")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContactAccountRelationships) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContactAccountRelationships) UnmarshalBinary(b []byte) error {
	var res ContactAccountRelationships
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *ContactAccountRelationships) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
