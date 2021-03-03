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

	models "github.com/form3tech/go-paymentapi-client/models"
)

// NewPostPaymentsParams creates a new PostPaymentsParams object
// with the default values initialized.
func NewPostPaymentsParams() *PostPaymentsParams {
	var ()
	return &PostPaymentsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostPaymentsParamsWithTimeout creates a new PostPaymentsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostPaymentsParamsWithTimeout(timeout time.Duration) *PostPaymentsParams {
	var ()
	return &PostPaymentsParams{

		timeout: timeout,
	}
}

// NewPostPaymentsParamsWithContext creates a new PostPaymentsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostPaymentsParamsWithContext(ctx context.Context) *PostPaymentsParams {
	var ()
	return &PostPaymentsParams{

		Context: ctx,
	}
}

// NewPostPaymentsParamsWithHTTPClient creates a new PostPaymentsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostPaymentsParamsWithHTTPClient(client *http.Client) *PostPaymentsParams {
	var ()
	return &PostPaymentsParams{
		HTTPClient: client,
	}
}

/*PostPaymentsParams contains all the parameters to send to the API endpoint
for the post payments operation typically these are written to a http.Request
*/
type PostPaymentsParams struct {

	/*PaymentCreationRequest*/
	PaymentCreationRequest *models.PaymentCreation
	/*XForm3TransactionTime*/
	XForm3TransactionTime *string
	/*XForm3CorrelationID*/
	XForm3CorrelationID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post payments params
func (o *PostPaymentsParams) WithTimeout(timeout time.Duration) *PostPaymentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post payments params
func (o *PostPaymentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post payments params
func (o *PostPaymentsParams) WithContext(ctx context.Context) *PostPaymentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post payments params
func (o *PostPaymentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post payments params
func (o *PostPaymentsParams) WithHTTPClient(client *http.Client) *PostPaymentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post payments params
func (o *PostPaymentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPaymentCreationRequest adds the paymentCreationRequest to the post payments params
func (o *PostPaymentsParams) WithPaymentCreationRequest(paymentCreationRequest *models.PaymentCreation) *PostPaymentsParams {
	o.SetPaymentCreationRequest(paymentCreationRequest)
	return o
}

// SetPaymentCreationRequest adds the paymentCreationRequest to the post payments params
func (o *PostPaymentsParams) SetPaymentCreationRequest(paymentCreationRequest *models.PaymentCreation) {
	o.PaymentCreationRequest = paymentCreationRequest
}

// WithXForm3TransactionTime adds the xForm3TransactionTime to the post payments params
func (o *PostPaymentsParams) WithXForm3TransactionTime(xForm3TransactionTime *string) *PostPaymentsParams {
	o.SetXForm3TransactionTime(xForm3TransactionTime)
	return o
}

// SetXForm3TransactionTime adds the xForm3TransactionTime to the post payments params
func (o *PostPaymentsParams) SetXForm3TransactionTime(xForm3TransactionTime *string) {
	o.XForm3TransactionTime = xForm3TransactionTime
}

// WithXForm3CorrelationID adds the xForm3CorrelationID to the post payments params
func (o *PostPaymentsParams) WithXForm3CorrelationID(xForm3CorrelationID *string) *PostPaymentsParams {
	o.SetXForm3CorrelationID(xForm3CorrelationID)
	return o
}

// SetXForm3CorrelationID adds the xForm3CorrelationId to the post payments params
func (o *PostPaymentsParams) SetXForm3CorrelationID(xForm3CorrelationID *string) {
	o.XForm3CorrelationID = xForm3CorrelationID
}

// WriteToRequest writes these params to a swagger request
func (o *PostPaymentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.PaymentCreationRequest != nil {
		if err := r.SetBodyParam(o.PaymentCreationRequest); err != nil {
			return err
		}
	}

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
