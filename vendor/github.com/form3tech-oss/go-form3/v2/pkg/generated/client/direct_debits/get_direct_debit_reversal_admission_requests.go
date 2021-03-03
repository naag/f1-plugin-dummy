// Code generated by go-swagger; DO NOT EDIT.

package direct_debits

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// Client.GetDirectDebitReversalAdmission creates a new GetDirectDebitReversalAdmissionRequest object
// with the default values initialized.
func (c *Client) GetDirectDebitReversalAdmission() *GetDirectDebitReversalAdmissionRequest {
	var ()
	return &GetDirectDebitReversalAdmissionRequest{

		AdmissionID: c.Defaults.GetStrfmtUUID("GetDirectDebitReversalAdmission", "admissionId"),

		ID: c.Defaults.GetStrfmtUUID("GetDirectDebitReversalAdmission", "id"),

		ReversalID: c.Defaults.GetStrfmtUUID("GetDirectDebitReversalAdmission", "reversalId"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type GetDirectDebitReversalAdmissionRequest struct {

	/*AdmissionID      Admission Id      */

	AdmissionID strfmt.UUID

	/*ID      DirectDebit Id      */

	ID strfmt.UUID

	/*ReversalID      Reversal Id      */

	ReversalID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *GetDirectDebitReversalAdmissionRequest) FromJson(j string) *GetDirectDebitReversalAdmissionRequest {

	return o
}

func (o *GetDirectDebitReversalAdmissionRequest) WithAdmissionID(admissionID strfmt.UUID) *GetDirectDebitReversalAdmissionRequest {

	o.AdmissionID = admissionID

	return o
}

func (o *GetDirectDebitReversalAdmissionRequest) WithID(id strfmt.UUID) *GetDirectDebitReversalAdmissionRequest {

	o.ID = id

	return o
}

func (o *GetDirectDebitReversalAdmissionRequest) WithReversalID(reversalID strfmt.UUID) *GetDirectDebitReversalAdmissionRequest {

	o.ReversalID = reversalID

	return o
}

//////////////////
// WithContext adds the context to the get direct debit reversal admission Request
func (o *GetDirectDebitReversalAdmissionRequest) WithContext(ctx context.Context) *GetDirectDebitReversalAdmissionRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the get direct debit reversal admission Request
func (o *GetDirectDebitReversalAdmissionRequest) WithHTTPClient(client *http.Client) *GetDirectDebitReversalAdmissionRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *GetDirectDebitReversalAdmissionRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param admissionId
	if err := r.SetPathParam("admissionId", o.AdmissionID.String()); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
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