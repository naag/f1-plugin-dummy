// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// ClaimResourceType claim resource type
// swagger:model ClaimResourceType
type ClaimResourceType string

const (

	// ClaimResourceTypeClaims captures enum value "claims"
	ClaimResourceTypeClaims ClaimResourceType = "claims"

	// ClaimResourceTypeClaimSubmissions captures enum value "claim_submissions"
	ClaimResourceTypeClaimSubmissions ClaimResourceType = "claim_submissions"

	// ClaimResourceTypeClaimSubmissionValidations captures enum value "claim_submission_validations"
	ClaimResourceTypeClaimSubmissionValidations ClaimResourceType = "claim_submission_validations"

	// ClaimResourceTypeClaimReversals captures enum value "claim_reversals"
	ClaimResourceTypeClaimReversals ClaimResourceType = "claim_reversals"

	// ClaimResourceTypeClaimReversalSubmissions captures enum value "claim_reversal_submissions"
	ClaimResourceTypeClaimReversalSubmissions ClaimResourceType = "claim_reversal_submissions"

	// ClaimResourceTypeClaimReversalSubmissionValidations captures enum value "claim_reversal_submission_validations"
	ClaimResourceTypeClaimReversalSubmissionValidations ClaimResourceType = "claim_reversal_submission_validations"
)

// for schema
var claimResourceTypeEnum []interface{}

func init() {
	var res []ClaimResourceType
	if err := json.Unmarshal([]byte(`["claims","claim_submissions","claim_submission_validations","claim_reversals","claim_reversal_submissions","claim_reversal_submission_validations"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		claimResourceTypeEnum = append(claimResourceTypeEnum, v)
	}
}

func (m ClaimResourceType) validateClaimResourceTypeEnum(path, location string, value ClaimResourceType) error {
	if err := validate.Enum(path, location, value, claimResourceTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this claim resource type
func (m ClaimResourceType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateClaimResourceTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClaimResourceType) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
