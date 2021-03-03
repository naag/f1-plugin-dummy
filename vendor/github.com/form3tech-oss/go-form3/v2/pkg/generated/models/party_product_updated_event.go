// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/form3tech-oss/go-form3/v2/pkg/client"
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PartyProductUpdatedEvent party product updated event
// swagger:model PartyProductUpdatedEvent
type PartyProductUpdatedEvent struct {

	// event type
	// Required: true
	// Enum: [updated]
	EventType *string `json:"event_type"`

	// id
	// Required: true
	// Format: uuid
	ID *strfmt.UUID `json:"id"`

	// organisation id
	// Required: true
	// Format: uuid
	OrganisationID *strfmt.UUID `json:"organisation_id"`

	// party id
	// Format: uuid
	PartyID strfmt.UUID `json:"party_id,omitempty"`

	// party product status
	PartyProductStatus PartyProductStatus `json:"party_product_status,omitempty"`

	// product id
	// Format: uuid
	ProductID strfmt.UUID `json:"product_id,omitempty"`

	// record type
	// Required: true
	RecordType ResourceType `json:"record_type"`

	// status items
	StatusItems []*PartyProductStatusItem `json:"status_items"`

	// user id
	UserID string `json:"user_id,omitempty"`

	// vendor status
	VendorStatus string `json:"vendor_status,omitempty"`
}

func PartyProductUpdatedEventWithDefaults(defaults client.Defaults) *PartyProductUpdatedEvent {
	return &PartyProductUpdatedEvent{

		EventType: defaults.GetStringPtr("PartyProductUpdatedEvent", "event_type"),

		ID: defaults.GetStrfmtUUIDPtr("PartyProductUpdatedEvent", "id"),

		OrganisationID: defaults.GetStrfmtUUIDPtr("PartyProductUpdatedEvent", "organisation_id"),

		PartyID: defaults.GetStrfmtUUID("PartyProductUpdatedEvent", "party_id"),

		// TODO PartyProductStatus: PartyProductStatus,

		ProductID: defaults.GetStrfmtUUID("PartyProductUpdatedEvent", "product_id"),

		// TODO RecordType: ResourceType,

		StatusItems: make([]*PartyProductStatusItem, 0),

		UserID: defaults.GetString("PartyProductUpdatedEvent", "user_id"),

		VendorStatus: defaults.GetString("PartyProductUpdatedEvent", "vendor_status"),
	}
}

func (m *PartyProductUpdatedEvent) WithEventType(eventType string) *PartyProductUpdatedEvent {

	m.EventType = &eventType

	return m
}

func (m *PartyProductUpdatedEvent) WithoutEventType() *PartyProductUpdatedEvent {
	m.EventType = nil
	return m
}

func (m *PartyProductUpdatedEvent) WithID(id strfmt.UUID) *PartyProductUpdatedEvent {

	m.ID = &id

	return m
}

func (m *PartyProductUpdatedEvent) WithoutID() *PartyProductUpdatedEvent {
	m.ID = nil
	return m
}

func (m *PartyProductUpdatedEvent) WithOrganisationID(organisationID strfmt.UUID) *PartyProductUpdatedEvent {

	m.OrganisationID = &organisationID

	return m
}

func (m *PartyProductUpdatedEvent) WithoutOrganisationID() *PartyProductUpdatedEvent {
	m.OrganisationID = nil
	return m
}

func (m *PartyProductUpdatedEvent) WithPartyID(partyID strfmt.UUID) *PartyProductUpdatedEvent {

	m.PartyID = partyID

	return m
}

func (m *PartyProductUpdatedEvent) WithPartyProductStatus(partyProductStatus PartyProductStatus) *PartyProductUpdatedEvent {

	m.PartyProductStatus = partyProductStatus

	return m
}

func (m *PartyProductUpdatedEvent) WithProductID(productID strfmt.UUID) *PartyProductUpdatedEvent {

	m.ProductID = productID

	return m
}

func (m *PartyProductUpdatedEvent) WithRecordType(recordType ResourceType) *PartyProductUpdatedEvent {

	m.RecordType = recordType

	return m
}

func (m *PartyProductUpdatedEvent) WithStatusItems(statusItems []*PartyProductStatusItem) *PartyProductUpdatedEvent {

	m.StatusItems = statusItems

	return m
}

func (m *PartyProductUpdatedEvent) WithUserID(userID string) *PartyProductUpdatedEvent {

	m.UserID = userID

	return m
}

func (m *PartyProductUpdatedEvent) WithVendorStatus(vendorStatus string) *PartyProductUpdatedEvent {

	m.VendorStatus = vendorStatus

	return m
}

// Validate validates this party product updated event
func (m *PartyProductUpdatedEvent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEventType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganisationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePartyID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePartyProductStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProductID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecordType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatusItems(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var partyProductUpdatedEventTypeEventTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["updated"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		partyProductUpdatedEventTypeEventTypePropEnum = append(partyProductUpdatedEventTypeEventTypePropEnum, v)
	}
}

const (

	// PartyProductUpdatedEventEventTypeUpdated captures enum value "updated"
	PartyProductUpdatedEventEventTypeUpdated string = "updated"
)

// prop value enum
func (m *PartyProductUpdatedEvent) validateEventTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, partyProductUpdatedEventTypeEventTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PartyProductUpdatedEvent) validateEventType(formats strfmt.Registry) error {

	if err := validate.Required("event_type", "body", m.EventType); err != nil {
		return err
	}

	// value enum
	if err := m.validateEventTypeEnum("event_type", "body", *m.EventType); err != nil {
		return err
	}

	return nil
}

func (m *PartyProductUpdatedEvent) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PartyProductUpdatedEvent) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", m.OrganisationID); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PartyProductUpdatedEvent) validatePartyID(formats strfmt.Registry) error {

	if swag.IsZero(m.PartyID) { // not required
		return nil
	}

	if err := validate.FormatOf("party_id", "body", "uuid", m.PartyID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PartyProductUpdatedEvent) validatePartyProductStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.PartyProductStatus) { // not required
		return nil
	}

	if err := m.PartyProductStatus.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("party_product_status")
		}
		return err
	}

	return nil
}

func (m *PartyProductUpdatedEvent) validateProductID(formats strfmt.Registry) error {

	if swag.IsZero(m.ProductID) { // not required
		return nil
	}

	if err := validate.FormatOf("product_id", "body", "uuid", m.ProductID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PartyProductUpdatedEvent) validateRecordType(formats strfmt.Registry) error {

	if err := m.RecordType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("record_type")
		}
		return err
	}

	return nil
}

func (m *PartyProductUpdatedEvent) validateStatusItems(formats strfmt.Registry) error {

	if swag.IsZero(m.StatusItems) { // not required
		return nil
	}

	for i := 0; i < len(m.StatusItems); i++ {
		if swag.IsZero(m.StatusItems[i]) { // not required
			continue
		}

		if m.StatusItems[i] != nil {
			if err := m.StatusItems[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("status_items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PartyProductUpdatedEvent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PartyProductUpdatedEvent) UnmarshalBinary(b []byte) error {
	var res PartyProductUpdatedEvent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *PartyProductUpdatedEvent) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}