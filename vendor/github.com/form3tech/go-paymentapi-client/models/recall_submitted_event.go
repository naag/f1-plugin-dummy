// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// RecallSubmittedEvent recall submitted event
// swagger:model RecallSubmittedEvent
type RecallSubmittedEvent struct {

	// payment
	Payment *Payment `json:"payment,omitempty"`

	// recall
	Recall *Recall `json:"recall,omitempty"`

	// recall submission
	RecallSubmission *RecallSubmission `json:"recall_submission,omitempty"`
}

// Validate validates this recall submitted event
func (m *RecallSubmittedEvent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePayment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecall(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecallSubmission(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecallSubmittedEvent) validatePayment(formats strfmt.Registry) error {

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

func (m *RecallSubmittedEvent) validateRecall(formats strfmt.Registry) error {

	if swag.IsZero(m.Recall) { // not required
		return nil
	}

	if m.Recall != nil {
		if err := m.Recall.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("recall")
			}
			return err
		}
	}

	return nil
}

func (m *RecallSubmittedEvent) validateRecallSubmission(formats strfmt.Registry) error {

	if swag.IsZero(m.RecallSubmission) { // not required
		return nil
	}

	if m.RecallSubmission != nil {
		if err := m.RecallSubmission.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("recall_submission")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RecallSubmittedEvent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecallSubmittedEvent) UnmarshalBinary(b []byte) error {
	var res RecallSubmittedEvent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
