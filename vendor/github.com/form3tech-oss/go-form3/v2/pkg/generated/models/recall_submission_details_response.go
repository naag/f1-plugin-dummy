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

// RecallSubmissionDetailsResponse recall submission details response
// swagger:model RecallSubmissionDetailsResponse
type RecallSubmissionDetailsResponse struct {

	// data
	Data *RecallSubmission `json:"data,omitempty"`

	// links
	Links *Links `json:"links,omitempty"`
}

func RecallSubmissionDetailsResponseWithDefaults(defaults client.Defaults) *RecallSubmissionDetailsResponse {
	return &RecallSubmissionDetailsResponse{

		Data: RecallSubmissionWithDefaults(defaults),

		Links: LinksWithDefaults(defaults),
	}
}

func (m *RecallSubmissionDetailsResponse) WithData(data RecallSubmission) *RecallSubmissionDetailsResponse {

	m.Data = &data

	return m
}

func (m *RecallSubmissionDetailsResponse) WithoutData() *RecallSubmissionDetailsResponse {
	m.Data = nil
	return m
}

func (m *RecallSubmissionDetailsResponse) WithLinks(links Links) *RecallSubmissionDetailsResponse {

	m.Links = &links

	return m
}

func (m *RecallSubmissionDetailsResponse) WithoutLinks() *RecallSubmissionDetailsResponse {
	m.Links = nil
	return m
}

// Validate validates this recall submission details response
func (m *RecallSubmissionDetailsResponse) Validate(formats strfmt.Registry) error {
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

func (m *RecallSubmissionDetailsResponse) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
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

func (m *RecallSubmissionDetailsResponse) validateLinks(formats strfmt.Registry) error {

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
func (m *RecallSubmissionDetailsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecallSubmissionDetailsResponse) UnmarshalBinary(b []byte) error {
	var res RecallSubmissionDetailsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *RecallSubmissionDetailsResponse) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}