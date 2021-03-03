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

// Client.GetTransactionDirectdebitsIDRecallsRecallIDSubmissionsSubmissionID creates a new GetTransactionDirectdebitsIDRecallsRecallIDSubmissionsSubmissionIDRequest object
// with the default values initialized.
func (c *Client) GetTransactionDirectdebitsIDRecallsRecallIDSubmissionsSubmissionID() *GetTransactionDirectdebitsIDRecallsRecallIDSubmissionsSubmissionIDRequest {
	var ()
	return &GetTransactionDirectdebitsIDRecallsRecallIDSubmissionsSubmissionIDRequest{

		ID: c.Defaults.GetStrfmtUUID("GetTransactionDirectdebitsIDRecallsRecallIDSubmissionsSubmissionID", "id"),

		RecallID: c.Defaults.GetStrfmtUUID("GetTransactionDirectdebitsIDRecallsRecallIDSubmissionsSubmissionID", "recallId"),

		SubmissionID: c.Defaults.GetStrfmtUUID("GetTransactionDirectdebitsIDRecallsRecallIDSubmissionsSubmissionID", "submissionId"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type GetTransactionDirectdebitsIDRecallsRecallIDSubmissionsSubmissionIDRequest struct {

	/*ID      Direct Debit Id      */

	ID strfmt.UUID

	/*RecallID      Recall Id      */

	RecallID strfmt.UUID

	/*SubmissionID      Submission Id      */

	SubmissionID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *GetTransactionDirectdebitsIDRecallsRecallIDSubmissionsSubmissionIDRequest) FromJson(j string) *GetTransactionDirectdebitsIDRecallsRecallIDSubmissionsSubmissionIDRequest {

	return o
}

func (o *GetTransactionDirectdebitsIDRecallsRecallIDSubmissionsSubmissionIDRequest) WithID(id strfmt.UUID) *GetTransactionDirectdebitsIDRecallsRecallIDSubmissionsSubmissionIDRequest {

	o.ID = id

	return o
}

func (o *GetTransactionDirectdebitsIDRecallsRecallIDSubmissionsSubmissionIDRequest) WithRecallID(recallID strfmt.UUID) *GetTransactionDirectdebitsIDRecallsRecallIDSubmissionsSubmissionIDRequest {

	o.RecallID = recallID

	return o
}

func (o *GetTransactionDirectdebitsIDRecallsRecallIDSubmissionsSubmissionIDRequest) WithSubmissionID(submissionID strfmt.UUID) *GetTransactionDirectdebitsIDRecallsRecallIDSubmissionsSubmissionIDRequest {

	o.SubmissionID = submissionID

	return o
}

//////////////////
// WithContext adds the context to the get transaction directdebits ID recalls recall ID submissions submission ID Request
func (o *GetTransactionDirectdebitsIDRecallsRecallIDSubmissionsSubmissionIDRequest) WithContext(ctx context.Context) *GetTransactionDirectdebitsIDRecallsRecallIDSubmissionsSubmissionIDRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the get transaction directdebits ID recalls recall ID submissions submission ID Request
func (o *GetTransactionDirectdebitsIDRecallsRecallIDSubmissionsSubmissionIDRequest) WithHTTPClient(client *http.Client) *GetTransactionDirectdebitsIDRecallsRecallIDSubmissionsSubmissionIDRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *GetTransactionDirectdebitsIDRecallsRecallIDSubmissionsSubmissionIDRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	// path param recallId
	if err := r.SetPathParam("recallId", o.RecallID.String()); err != nil {
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