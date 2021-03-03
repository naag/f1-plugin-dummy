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

// NewPatchPaymentsIDAdvicesAdviceIDSubmissionsSubmissionIDParams creates a new PatchPaymentsIDAdvicesAdviceIDSubmissionsSubmissionIDParams object
// with the default values initialized.
func NewPatchPaymentsIDAdvicesAdviceIDSubmissionsSubmissionIDParams() *PatchPaymentsIDAdvicesAdviceIDSubmissionsSubmissionIDParams {
	var ()
	return &PatchPaymentsIDAdvicesAdviceIDSubmissionsSubmissionIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchPaymentsIDAdvicesAdviceIDSubmissionsSubmissionIDParamsWithTimeout creates a new PatchPaymentsIDAdvicesAdviceIDSubmissionsSubmissionIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchPaymentsIDAdvicesAdviceIDSubmissionsSubmissionIDParamsWithTimeout(timeout time.Duration) *PatchPaymentsIDAdvicesAdviceIDSubmissionsSubmissionIDParams {
	var ()
	return &PatchPaymentsIDAdvicesAdviceIDSubmissionsSubmissionIDParams{

		timeout: timeout,
	}
}

// NewPatchPaymentsIDAdvicesAdviceIDSubmissionsSubmissionIDParamsWithContext creates a new PatchPaymentsIDAdvicesAdviceIDSubmissionsSubmissionIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchPaymentsIDAdvicesAdviceIDSubmissionsSubmissionIDParamsWithContext(ctx context.Context) *PatchPaymentsIDAdvicesAdviceIDSubmissionsSubmissionIDParams {
	var ()
	return &PatchPaymentsIDAdvicesAdviceIDSubmissionsSubmissionIDParams{

		Context: ctx,
	}
}

// NewPatchPaymentsIDAdvicesAdviceIDSubmissionsSubmissionIDParamsWithHTTPClient creates a new PatchPaymentsIDAdvicesAdviceIDSubmissionsSubmissionIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchPaymentsIDAdvicesAdviceIDSubmissionsSubmissionIDParamsWithHTTPClient(client *http.Client) *PatchPaymentsIDAdvicesAdviceIDSubmissionsSubmissionIDParams {
	var ()
	return &PatchPaymentsIDAdvicesAdviceIDSubmissionsSubmissionIDParams{
		HTTPClient: client,
	}
}

/*PatchPaymentsIDAdvicesAdviceIDSubmissionsSubmissionIDParams contains all the parameters to send to the API endpoint
for the patch payments ID advices advice ID submissions submission ID operation typically these are written to a http.Request
*/
type PatchPaymentsIDAdvicesAdviceIDSubmissionsSubmissionIDParams struct {

	/*AdviceSubmissionUpdateRequest*/
	AdviceSubmissionUpdateRequest *models.AdviceSubmissionAmendment
	/*XForm3TransactionTime*/
	XForm3TransactionTime *string
	/*XForm3CorrelationID*/
	XForm3CorrelationID *string
	/*AdviceID
	  Advice Id

	*/
	AdviceID strfmt.UUID
	/*ID
	  Payment Id

	*/
	ID strfmt.UUID
	/*SubmissionID
	  Submission Id

	*/
	SubmissionID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch payments ID advices advice ID submissions submission ID params
func (o *PatchPaymentsIDAdvicesAdviceIDSubmissionsSubmissionIDParams) WithTimeout(timeout time.Duration) *PatchPaymentsIDAdvicesAdviceIDSubmissionsSubmissionIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch payments ID advices advice ID submissions submission ID params
func (o *PatchPaymentsIDAdvicesAdviceIDSubmissionsSubmissionIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch payments ID advices advice ID submissions submission ID params
func (o *PatchPaymentsIDAdvicesAdviceIDSubmissionsSubmissionIDParams) WithContext(ctx context.Context) *PatchPaymentsIDAdvicesAdviceIDSubmissionsSubmissionIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch payments ID advices advice ID submissions submission ID params
func (o *PatchPaymentsIDAdvicesAdviceIDSubmissionsSubmissionIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch payments ID advices advice ID submissions submission ID params
func (o *PatchPaymentsIDAdvicesAdviceIDSubmissionsSubmissionIDParams) WithHTTPClient(client *http.Client) *PatchPaymentsIDAdvicesAdviceIDSubmissionsSubmissionIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch payments ID advices advice ID submissions submission ID params
func (o *PatchPaymentsIDAdvicesAdviceIDSubmissionsSubmissionIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAdviceSubmissionUpdateRequest adds the adviceSubmissionUpdateRequest to the patch payments ID advices advice ID submissions submission ID params
func (o *PatchPaymentsIDAdvicesAdviceIDSubmissionsSubmissionIDParams) WithAdviceSubmissionUpdateRequest(adviceSubmissionUpdateRequest *models.AdviceSubmissionAmendment) *PatchPaymentsIDAdvicesAdviceIDSubmissionsSubmissionIDParams {
	o.SetAdviceSubmissionUpdateRequest(adviceSubmissionUpdateRequest)
	return o
}

// SetAdviceSubmissionUpdateRequest adds the adviceSubmissionUpdateRequest to the patch payments ID advices advice ID submissions submission ID params
func (o *PatchPaymentsIDAdvicesAdviceIDSubmissionsSubmissionIDParams) SetAdviceSubmissionUpdateRequest(adviceSubmissionUpdateRequest *models.AdviceSubmissionAmendment) {
	o.AdviceSubmissionUpdateRequest = adviceSubmissionUpdateRequest
}

// WithXForm3TransactionTime adds the xForm3TransactionTime to the patch payments ID advices advice ID submissions submission ID params
func (o *PatchPaymentsIDAdvicesAdviceIDSubmissionsSubmissionIDParams) WithXForm3TransactionTime(xForm3TransactionTime *string) *PatchPaymentsIDAdvicesAdviceIDSubmissionsSubmissionIDParams {
	o.SetXForm3TransactionTime(xForm3TransactionTime)
	return o
}

// SetXForm3TransactionTime adds the xForm3TransactionTime to the patch payments ID advices advice ID submissions submission ID params
func (o *PatchPaymentsIDAdvicesAdviceIDSubmissionsSubmissionIDParams) SetXForm3TransactionTime(xForm3TransactionTime *string) {
	o.XForm3TransactionTime = xForm3TransactionTime
}

// WithXForm3CorrelationID adds the xForm3CorrelationID to the patch payments ID advices advice ID submissions submission ID params
func (o *PatchPaymentsIDAdvicesAdviceIDSubmissionsSubmissionIDParams) WithXForm3CorrelationID(xForm3CorrelationID *string) *PatchPaymentsIDAdvicesAdviceIDSubmissionsSubmissionIDParams {
	o.SetXForm3CorrelationID(xForm3CorrelationID)
	return o
}

// SetXForm3CorrelationID adds the xForm3CorrelationId to the patch payments ID advices advice ID submissions submission ID params
func (o *PatchPaymentsIDAdvicesAdviceIDSubmissionsSubmissionIDParams) SetXForm3CorrelationID(xForm3CorrelationID *string) {
	o.XForm3CorrelationID = xForm3CorrelationID
}

// WithAdviceID adds the adviceID to the patch payments ID advices advice ID submissions submission ID params
func (o *PatchPaymentsIDAdvicesAdviceIDSubmissionsSubmissionIDParams) WithAdviceID(adviceID strfmt.UUID) *PatchPaymentsIDAdvicesAdviceIDSubmissionsSubmissionIDParams {
	o.SetAdviceID(adviceID)
	return o
}

// SetAdviceID adds the adviceId to the patch payments ID advices advice ID submissions submission ID params
func (o *PatchPaymentsIDAdvicesAdviceIDSubmissionsSubmissionIDParams) SetAdviceID(adviceID strfmt.UUID) {
	o.AdviceID = adviceID
}

// WithID adds the id to the patch payments ID advices advice ID submissions submission ID params
func (o *PatchPaymentsIDAdvicesAdviceIDSubmissionsSubmissionIDParams) WithID(id strfmt.UUID) *PatchPaymentsIDAdvicesAdviceIDSubmissionsSubmissionIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch payments ID advices advice ID submissions submission ID params
func (o *PatchPaymentsIDAdvicesAdviceIDSubmissionsSubmissionIDParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WithSubmissionID adds the submissionID to the patch payments ID advices advice ID submissions submission ID params
func (o *PatchPaymentsIDAdvicesAdviceIDSubmissionsSubmissionIDParams) WithSubmissionID(submissionID strfmt.UUID) *PatchPaymentsIDAdvicesAdviceIDSubmissionsSubmissionIDParams {
	o.SetSubmissionID(submissionID)
	return o
}

// SetSubmissionID adds the submissionId to the patch payments ID advices advice ID submissions submission ID params
func (o *PatchPaymentsIDAdvicesAdviceIDSubmissionsSubmissionIDParams) SetSubmissionID(submissionID strfmt.UUID) {
	o.SubmissionID = submissionID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchPaymentsIDAdvicesAdviceIDSubmissionsSubmissionIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AdviceSubmissionUpdateRequest != nil {
		if err := r.SetBodyParam(o.AdviceSubmissionUpdateRequest); err != nil {
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

	// path param adviceId
	if err := r.SetPathParam("adviceId", o.AdviceID.String()); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
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