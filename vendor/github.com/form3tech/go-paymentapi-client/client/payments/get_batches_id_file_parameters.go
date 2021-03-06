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

// NewGetBatchesIDFileParams creates a new GetBatchesIDFileParams object
// with the default values initialized.
func NewGetBatchesIDFileParams() *GetBatchesIDFileParams {
	var ()
	return &GetBatchesIDFileParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetBatchesIDFileParamsWithTimeout creates a new GetBatchesIDFileParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetBatchesIDFileParamsWithTimeout(timeout time.Duration) *GetBatchesIDFileParams {
	var ()
	return &GetBatchesIDFileParams{

		timeout: timeout,
	}
}

// NewGetBatchesIDFileParamsWithContext creates a new GetBatchesIDFileParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetBatchesIDFileParamsWithContext(ctx context.Context) *GetBatchesIDFileParams {
	var ()
	return &GetBatchesIDFileParams{

		Context: ctx,
	}
}

// NewGetBatchesIDFileParamsWithHTTPClient creates a new GetBatchesIDFileParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetBatchesIDFileParamsWithHTTPClient(client *http.Client) *GetBatchesIDFileParams {
	var ()
	return &GetBatchesIDFileParams{
		HTTPClient: client,
	}
}

/*GetBatchesIDFileParams contains all the parameters to send to the API endpoint
for the get batches ID file operation typically these are written to a http.Request
*/
type GetBatchesIDFileParams struct {

	/*XForm3TransactionTime*/
	XForm3TransactionTime *string
	/*XForm3CorrelationID*/
	XForm3CorrelationID *string
	/*ID
	  Batch Id

	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get batches ID file params
func (o *GetBatchesIDFileParams) WithTimeout(timeout time.Duration) *GetBatchesIDFileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get batches ID file params
func (o *GetBatchesIDFileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get batches ID file params
func (o *GetBatchesIDFileParams) WithContext(ctx context.Context) *GetBatchesIDFileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get batches ID file params
func (o *GetBatchesIDFileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get batches ID file params
func (o *GetBatchesIDFileParams) WithHTTPClient(client *http.Client) *GetBatchesIDFileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get batches ID file params
func (o *GetBatchesIDFileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXForm3TransactionTime adds the xForm3TransactionTime to the get batches ID file params
func (o *GetBatchesIDFileParams) WithXForm3TransactionTime(xForm3TransactionTime *string) *GetBatchesIDFileParams {
	o.SetXForm3TransactionTime(xForm3TransactionTime)
	return o
}

// SetXForm3TransactionTime adds the xForm3TransactionTime to the get batches ID file params
func (o *GetBatchesIDFileParams) SetXForm3TransactionTime(xForm3TransactionTime *string) {
	o.XForm3TransactionTime = xForm3TransactionTime
}

// WithXForm3CorrelationID adds the xForm3CorrelationID to the get batches ID file params
func (o *GetBatchesIDFileParams) WithXForm3CorrelationID(xForm3CorrelationID *string) *GetBatchesIDFileParams {
	o.SetXForm3CorrelationID(xForm3CorrelationID)
	return o
}

// SetXForm3CorrelationID adds the xForm3CorrelationId to the get batches ID file params
func (o *GetBatchesIDFileParams) SetXForm3CorrelationID(xForm3CorrelationID *string) {
	o.XForm3CorrelationID = xForm3CorrelationID
}

// WithID adds the id to the get batches ID file params
func (o *GetBatchesIDFileParams) WithID(id strfmt.UUID) *GetBatchesIDFileParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get batches ID file params
func (o *GetBatchesIDFileParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetBatchesIDFileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
