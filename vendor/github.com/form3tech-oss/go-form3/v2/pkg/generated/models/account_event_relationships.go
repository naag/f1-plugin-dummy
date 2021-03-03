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

// AccountEventRelationships account event relationships
// swagger:model AccountEventRelationships
type AccountEventRelationships struct {

	// Account this event relates to
	Account *RelationshipAccount `json:"account,omitempty"`
}

func AccountEventRelationshipsWithDefaults(defaults client.Defaults) *AccountEventRelationships {
	return &AccountEventRelationships{

		Account: RelationshipAccountWithDefaults(defaults),
	}
}

func (m *AccountEventRelationships) WithAccount(account RelationshipAccount) *AccountEventRelationships {

	m.Account = &account

	return m
}

func (m *AccountEventRelationships) WithoutAccount() *AccountEventRelationships {
	m.Account = nil
	return m
}

// Validate validates this account event relationships
func (m *AccountEventRelationships) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountEventRelationships) validateAccount(formats strfmt.Registry) error {

	if swag.IsZero(m.Account) { // not required
		return nil
	}

	if m.Account != nil {
		if err := m.Account.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("account")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccountEventRelationships) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountEventRelationships) UnmarshalBinary(b []byte) error {
	var res AccountEventRelationships
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *AccountEventRelationships) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}