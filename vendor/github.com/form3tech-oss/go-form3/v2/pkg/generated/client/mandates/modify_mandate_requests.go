// Code generated by go-swagger; DO NOT EDIT.

package mandates

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v2/pkg/generated/models"
)

// Client.ModifyMandate creates a new ModifyMandateRequest object
// with the default values initialized.
func (c *Client) ModifyMandate() *ModifyMandateRequest {
	var ()
	return &ModifyMandateRequest{

		MandateAmendment: models.MandateAmendmentWithDefaults(c.Defaults),

		ID: c.Defaults.GetStrfmtUUID("ModifyMandate", "id"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type ModifyMandateRequest struct {

	/*MandateAmendment*/

	*models.MandateAmendment

	/*ID      Mandate Id      */

	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *ModifyMandateRequest) FromJson(j string) *ModifyMandateRequest {

	var m models.MandateAmendment
	if err := json.Unmarshal([]byte(j), &m); err != nil {
		log.Fatal(err)
	}

	o.MandateAmendment = &m

	return o
}

func (o *ModifyMandateRequest) WithMandateAmendment(mandateAmendment models.MandateAmendment) *ModifyMandateRequest {

	o.MandateAmendment = &mandateAmendment

	return o
}

func (o *ModifyMandateRequest) WithoutMandateAmendment() *ModifyMandateRequest {

	o.MandateAmendment = &models.MandateAmendment{}

	return o
}

func (o *ModifyMandateRequest) WithID(id strfmt.UUID) *ModifyMandateRequest {

	o.ID = id

	return o
}

//////////////////
// WithContext adds the context to the modify mandate Request
func (o *ModifyMandateRequest) WithContext(ctx context.Context) *ModifyMandateRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the modify mandate Request
func (o *ModifyMandateRequest) WithHTTPClient(client *http.Client) *ModifyMandateRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *ModifyMandateRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// ISBODYPARAM
	if o.MandateAmendment != nil {
		if err := r.SetBodyParam(o.MandateAmendment); err != nil {
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