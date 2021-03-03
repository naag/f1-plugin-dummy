// Code generated by go-swagger; DO NOT EDIT.

package reports

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// Client.GetReport creates a new GetReportRequest object
// with the default values initialized.
func (c *Client) GetReport() *GetReportRequest {
	var ()
	return &GetReportRequest{

		Accept: c.Defaults.GetString("GetReport", "Accept"),

		ID: c.Defaults.GetStrfmtUUID("GetReport", "id"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type GetReportRequest struct {

	/*Accept      Acceptable Format      */

	Accept string

	/*ID      Report ID      */

	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *GetReportRequest) FromJson(j string) *GetReportRequest {

	return o
}

func (o *GetReportRequest) WithAccept(accept string) *GetReportRequest {

	o.Accept = accept

	return o
}

func (o *GetReportRequest) WithID(id strfmt.UUID) *GetReportRequest {

	o.ID = id

	return o
}

//////////////////
// WithContext adds the context to the get report Request
func (o *GetReportRequest) WithContext(ctx context.Context) *GetReportRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the get report Request
func (o *GetReportRequest) WithHTTPClient(client *http.Client) *GetReportRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *GetReportRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Accept
	if err := r.SetHeaderParam("Accept", o.Accept); err != nil {
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