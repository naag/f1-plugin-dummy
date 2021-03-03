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

// NewPostPaymentsIDRecallsRecallIDReversalsReversalIDAdmissionsParams creates a new PostPaymentsIDRecallsRecallIDReversalsReversalIDAdmissionsParams object
// with the default values initialized.
func NewPostPaymentsIDRecallsRecallIDReversalsReversalIDAdmissionsParams() *PostPaymentsIDRecallsRecallIDReversalsReversalIDAdmissionsParams {
	var ()
	return &PostPaymentsIDRecallsRecallIDReversalsReversalIDAdmissionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostPaymentsIDRecallsRecallIDReversalsReversalIDAdmissionsParamsWithTimeout creates a new PostPaymentsIDRecallsRecallIDReversalsReversalIDAdmissionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostPaymentsIDRecallsRecallIDReversalsReversalIDAdmissionsParamsWithTimeout(timeout time.Duration) *PostPaymentsIDRecallsRecallIDReversalsReversalIDAdmissionsParams {
	var ()
	return &PostPaymentsIDRecallsRecallIDReversalsReversalIDAdmissionsParams{

		timeout: timeout,
	}
}

// NewPostPaymentsIDRecallsRecallIDReversalsReversalIDAdmissionsParamsWithContext creates a new PostPaymentsIDRecallsRecallIDReversalsReversalIDAdmissionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostPaymentsIDRecallsRecallIDReversalsReversalIDAdmissionsParamsWithContext(ctx context.Context) *PostPaymentsIDRecallsRecallIDReversalsReversalIDAdmissionsParams {
	var ()
	return &PostPaymentsIDRecallsRecallIDReversalsReversalIDAdmissionsParams{

		Context: ctx,
	}
}

// NewPostPaymentsIDRecallsRecallIDReversalsReversalIDAdmissionsParamsWithHTTPClient creates a new PostPaymentsIDRecallsRecallIDReversalsReversalIDAdmissionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostPaymentsIDRecallsRecallIDReversalsReversalIDAdmissionsParamsWithHTTPClient(client *http.Client) *PostPaymentsIDRecallsRecallIDReversalsReversalIDAdmissionsParams {
	var ()
	return &PostPaymentsIDRecallsRecallIDReversalsReversalIDAdmissionsParams{
		HTTPClient: client,
	}
}

/*PostPaymentsIDRecallsRecallIDReversalsReversalIDAdmissionsParams contains all the parameters to send to the API endpoint
for the post payments ID recalls recall ID reversals reversal ID admissions operation typically these are written to a http.Request
*/
type PostPaymentsIDRecallsRecallIDReversalsReversalIDAdmissionsParams struct {

	/*RecallReversalAdmissionCreationRequest*/
	RecallReversalAdmissionCreationRequest *models.RecallReversalAdmissionCreation
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
	/*ReversalID
	  Reversal Id

	*/
	ReversalID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post payments ID recalls recall ID reversals reversal ID admissions params
func (o *PostPaymentsIDRecallsRecallIDReversalsReversalIDAdmissionsParams) WithTimeout(timeout time.Duration) *PostPaymentsIDRecallsRecallIDReversalsReversalIDAdmissionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post payments ID recalls recall ID reversals reversal ID admissions params
func (o *PostPaymentsIDRecallsRecallIDReversalsReversalIDAdmissionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post payments ID recalls recall ID reversals reversal ID admissions params
func (o *PostPaymentsIDRecallsRecallIDReversalsReversalIDAdmissionsParams) WithContext(ctx context.Context) *PostPaymentsIDRecallsRecallIDReversalsReversalIDAdmissionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post payments ID recalls recall ID reversals reversal ID admissions params
func (o *PostPaymentsIDRecallsRecallIDReversalsReversalIDAdmissionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post payments ID recalls recall ID reversals reversal ID admissions params
func (o *PostPaymentsIDRecallsRecallIDReversalsReversalIDAdmissionsParams) WithHTTPClient(client *http.Client) *PostPaymentsIDRecallsRecallIDReversalsReversalIDAdmissionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post payments ID recalls recall ID reversals reversal ID admissions params
func (o *PostPaymentsIDRecallsRecallIDReversalsReversalIDAdmissionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRecallReversalAdmissionCreationRequest adds the recallReversalAdmissionCreationRequest to the post payments ID recalls recall ID reversals reversal ID admissions params
func (o *PostPaymentsIDRecallsRecallIDReversalsReversalIDAdmissionsParams) WithRecallReversalAdmissionCreationRequest(recallReversalAdmissionCreationRequest *models.RecallReversalAdmissionCreation) *PostPaymentsIDRecallsRecallIDReversalsReversalIDAdmissionsParams {
	o.SetRecallReversalAdmissionCreationRequest(recallReversalAdmissionCreationRequest)
	return o
}

// SetRecallReversalAdmissionCreationRequest adds the recallReversalAdmissionCreationRequest to the post payments ID recalls recall ID reversals reversal ID admissions params
func (o *PostPaymentsIDRecallsRecallIDReversalsReversalIDAdmissionsParams) SetRecallReversalAdmissionCreationRequest(recallReversalAdmissionCreationRequest *models.RecallReversalAdmissionCreation) {
	o.RecallReversalAdmissionCreationRequest = recallReversalAdmissionCreationRequest
}

// WithXForm3TransactionTime adds the xForm3TransactionTime to the post payments ID recalls recall ID reversals reversal ID admissions params
func (o *PostPaymentsIDRecallsRecallIDReversalsReversalIDAdmissionsParams) WithXForm3TransactionTime(xForm3TransactionTime *string) *PostPaymentsIDRecallsRecallIDReversalsReversalIDAdmissionsParams {
	o.SetXForm3TransactionTime(xForm3TransactionTime)
	return o
}

// SetXForm3TransactionTime adds the xForm3TransactionTime to the post payments ID recalls recall ID reversals reversal ID admissions params
func (o *PostPaymentsIDRecallsRecallIDReversalsReversalIDAdmissionsParams) SetXForm3TransactionTime(xForm3TransactionTime *string) {
	o.XForm3TransactionTime = xForm3TransactionTime
}

// WithXForm3CorrelationID adds the xForm3CorrelationID to the post payments ID recalls recall ID reversals reversal ID admissions params
func (o *PostPaymentsIDRecallsRecallIDReversalsReversalIDAdmissionsParams) WithXForm3CorrelationID(xForm3CorrelationID *string) *PostPaymentsIDRecallsRecallIDReversalsReversalIDAdmissionsParams {
	o.SetXForm3CorrelationID(xForm3CorrelationID)
	return o
}

// SetXForm3CorrelationID adds the xForm3CorrelationId to the post payments ID recalls recall ID reversals reversal ID admissions params
func (o *PostPaymentsIDRecallsRecallIDReversalsReversalIDAdmissionsParams) SetXForm3CorrelationID(xForm3CorrelationID *string) {
	o.XForm3CorrelationID = xForm3CorrelationID
}

// WithID adds the id to the post payments ID recalls recall ID reversals reversal ID admissions params
func (o *PostPaymentsIDRecallsRecallIDReversalsReversalIDAdmissionsParams) WithID(id strfmt.UUID) *PostPaymentsIDRecallsRecallIDReversalsReversalIDAdmissionsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post payments ID recalls recall ID reversals reversal ID admissions params
func (o *PostPaymentsIDRecallsRecallIDReversalsReversalIDAdmissionsParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WithRecallID adds the recallID to the post payments ID recalls recall ID reversals reversal ID admissions params
func (o *PostPaymentsIDRecallsRecallIDReversalsReversalIDAdmissionsParams) WithRecallID(recallID strfmt.UUID) *PostPaymentsIDRecallsRecallIDReversalsReversalIDAdmissionsParams {
	o.SetRecallID(recallID)
	return o
}

// SetRecallID adds the recallId to the post payments ID recalls recall ID reversals reversal ID admissions params
func (o *PostPaymentsIDRecallsRecallIDReversalsReversalIDAdmissionsParams) SetRecallID(recallID strfmt.UUID) {
	o.RecallID = recallID
}

// WithReversalID adds the reversalID to the post payments ID recalls recall ID reversals reversal ID admissions params
func (o *PostPaymentsIDRecallsRecallIDReversalsReversalIDAdmissionsParams) WithReversalID(reversalID strfmt.UUID) *PostPaymentsIDRecallsRecallIDReversalsReversalIDAdmissionsParams {
	o.SetReversalID(reversalID)
	return o
}

// SetReversalID adds the reversalId to the post payments ID recalls recall ID reversals reversal ID admissions params
func (o *PostPaymentsIDRecallsRecallIDReversalsReversalIDAdmissionsParams) SetReversalID(reversalID strfmt.UUID) {
	o.ReversalID = reversalID
}

// WriteToRequest writes these params to a swagger request
func (o *PostPaymentsIDRecallsRecallIDReversalsReversalIDAdmissionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.RecallReversalAdmissionCreationRequest != nil {
		if err := r.SetBodyParam(o.RecallReversalAdmissionCreationRequest); err != nil {
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

	// path param recallId
	if err := r.SetPathParam("recallId", o.RecallID.String()); err != nil {
		return err
	}

	// path param reversalId
	if err := r.SetPathParam("reversalId", o.ReversalID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}