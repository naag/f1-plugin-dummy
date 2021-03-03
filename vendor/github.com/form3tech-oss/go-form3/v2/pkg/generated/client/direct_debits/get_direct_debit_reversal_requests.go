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

// Client.GetDirectDebitReversal creates a new GetDirectDebitReversalRequest object
// with the default values initialized.
func (c *Client) GetDirectDebitReversal() *GetDirectDebitReversalRequest {
	var ()
	return &GetDirectDebitReversalRequest{

		ID: c.Defaults.GetStrfmtUUID("GetDirectDebitReversal", "id"),

		ReversalID: c.Defaults.GetStrfmtUUID("GetDirectDebitReversal", "reversalId"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type GetDirectDebitReversalRequest struct {

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

func (o *GetDirectDebitReversalRequest) FromJson(j string) *GetDirectDebitReversalRequest {

	return o
}

func (o *GetDirectDebitReversalRequest) WithID(id strfmt.UUID) *GetDirectDebitReversalRequest {

	o.ID = id

	return o
}

func (o *GetDirectDebitReversalRequest) WithReversalID(reversalID strfmt.UUID) *GetDirectDebitReversalRequest {

	o.ReversalID = reversalID

	return o
}

//////////////////
// WithContext adds the context to the get direct debit reversal Request
func (o *GetDirectDebitReversalRequest) WithContext(ctx context.Context) *GetDirectDebitReversalRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the get direct debit reversal Request
func (o *GetDirectDebitReversalRequest) WithHTTPClient(client *http.Client) *GetDirectDebitReversalRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *GetDirectDebitReversalRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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