// Code generated by go-swagger; DO NOT EDIT.

package mandates

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// Client.GetMandateAdmission creates a new GetMandateAdmissionRequest object
// with the default values initialized.
func (c *Client) GetMandateAdmission() *GetMandateAdmissionRequest {
	var ()
	return &GetMandateAdmissionRequest{

		AdmissionID: c.Defaults.GetStrfmtUUID("GetMandateAdmission", "admissionId"),

		ID: c.Defaults.GetStrfmtUUID("GetMandateAdmission", "id"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type GetMandateAdmissionRequest struct {

	/*AdmissionID      Mandate Admission Id      */

	AdmissionID strfmt.UUID

	/*ID      Mandate Id      */

	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *GetMandateAdmissionRequest) FromJson(j string) *GetMandateAdmissionRequest {

	return o
}

func (o *GetMandateAdmissionRequest) WithAdmissionID(admissionID strfmt.UUID) *GetMandateAdmissionRequest {

	o.AdmissionID = admissionID

	return o
}

func (o *GetMandateAdmissionRequest) WithID(id strfmt.UUID) *GetMandateAdmissionRequest {

	o.ID = id

	return o
}

//////////////////
// WithContext adds the context to the get mandate admission Request
func (o *GetMandateAdmissionRequest) WithContext(ctx context.Context) *GetMandateAdmissionRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the get mandate admission Request
func (o *GetMandateAdmissionRequest) WithHTTPClient(client *http.Client) *GetMandateAdmissionRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *GetMandateAdmissionRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}