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

// RecallDecision recall decision
// swagger:model RecallDecision
type RecallDecision struct {

	// attributes
	Attributes *RecallDecisionAttributes `json:"attributes,omitempty"`

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
	Relationships *RecallDecisionRelationships `json:"relationships,omitempty"`

	// Name of the resource type
	// Pattern: ^[A-Za-z_]*$
	Type string `json:"type,omitempty"`

	// Version number
	// Minimum: 0
	Version *int64 `json:"version,omitempty"`
}

func RecallDecisionWithDefaults(defaults client.Defaults) *RecallDecision {
	return &RecallDecision{

		Attributes: RecallDecisionAttributesWithDefaults(defaults),

		CreatedOn: defaults.GetStrfmtDateTimePtr("RecallDecision", "created_on"),

		ID: defaults.GetStrfmtUUIDPtr("RecallDecision", "id"),

		ModifiedOn: defaults.GetStrfmtDateTimePtr("RecallDecision", "modified_on"),

		OrganisationID: defaults.GetStrfmtUUIDPtr("RecallDecision", "organisation_id"),

		Relationships: RecallDecisionRelationshipsWithDefaults(defaults),

		Type: defaults.GetString("RecallDecision", "type"),

		Version: defaults.GetInt64Ptr("RecallDecision", "version"),
	}
}

func (m *RecallDecision) WithAttributes(attributes RecallDecisionAttributes) *RecallDecision {

	m.Attributes = &attributes

	return m
}

func (m *RecallDecision) WithoutAttributes() *RecallDecision {
	m.Attributes = nil
	return m
}

func (m *RecallDecision) WithCreatedOn(createdOn strfmt.DateTime) *RecallDecision {

	m.CreatedOn = &createdOn

	return m
}

func (m *RecallDecision) WithoutCreatedOn() *RecallDecision {
	m.CreatedOn = nil
	return m
}

func (m *RecallDecision) WithID(id strfmt.UUID) *RecallDecision {

	m.ID = &id

	return m
}

func (m *RecallDecision) WithoutID() *RecallDecision {
	m.ID = nil
	return m
}

func (m *RecallDecision) WithModifiedOn(modifiedOn strfmt.DateTime) *RecallDecision {

	m.ModifiedOn = &modifiedOn

	return m
}

func (m *RecallDecision) WithoutModifiedOn() *RecallDecision {
	m.ModifiedOn = nil
	return m
}

func (m *RecallDecision) WithOrganisationID(organisationID strfmt.UUID) *RecallDecision {

	m.OrganisationID = &organisationID

	return m
}

func (m *RecallDecision) WithoutOrganisationID() *RecallDecision {
	m.OrganisationID = nil
	return m
}

func (m *RecallDecision) WithRelationships(relationships RecallDecisionRelationships) *RecallDecision {

	m.Relationships = &relationships

	return m
}

func (m *RecallDecision) WithoutRelationships() *RecallDecision {
	m.Relationships = nil
	return m
}

func (m *RecallDecision) WithType(typeVar string) *RecallDecision {

	m.Type = typeVar

	return m
}

func (m *RecallDecision) WithVersion(version int64) *RecallDecision {

	m.Version = &version

	return m
}

func (m *RecallDecision) WithoutVersion() *RecallDecision {
	m.Version = nil
	return m
}

// Validate validates this recall decision
func (m *RecallDecision) Validate(formats strfmt.Registry) error {
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

func (m *RecallDecision) validateAttributes(formats strfmt.Registry) error {

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

func (m *RecallDecision) validateCreatedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("created_on", "body", "date-time", m.CreatedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RecallDecision) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RecallDecision) validateModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("modified_on", "body", "date-time", m.ModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RecallDecision) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", m.OrganisationID); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RecallDecision) validateRelationships(formats strfmt.Registry) error {

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

func (m *RecallDecision) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := validate.Pattern("type", "body", string(m.Type), `^[A-Za-z_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *RecallDecision) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RecallDecision) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecallDecision) UnmarshalBinary(b []byte) error {
	var res RecallDecision
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *RecallDecision) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// RecallDecisionAttributes recall decision attributes
// swagger:model RecallDecisionAttributes
type RecallDecisionAttributes struct {

	// answer
	Answer RecallDecisionAnswer `json:"answer,omitempty"`

	// charges amount
	ChargesAmount *CurrencyAndAmount `json:"charges_amount,omitempty"`

	// Optional free text reason in addition to `reason_code`
	Reason string `json:"reason,omitempty"`

	// Reason for a rejected decision. Required when answer is rejected, ignored otherwise. Has to be a valid [rejected recall decision reason code](http://api-docs.form3.tech/api.html#enumerations-rejected-recall-decision-reason-codes)
	ReasonCode string `json:"reason_code,omitempty"`

	// recall amount
	RecallAmount *CurrencyAndAmount `json:"recall_amount,omitempty"`
}

func RecallDecisionAttributesWithDefaults(defaults client.Defaults) *RecallDecisionAttributes {
	return &RecallDecisionAttributes{

		// TODO Answer: RecallDecisionAnswer,

		ChargesAmount: CurrencyAndAmountWithDefaults(defaults),

		Reason: defaults.GetString("RecallDecisionAttributes", "reason"),

		ReasonCode: defaults.GetString("RecallDecisionAttributes", "reason_code"),

		RecallAmount: CurrencyAndAmountWithDefaults(defaults),
	}
}

func (m *RecallDecisionAttributes) WithAnswer(answer RecallDecisionAnswer) *RecallDecisionAttributes {

	m.Answer = answer

	return m
}

func (m *RecallDecisionAttributes) WithChargesAmount(chargesAmount CurrencyAndAmount) *RecallDecisionAttributes {

	m.ChargesAmount = &chargesAmount

	return m
}

func (m *RecallDecisionAttributes) WithoutChargesAmount() *RecallDecisionAttributes {
	m.ChargesAmount = nil
	return m
}

func (m *RecallDecisionAttributes) WithReason(reason string) *RecallDecisionAttributes {

	m.Reason = reason

	return m
}

func (m *RecallDecisionAttributes) WithReasonCode(reasonCode string) *RecallDecisionAttributes {

	m.ReasonCode = reasonCode

	return m
}

func (m *RecallDecisionAttributes) WithRecallAmount(recallAmount CurrencyAndAmount) *RecallDecisionAttributes {

	m.RecallAmount = &recallAmount

	return m
}

func (m *RecallDecisionAttributes) WithoutRecallAmount() *RecallDecisionAttributes {
	m.RecallAmount = nil
	return m
}

// Validate validates this recall decision attributes
func (m *RecallDecisionAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAnswer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChargesAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecallAmount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecallDecisionAttributes) validateAnswer(formats strfmt.Registry) error {

	if swag.IsZero(m.Answer) { // not required
		return nil
	}

	if err := m.Answer.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("attributes" + "." + "answer")
		}
		return err
	}

	return nil
}

func (m *RecallDecisionAttributes) validateChargesAmount(formats strfmt.Registry) error {

	if swag.IsZero(m.ChargesAmount) { // not required
		return nil
	}

	if m.ChargesAmount != nil {
		if err := m.ChargesAmount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes" + "." + "charges_amount")
			}
			return err
		}
	}

	return nil
}

func (m *RecallDecisionAttributes) validateRecallAmount(formats strfmt.Registry) error {

	if swag.IsZero(m.RecallAmount) { // not required
		return nil
	}

	if m.RecallAmount != nil {
		if err := m.RecallAmount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes" + "." + "recall_amount")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RecallDecisionAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecallDecisionAttributes) UnmarshalBinary(b []byte) error {
	var res RecallDecisionAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *RecallDecisionAttributes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// RecallDecisionRelationships recall decision relationships
// swagger:model RecallDecisionRelationships
type RecallDecisionRelationships struct {

	// ID of the recall decision admission resource related to the recall decision
	DecisionAdmission *RelationshipLinks `json:"decision_admission,omitempty"`

	// ID of the recall decision submission resource related to the recall decision
	DecisionSubmission *RelationshipLinks `json:"decision_submission,omitempty"`

	// ID of the payment resource related to the recall decision
	Payment *RelationshipLinks `json:"payment,omitempty"`

	// ID of the recall resource related to the recall decision
	Recall *RelationshipLinks `json:"recall,omitempty"`
}

func RecallDecisionRelationshipsWithDefaults(defaults client.Defaults) *RecallDecisionRelationships {
	return &RecallDecisionRelationships{

		DecisionAdmission: RelationshipLinksWithDefaults(defaults),

		DecisionSubmission: RelationshipLinksWithDefaults(defaults),

		Payment: RelationshipLinksWithDefaults(defaults),

		Recall: RelationshipLinksWithDefaults(defaults),
	}
}

func (m *RecallDecisionRelationships) WithDecisionAdmission(decisionAdmission RelationshipLinks) *RecallDecisionRelationships {

	m.DecisionAdmission = &decisionAdmission

	return m
}

func (m *RecallDecisionRelationships) WithoutDecisionAdmission() *RecallDecisionRelationships {
	m.DecisionAdmission = nil
	return m
}

func (m *RecallDecisionRelationships) WithDecisionSubmission(decisionSubmission RelationshipLinks) *RecallDecisionRelationships {

	m.DecisionSubmission = &decisionSubmission

	return m
}

func (m *RecallDecisionRelationships) WithoutDecisionSubmission() *RecallDecisionRelationships {
	m.DecisionSubmission = nil
	return m
}

func (m *RecallDecisionRelationships) WithPayment(payment RelationshipLinks) *RecallDecisionRelationships {

	m.Payment = &payment

	return m
}

func (m *RecallDecisionRelationships) WithoutPayment() *RecallDecisionRelationships {
	m.Payment = nil
	return m
}

func (m *RecallDecisionRelationships) WithRecall(recall RelationshipLinks) *RecallDecisionRelationships {

	m.Recall = &recall

	return m
}

func (m *RecallDecisionRelationships) WithoutRecall() *RecallDecisionRelationships {
	m.Recall = nil
	return m
}

// Validate validates this recall decision relationships
func (m *RecallDecisionRelationships) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDecisionAdmission(formats); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecallDecisionRelationships) validateDecisionAdmission(formats strfmt.Registry) error {

	if swag.IsZero(m.DecisionAdmission) { // not required
		return nil
	}

	if m.DecisionAdmission != nil {
		if err := m.DecisionAdmission.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("relationships" + "." + "decision_admission")
			}
			return err
		}
	}

	return nil
}

func (m *RecallDecisionRelationships) validateDecisionSubmission(formats strfmt.Registry) error {

	if swag.IsZero(m.DecisionSubmission) { // not required
		return nil
	}

	if m.DecisionSubmission != nil {
		if err := m.DecisionSubmission.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("relationships" + "." + "decision_submission")
			}
			return err
		}
	}

	return nil
}

func (m *RecallDecisionRelationships) validatePayment(formats strfmt.Registry) error {

	if swag.IsZero(m.Payment) { // not required
		return nil
	}

	if m.Payment != nil {
		if err := m.Payment.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("relationships" + "." + "payment")
			}
			return err
		}
	}

	return nil
}

func (m *RecallDecisionRelationships) validateRecall(formats strfmt.Registry) error {

	if swag.IsZero(m.Recall) { // not required
		return nil
	}

	if m.Recall != nil {
		if err := m.Recall.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("relationships" + "." + "recall")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RecallDecisionRelationships) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecallDecisionRelationships) UnmarshalBinary(b []byte) error {
	var res RecallDecisionRelationships
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *RecallDecisionRelationships) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}