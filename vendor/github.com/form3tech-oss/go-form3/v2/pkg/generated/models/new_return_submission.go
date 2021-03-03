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

// NewReturnSubmission new return submission
// swagger:model NewReturnSubmission
type NewReturnSubmission struct {

	// Unique resource ID
	// Required: true
	// Format: uuid
	ID *strfmt.UUID `json:"id"`

	// Unique ID of the organisation this resource is created by
	// Required: true
	// Format: uuid
	OrganisationID *strfmt.UUID `json:"organisation_id"`

	// relationships
	Relationships *ReturnSubmissionRelationships `json:"relationships,omitempty"`

	// Name of the resource type
	// Pattern: ^[A-Za-z_]*$
	Type string `json:"type,omitempty"`

	// Version number
	// Minimum: 0
	Version *int64 `json:"version,omitempty"`
}

func NewReturnSubmissionWithDefaults(defaults client.Defaults) *NewReturnSubmission {
	return &NewReturnSubmission{

		ID: defaults.GetStrfmtUUIDPtr("NewReturnSubmission", "id"),

		OrganisationID: defaults.GetStrfmtUUIDPtr("NewReturnSubmission", "organisation_id"),

		Relationships: ReturnSubmissionRelationshipsWithDefaults(defaults),

		Type: defaults.GetString("NewReturnSubmission", "type"),

		Version: defaults.GetInt64Ptr("NewReturnSubmission", "version"),
	}
}

func (m *NewReturnSubmission) WithID(id strfmt.UUID) *NewReturnSubmission {

	m.ID = &id

	return m
}

func (m *NewReturnSubmission) WithoutID() *NewReturnSubmission {
	m.ID = nil
	return m
}

func (m *NewReturnSubmission) WithOrganisationID(organisationID strfmt.UUID) *NewReturnSubmission {

	m.OrganisationID = &organisationID

	return m
}

func (m *NewReturnSubmission) WithoutOrganisationID() *NewReturnSubmission {
	m.OrganisationID = nil
	return m
}

func (m *NewReturnSubmission) WithRelationships(relationships ReturnSubmissionRelationships) *NewReturnSubmission {

	m.Relationships = &relationships

	return m
}

func (m *NewReturnSubmission) WithoutRelationships() *NewReturnSubmission {
	m.Relationships = nil
	return m
}

func (m *NewReturnSubmission) WithType(typeVar string) *NewReturnSubmission {

	m.Type = typeVar

	return m
}

func (m *NewReturnSubmission) WithVersion(version int64) *NewReturnSubmission {

	m.Version = &version

	return m
}

func (m *NewReturnSubmission) WithoutVersion() *NewReturnSubmission {
	m.Version = nil
	return m
}

// Validate validates this new return submission
func (m *NewReturnSubmission) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganisationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRelationships(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewReturnSubmission) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NewReturnSubmission) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", m.OrganisationID); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NewReturnSubmission) validateRelationships(formats strfmt.Registry) error {

	if swag.IsZero(m.Relationships) { // not required
		return nil
	}

	if m.Relationships != nil {
		if err := m.Relationships.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("relationships")
			}
			return err
		}
	}

	return nil
}

func (m *NewReturnSubmission) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := validate.Pattern("type", "body", string(m.Type), `^[A-Za-z_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *NewReturnSubmission) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NewReturnSubmission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewReturnSubmission) UnmarshalBinary(b []byte) error {
	var res NewReturnSubmission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *NewReturnSubmission) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}