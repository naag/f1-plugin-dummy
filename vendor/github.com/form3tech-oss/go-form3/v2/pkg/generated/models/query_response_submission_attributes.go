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

// QueryResponseSubmissionAttributes query response submission attributes
// swagger:model QueryResponseSubmissionAttributes
type QueryResponseSubmissionAttributes struct {

	// auto
	Auto bool `json:"auto,omitempty"`

	// status
	Status QueryResponseSubmissionStatus `json:"status,omitempty"`
}

func QueryResponseSubmissionAttributesWithDefaults(defaults client.Defaults) *QueryResponseSubmissionAttributes {
	return &QueryResponseSubmissionAttributes{

		Auto: defaults.GetBool("QueryResponseSubmissionAttributes", "auto"),

		// TODO Status: QueryResponseSubmissionStatus,

	}
}

func (m *QueryResponseSubmissionAttributes) WithAuto(auto bool) *QueryResponseSubmissionAttributes {

	m.Auto = auto

	return m
}

func (m *QueryResponseSubmissionAttributes) WithStatus(status QueryResponseSubmissionStatus) *QueryResponseSubmissionAttributes {

	m.Status = status

	return m
}

// Validate validates this query response submission attributes
func (m *QueryResponseSubmissionAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QueryResponseSubmissionAttributes) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.Status.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *QueryResponseSubmissionAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QueryResponseSubmissionAttributes) UnmarshalBinary(b []byte) error {
	var res QueryResponseSubmissionAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *QueryResponseSubmissionAttributes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
