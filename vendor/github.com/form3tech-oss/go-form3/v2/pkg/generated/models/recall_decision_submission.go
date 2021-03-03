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

// RecallDecisionSubmission recall decision submission
// swagger:model RecallDecisionSubmission
type RecallDecisionSubmission struct {

	// attributes
	Attributes *RecallDecisionSubmissionAttributes `json:"attributes,omitempty"`

	// created on
	// Format: date-time
	CreatedOn *strfmt.DateTime `json:"created_on,omitempty"`

	// Unique resource ID
	// Required: true
	// Format: uuid
	ID *strfmt.UUID `json:"id"`

	// modified on
	// Format: date-time
	ModifiedOn *strfmt.DateTime `json:"modified_on,omitempty"`

	// Unique ID of the organisation this resource is created by
	// Required: true
	// Format: uuid
	OrganisationID *strfmt.UUID `json:"organisation_id"`

	// relationships
	Relationships *RecallDecisionSubmissionRelationships `json:"relationships,omitempty"`

	// Name of the resource type
	// Pattern: ^[A-Za-z_]*$
	Type string `json:"type,omitempty"`

	// Version number
	// Minimum: 0
	Version *int64 `json:"version,omitempty"`
}

func RecallDecisionSubmissionWithDefaults(defaults client.Defaults) *RecallDecisionSubmission {
	return &RecallDecisionSubmission{

		Attributes: RecallDecisionSubmissionAttributesWithDefaults(defaults),

		CreatedOn: defaults.GetStrfmtDateTimePtr("RecallDecisionSubmission", "created_on"),

		ID: defaults.GetStrfmtUUIDPtr("RecallDecisionSubmission", "id"),

		ModifiedOn: defaults.GetStrfmtDateTimePtr("RecallDecisionSubmission", "modified_on"),

		OrganisationID: defaults.GetStrfmtUUIDPtr("RecallDecisionSubmission", "organisation_id"),

		Relationships: RecallDecisionSubmissionRelationshipsWithDefaults(defaults),

		Type: defaults.GetString("RecallDecisionSubmission", "type"),

		Version: defaults.GetInt64Ptr("RecallDecisionSubmission", "version"),
	}
}

func (m *RecallDecisionSubmission) WithAttributes(attributes RecallDecisionSubmissionAttributes) *RecallDecisionSubmission {

	m.Attributes = &attributes

	return m
}

func (m *RecallDecisionSubmission) WithoutAttributes() *RecallDecisionSubmission {
	m.Attributes = nil
	return m
}

func (m *RecallDecisionSubmission) WithCreatedOn(createdOn strfmt.DateTime) *RecallDecisionSubmission {

	m.CreatedOn = &createdOn

	return m
}

func (m *RecallDecisionSubmission) WithoutCreatedOn() *RecallDecisionSubmission {
	m.CreatedOn = nil
	return m
}

func (m *RecallDecisionSubmission) WithID(id strfmt.UUID) *RecallDecisionSubmission {

	m.ID = &id

	return m
}

func (m *RecallDecisionSubmission) WithoutID() *RecallDecisionSubmission {
	m.ID = nil
	return m
}

func (m *RecallDecisionSubmission) WithModifiedOn(modifiedOn strfmt.DateTime) *RecallDecisionSubmission {

	m.ModifiedOn = &modifiedOn

	return m
}

func (m *RecallDecisionSubmission) WithoutModifiedOn() *RecallDecisionSubmission {
	m.ModifiedOn = nil
	return m
}

func (m *RecallDecisionSubmission) WithOrganisationID(organisationID strfmt.UUID) *RecallDecisionSubmission {

	m.OrganisationID = &organisationID

	return m
}

func (m *RecallDecisionSubmission) WithoutOrganisationID() *RecallDecisionSubmission {
	m.OrganisationID = nil
	return m
}

func (m *RecallDecisionSubmission) WithRelationships(relationships RecallDecisionSubmissionRelationships) *RecallDecisionSubmission {

	m.Relationships = &relationships

	return m
}

func (m *RecallDecisionSubmission) WithoutRelationships() *RecallDecisionSubmission {
	m.Relationships = nil
	return m
}

func (m *RecallDecisionSubmission) WithType(typeVar string) *RecallDecisionSubmission {

	m.Type = typeVar

	return m
}

func (m *RecallDecisionSubmission) WithVersion(version int64) *RecallDecisionSubmission {

	m.Version = &version

	return m
}

func (m *RecallDecisionSubmission) WithoutVersion() *RecallDecisionSubmission {
	m.Version = nil
	return m
}

// Validate validates this recall decision submission
func (m *RecallDecisionSubmission) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttributes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModifiedOn(formats); err != nil {
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

func (m *RecallDecisionSubmission) validateAttributes(formats strfmt.Registry) error {

	if swag.IsZero(m.Attributes) { // not required
		return nil
	}

	if m.Attributes != nil {
		if err := m.Attributes.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes")
			}
			return err
		}
	}

	return nil
}

func (m *RecallDecisionSubmission) validateCreatedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("created_on", "body", "date-time", m.CreatedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RecallDecisionSubmission) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RecallDecisionSubmission) validateModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("modified_on", "body", "date-time", m.ModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RecallDecisionSubmission) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", m.OrganisationID); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RecallDecisionSubmission) validateRelationships(formats strfmt.Registry) error {

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

func (m *RecallDecisionSubmission) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := validate.Pattern("type", "body", string(m.Type), `^[A-Za-z_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *RecallDecisionSubmission) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RecallDecisionSubmission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecallDecisionSubmission) UnmarshalBinary(b []byte) error {
	var res RecallDecisionSubmission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *RecallDecisionSubmission) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}