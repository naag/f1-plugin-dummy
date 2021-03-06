// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ReturnReversalAdmissionRelationships return reversal admission relationships
// swagger:model ReturnReversalAdmissionRelationships
type ReturnReversalAdmissionRelationships struct {

	// payment
	Payment *RelationshipPayments `json:"payment,omitempty"`

	// payment return
	PaymentReturn *RelationshipReturns `json:"payment_return,omitempty"`

	// payment return reversal
	PaymentReturnReversal *RelationshipReturnReversals `json:"payment_return_reversal,omitempty"`
}

// Validate validates this return reversal admission relationships
func (m *ReturnReversalAdmissionRelationships) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePayment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePaymentReturn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePaymentReturnReversal(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReturnReversalAdmissionRelationships) validatePayment(formats strfmt.Registry) error {

	if swag.IsZero(m.Payment) { // not required
		return nil
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

func (m *ReturnReversalAdmissionRelationships) validatePaymentReturn(formats strfmt.Registry) error {

	if swag.IsZero(m.PaymentReturn) { // not required
		return nil
	}

	if m.PaymentReturn != nil {
		if err := m.PaymentReturn.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("payment_return")
			}
			return err
		}
	}

	return nil
}

func (m *ReturnReversalAdmissionRelationships) validatePaymentReturnReversal(formats strfmt.Registry) error {

	if swag.IsZero(m.PaymentReturnReversal) { // not required
		return nil
	}

	if m.PaymentReturnReversal != nil {
		if err := m.PaymentReturnReversal.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("payment_return_reversal")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReturnReversalAdmissionRelationships) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReturnReversalAdmissionRelationships) UnmarshalBinary(b []byte) error {
	var res ReturnReversalAdmissionRelationships
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
