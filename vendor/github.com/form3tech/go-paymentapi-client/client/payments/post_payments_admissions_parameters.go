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

// NewPostPaymentsAdmissionsParams creates a new PostPaymentsAdmissionsParams object
// with the default values initialized.
func NewPostPaymentsAdmissionsParams() *PostPaymentsAdmissionsParams {
	var ()
	return &PostPaymentsAdmissionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostPaymentsAdmissionsParamsWithTimeout creates a new PostPaymentsAdmissionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostPaymentsAdmissionsParamsWithTimeout(timeout time.Duration) *PostPaymentsAdmissionsParams {
	var ()
	return &PostPaymentsAdmissionsParams{

		timeout: timeout,
	}
}

// NewPostPaymentsAdmissionsParamsWithContext creates a new PostPaymentsAdmissionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostPaymentsAdmissionsParamsWithContext(ctx context.Context) *PostPaymentsAdmissionsParams {
	var ()
	return &PostPaymentsAdmissionsParams{

		Context: ctx,
	}
}

// NewPostPaymentsAdmissionsParamsWithHTTPClient creates a new PostPaymentsAdmissionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostPaymentsAdmissionsParamsWithHTTPClient(client *http.Client) *PostPaymentsAdmissionsParams {
	var ()
	return &PostPaymentsAdmissionsParams{
		HTTPClient: client,
	}
}

/*PostPaymentsAdmissionsParams contains all the parameters to send to the API endpoint
for the post payments admissions operation typically these are written to a http.Request
*/
type PostPaymentsAdmissionsParams struct {

	/*PaymentWithAdmissionCreationRequest*/
	PaymentWithAdmissionCreationRequest *models.NewPaymentWithAdmissionCreation
	/*XForm3TransactionTime*/
	XForm3TransactionTime *string
	/*XForm3CorrelationID*/
	XForm3CorrelationID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post payments admissions params
func (o *PostPaymentsAdmissionsParams) WithTimeout(timeout time.Duration) *PostPaymentsAdmissionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post payments admissions params
func (o *PostPaymentsAdmissionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post payments admissions params
func (o *PostPaymentsAdmissionsParams) WithContext(ctx context.Context) *PostPaymentsAdmissionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post payments admissions params
func (o *PostPaymentsAdmissionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post payments admissions params
func (o *PostPaymentsAdmissionsParams) WithHTTPClient(client *http.Client) *PostPaymentsAdmissionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post payments admissions params
func (o *PostPaymentsAdmissionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPaymentWithAdmissionCreationRequest adds the paymentWithAdmissionCreationRequest to the post payments admissions params
func (o *PostPaymentsAdmissionsParams) WithPaymentWithAdmissionCreationRequest(paymentWithAdmissionCreationRequest *models.NewPaymentWithAdmissionCreation) *PostPaymentsAdmissionsParams {
	o.SetPaymentWithAdmissionCreationRequest(paymentWithAdmissionCreationRequest)
	return o
}

// SetPaymentWithAdmissionCreationRequest adds the paymentWithAdmissionCreationRequest to the post payments admissions params
func (o *PostPaymentsAdmissionsParams) SetPaymentWithAdmissionCreationRequest(paymentWithAdmissionCreationRequest *models.NewPaymentWithAdmissionCreation) {
	o.PaymentWithAdmissionCreationRequest = paymentWithAdmissionCreationRequest
}

// WithXForm3TransactionTime adds the xForm3TransactionTime to the post payments admissions params
func (o *PostPaymentsAdmissionsParams) WithXForm3TransactionTime(xForm3TransactionTime *string) *PostPaymentsAdmissionsParams {
	o.SetXForm3TransactionTime(xForm3TransactionTime)
	return o
}

// SetXForm3TransactionTime adds the xForm3TransactionTime to the post payments admissions params
func (o *PostPaymentsAdmissionsParams) SetXForm3TransactionTime(xForm3TransactionTime *string) {
	o.XForm3TransactionTime = xForm3TransactionTime
}

// WithXForm3CorrelationID adds the xForm3CorrelationID to the post payments admissions params
func (o *PostPaymentsAdmissionsParams) WithXForm3CorrelationID(xForm3CorrelationID *string) *PostPaymentsAdmissionsParams {
	o.SetXForm3CorrelationID(xForm3CorrelationID)
	return o
}

// SetXForm3CorrelationID adds the xForm3CorrelationId to the post payments admissions params
func (o *PostPaymentsAdmissionsParams) SetXForm3CorrelationID(xForm3CorrelationID *string) {
	o.XForm3CorrelationID = xForm3CorrelationID
}

// WriteToRequest writes these params to a swagger request
func (o *PostPaymentsAdmissionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.PaymentWithAdmissionCreationRequest != nil {
		if err := r.SetBodyParam(o.PaymentWithAdmissionCreationRequest); err != nil {
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
