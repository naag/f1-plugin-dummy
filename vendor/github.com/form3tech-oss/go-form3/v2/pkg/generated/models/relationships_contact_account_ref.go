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
)

// RelationshipsContactAccountRef relationships contact account ref
// swagger:model RelationshipsContactAccountRef
type RelationshipsContactAccountRef struct {

	// data
	Data []*RelationshipsContactAccountRefData `json:"data"`
}

func RelationshipsContactAccountRefWithDefaults(defaults client.Defaults) *RelationshipsContactAccountRef {
	return &RelationshipsContactAccountRef{

		Data: make([]*RelationshipsContactAccountRefData, 0),
	}
}

func (m *RelationshipsContactAccountRef) WithData(data []*RelationshipsContactAccountRefData) *RelationshipsContactAccountRef {

	m.Data = data

	return m
}

// Validate validates this relationships contact account ref
func (m *RelationshipsContactAccountRef) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RelationshipsContactAccountRef) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	for i := 0; i < len(m.Data); i++ {
		if swag.IsZero(m.Data[i]) { // not required
			continue
		}

		if m.Data[i] != nil {
			if err := m.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RelationshipsContactAccountRef) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RelationshipsContactAccountRef) UnmarshalBinary(b []byte) error {
	var res RelationshipsContactAccountRef
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *RelationshipsContactAccountRef) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}