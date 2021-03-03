// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ReceiversCorrespondentAccountHoldingEntity receivers correspondent account holding entity
// swagger:model ReceiversCorrespondentAccountHoldingEntity
type ReceiversCorrespondentAccountHoldingEntity struct {

	// Receiver's correspondent's address
	BankAddress []string `json:"bank_address,omitempty"`

	// SWIFT BIC for receiver's correspondent
	BankID string `json:"bank_id,omitempty"`

	// bank id code
	BankIDCode BankIDCode `json:"bank_id_code,omitempty"`

	// Receiver's correspondent's name
	BankName string `json:"bank_name,omitempty"`

	// Receiver's correspondent party identifier
	BankPartyID string `json:"bank_party_id,omitempty"`
}

// Validate validates this receivers correspondent account holding entity
func (m *ReceiversCorrespondentAccountHoldingEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBankIDCode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReceiversCorrespondentAccountHoldingEntity) validateBankIDCode(formats strfmt.Registry) error {

	if swag.IsZero(m.BankIDCode) { // not required
		return nil
	}

	if err := m.BankIDCode.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("bank_id_code")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReceiversCorrespondentAccountHoldingEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReceiversCorrespondentAccountHoldingEntity) UnmarshalBinary(b []byte) error {
	var res ReceiversCorrespondentAccountHoldingEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
