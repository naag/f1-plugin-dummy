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

// SchemeMessageAdmissionStatus scheme message admission status
// swagger:model SchemeMessageAdmissionStatus
type SchemeMessageAdmissionStatus string

const (

	// SchemeMessageAdmissionStatusDeliveryConfirmed captures enum value "delivery_confirmed"
	SchemeMessageAdmissionStatusDeliveryConfirmed SchemeMessageAdmissionStatus = "delivery_confirmed"

	// SchemeMessageAdmissionStatusFailed captures enum value "failed"
	SchemeMessageAdmissionStatusFailed SchemeMessageAdmissionStatus = "failed"
)

// for schema
var schemeMessageAdmissionStatusEnum []interface{}

func init() {
	var res []SchemeMessageAdmissionStatus
	if err := json.Unmarshal([]byte(`["delivery_confirmed","failed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		schemeMessageAdmissionStatusEnum = append(schemeMessageAdmissionStatusEnum, v)
	}
}

func (m SchemeMessageAdmissionStatus) validateSchemeMessageAdmissionStatusEnum(path, location string, value SchemeMessageAdmissionStatus) error {
	if err := validate.Enum(path, location, value, schemeMessageAdmissionStatusEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this scheme message admission status
func (m SchemeMessageAdmissionStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateSchemeMessageAdmissionStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SchemeMessageAdmissionStatus) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}