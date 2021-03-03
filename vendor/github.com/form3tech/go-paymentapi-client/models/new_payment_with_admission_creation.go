// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NewPaymentWithAdmissionCreation new payment with admission creation
// swagger:model NewPaymentWithAdmissionCreation
type NewPaymentWithAdmissionCreation struct {

	// admission
	// Required: true
	Admission *NewPaymentAdmission `json:"admission"`

	// payment
	// Required: true
	Payment *PaymentCreation `json:"payment"`

	// Name of the resource type
	// Pattern: ^[A-Za-z_]*$
	Type string `json:"type,omitempty"`
}

// Validate validates this new payment with admission creation
func (m *NewPaymentWithAdmissionCreation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdmission(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePayment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewPaymentWithAdmissionCreation) validateAdmission(formats strfmt.Registry) error {

	if err := validate.Required("admission", "body", m.Admission); err != nil {
		return err
	}

	if m.Admission != nil {
		if err := m.Admission.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("admission")
			}
			return err
		}
	}

	return nil
}

func (m *NewPaymentWithAdmissionCreation) validatePayment(formats strfmt.Registry) error {

	if err := validate.Required("payment", "body", m.Payment); err != nil {
		return err
	}

	if m.Payment != nil {
		if err := m.Payment.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("payment")
			}
			return err
		}
	}

	return nil
}

func (m *NewPaymentWithAdmissionCreation) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := validate.Pattern("type", "body", string(m.Type), `^[A-Za-z_]*$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NewPaymentWithAdmissionCreation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewPaymentWithAdmissionCreation) UnmarshalBinary(b []byte) error {
	var res NewPaymentWithAdmissionCreation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}