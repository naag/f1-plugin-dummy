// Code generated by go-swagger; DO NOT EDIT.

package payments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetPaymentsIDRecallsRecallIDDecisionsDecisionIDAdmissionsAdmissionIDParams creates a new GetPaymentsIDRecallsRecallIDDecisionsDecisionIDAdmissionsAdmissionIDParams object
// with the default values initialized.
func NewGetPaymentsIDRecallsRecallIDDecisionsDecisionIDAdmissionsAdmissionIDParams() *GetPaymentsIDRecallsRecallIDDecisionsDecisionIDAdmissionsAdmissionIDParams {
	var ()
	return &GetPaymentsIDRecallsRecallIDDecisionsDecisionIDAdmissionsAdmissionIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPaymentsIDRecallsRecallIDDecisionsDecisionIDAdmissionsAdmissionIDParamsWithTimeout creates a new GetPaymentsIDRecallsRecallIDDecisionsDecisionIDAdmissionsAdmissionIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPaymentsIDRecallsRecallIDDecisionsDecisionIDAdmissionsAdmissionIDParamsWithTimeout(timeout time.Duration) *GetPaymentsIDRecallsRecallIDDecisionsDecisionIDAdmissionsAdmissionIDParams {
	var ()
	return &GetPaymentsIDRecallsRecallIDDecisionsDecisionIDAdmissionsAdmissionIDParams{

		timeout: timeout,
	}
}

// NewGetPaymentsIDRecallsRecallIDDecisionsDecisionIDAdmissionsAdmissionIDParamsWithContext creates a new GetPaymentsIDRecallsRecallIDDecisionsDecisionIDAdmissionsAdmissionIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPaymentsIDRecallsRecallIDDecisionsDecisionIDAdmissionsAdmissionIDParamsWithContext(ctx context.Context) *GetPaymentsIDRecallsRecallIDDecisionsDecisionIDAdmissionsAdmissionIDParams {
	var ()
	return &GetPaymentsIDRecallsRecallIDDecisionsDecisionIDAdmissionsAdmissionIDParams{

		Context: ctx,
	}
}

// NewGetPaymentsIDRecallsRecallIDDecisionsDecisionIDAdmissionsAdmissionIDParamsWithHTTPClient creates a new GetPaymentsIDRecallsRecallIDDecisionsDecisionIDAdmissionsAdmissionIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPaymentsIDRecallsRecallIDDecisionsDecisionIDAdmissionsAdmissionIDParamsWithHTTPClient(client *http.Client) *GetPaymentsIDRecallsRecallIDDecisionsDecisionIDAdmissionsAdmissionIDParams {
	var ()
	return &GetPaymentsIDRecallsRecallIDDecisionsDecisionIDAdmissionsAdmissionIDParams{
		HTTPClient: client,
	}
}

/*GetPaymentsIDRecallsRecallIDDecisionsDecisionIDAdmissionsAdmissionIDParams contains all the parameters to send to the API endpoint
for the get payments ID recalls recall ID decisions decision ID admissions admission ID operation typically these are written to a http.Request
*/
type GetPaymentsIDRecallsRecallIDDecisionsDecisionIDAdmissionsAdmissionIDParams struct {

	/*XForm3TransactionTime*/
	XForm3TransactionTime *string
	/*XForm3CorrelationID*/
	XForm3CorrelationID *string
	/*AdmissionID
	  Admission Id

	*/
	AdmissionID strfmt.UUID
	/*DecisionID
	  Decision Id

	*/
	DecisionID strfmt.UUID
	/*ID
	  Payment Id

	*/
	ID strfmt.UUID
	/*RecallID
	  Recall Id

	*/
	RecallID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get payments ID recalls recall ID decisions decision ID admissions admission ID params
func (o *GetPaymentsIDRecallsRecallIDDecisionsDecisionIDAdmissionsAdmissionIDParams) WithTimeout(timeout time.Duration) *GetPaymentsIDRecallsRecallIDDecisionsDecisionIDAdmissionsAdmissionIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get payments ID recalls recall ID decisions decision ID admissions admission ID params
func (o *GetPaymentsIDRecallsRecallIDDecisionsDecisionIDAdmissionsAdmissionIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get payments ID recalls recall ID decisions decision ID admissions admission ID params
func (o *GetPaymentsIDRecallsRecallIDDecisionsDecisionIDAdmissionsAdmissionIDParams) WithContext(ctx context.Context) *GetPaymentsIDRecallsRecallIDDecisionsDecisionIDAdmissionsAdmissionIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get payments ID recalls recall ID decisions decision ID admissions admission ID params
func (o *GetPaymentsIDRecallsRecallIDDecisionsDecisionIDAdmissionsAdmissionIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get payments ID recalls recall ID decisions decision ID admissions admission ID params
func (o *GetPaymentsIDRecallsRecallIDDecisionsDecisionIDAdmissionsAdmissionIDParams) WithHTTPClient(client *http.Client) *GetPaymentsIDRecallsRecallIDDecisionsDecisionIDAdmissionsAdmissionIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get payments ID recalls recall ID decisions decision ID admissions admission ID params
func (o *GetPaymentsIDRecallsRecallIDDecisionsDecisionIDAdmissionsAdmissionIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXForm3TransactionTime adds the xForm3TransactionTime to the get payments ID recalls recall ID decisions decision ID admissions admission ID params
func (o *GetPaymentsIDRecallsRecallIDDecisionsDecisionIDAdmissionsAdmissionIDParams) WithXForm3TransactionTime(xForm3TransactionTime *string) *GetPaymentsIDRecallsRecallIDDecisionsDecisionIDAdmissionsAdmissionIDParams {
	o.SetXForm3TransactionTime(xForm3TransactionTime)
	return o
}

// SetXForm3TransactionTime adds the xForm3TransactionTime to the get payments ID recalls recall ID decisions decision ID admissions admission ID params
func (o *GetPaymentsIDRecallsRecallIDDecisionsDecisionIDAdmissionsAdmissionIDParams) SetXForm3TransactionTime(xForm3TransactionTime *string) {
	o.XForm3TransactionTime = xForm3TransactionTime
}

// WithXForm3CorrelationID adds the xForm3CorrelationID to the get payments ID recalls recall ID decisions decision ID admissions admission ID params
func (o *GetPaymentsIDRecallsRecallIDDecisionsDecisionIDAdmissionsAdmissionIDParams) WithXForm3CorrelationID(xForm3CorrelationID *string) *GetPaymentsIDRecallsRecallIDDecisionsDecisionIDAdmissionsAdmissionIDParams {
	o.SetXForm3CorrelationID(xForm3CorrelationID)
	return o
}

// SetXForm3CorrelationID adds the xForm3CorrelationId to the get payments ID recalls recall ID decisions decision ID admissions admission ID params
func (o *GetPaymentsIDRecallsRecallIDDecisionsDecisionIDAdmissionsAdmissionIDParams) SetXForm3CorrelationID(xForm3CorrelationID *string) {
	o.XForm3CorrelationID = xForm3CorrelationID
}

// WithAdmissionID adds the admissionID to the get payments ID recalls recall ID decisions decision ID admissions admission ID params
func (o *GetPaymentsIDRecallsRecallIDDecisionsDecisionIDAdmissionsAdmissionIDParams) WithAdmissionID(admissionID strfmt.UUID) *GetPaymentsIDRecallsRecallIDDecisionsDecisionIDAdmissionsAdmissionIDParams {
	o.SetAdmissionID(admissionID)
	return o
}

// SetAdmissionID adds the admissionId to the get payments ID recalls recall ID decisions decision ID admissions admission ID params
func (o *GetPaymentsIDRecallsRecallIDDecisionsDecisionIDAdmissionsAdmissionIDParams) SetAdmissionID(admissionID strfmt.UUID) {
	o.AdmissionID = admissionID
}

// WithDecisionID adds the decisionID to the get payments ID recalls recall ID decisions decision ID admissions admission ID params
func (o *GetPaymentsIDRecallsRecallIDDecisionsDecisionIDAdmissionsAdmissionIDParams) WithDecisionID(decisionID strfmt.UUID) *GetPaymentsIDRecallsRecallIDDecisionsDecisionIDAdmissionsAdmissionIDParams {
	o.SetDecisionID(decisionID)
	return o
}

// SetDecisionID adds the decisionId to the get payments ID recalls recall ID decisions decision ID admissions admission ID params
func (o *GetPaymentsIDRecallsRecallIDDecisionsDecisionIDAdmissionsAdmissionIDParams) SetDecisionID(decisionID strfmt.UUID) {
	o.DecisionID = decisionID
}

// WithID adds the id to the get payments ID recalls recall ID decisions decision ID admissions admission ID params
func (o *GetPaymentsIDRecallsRecallIDDecisionsDecisionIDAdmissionsAdmissionIDParams) WithID(id strfmt.UUID) *GetPaymentsIDRecallsRecallIDDecisionsDecisionIDAdmissionsAdmissionIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get payments ID recalls recall ID decisions decision ID admissions admission ID params
func (o *GetPaymentsIDRecallsRecallIDDecisionsDecisionIDAdmissionsAdmissionIDParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WithRecallID adds the recallID to the get payments ID recalls recall ID decisions decision ID admissions admission ID params
func (o *GetPaymentsIDRecallsRecallIDDecisionsDecisionIDAdmissionsAdmissionIDParams) WithRecallID(recallID strfmt.UUID) *GetPaymentsIDRecallsRecallIDDecisionsDecisionIDAdmissionsAdmissionIDParams {
	o.SetRecallID(recallID)
	return o
}

// SetRecallID adds the recallId to the get payments ID recalls recall ID decisions decision ID admissions admission ID params
func (o *GetPaymentsIDRecallsRecallIDDecisionsDecisionIDAdmissionsAdmissionIDParams) SetRecallID(recallID strfmt.UUID) {
	o.RecallID = recallID
}

// WriteToRequest writes these params to a swagger request
func (o *GetPaymentsIDRecallsRecallIDDecisionsDecisionIDAdmissionsAdmissionIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XForm3TransactionTime != nil {

		// header param X-Form3-Transaction-Time
		if err := r.SetHeaderParam("X-Form3-Transaction-Time", *o.XForm3TransactionTime); err != nil {
			return err
		}

	}

	if o.XForm3CorrelationID != nil {

		// header param X-Form3-correlation-id
		if err := r.SetHeaderParam("X-Form3-correlation-id", *o.XForm3CorrelationID); err != nil {
			return err
		}

	}

	// path param admissionId
	if err := r.SetPathParam("admissionId", o.AdmissionID.String()); err != nil {
		return err
	}

	// path param decisionId
	if err := r.SetPathParam("decisionId", o.DecisionID.String()); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	// path param recallId
	if err := r.SetPathParam("recallId", o.RecallID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}