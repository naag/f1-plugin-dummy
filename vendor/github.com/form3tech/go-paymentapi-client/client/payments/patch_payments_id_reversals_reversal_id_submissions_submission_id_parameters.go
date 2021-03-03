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

// NewPatchPaymentsIDReversalsReversalIDSubmissionsSubmissionIDParams creates a new PatchPaymentsIDReversalsReversalIDSubmissionsSubmissionIDParams object
// with the default values initialized.
func NewPatchPaymentsIDReversalsReversalIDSubmissionsSubmissionIDParams() *PatchPaymentsIDReversalsReversalIDSubmissionsSubmissionIDParams {
	var ()
	return &PatchPaymentsIDReversalsReversalIDSubmissionsSubmissionIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchPaymentsIDReversalsReversalIDSubmissionsSubmissionIDParamsWithTimeout creates a new PatchPaymentsIDReversalsReversalIDSubmissionsSubmissionIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchPaymentsIDReversalsReversalIDSubmissionsSubmissionIDParamsWithTimeout(timeout time.Duration) *PatchPaymentsIDReversalsReversalIDSubmissionsSubmissionIDParams {
	var ()
	return &PatchPaymentsIDReversalsReversalIDSubmissionsSubmissionIDParams{

		timeout: timeout,
	}
}

// NewPatchPaymentsIDReversalsReversalIDSubmissionsSubmissionIDParamsWithContext creates a new PatchPaymentsIDReversalsReversalIDSubmissionsSubmissionIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchPaymentsIDReversalsReversalIDSubmissionsSubmissionIDParamsWithContext(ctx context.Context) *PatchPaymentsIDReversalsReversalIDSubmissionsSubmissionIDParams {
	var ()
	return &PatchPaymentsIDReversalsReversalIDSubmissionsSubmissionIDParams{

		Context: ctx,
	}
}

// NewPatchPaymentsIDReversalsReversalIDSubmissionsSubmissionIDParamsWithHTTPClient creates a new PatchPaymentsIDReversalsReversalIDSubmissionsSubmissionIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchPaymentsIDReversalsReversalIDSubmissionsSubmissionIDParamsWithHTTPClient(client *http.Client) *PatchPaymentsIDReversalsReversalIDSubmissionsSubmissionIDParams {
	var ()
	return &PatchPaymentsIDReversalsReversalIDSubmissionsSubmissionIDParams{
		HTTPClient: client,
	}
}

/*PatchPaymentsIDReversalsReversalIDSubmissionsSubmissionIDParams contains all the parameters to send to the API endpoint
for the patch payments ID reversals reversal ID submissions submission ID operation typically these are written to a http.Request
*/
type PatchPaymentsIDReversalsReversalIDSubmissionsSubmissionIDParams struct {

	/*ReversalSubmissionUpdateRequest*/
	ReversalSubmissionUpdateRequest *models.ReversalSubmissionAmendment
	/*XForm3TransactionTime*/
	XForm3TransactionTime *string
	/*XForm3CorrelationID*/
	XForm3CorrelationID *string
	/*ID
	  Payment Id

	*/
	ID strfmt.UUID
	/*ReversalID
	  Reversal Id

	*/
	ReversalID strfmt.UUID
	/*SubmissionID
	  Submission Id

	*/
	SubmissionID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch payments ID reversals reversal ID submissions submission ID params
func (o *PatchPaymentsIDReversalsReversalIDSubmissionsSubmissionIDParams) WithTimeout(timeout time.Duration) *PatchPaymentsIDReversalsReversalIDSubmissionsSubmissionIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch payments ID reversals reversal ID submissions submission ID params
func (o *PatchPaymentsIDReversalsReversalIDSubmissionsSubmissionIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch payments ID reversals reversal ID submissions submission ID params
func (o *PatchPaymentsIDReversalsReversalIDSubmissionsSubmissionIDParams) WithContext(ctx context.Context) *PatchPaymentsIDReversalsReversalIDSubmissionsSubmissionIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch payments ID reversals reversal ID submissions submission ID params
func (o *PatchPaymentsIDReversalsReversalIDSubmissionsSubmissionIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch payments ID reversals reversal ID submissions submission ID params
func (o *PatchPaymentsIDReversalsReversalIDSubmissionsSubmissionIDParams) WithHTTPClient(client *http.Client) *PatchPaymentsIDReversalsReversalIDSubmissionsSubmissionIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch payments ID reversals reversal ID submissions submission ID params
func (o *PatchPaymentsIDReversalsReversalIDSubmissionsSubmissionIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithReversalSubmissionUpdateRequest adds the reversalSubmissionUpdateRequest to the patch payments ID reversals reversal ID submissions submission ID params
func (o *PatchPaymentsIDReversalsReversalIDSubmissionsSubmissionIDParams) WithReversalSubmissionUpdateRequest(reversalSubmissionUpdateRequest *models.ReversalSubmissionAmendment) *PatchPaymentsIDReversalsReversalIDSubmissionsSubmissionIDParams {
	o.SetReversalSubmissionUpdateRequest(reversalSubmissionUpdateRequest)
	return o
}

// SetReversalSubmissionUpdateRequest adds the reversalSubmissionUpdateRequest to the patch payments ID reversals reversal ID submissions submission ID params
func (o *PatchPaymentsIDReversalsReversalIDSubmissionsSubmissionIDParams) SetReversalSubmissionUpdateRequest(reversalSubmissionUpdateRequest *models.ReversalSubmissionAmendment) {
	o.ReversalSubmissionUpdateRequest = reversalSubmissionUpdateRequest
}

// WithXForm3TransactionTime adds the xForm3TransactionTime to the patch payments ID reversals reversal ID submissions submission ID params
func (o *PatchPaymentsIDReversalsReversalIDSubmissionsSubmissionIDParams) WithXForm3TransactionTime(xForm3TransactionTime *string) *PatchPaymentsIDReversalsReversalIDSubmissionsSubmissionIDParams {
	o.SetXForm3TransactionTime(xForm3TransactionTime)
	return o
}

// SetXForm3TransactionTime adds the xForm3TransactionTime to the patch payments ID reversals reversal ID submissions submission ID params
func (o *PatchPaymentsIDReversalsReversalIDSubmissionsSubmissionIDParams) SetXForm3TransactionTime(xForm3TransactionTime *string) {
	o.XForm3TransactionTime = xForm3TransactionTime
}

// WithXForm3CorrelationID adds the xForm3CorrelationID to the patch payments ID reversals reversal ID submissions submission ID params
func (o *PatchPaymentsIDReversalsReversalIDSubmissionsSubmissionIDParams) WithXForm3CorrelationID(xForm3CorrelationID *string) *PatchPaymentsIDReversalsReversalIDSubmissionsSubmissionIDParams {
	o.SetXForm3CorrelationID(xForm3CorrelationID)
	return o
}

// SetXForm3CorrelationID adds the xForm3CorrelationId to the patch payments ID reversals reversal ID submissions submission ID params
func (o *PatchPaymentsIDReversalsReversalIDSubmissionsSubmissionIDParams) SetXForm3CorrelationID(xForm3CorrelationID *string) {
	o.XForm3CorrelationID = xForm3CorrelationID
}

// WithID adds the id to the patch payments ID reversals reversal ID submissions submission ID params
func (o *PatchPaymentsIDReversalsReversalIDSubmissionsSubmissionIDParams) WithID(id strfmt.UUID) *PatchPaymentsIDReversalsReversalIDSubmissionsSubmissionIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch payments ID reversals reversal ID submissions submission ID params
func (o *PatchPaymentsIDReversalsReversalIDSubmissionsSubmissionIDParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WithReversalID adds the reversalID to the patch payments ID reversals reversal ID submissions submission ID params
func (o *PatchPaymentsIDReversalsReversalIDSubmissionsSubmissionIDParams) WithReversalID(reversalID strfmt.UUID) *PatchPaymentsIDReversalsReversalIDSubmissionsSubmissionIDParams {
	o.SetReversalID(reversalID)
	return o
}

// SetReversalID adds the reversalId to the patch payments ID reversals reversal ID submissions submission ID params
func (o *PatchPaymentsIDReversalsReversalIDSubmissionsSubmissionIDParams) SetReversalID(reversalID strfmt.UUID) {
	o.ReversalID = reversalID
}

// WithSubmissionID adds the submissionID to the patch payments ID reversals reversal ID submissions submission ID params
func (o *PatchPaymentsIDReversalsReversalIDSubmissionsSubmissionIDParams) WithSubmissionID(submissionID strfmt.UUID) *PatchPaymentsIDReversalsReversalIDSubmissionsSubmissionIDParams {
	o.SetSubmissionID(submissionID)
	return o
}

// SetSubmissionID adds the submissionId to the patch payments ID reversals reversal ID submissions submission ID params
func (o *PatchPaymentsIDReversalsReversalIDSubmissionsSubmissionIDParams) SetSubmissionID(submissionID strfmt.UUID) {
	o.SubmissionID = submissionID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchPaymentsIDReversalsReversalIDSubmissionsSubmissionIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ReversalSubmissionUpdateRequest != nil {
		if err := r.SetBodyParam(o.ReversalSubmissionUpdateRequest); err != nil {
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

	// path param reversalId
	if err := r.SetPathParam("reversalId", o.ReversalID.String()); err != nil {
		return err
	}

	// path param submissionId
	if err := r.SetPathParam("submissionId", o.SubmissionID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}