// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// AdviceSubmittedEvent advice submitted event
// swagger:model AdviceSubmittedEvent
type AdviceSubmittedEvent struct {

	// advice
	Advice *PaymentAdvice `json:"advice,omitempty"`

	// advice submission
	AdviceSubmission *AdviceSubmission `json:"advice_submission,omitempty"`

	// payment
	Payment *Payment `json:"payment,omitempty"`
}

// Validate validates this advice submitted event
func (m *AdviceSubmittedEvent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdvice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAdviceSubmission(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePayment(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AdviceSubmittedEvent) validateAdvice(formats strfmt.Registry) error {

	if swag.IsZero(m.Advice) { // not required
		return nil
	}

	if m.Advice != nil {
		if err := m.Advice.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("advice")
			}
			return err
		}
	}

	return nil
}

func (m *AdviceSubmittedEvent) validateAdviceSubmission(formats strfmt.Registry) error {

	if swag.IsZero(m.AdviceSubmission) { // not required
		return nil
	}

	if m.AdviceSubmission != nil {
		if err := m.AdviceSubmission.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("advice_submission")
			}
			return err
		}
	}

	return nil
}

func (m *AdviceSubmittedEvent) validatePayment(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *AdviceSubmittedEvent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AdviceSubmittedEvent) UnmarshalBinary(b []byte) error {
	var res AdviceSubmittedEvent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
