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
	"github.com/go-openapi/validate"
)

// QueryResponseSubmissionResponse query response submission response
// swagger:model QueryResponseSubmissionResponse
type QueryResponseSubmissionResponse struct {

	// data
	// Required: true
	Data *QueryResponseSubmission `json:"data"`

	// links
	Links *Links `json:"links,omitempty"`
}

func QueryResponseSubmissionResponseWithDefaults(defaults client.Defaults) *QueryResponseSubmissionResponse {
	return &QueryResponseSubmissionResponse{

		Data: QueryResponseSubmissionWithDefaults(defaults),

		Links: LinksWithDefaults(defaults),
	}
}

func (m *QueryResponseSubmissionResponse) WithData(data QueryResponseSubmission) *QueryResponseSubmissionResponse {

	m.Data = &data

	return m
}

func (m *QueryResponseSubmissionResponse) WithoutData() *QueryResponseSubmissionResponse {
	m.Data = nil
	return m
}

func (m *QueryResponseSubmissionResponse) WithLinks(links Links) *QueryResponseSubmissionResponse {

	m.Links = &links

	return m
}

func (m *QueryResponseSubmissionResponse) WithoutLinks() *QueryResponseSubmissionResponse {
	m.Links = nil
	return m
}

// Validate validates this query response submission response
func (m *QueryResponseSubmissionResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QueryResponseSubmissionResponse) validateData(formats strfmt.Registry) error {

	if err := validate.Required("data", "body", m.Data); err != nil {
		return err
	}

	if m.Data != nil {
		if err := m.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data")
			}
			return err
		}
	}

	return nil
}

func (m *QueryResponseSubmissionResponse) validateLinks(formats strfmt.Registry) error {

	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *QueryResponseSubmissionResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QueryResponseSubmissionResponse) UnmarshalBinary(b []byte) error {
	var res QueryResponseSubmissionResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *QueryResponseSubmissionResponse) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}