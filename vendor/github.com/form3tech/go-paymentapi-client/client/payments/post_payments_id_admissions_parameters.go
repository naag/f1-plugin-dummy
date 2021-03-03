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

// NewPostPaymentsIDAdmissionsParams creates a new PostPaymentsIDAdmissionsParams object
// with the default values initialized.
func NewPostPaymentsIDAdmissionsParams() *PostPaymentsIDAdmissionsParams {
	var ()
	return &PostPaymentsIDAdmissionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostPaymentsIDAdmissionsParamsWithTimeout creates a new PostPaymentsIDAdmissionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostPaymentsIDAdmissionsParamsWithTimeout(timeout time.Duration) *PostPaymentsIDAdmissionsParams {
	var ()
	return &PostPaymentsIDAdmissionsParams{

		timeout: timeout,
	}
}

// NewPostPaymentsIDAdmissionsParamsWithContext creates a new PostPaymentsIDAdmissionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostPaymentsIDAdmissionsParamsWithContext(ctx context.Context) *PostPaymentsIDAdmissionsParams {
	var ()
	return &PostPaymentsIDAdmissionsParams{

		Context: ctx,
	}
}

// NewPostPaymentsIDAdmissionsParamsWithHTTPClient creates a new PostPaymentsIDAdmissionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostPaymentsIDAdmissionsParamsWithHTTPClient(client *http.Client) *PostPaymentsIDAdmissionsParams {
	var ()
	return &PostPaymentsIDAdmissionsParams{
		HTTPClient: client,
	}
}

/*PostPaymentsIDAdmissionsParams contains all the parameters to send to the API endpoint
for the post payments ID admissions operation typically these are written to a http.Request
*/
type PostPaymentsIDAdmissionsParams struct {

	/*AdmissionCreationRequest*/
	AdmissionCreationRequest *models.PaymentAdmissionCreation
	/*XForm3TransactionTime*/
	XForm3TransactionTime *string
	/*XForm3CorrelationID*/
	XForm3CorrelationID *string
	/*ID
	  Payment Id

	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post payments ID admissions params
func (o *PostPaymentsIDAdmissionsParams) WithTimeout(timeout time.Duration) *PostPaymentsIDAdmissionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post payments ID admissions params
func (o *PostPaymentsIDAdmissionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post payments ID admissions params
func (o *PostPaymentsIDAdmissionsParams) WithContext(ctx context.Context) *PostPaymentsIDAdmissionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post payments ID admissions params
func (o *PostPaymentsIDAdmissionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post payments ID admissions params
func (o *PostPaymentsIDAdmissionsParams) WithHTTPClient(client *http.Client) *PostPaymentsIDAdmissionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post payments ID admissions params
func (o *PostPaymentsIDAdmissionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAdmissionCreationRequest adds the admissionCreationRequest to the post payments ID admissions params
func (o *PostPaymentsIDAdmissionsParams) WithAdmissionCreationRequest(admissionCreationRequest *models.PaymentAdmissionCreation) *PostPaymentsIDAdmissionsParams {
	o.SetAdmissionCreationRequest(admissionCreationRequest)
	return o
}

// SetAdmissionCreationRequest adds the admissionCreationRequest to the post payments ID admissions params
func (o *PostPaymentsIDAdmissionsParams) SetAdmissionCreationRequest(admissionCreationRequest *models.PaymentAdmissionCreation) {
	o.AdmissionCreationRequest = admissionCreationRequest
}

// WithXForm3TransactionTime adds the xForm3TransactionTime to the post payments ID admissions params
func (o *PostPaymentsIDAdmissionsParams) WithXForm3TransactionTime(xForm3TransactionTime *string) *PostPaymentsIDAdmissionsParams {
	o.SetXForm3TransactionTime(xForm3TransactionTime)
	return o
}

// SetXForm3TransactionTime adds the xForm3TransactionTime to the post payments ID admissions params
func (o *PostPaymentsIDAdmissionsParams) SetXForm3TransactionTime(xForm3TransactionTime *string) {
	o.XForm3TransactionTime = xForm3TransactionTime
}

// WithXForm3CorrelationID adds the xForm3CorrelationID to the post payments ID admissions params
func (o *PostPaymentsIDAdmissionsParams) WithXForm3CorrelationID(xForm3CorrelationID *string) *PostPaymentsIDAdmissionsParams {
	o.SetXForm3CorrelationID(xForm3CorrelationID)
	return o
}

// SetXForm3CorrelationID adds the xForm3CorrelationId to the post payments ID admissions params
func (o *PostPaymentsIDAdmissionsParams) SetXForm3CorrelationID(xForm3CorrelationID *string) {
	o.XForm3CorrelationID = xForm3CorrelationID
}

// WithID adds the id to the post payments ID admissions params
func (o *PostPaymentsIDAdmissionsParams) WithID(id strfmt.UUID) *PostPaymentsIDAdmissionsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post payments ID admissions params
func (o *PostPaymentsIDAdmissionsParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PostPaymentsIDAdmissionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AdmissionCreationRequest != nil {
		if err := r.SetBodyParam(o.AdmissionCreationRequest); err != nil {
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

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}