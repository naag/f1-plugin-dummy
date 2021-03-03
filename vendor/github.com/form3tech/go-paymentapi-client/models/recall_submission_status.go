// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// RecallSubmissionStatus [Status](http://draft-api-docs.form3.tech/api.html#enumerations-payment-status-codes-payment-submission-status) of the submission
// swagger:model RecallSubmissionStatus
type RecallSubmissionStatus string

const (

	// RecallSubmissionStatusAccepted captures enum value "accepted"
	RecallSubmissionStatusAccepted RecallSubmissionStatus = "accepted"

	// RecallSubmissionStatusValidationPending captures enum value "validation_pending"
	RecallSubmissionStatusValidationPending RecallSubmissionStatus = "validation_pending"

	// RecallSubmissionStatusValidationPassed captures enum value "validation_passed"
	RecallSubmissionStatusValidationPassed RecallSubmissionStatus = "validation_passed"

	// RecallSubmissionStatusReleasedToGateway captures enum value "released_to_gateway"
	RecallSubmissionStatusReleasedToGateway RecallSubmissionStatus = "released_to_gateway"

	// RecallSubmissionStatusDeliveryConfirmed captures enum value "delivery_confirmed"
	RecallSubmissionStatusDeliveryConfirmed RecallSubmissionStatus = "delivery_confirmed"

	// RecallSubmissionStatusDeliveryFailed captures enum value "delivery_failed"
	RecallSubmissionStatusDeliveryFailed RecallSubmissionStatus = "delivery_failed"

	// RecallSubmissionStatusQueuedForDelivery captures enum value "queued_for_delivery"
	RecallSubmissionStatusQueuedForDelivery RecallSubmissionStatus = "queued_for_delivery"

	// RecallSubmissionStatusSubmitted captures enum value "submitted"
	RecallSubmissionStatusSubmitted RecallSubmissionStatus = "submitted"
)

// for schema
var recallSubmissionStatusEnum []interface{}

func init() {
	var res []RecallSubmissionStatus
	if err := json.Unmarshal([]byte(`["accepted","validation_pending","validation_passed","released_to_gateway","delivery_confirmed","delivery_failed","queued_for_delivery","submitted"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		recallSubmissionStatusEnum = append(recallSubmissionStatusEnum, v)
	}
}

func (m RecallSubmissionStatus) validateRecallSubmissionStatusEnum(path, location string, value RecallSubmissionStatus) error {
	if err := validate.Enum(path, location, value, recallSubmissionStatusEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this recall submission status
func (m RecallSubmissionStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateRecallSubmissionStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}