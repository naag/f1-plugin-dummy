// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// RecallDecisionSubmittedEvent recall decision submitted event
// swagger:model RecallDecisionSubmittedEvent
type RecallDecisionSubmittedEvent struct {

	// decision
	Decision *RecallDecision `json:"decision,omitempty"`

	// decision submission
	DecisionSubmission *RecallDecisionSubmission `json:"decision_submission,omitempty"`

	// payment
	Payment *Payment `json:"payment,omitempty"`

	// recall
	Recall *Recall `json:"recall,omitempty"`

	// recall admission
	RecallAdmission *RecallAdmission `json:"recall_admission,omitempty"`
}

// Validate validates this recall decision submitted event
func (m *RecallDecisionSubmittedEvent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDecision(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDecisionSubmission(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePayment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecall(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecallAdmission(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecallDecisionSubmittedEvent) validateDecision(formats strfmt.Registry) error {

	if swag.IsZero(m.Decision) { // not required
		return nil
	}

	if m.Decision != nil {
		if err := m.Decision.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("decision")
			}
			return err
		}
	}

	return nil
}

func (m *RecallDecisionSubmittedEvent) validateDecisionSubmission(formats strfmt.Registry) error {

	if swag.IsZero(m.DecisionSubmission) { // not required
		return nil
	}

	if m.DecisionSubmission != nil {
		if err := m.DecisionSubmission.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("decision_submission")
			}
			return err
		}
	}

	return nil
}

func (m *RecallDecisionSubmittedEvent) validatePayment(formats strfmt.Registry) error {

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

func (m *RecallDecisionSubmittedEvent) validateRecall(formats strfmt.Registry) error {

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

func (m *RecallDecisionSubmittedEvent) validateRecallAdmission(formats strfmt.Registry) error {

	if swag.IsZero(m.RecallAdmission) { // not required
		return nil
	}

	if m.RecallAdmission != nil {
		if err := m.RecallAdmission.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("recall_admission")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RecallDecisionSubmittedEvent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecallDecisionSubmittedEvent) UnmarshalBinary(b []byte) error {
	var res RecallDecisionSubmittedEvent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
