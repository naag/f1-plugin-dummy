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

// PaymentRelationships payment relationships
// swagger:model PaymentRelationships
type PaymentRelationships struct {

	// beneficiary
	Beneficiary *PaymentRelationshipsBeneficiary `json:"beneficiary,omitempty"`

	// beneficiary account
	BeneficiaryAccount *PaymentRelationshipsBeneficiaryAccount `json:"beneficiary_account,omitempty"`

	// debtor
	Debtor *PaymentRelationshipsDebtor `json:"debtor,omitempty"`

	// debtor account
	DebtorAccount *PaymentRelationshipsDebtorAccount `json:"debtor_account,omitempty"`

	// direct debit
	DirectDebit *PaymentRelationshipsDirectDebit `json:"direct_debit,omitempty"`

	// fx deal
	FxDeal *PaymentRelationshipsFxDeal `json:"fx_deal,omitempty"`

	// payment admission
	PaymentAdmission *PaymentRelationshipsPaymentAdmission `json:"payment_admission,omitempty"`

	// payment advice
	PaymentAdvice *PaymentRelationshipsPaymentAdvice `json:"payment_advice,omitempty"`

	// payment recall
	PaymentRecall *PaymentRelationshipsPaymentRecall `json:"payment_recall,omitempty"`

	// payment return
	PaymentReturn *PaymentRelationshipsPaymentReturn `json:"payment_return,omitempty"`

	// payment reversal
	PaymentReversal *PaymentRelationshipsPaymentReversal `json:"payment_reversal,omitempty"`

	// payment submission
	PaymentSubmission *PaymentRelationshipsPaymentSubmission `json:"payment_submission,omitempty"`
}

func PaymentRelationshipsWithDefaults(defaults client.Defaults) *PaymentRelationships {
	return &PaymentRelationships{

		Beneficiary: PaymentRelationshipsBeneficiaryWithDefaults(defaults),

		BeneficiaryAccount: PaymentRelationshipsBeneficiaryAccountWithDefaults(defaults),

		Debtor: PaymentRelationshipsDebtorWithDefaults(defaults),

		DebtorAccount: PaymentRelationshipsDebtorAccountWithDefaults(defaults),

		DirectDebit: PaymentRelationshipsDirectDebitWithDefaults(defaults),

		FxDeal: PaymentRelationshipsFxDealWithDefaults(defaults),

		PaymentAdmission: PaymentRelationshipsPaymentAdmissionWithDefaults(defaults),

		PaymentAdvice: PaymentRelationshipsPaymentAdviceWithDefaults(defaults),

		PaymentRecall: PaymentRelationshipsPaymentRecallWithDefaults(defaults),

		PaymentReturn: PaymentRelationshipsPaymentReturnWithDefaults(defaults),

		PaymentReversal: PaymentRelationshipsPaymentReversalWithDefaults(defaults),

		PaymentSubmission: PaymentRelationshipsPaymentSubmissionWithDefaults(defaults),
	}
}

func (m *PaymentRelationships) WithBeneficiary(beneficiary PaymentRelationshipsBeneficiary) *PaymentRelationships {

	m.Beneficiary = &beneficiary

	return m
}

func (m *PaymentRelationships) WithoutBeneficiary() *PaymentRelationships {
	m.Beneficiary = nil
	return m
}

func (m *PaymentRelationships) WithBeneficiaryAccount(beneficiaryAccount PaymentRelationshipsBeneficiaryAccount) *PaymentRelationships {

	m.BeneficiaryAccount = &beneficiaryAccount

	return m
}

func (m *PaymentRelationships) WithoutBeneficiaryAccount() *PaymentRelationships {
	m.BeneficiaryAccount = nil
	return m
}

func (m *PaymentRelationships) WithDebtor(debtor PaymentRelationshipsDebtor) *PaymentRelationships {

	m.Debtor = &debtor

	return m
}

func (m *PaymentRelationships) WithoutDebtor() *PaymentRelationships {
	m.Debtor = nil
	return m
}

func (m *PaymentRelationships) WithDebtorAccount(debtorAccount PaymentRelationshipsDebtorAccount) *PaymentRelationships {

	m.DebtorAccount = &debtorAccount

	return m
}

func (m *PaymentRelationships) WithoutDebtorAccount() *PaymentRelationships {
	m.DebtorAccount = nil
	return m
}

func (m *PaymentRelationships) WithDirectDebit(directDebit PaymentRelationshipsDirectDebit) *PaymentRelationships {

	m.DirectDebit = &directDebit

	return m
}

func (m *PaymentRelationships) WithoutDirectDebit() *PaymentRelationships {
	m.DirectDebit = nil
	return m
}

func (m *PaymentRelationships) WithFxDeal(fxDeal PaymentRelationshipsFxDeal) *PaymentRelationships {

	m.FxDeal = &fxDeal

	return m
}

func (m *PaymentRelationships) WithoutFxDeal() *PaymentRelationships {
	m.FxDeal = nil
	return m
}

func (m *PaymentRelationships) WithPaymentAdmission(paymentAdmission PaymentRelationshipsPaymentAdmission) *PaymentRelationships {

	m.PaymentAdmission = &paymentAdmission

	return m
}

func (m *PaymentRelationships) WithoutPaymentAdmission() *PaymentRelationships {
	m.PaymentAdmission = nil
	return m
}

func (m *PaymentRelationships) WithPaymentAdvice(paymentAdvice PaymentRelationshipsPaymentAdvice) *PaymentRelationships {

	m.PaymentAdvice = &paymentAdvice

	return m
}

func (m *PaymentRelationships) WithoutPaymentAdvice() *PaymentRelationships {
	m.PaymentAdvice = nil
	return m
}

func (m *PaymentRelationships) WithPaymentRecall(paymentRecall PaymentRelationshipsPaymentRecall) *PaymentRelationships {

	m.PaymentRecall = &paymentRecall

	return m
}

func (m *PaymentRelationships) WithoutPaymentRecall() *PaymentRelationships {
	m.PaymentRecall = nil
	return m
}

func (m *PaymentRelationships) WithPaymentReturn(paymentReturn PaymentRelationshipsPaymentReturn) *PaymentRelationships {

	m.PaymentReturn = &paymentReturn

	return m
}

func (m *PaymentRelationships) WithoutPaymentReturn() *PaymentRelationships {
	m.PaymentReturn = nil
	return m
}

func (m *PaymentRelationships) WithPaymentReversal(paymentReversal PaymentRelationshipsPaymentReversal) *PaymentRelationships {

	m.PaymentReversal = &paymentReversal

	return m
}

func (m *PaymentRelationships) WithoutPaymentReversal() *PaymentRelationships {
	m.PaymentReversal = nil
	return m
}

func (m *PaymentRelationships) WithPaymentSubmission(paymentSubmission PaymentRelationshipsPaymentSubmission) *PaymentRelationships {

	m.PaymentSubmission = &paymentSubmission

	return m
}

func (m *PaymentRelationships) WithoutPaymentSubmission() *PaymentRelationships {
	m.PaymentSubmission = nil
	return m
}

// Validate validates this payment relationships
func (m *PaymentRelationships) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBeneficiary(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBeneficiaryAccount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDebtor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDebtorAccount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDirectDebit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFxDeal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePaymentAdmission(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePaymentAdvice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePaymentRecall(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePaymentReturn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePaymentReversal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePaymentSubmission(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentRelationships) validateBeneficiary(formats strfmt.Registry) error {

	if swag.IsZero(m.Beneficiary) { // not required
		return nil
	}

	if m.Beneficiary != nil {
		if err := m.Beneficiary.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("beneficiary")
			}
			return err
		}
	}

	return nil
}

func (m *PaymentRelationships) validateBeneficiaryAccount(formats strfmt.Registry) error {

	if swag.IsZero(m.BeneficiaryAccount) { // not required
		return nil
	}

	if m.BeneficiaryAccount != nil {
		if err := m.BeneficiaryAccount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("beneficiary_account")
			}
			return err
		}
	}

	return nil
}

func (m *PaymentRelationships) validateDebtor(formats strfmt.Registry) error {

	if swag.IsZero(m.Debtor) { // not required
		return nil
	}

	if m.Debtor != nil {
		if err := m.Debtor.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("debtor")
			}
			return err
		}
	}

	return nil
}

func (m *PaymentRelationships) validateDebtorAccount(formats strfmt.Registry) error {

	if swag.IsZero(m.DebtorAccount) { // not required
		return nil
	}

	if m.DebtorAccount != nil {
		if err := m.DebtorAccount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("debtor_account")
			}
			return err
		}
	}

	return nil
}

func (m *PaymentRelationships) validateDirectDebit(formats strfmt.Registry) error {

	if swag.IsZero(m.DirectDebit) { // not required
		return nil
	}

	if m.DirectDebit != nil {
		if err := m.DirectDebit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("direct_debit")
			}
			return err
		}
	}

	return nil
}

func (m *PaymentRelationships) validateFxDeal(formats strfmt.Registry) error {

	if swag.IsZero(m.FxDeal) { // not required
		return nil
	}

	if m.FxDeal != nil {
		if err := m.FxDeal.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fx_deal")
			}
			return err
		}
	}

	return nil
}

func (m *PaymentRelationships) validatePaymentAdmission(formats strfmt.Registry) error {

	if swag.IsZero(m.PaymentAdmission) { // not required
		return nil
	}

	if m.PaymentAdmission != nil {
		if err := m.PaymentAdmission.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("payment_admission")
			}
			return err
		}
	}

	return nil
}

func (m *PaymentRelationships) validatePaymentAdvice(formats strfmt.Registry) error {

	if swag.IsZero(m.PaymentAdvice) { // not required
		return nil
	}

	if m.PaymentAdvice != nil {
		if err := m.PaymentAdvice.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("payment_advice")
			}
			return err
		}
	}

	return nil
}

func (m *PaymentRelationships) validatePaymentRecall(formats strfmt.Registry) error {

	if swag.IsZero(m.PaymentRecall) { // not required
		return nil
	}

	if m.PaymentRecall != nil {
		if err := m.PaymentRecall.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("payment_recall")
			}
			return err
		}
	}

	return nil
}

func (m *PaymentRelationships) validatePaymentReturn(formats strfmt.Registry) error {

	if swag.IsZero(m.PaymentReturn) { // not required
		return nil
	}

	if m.PaymentReturn != nil {
		if err := m.PaymentReturn.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("payment_return")
			}
			return err
		}
	}

	return nil
}

func (m *PaymentRelationships) validatePaymentReversal(formats strfmt.Registry) error {

	if swag.IsZero(m.PaymentReversal) { // not required
		return nil
	}

	if m.PaymentReversal != nil {
		if err := m.PaymentReversal.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("payment_reversal")
			}
			return err
		}
	}

	return nil
}

func (m *PaymentRelationships) validatePaymentSubmission(formats strfmt.Registry) error {

	if swag.IsZero(m.PaymentSubmission) { // not required
		return nil
	}

	if m.PaymentSubmission != nil {
		if err := m.PaymentSubmission.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("payment_submission")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PaymentRelationships) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentRelationships) UnmarshalBinary(b []byte) error {
	var res PaymentRelationships
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *PaymentRelationships) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// PaymentRelationshipsBeneficiary payment relationships beneficiary
// swagger:model PaymentRelationshipsBeneficiary
type PaymentRelationshipsBeneficiary struct {

	// Array of beneficiary resources related to the payment
	Data []*RelationshipData `json:"data"`
}

func PaymentRelationshipsBeneficiaryWithDefaults(defaults client.Defaults) *PaymentRelationshipsBeneficiary {
	return &PaymentRelationshipsBeneficiary{

		Data: make([]*RelationshipData, 0),
	}
}

func (m *PaymentRelationshipsBeneficiary) WithData(data []*RelationshipData) *PaymentRelationshipsBeneficiary {

	m.Data = data

	return m
}

// Validate validates this payment relationships beneficiary
func (m *PaymentRelationshipsBeneficiary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentRelationshipsBeneficiary) validateData(formats strfmt.Registry) error {

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
					return ve.ValidateName("beneficiary" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PaymentRelationshipsBeneficiary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentRelationshipsBeneficiary) UnmarshalBinary(b []byte) error {
	var res PaymentRelationshipsBeneficiary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *PaymentRelationshipsBeneficiary) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// PaymentRelationshipsBeneficiaryAccount payment relationships beneficiary account
// swagger:model PaymentRelationshipsBeneficiaryAccount
type PaymentRelationshipsBeneficiaryAccount struct {

	// Array of beneficiary account resources related to the payment
	Data []*RelationshipData `json:"data"`
}

func PaymentRelationshipsBeneficiaryAccountWithDefaults(defaults client.Defaults) *PaymentRelationshipsBeneficiaryAccount {
	return &PaymentRelationshipsBeneficiaryAccount{

		Data: make([]*RelationshipData, 0),
	}
}

func (m *PaymentRelationshipsBeneficiaryAccount) WithData(data []*RelationshipData) *PaymentRelationshipsBeneficiaryAccount {

	m.Data = data

	return m
}

// Validate validates this payment relationships beneficiary account
func (m *PaymentRelationshipsBeneficiaryAccount) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentRelationshipsBeneficiaryAccount) validateData(formats strfmt.Registry) error {

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
					return ve.ValidateName("beneficiary_account" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PaymentRelationshipsBeneficiaryAccount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentRelationshipsBeneficiaryAccount) UnmarshalBinary(b []byte) error {
	var res PaymentRelationshipsBeneficiaryAccount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *PaymentRelationshipsBeneficiaryAccount) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// PaymentRelationshipsDebtor payment relationships debtor
// swagger:model PaymentRelationshipsDebtor
type PaymentRelationshipsDebtor struct {

	// Array of debtor resources related to the payment
	Data []*RelationshipData `json:"data"`
}

func PaymentRelationshipsDebtorWithDefaults(defaults client.Defaults) *PaymentRelationshipsDebtor {
	return &PaymentRelationshipsDebtor{

		Data: make([]*RelationshipData, 0),
	}
}

func (m *PaymentRelationshipsDebtor) WithData(data []*RelationshipData) *PaymentRelationshipsDebtor {

	m.Data = data

	return m
}

// Validate validates this payment relationships debtor
func (m *PaymentRelationshipsDebtor) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentRelationshipsDebtor) validateData(formats strfmt.Registry) error {

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
					return ve.ValidateName("debtor" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PaymentRelationshipsDebtor) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentRelationshipsDebtor) UnmarshalBinary(b []byte) error {
	var res PaymentRelationshipsDebtor
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *PaymentRelationshipsDebtor) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// PaymentRelationshipsDebtorAccount payment relationships debtor account
// swagger:model PaymentRelationshipsDebtorAccount
type PaymentRelationshipsDebtorAccount struct {

	// Array of debtor account resources related to the payment
	Data []*RelationshipData `json:"data"`
}

func PaymentRelationshipsDebtorAccountWithDefaults(defaults client.Defaults) *PaymentRelationshipsDebtorAccount {
	return &PaymentRelationshipsDebtorAccount{

		Data: make([]*RelationshipData, 0),
	}
}

func (m *PaymentRelationshipsDebtorAccount) WithData(data []*RelationshipData) *PaymentRelationshipsDebtorAccount {

	m.Data = data

	return m
}

// Validate validates this payment relationships debtor account
func (m *PaymentRelationshipsDebtorAccount) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentRelationshipsDebtorAccount) validateData(formats strfmt.Registry) error {

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
					return ve.ValidateName("debtor_account" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PaymentRelationshipsDebtorAccount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentRelationshipsDebtorAccount) UnmarshalBinary(b []byte) error {
	var res PaymentRelationshipsDebtorAccount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *PaymentRelationshipsDebtorAccount) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// PaymentRelationshipsDirectDebit payment relationships direct debit
// swagger:model PaymentRelationshipsDirectDebit
type PaymentRelationshipsDirectDebit struct {

	// Array of Direct Debit resources related to the payment
	Data []*RelationshipData `json:"data"`
}

func PaymentRelationshipsDirectDebitWithDefaults(defaults client.Defaults) *PaymentRelationshipsDirectDebit {
	return &PaymentRelationshipsDirectDebit{

		Data: make([]*RelationshipData, 0),
	}
}

func (m *PaymentRelationshipsDirectDebit) WithData(data []*RelationshipData) *PaymentRelationshipsDirectDebit {

	m.Data = data

	return m
}

// Validate validates this payment relationships direct debit
func (m *PaymentRelationshipsDirectDebit) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentRelationshipsDirectDebit) validateData(formats strfmt.Registry) error {

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
					return ve.ValidateName("direct_debit" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PaymentRelationshipsDirectDebit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentRelationshipsDirectDebit) UnmarshalBinary(b []byte) error {
	var res PaymentRelationshipsDirectDebit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *PaymentRelationshipsDirectDebit) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// PaymentRelationshipsFxDeal payment relationships fx deal
// swagger:model PaymentRelationshipsFxDeal
type PaymentRelationshipsFxDeal struct {

	// Array of FXDeal resources related to the payment
	Data []*RelationshipData `json:"data"`
}

func PaymentRelationshipsFxDealWithDefaults(defaults client.Defaults) *PaymentRelationshipsFxDeal {
	return &PaymentRelationshipsFxDeal{

		Data: make([]*RelationshipData, 0),
	}
}

func (m *PaymentRelationshipsFxDeal) WithData(data []*RelationshipData) *PaymentRelationshipsFxDeal {

	m.Data = data

	return m
}

// Validate validates this payment relationships fx deal
func (m *PaymentRelationshipsFxDeal) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentRelationshipsFxDeal) validateData(formats strfmt.Registry) error {

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
					return ve.ValidateName("fx_deal" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PaymentRelationshipsFxDeal) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentRelationshipsFxDeal) UnmarshalBinary(b []byte) error {
	var res PaymentRelationshipsFxDeal
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *PaymentRelationshipsFxDeal) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// PaymentRelationshipsPaymentAdmission payment relationships payment admission
// swagger:model PaymentRelationshipsPaymentAdmission
type PaymentRelationshipsPaymentAdmission struct {

	// Array of Payment Admission resources related to the payment
	Data []*PaymentAdmission `json:"data"`
}

func PaymentRelationshipsPaymentAdmissionWithDefaults(defaults client.Defaults) *PaymentRelationshipsPaymentAdmission {
	return &PaymentRelationshipsPaymentAdmission{

		Data: make([]*PaymentAdmission, 0),
	}
}

func (m *PaymentRelationshipsPaymentAdmission) WithData(data []*PaymentAdmission) *PaymentRelationshipsPaymentAdmission {

	m.Data = data

	return m
}

// Validate validates this payment relationships payment admission
func (m *PaymentRelationshipsPaymentAdmission) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentRelationshipsPaymentAdmission) validateData(formats strfmt.Registry) error {

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
					return ve.ValidateName("payment_admission" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PaymentRelationshipsPaymentAdmission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentRelationshipsPaymentAdmission) UnmarshalBinary(b []byte) error {
	var res PaymentRelationshipsPaymentAdmission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *PaymentRelationshipsPaymentAdmission) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// PaymentRelationshipsPaymentAdvice payment relationships payment advice
// swagger:model PaymentRelationshipsPaymentAdvice
type PaymentRelationshipsPaymentAdvice struct {

	// Array of Payment Advice resources related to the payment
	Data []*PaymentAdvice `json:"data"`
}

func PaymentRelationshipsPaymentAdviceWithDefaults(defaults client.Defaults) *PaymentRelationshipsPaymentAdvice {
	return &PaymentRelationshipsPaymentAdvice{

		Data: make([]*PaymentAdvice, 0),
	}
}

func (m *PaymentRelationshipsPaymentAdvice) WithData(data []*PaymentAdvice) *PaymentRelationshipsPaymentAdvice {

	m.Data = data

	return m
}

// Validate validates this payment relationships payment advice
func (m *PaymentRelationshipsPaymentAdvice) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentRelationshipsPaymentAdvice) validateData(formats strfmt.Registry) error {

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
					return ve.ValidateName("payment_advice" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PaymentRelationshipsPaymentAdvice) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentRelationshipsPaymentAdvice) UnmarshalBinary(b []byte) error {
	var res PaymentRelationshipsPaymentAdvice
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *PaymentRelationshipsPaymentAdvice) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// PaymentRelationshipsPaymentRecall payment relationships payment recall
// swagger:model PaymentRelationshipsPaymentRecall
type PaymentRelationshipsPaymentRecall struct {

	// Array of Payment Recall resources related to the payment
	Data []*Recall `json:"data"`
}

func PaymentRelationshipsPaymentRecallWithDefaults(defaults client.Defaults) *PaymentRelationshipsPaymentRecall {
	return &PaymentRelationshipsPaymentRecall{

		Data: make([]*Recall, 0),
	}
}

func (m *PaymentRelationshipsPaymentRecall) WithData(data []*Recall) *PaymentRelationshipsPaymentRecall {

	m.Data = data

	return m
}

// Validate validates this payment relationships payment recall
func (m *PaymentRelationshipsPaymentRecall) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentRelationshipsPaymentRecall) validateData(formats strfmt.Registry) error {

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
					return ve.ValidateName("payment_recall" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PaymentRelationshipsPaymentRecall) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentRelationshipsPaymentRecall) UnmarshalBinary(b []byte) error {
	var res PaymentRelationshipsPaymentRecall
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *PaymentRelationshipsPaymentRecall) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// PaymentRelationshipsPaymentReturn payment relationships payment return
// swagger:model PaymentRelationshipsPaymentReturn
type PaymentRelationshipsPaymentReturn struct {

	// Array of Payment Return resources related to the payment
	Data []*ReturnPayment `json:"data"`
}

func PaymentRelationshipsPaymentReturnWithDefaults(defaults client.Defaults) *PaymentRelationshipsPaymentReturn {
	return &PaymentRelationshipsPaymentReturn{

		Data: make([]*ReturnPayment, 0),
	}
}

func (m *PaymentRelationshipsPaymentReturn) WithData(data []*ReturnPayment) *PaymentRelationshipsPaymentReturn {

	m.Data = data

	return m
}

// Validate validates this payment relationships payment return
func (m *PaymentRelationshipsPaymentReturn) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentRelationshipsPaymentReturn) validateData(formats strfmt.Registry) error {

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
					return ve.ValidateName("payment_return" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PaymentRelationshipsPaymentReturn) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentRelationshipsPaymentReturn) UnmarshalBinary(b []byte) error {
	var res PaymentRelationshipsPaymentReturn
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *PaymentRelationshipsPaymentReturn) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// PaymentRelationshipsPaymentReversal payment relationships payment reversal
// swagger:model PaymentRelationshipsPaymentReversal
type PaymentRelationshipsPaymentReversal struct {

	// Array of Payment Reversal resources related to the payment
	Data []*ReversalPayment `json:"data"`
}

func PaymentRelationshipsPaymentReversalWithDefaults(defaults client.Defaults) *PaymentRelationshipsPaymentReversal {
	return &PaymentRelationshipsPaymentReversal{

		Data: make([]*ReversalPayment, 0),
	}
}

func (m *PaymentRelationshipsPaymentReversal) WithData(data []*ReversalPayment) *PaymentRelationshipsPaymentReversal {

	m.Data = data

	return m
}

// Validate validates this payment relationships payment reversal
func (m *PaymentRelationshipsPaymentReversal) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentRelationshipsPaymentReversal) validateData(formats strfmt.Registry) error {

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
					return ve.ValidateName("payment_reversal" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PaymentRelationshipsPaymentReversal) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentRelationshipsPaymentReversal) UnmarshalBinary(b []byte) error {
	var res PaymentRelationshipsPaymentReversal
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *PaymentRelationshipsPaymentReversal) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// PaymentRelationshipsPaymentSubmission payment relationships payment submission
// swagger:model PaymentRelationshipsPaymentSubmission
type PaymentRelationshipsPaymentSubmission struct {

	// Array of Payment Submission resources related to the payment
	Data []*PaymentSubmission `json:"data"`
}

func PaymentRelationshipsPaymentSubmissionWithDefaults(defaults client.Defaults) *PaymentRelationshipsPaymentSubmission {
	return &PaymentRelationshipsPaymentSubmission{

		Data: make([]*PaymentSubmission, 0),
	}
}

func (m *PaymentRelationshipsPaymentSubmission) WithData(data []*PaymentSubmission) *PaymentRelationshipsPaymentSubmission {

	m.Data = data

	return m
}

// Validate validates this payment relationships payment submission
func (m *PaymentRelationshipsPaymentSubmission) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentRelationshipsPaymentSubmission) validateData(formats strfmt.Registry) error {

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
					return ve.ValidateName("payment_submission" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PaymentRelationshipsPaymentSubmission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentRelationshipsPaymentSubmission) UnmarshalBinary(b []byte) error {
	var res PaymentRelationshipsPaymentSubmission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *PaymentRelationshipsPaymentSubmission) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
