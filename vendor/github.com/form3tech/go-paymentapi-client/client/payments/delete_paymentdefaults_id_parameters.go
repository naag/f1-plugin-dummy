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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeletePaymentdefaultsIDParams creates a new DeletePaymentdefaultsIDParams object
// with the default values initialized.
func NewDeletePaymentdefaultsIDParams() *DeletePaymentdefaultsIDParams {
	var ()
	return &DeletePaymentdefaultsIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeletePaymentdefaultsIDParamsWithTimeout creates a new DeletePaymentdefaultsIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeletePaymentdefaultsIDParamsWithTimeout(timeout time.Duration) *DeletePaymentdefaultsIDParams {
	var ()
	return &DeletePaymentdefaultsIDParams{

		timeout: timeout,
	}
}

// NewDeletePaymentdefaultsIDParamsWithContext creates a new DeletePaymentdefaultsIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeletePaymentdefaultsIDParamsWithContext(ctx context.Context) *DeletePaymentdefaultsIDParams {
	var ()
	return &DeletePaymentdefaultsIDParams{

		Context: ctx,
	}
}

// NewDeletePaymentdefaultsIDParamsWithHTTPClient creates a new DeletePaymentdefaultsIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeletePaymentdefaultsIDParamsWithHTTPClient(client *http.Client) *DeletePaymentdefaultsIDParams {
	var ()
	return &DeletePaymentdefaultsIDParams{
		HTTPClient: client,
	}
}

/*DeletePaymentdefaultsIDParams contains all the parameters to send to the API endpoint
for the delete paymentdefaults ID operation typically these are written to a http.Request
*/
type DeletePaymentdefaultsIDParams struct {

	/*XForm3TransactionTime*/
	XForm3TransactionTime *string
	/*XForm3CorrelationID*/
	XForm3CorrelationID *string
	/*ID
	  Limit Id

	*/
	ID strfmt.UUID
	/*Version
	  Version

	*/
	Version int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete paymentdefaults ID params
func (o *DeletePaymentdefaultsIDParams) WithTimeout(timeout time.Duration) *DeletePaymentdefaultsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete paymentdefaults ID params
func (o *DeletePaymentdefaultsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete paymentdefaults ID params
func (o *DeletePaymentdefaultsIDParams) WithContext(ctx context.Context) *DeletePaymentdefaultsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete paymentdefaults ID params
func (o *DeletePaymentdefaultsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete paymentdefaults ID params
func (o *DeletePaymentdefaultsIDParams) WithHTTPClient(client *http.Client) *DeletePaymentdefaultsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete paymentdefaults ID params
func (o *DeletePaymentdefaultsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXForm3TransactionTime adds the xForm3TransactionTime to the delete paymentdefaults ID params
func (o *DeletePaymentdefaultsIDParams) WithXForm3TransactionTime(xForm3TransactionTime *string) *DeletePaymentdefaultsIDParams {
	o.SetXForm3TransactionTime(xForm3TransactionTime)
	return o
}

// SetXForm3TransactionTime adds the xForm3TransactionTime to the delete paymentdefaults ID params
func (o *DeletePaymentdefaultsIDParams) SetXForm3TransactionTime(xForm3TransactionTime *string) {
	o.XForm3TransactionTime = xForm3TransactionTime
}

// WithXForm3CorrelationID adds the xForm3CorrelationID to the delete paymentdefaults ID params
func (o *DeletePaymentdefaultsIDParams) WithXForm3CorrelationID(xForm3CorrelationID *string) *DeletePaymentdefaultsIDParams {
	o.SetXForm3CorrelationID(xForm3CorrelationID)
	return o
}

// SetXForm3CorrelationID adds the xForm3CorrelationId to the delete paymentdefaults ID params
func (o *DeletePaymentdefaultsIDParams) SetXForm3CorrelationID(xForm3CorrelationID *string) {
	o.XForm3CorrelationID = xForm3CorrelationID
}

// WithID adds the id to the delete paymentdefaults ID params
func (o *DeletePaymentdefaultsIDParams) WithID(id strfmt.UUID) *DeletePaymentdefaultsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete paymentdefaults ID params
func (o *DeletePaymentdefaultsIDParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WithVersion adds the version to the delete paymentdefaults ID params
func (o *DeletePaymentdefaultsIDParams) WithVersion(version int64) *DeletePaymentdefaultsIDParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the delete paymentdefaults ID params
func (o *DeletePaymentdefaultsIDParams) SetVersion(version int64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *DeletePaymentdefaultsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// query param version
	qrVersion := o.Version
	qVersion := swag.FormatInt64(qrVersion)
	if qVersion != "" {
		if err := r.SetQueryParam("version", qVersion); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
