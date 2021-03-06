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

// NewGetPaymentsIDRecallsRecallIDParams creates a new GetPaymentsIDRecallsRecallIDParams object
// with the default values initialized.
func NewGetPaymentsIDRecallsRecallIDParams() *GetPaymentsIDRecallsRecallIDParams {
	var ()
	return &GetPaymentsIDRecallsRecallIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPaymentsIDRecallsRecallIDParamsWithTimeout creates a new GetPaymentsIDRecallsRecallIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPaymentsIDRecallsRecallIDParamsWithTimeout(timeout time.Duration) *GetPaymentsIDRecallsRecallIDParams {
	var ()
	return &GetPaymentsIDRecallsRecallIDParams{

		timeout: timeout,
	}
}

// NewGetPaymentsIDRecallsRecallIDParamsWithContext creates a new GetPaymentsIDRecallsRecallIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPaymentsIDRecallsRecallIDParamsWithContext(ctx context.Context) *GetPaymentsIDRecallsRecallIDParams {
	var ()
	return &GetPaymentsIDRecallsRecallIDParams{

		Context: ctx,
	}
}

// NewGetPaymentsIDRecallsRecallIDParamsWithHTTPClient creates a new GetPaymentsIDRecallsRecallIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPaymentsIDRecallsRecallIDParamsWithHTTPClient(client *http.Client) *GetPaymentsIDRecallsRecallIDParams {
	var ()
	return &GetPaymentsIDRecallsRecallIDParams{
		HTTPClient: client,
	}
}

/*GetPaymentsIDRecallsRecallIDParams contains all the parameters to send to the API endpoint
for the get payments ID recalls recall ID operation typically these are written to a http.Request
*/
type GetPaymentsIDRecallsRecallIDParams struct {

	/*XForm3TransactionTime*/
	XForm3TransactionTime *string
	/*XForm3CorrelationID*/
	XForm3CorrelationID *string
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

// WithTimeout adds the timeout to the get payments ID recalls recall ID params
func (o *GetPaymentsIDRecallsRecallIDParams) WithTimeout(timeout time.Duration) *GetPaymentsIDRecallsRecallIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get payments ID recalls recall ID params
func (o *GetPaymentsIDRecallsRecallIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get payments ID recalls recall ID params
func (o *GetPaymentsIDRecallsRecallIDParams) WithContext(ctx context.Context) *GetPaymentsIDRecallsRecallIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get payments ID recalls recall ID params
func (o *GetPaymentsIDRecallsRecallIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get payments ID recalls recall ID params
func (o *GetPaymentsIDRecallsRecallIDParams) WithHTTPClient(client *http.Client) *GetPaymentsIDRecallsRecallIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get payments ID recalls recall ID params
func (o *GetPaymentsIDRecallsRecallIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXForm3TransactionTime adds the xForm3TransactionTime to the get payments ID recalls recall ID params
func (o *GetPaymentsIDRecallsRecallIDParams) WithXForm3TransactionTime(xForm3TransactionTime *string) *GetPaymentsIDRecallsRecallIDParams {
	o.SetXForm3TransactionTime(xForm3TransactionTime)
	return o
}

// SetXForm3TransactionTime adds the xForm3TransactionTime to the get payments ID recalls recall ID params
func (o *GetPaymentsIDRecallsRecallIDParams) SetXForm3TransactionTime(xForm3TransactionTime *string) {
	o.XForm3TransactionTime = xForm3TransactionTime
}

// WithXForm3CorrelationID adds the xForm3CorrelationID to the get payments ID recalls recall ID params
func (o *GetPaymentsIDRecallsRecallIDParams) WithXForm3CorrelationID(xForm3CorrelationID *string) *GetPaymentsIDRecallsRecallIDParams {
	o.SetXForm3CorrelationID(xForm3CorrelationID)
	return o
}

// SetXForm3CorrelationID adds the xForm3CorrelationId to the get payments ID recalls recall ID params
func (o *GetPaymentsIDRecallsRecallIDParams) SetXForm3CorrelationID(xForm3CorrelationID *string) {
	o.XForm3CorrelationID = xForm3CorrelationID
}

// WithID adds the id to the get payments ID recalls recall ID params
func (o *GetPaymentsIDRecallsRecallIDParams) WithID(id strfmt.UUID) *GetPaymentsIDRecallsRecallIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get payments ID recalls recall ID params
func (o *GetPaymentsIDRecallsRecallIDParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WithRecallID adds the recallID to the get payments ID recalls recall ID params
func (o *GetPaymentsIDRecallsRecallIDParams) WithRecallID(recallID strfmt.UUID) *GetPaymentsIDRecallsRecallIDParams {
	o.SetRecallID(recallID)
	return o
}

// SetRecallID adds the recallId to the get payments ID recalls recall ID params
func (o *GetPaymentsIDRecallsRecallIDParams) SetRecallID(recallID strfmt.UUID) {
	o.RecallID = recallID
}

// WriteToRequest writes these params to a swagger request
func (o *GetPaymentsIDRecallsRecallIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
