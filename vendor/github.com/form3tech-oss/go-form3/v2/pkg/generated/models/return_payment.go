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

// ReturnPayment return payment
// swagger:model ReturnPayment
type ReturnPayment struct {

	// attributes
	Attributes *ReturnPaymentAttributes `json:"attributes,omitempty"`

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
	Relationships *ReturnPaymentRelationships `json:"relationships,omitempty"`

	// Name of the resource type
	// Pattern: ^[A-Za-z_]*$
	Type string `json:"type,omitempty"`

	// Version number
	// Minimum: 0
	Version *int64 `json:"version,omitempty"`
}

func ReturnPaymentWithDefaults(defaults client.Defaults) *ReturnPayment {
	return &ReturnPayment{

		Attributes: ReturnPaymentAttributesWithDefaults(defaults),

		CreatedOn: defaults.GetStrfmtDateTimePtr("ReturnPayment", "created_on"),

		ID: defaults.GetStrfmtUUIDPtr("ReturnPayment", "id"),

		ModifiedOn: defaults.GetStrfmtDateTimePtr("ReturnPayment", "modified_on"),

		OrganisationID: defaults.GetStrfmtUUIDPtr("ReturnPayment", "organisation_id"),

		Relationships: ReturnPaymentRelationshipsWithDefaults(defaults),

		Type: defaults.GetString("ReturnPayment", "type"),

		Version: defaults.GetInt64Ptr("ReturnPayment", "version"),
	}
}

func (m *ReturnPayment) WithAttributes(attributes ReturnPaymentAttributes) *ReturnPayment {

	m.Attributes = &attributes

	return m
}

func (m *ReturnPayment) WithoutAttributes() *ReturnPayment {
	m.Attributes = nil
	return m
}

func (m *ReturnPayment) WithCreatedOn(createdOn strfmt.DateTime) *ReturnPayment {

	m.CreatedOn = &createdOn

	return m
}

func (m *ReturnPayment) WithoutCreatedOn() *ReturnPayment {
	m.CreatedOn = nil
	return m
}

func (m *ReturnPayment) WithID(id strfmt.UUID) *ReturnPayment {

	m.ID = &id

	return m
}

func (m *ReturnPayment) WithoutID() *ReturnPayment {
	m.ID = nil
	return m
}

func (m *ReturnPayment) WithModifiedOn(modifiedOn strfmt.DateTime) *ReturnPayment {

	m.ModifiedOn = &modifiedOn

	return m
}

func (m *ReturnPayment) WithoutModifiedOn() *ReturnPayment {
	m.ModifiedOn = nil
	return m
}

func (m *ReturnPayment) WithOrganisationID(organisationID strfmt.UUID) *ReturnPayment {

	m.OrganisationID = &organisationID

	return m
}

func (m *ReturnPayment) WithoutOrganisationID() *ReturnPayment {
	m.OrganisationID = nil
	return m
}

func (m *ReturnPayment) WithRelationships(relationships ReturnPaymentRelationships) *ReturnPayment {

	m.Relationships = &relationships

	return m
}

func (m *ReturnPayment) WithoutRelationships() *ReturnPayment {
	m.Relationships = nil
	return m
}

func (m *ReturnPayment) WithType(typeVar string) *ReturnPayment {

	m.Type = typeVar

	return m
}

func (m *ReturnPayment) WithVersion(version int64) *ReturnPayment {

	m.Version = &version

	return m
}

func (m *ReturnPayment) WithoutVersion() *ReturnPayment {
	m.Version = nil
	return m
}

// Validate validates this return payment
func (m *ReturnPayment) Validate(formats strfmt.Registry) error {
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

func (m *ReturnPayment) validateAttributes(formats strfmt.Registry) error {

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

func (m *ReturnPayment) validateCreatedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("created_on", "body", "date-time", m.CreatedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReturnPayment) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReturnPayment) validateModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("modified_on", "body", "date-time", m.ModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReturnPayment) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", m.OrganisationID); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReturnPayment) validateRelationships(formats strfmt.Registry) error {

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

func (m *ReturnPayment) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := validate.Pattern("type", "body", string(m.Type), `^[A-Za-z_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *ReturnPayment) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReturnPayment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReturnPayment) UnmarshalBinary(b []byte) error {
	var res ReturnPayment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *ReturnPayment) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// ReturnPaymentAttributes return payment attributes
// swagger:model ReturnPaymentAttributes
type ReturnPaymentAttributes struct {

	// The amount to return which should correspond to the amount of the original payment
	// Pattern: ^[0-9.]{0,20}$
	Amount string `json:"amount,omitempty"`

	// Unique identifier for organisations collecting payments
	ClearingID string `json:"clearing_id,omitempty"`

	// ISO currency code for transaction amount
	Currency string `json:"currency,omitempty"`

	// Time a payment was released from being held due to a limit breach
	// Read Only: true
	// Format: date-time
	LimitBreachEndDatetime *strfmt.DateTime `json:"limit_breach_end_datetime,omitempty"`

	// Start time a payment was held due to a limit breach
	// Read Only: true
	// Format: date-time
	LimitBreachStartDatetime *strfmt.DateTime `json:"limit_breach_start_datetime,omitempty"`

	// reason
	Reason *string `json:"reason,omitempty"`

	// The return [reason code](http://draft-api-docs.form3.tech/api.html#enumerations-payment-return-codes)
	ReturnCode string `json:"return_code,omitempty"`

	// A unique reference to the return payment instruction. If not supplied, the value is generated by Form3.
	SchemeTransactionID string `json:"scheme_transaction_id,omitempty"`
}

func ReturnPaymentAttributesWithDefaults(defaults client.Defaults) *ReturnPaymentAttributes {
	return &ReturnPaymentAttributes{

		Amount: defaults.GetString("ReturnPaymentAttributes", "amount"),

		ClearingID: defaults.GetString("ReturnPaymentAttributes", "clearing_id"),

		Currency: defaults.GetString("ReturnPaymentAttributes", "currency"),

		LimitBreachEndDatetime: defaults.GetStrfmtDateTimePtr("ReturnPaymentAttributes", "limit_breach_end_datetime"),

		LimitBreachStartDatetime: defaults.GetStrfmtDateTimePtr("ReturnPaymentAttributes", "limit_breach_start_datetime"),

		Reason: defaults.GetStringPtr("ReturnPaymentAttributes", "reason"),

		ReturnCode: defaults.GetString("ReturnPaymentAttributes", "return_code"),

		SchemeTransactionID: defaults.GetString("ReturnPaymentAttributes", "scheme_transaction_id"),
	}
}

func (m *ReturnPaymentAttributes) WithAmount(amount string) *ReturnPaymentAttributes {

	m.Amount = amount

	return m
}

func (m *ReturnPaymentAttributes) WithClearingID(clearingID string) *ReturnPaymentAttributes {

	m.ClearingID = clearingID

	return m
}

func (m *ReturnPaymentAttributes) WithCurrency(currency string) *ReturnPaymentAttributes {

	m.Currency = currency

	return m
}

func (m *ReturnPaymentAttributes) WithLimitBreachEndDatetime(limitBreachEndDatetime strfmt.DateTime) *ReturnPaymentAttributes {

	m.LimitBreachEndDatetime = &limitBreachEndDatetime

	return m
}

func (m *ReturnPaymentAttributes) WithoutLimitBreachEndDatetime() *ReturnPaymentAttributes {
	m.LimitBreachEndDatetime = nil
	return m
}

func (m *ReturnPaymentAttributes) WithLimitBreachStartDatetime(limitBreachStartDatetime strfmt.DateTime) *ReturnPaymentAttributes {

	m.LimitBreachStartDatetime = &limitBreachStartDatetime

	return m
}

func (m *ReturnPaymentAttributes) WithoutLimitBreachStartDatetime() *ReturnPaymentAttributes {
	m.LimitBreachStartDatetime = nil
	return m
}

func (m *ReturnPaymentAttributes) WithReason(reason string) *ReturnPaymentAttributes {

	m.Reason = &reason

	return m
}

func (m *ReturnPaymentAttributes) WithoutReason() *ReturnPaymentAttributes {
	m.Reason = nil
	return m
}

func (m *ReturnPaymentAttributes) WithReturnCode(returnCode string) *ReturnPaymentAttributes {

	m.ReturnCode = returnCode

	return m
}

func (m *ReturnPaymentAttributes) WithSchemeTransactionID(schemeTransactionID string) *ReturnPaymentAttributes {

	m.SchemeTransactionID = schemeTransactionID

	return m
}

// Validate validates this return payment attributes
func (m *ReturnPaymentAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLimitBreachEndDatetime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLimitBreachStartDatetime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReturnPaymentAttributes) validateAmount(formats strfmt.Registry) error {

	if swag.IsZero(m.Amount) { // not required
		return nil
	}

	if err := validate.Pattern("attributes"+"."+"amount", "body", string(m.Amount), `^[0-9.]{0,20}$`); err != nil {
		return err
	}

	return nil
}

func (m *ReturnPaymentAttributes) validateLimitBreachEndDatetime(formats strfmt.Registry) error {

	if swag.IsZero(m.LimitBreachEndDatetime) { // not required
		return nil
	}

	if err := validate.FormatOf("attributes"+"."+"limit_breach_end_datetime", "body", "date-time", m.LimitBreachEndDatetime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReturnPaymentAttributes) validateLimitBreachStartDatetime(formats strfmt.Registry) error {

	if swag.IsZero(m.LimitBreachStartDatetime) { // not required
		return nil
	}

	if err := validate.FormatOf("attributes"+"."+"limit_breach_start_datetime", "body", "date-time", m.LimitBreachStartDatetime.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReturnPaymentAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReturnPaymentAttributes) UnmarshalBinary(b []byte) error {
	var res ReturnPaymentAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *ReturnPaymentAttributes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// ReturnPaymentRelationships return payment relationships
// swagger:model ReturnPaymentRelationships
type ReturnPaymentRelationships struct {

	// ID of the payment resource related to the return
	Payment *RelationshipLinks `json:"payment,omitempty"`

	// ID of the return admission resource related to the return
	ReturnAdmission *RelationshipLinks `json:"return_admission,omitempty"`

	// ID of the return submission resource related to the return
	ReturnSubmission *RelationshipLinks `json:"return_submission,omitempty"`
}

func ReturnPaymentRelationshipsWithDefaults(defaults client.Defaults) *ReturnPaymentRelationships {
	return &ReturnPaymentRelationships{

		Payment: RelationshipLinksWithDefaults(defaults),

		ReturnAdmission: RelationshipLinksWithDefaults(defaults),

		ReturnSubmission: RelationshipLinksWithDefaults(defaults),
	}
}

func (m *ReturnPaymentRelationships) WithPayment(payment RelationshipLinks) *ReturnPaymentRelationships {

	m.Payment = &payment

	return m
}

func (m *ReturnPaymentRelationships) WithoutPayment() *ReturnPaymentRelationships {
	m.Payment = nil
	return m
}

func (m *ReturnPaymentRelationships) WithReturnAdmission(returnAdmission RelationshipLinks) *ReturnPaymentRelationships {

	m.ReturnAdmission = &returnAdmission

	return m
}

func (m *ReturnPaymentRelationships) WithoutReturnAdmission() *ReturnPaymentRelationships {
	m.ReturnAdmission = nil
	return m
}

func (m *ReturnPaymentRelationships) WithReturnSubmission(returnSubmission RelationshipLinks) *ReturnPaymentRelationships {

	m.ReturnSubmission = &returnSubmission

	return m
}

func (m *ReturnPaymentRelationships) WithoutReturnSubmission() *ReturnPaymentRelationships {
	m.ReturnSubmission = nil
	return m
}

// Validate validates this return payment relationships
func (m *ReturnPaymentRelationships) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePayment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReturnAdmission(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReturnSubmission(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReturnPaymentRelationships) validatePayment(formats strfmt.Registry) error {

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

func (m *ReturnPaymentRelationships) validateReturnAdmission(formats strfmt.Registry) error {

	if swag.IsZero(m.ReturnAdmission) { // not required
		return nil
	}

	if m.ReturnAdmission != nil {
		if err := m.ReturnAdmission.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("relationships" + "." + "return_admission")
			}
			return err
		}
	}

	return nil
}

func (m *ReturnPaymentRelationships) validateReturnSubmission(formats strfmt.Registry) error {

	if swag.IsZero(m.ReturnSubmission) { // not required
		return nil
	}

	if m.ReturnSubmission != nil {
		if err := m.ReturnSubmission.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("relationships" + "." + "return_submission")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReturnPaymentRelationships) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReturnPaymentRelationships) UnmarshalBinary(b []byte) error {
	var res ReturnPaymentRelationships
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *ReturnPaymentRelationships) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}