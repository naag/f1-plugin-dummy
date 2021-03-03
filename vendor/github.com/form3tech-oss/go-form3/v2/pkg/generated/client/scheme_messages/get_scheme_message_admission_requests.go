// Code generated by go-swagger; DO NOT EDIT.

package scheme_messages

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// Client.GetSchemeMessageAdmission creates a new GetSchemeMessageAdmissionRequest object
// with the default values initialized.
func (c *Client) GetSchemeMessageAdmission() *GetSchemeMessageAdmissionRequest {
	var ()
	return &GetSchemeMessageAdmissionRequest{

		AdmissionID: c.Defaults.GetStrfmtUUID("GetSchemeMessageAdmission", "admissionId"),

		ID: c.Defaults.GetStrfmtUUID("GetSchemeMessageAdmission", "id"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type GetSchemeMessageAdmissionRequest struct {

	/*AdmissionID      Scheme Message Admission Id      */

	AdmissionID strfmt.UUID

	/*ID      Scheme Message Id      */

	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *GetSchemeMessageAdmissionRequest) FromJson(j string) *GetSchemeMessageAdmissionRequest {

	return o
}

func (o *GetSchemeMessageAdmissionRequest) WithAdmissionID(admissionID strfmt.UUID) *GetSchemeMessageAdmissionRequest {

	o.AdmissionID = admissionID

	return o
}

func (o *GetSchemeMessageAdmissionRequest) WithID(id strfmt.UUID) *GetSchemeMessageAdmissionRequest {

	o.ID = id

	return o
}

//////////////////
// WithContext adds the context to the get scheme message admission Request
func (o *GetSchemeMessageAdmissionRequest) WithContext(ctx context.Context) *GetSchemeMessageAdmissionRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the get scheme message admission Request
func (o *GetSchemeMessageAdmissionRequest) WithHTTPClient(client *http.Client) *GetSchemeMessageAdmissionRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *GetSchemeMessageAdmissionRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
