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

// NewPostPaymentsIDRecallsRecallIDDecisionsParams creates a new PostPaymentsIDRecallsRecallIDDecisionsParams object
// with the default values initialized.
func NewPostPaymentsIDRecallsRecallIDDecisionsParams() *PostPaymentsIDRecallsRecallIDDecisionsParams {
	var ()
	return &PostPaymentsIDRecallsRecallIDDecisionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostPaymentsIDRecallsRecallIDDecisionsParamsWithTimeout creates a new PostPaymentsIDRecallsRecallIDDecisionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostPaymentsIDRecallsRecallIDDecisionsParamsWithTimeout(timeout time.Duration) *PostPaymentsIDRecallsRecallIDDecisionsParams {
	var ()
	return &PostPaymentsIDRecallsRecallIDDecisionsParams{

		timeout: timeout,
	}
}

// NewPostPaymentsIDRecallsRecallIDDecisionsParamsWithContext creates a new PostPaymentsIDRecallsRecallIDDecisionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostPaymentsIDRecallsRecallIDDecisionsParamsWithContext(ctx context.Context) *PostPaymentsIDRecallsRecallIDDecisionsParams {
	var ()
	return &PostPaymentsIDRecallsRecallIDDecisionsParams{

		Context: ctx,
	}
}

// NewPostPaymentsIDRecallsRecallIDDecisionsParamsWithHTTPClient creates a new PostPaymentsIDRecallsRecallIDDecisionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostPaymentsIDRecallsRecallIDDecisionsParamsWithHTTPClient(client *http.Client) *PostPaymentsIDRecallsRecallIDDecisionsParams {
	var ()
	return &PostPaymentsIDRecallsRecallIDDecisionsParams{
		HTTPClient: client,
	}
}

/*PostPaymentsIDRecallsRecallIDDecisionsParams contains all the parameters to send to the API endpoint
for the post payments ID recalls recall ID decisions operation typically these are written to a http.Request
*/
type PostPaymentsIDRecallsRecallIDDecisionsParams struct {

	/*RecallDecisionCreationRequest*/
	RecallDecisionCreationRequest *models.RecallDecisionCreation
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

// WithTimeout adds the timeout to the post payments ID recalls recall ID decisions params
func (o *PostPaymentsIDRecallsRecallIDDecisionsParams) WithTimeout(timeout time.Duration) *PostPaymentsIDRecallsRecallIDDecisionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post payments ID recalls recall ID decisions params
func (o *PostPaymentsIDRecallsRecallIDDecisionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post payments ID recalls recall ID decisions params
func (o *PostPaymentsIDRecallsRecallIDDecisionsParams) WithContext(ctx context.Context) *PostPaymentsIDRecallsRecallIDDecisionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post payments ID recalls recall ID decisions params
func (o *PostPaymentsIDRecallsRecallIDDecisionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post payments ID recalls recall ID decisions params
func (o *PostPaymentsIDRecallsRecallIDDecisionsParams) WithHTTPClient(client *http.Client) *PostPaymentsIDRecallsRecallIDDecisionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post payments ID recalls recall ID decisions params
func (o *PostPaymentsIDRecallsRecallIDDecisionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRecallDecisionCreationRequest adds the recallDecisionCreationRequest to the post payments ID recalls recall ID decisions params
func (o *PostPaymentsIDRecallsRecallIDDecisionsParams) WithRecallDecisionCreationRequest(recallDecisionCreationRequest *models.RecallDecisionCreation) *PostPaymentsIDRecallsRecallIDDecisionsParams {
	o.SetRecallDecisionCreationRequest(recallDecisionCreationRequest)
	return o
}

// SetRecallDecisionCreationRequest adds the recallDecisionCreationRequest to the post payments ID recalls recall ID decisions params
func (o *PostPaymentsIDRecallsRecallIDDecisionsParams) SetRecallDecisionCreationRequest(recallDecisionCreationRequest *models.RecallDecisionCreation) {
	o.RecallDecisionCreationRequest = recallDecisionCreationRequest
}

// WithXForm3TransactionTime adds the xForm3TransactionTime to the post payments ID recalls recall ID decisions params
func (o *PostPaymentsIDRecallsRecallIDDecisionsParams) WithXForm3TransactionTime(xForm3TransactionTime *string) *PostPaymentsIDRecallsRecallIDDecisionsParams {
	o.SetXForm3TransactionTime(xForm3TransactionTime)
	return o
}

// SetXForm3TransactionTime adds the xForm3TransactionTime to the post payments ID recalls recall ID decisions params
func (o *PostPaymentsIDRecallsRecallIDDecisionsParams) SetXForm3TransactionTime(xForm3TransactionTime *string) {
	o.XForm3TransactionTime = xForm3TransactionTime
}

// WithXForm3CorrelationID adds the xForm3CorrelationID to the post payments ID recalls recall ID decisions params
func (o *PostPaymentsIDRecallsRecallIDDecisionsParams) WithXForm3CorrelationID(xForm3CorrelationID *string) *PostPaymentsIDRecallsRecallIDDecisionsParams {
	o.SetXForm3CorrelationID(xForm3CorrelationID)
	return o
}

// SetXForm3CorrelationID adds the xForm3CorrelationId to the post payments ID recalls recall ID decisions params
func (o *PostPaymentsIDRecallsRecallIDDecisionsParams) SetXForm3CorrelationID(xForm3CorrelationID *string) {
	o.XForm3CorrelationID = xForm3CorrelationID
}

// WithID adds the id to the post payments ID recalls recall ID decisions params
func (o *PostPaymentsIDRecallsRecallIDDecisionsParams) WithID(id strfmt.UUID) *PostPaymentsIDRecallsRecallIDDecisionsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post payments ID recalls recall ID decisions params
func (o *PostPaymentsIDRecallsRecallIDDecisionsParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WithRecallID adds the recallID to the post payments ID recalls recall ID decisions params
func (o *PostPaymentsIDRecallsRecallIDDecisionsParams) WithRecallID(recallID strfmt.UUID) *PostPaymentsIDRecallsRecallIDDecisionsParams {
	o.SetRecallID(recallID)
	return o
}

// SetRecallID adds the recallId to the post payments ID recalls recall ID decisions params
func (o *PostPaymentsIDRecallsRecallIDDecisionsParams) SetRecallID(recallID strfmt.UUID) {
	o.RecallID = recallID
}

// WriteToRequest writes these params to a swagger request
func (o *PostPaymentsIDRecallsRecallIDDecisionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.RecallDecisionCreationRequest != nil {
		if err := r.SetBodyParam(o.RecallDecisionCreationRequest); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
