// Code generated by go-swagger; DO NOT EDIT.

package payments

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

// Client.CreatePaymentRecallDecisionSubmission creates a new CreatePaymentRecallDecisionSubmissionRequest object
// with the default values initialized.
func (c *Client) CreatePaymentRecallDecisionSubmission() *CreatePaymentRecallDecisionSubmissionRequest {
	var ()
	return &CreatePaymentRecallDecisionSubmissionRequest{

		RecallDecisionSubmissionCreation: models.RecallDecisionSubmissionCreationWithDefaults(c.Defaults),

		DecisionID: c.Defaults.GetStrfmtUUID("CreatePaymentRecallDecisionSubmission", "decisionId"),

		ID: c.Defaults.GetStrfmtUUID("CreatePaymentRecallDecisionSubmission", "id"),

		RecallID: c.Defaults.GetStrfmtUUID("CreatePaymentRecallDecisionSubmission", "recallId"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type CreatePaymentRecallDecisionSubmissionRequest struct {

	/*RecallDecisionSubmissionCreationRequest*/

	*models.RecallDecisionSubmissionCreation

	/*DecisionID      Decision Id      */

	DecisionID strfmt.UUID

	/*ID      Payment Id      */

	ID strfmt.UUID

	/*RecallID      Recall Id      */

	RecallID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *CreatePaymentRecallDecisionSubmissionRequest) FromJson(j string) *CreatePaymentRecallDecisionSubmissionRequest {

	var m models.RecallDecisionSubmissionCreation
	if err := json.Unmarshal([]byte(j), &m); err != nil {
		log.Fatal(err)
	}

	o.RecallDecisionSubmissionCreation = &m

	return o
}

func (o *CreatePaymentRecallDecisionSubmissionRequest) WithRecallDecisionSubmissionCreationRequest(recallDecisionSubmissionCreationRequest models.RecallDecisionSubmissionCreation) *CreatePaymentRecallDecisionSubmissionRequest {

	o.RecallDecisionSubmissionCreation = &recallDecisionSubmissionCreationRequest

	return o
}

func (o *CreatePaymentRecallDecisionSubmissionRequest) WithoutRecallDecisionSubmissionCreationRequest() *CreatePaymentRecallDecisionSubmissionRequest {

	o.RecallDecisionSubmissionCreation = &models.RecallDecisionSubmissionCreation{}

	return o
}

func (o *CreatePaymentRecallDecisionSubmissionRequest) WithDecisionID(decisionID strfmt.UUID) *CreatePaymentRecallDecisionSubmissionRequest {

	o.DecisionID = decisionID

	return o
}

func (o *CreatePaymentRecallDecisionSubmissionRequest) WithID(id strfmt.UUID) *CreatePaymentRecallDecisionSubmissionRequest {

	o.ID = id

	return o
}

func (o *CreatePaymentRecallDecisionSubmissionRequest) WithRecallID(recallID strfmt.UUID) *CreatePaymentRecallDecisionSubmissionRequest {

	o.RecallID = recallID

	return o
}

//////////////////
// WithContext adds the context to the create payment recall decision submission Request
func (o *CreatePaymentRecallDecisionSubmissionRequest) WithContext(ctx context.Context) *CreatePaymentRecallDecisionSubmissionRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the create payment recall decision submission Request
func (o *CreatePaymentRecallDecisionSubmissionRequest) WithHTTPClient(client *http.Client) *CreatePaymentRecallDecisionSubmissionRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *CreatePaymentRecallDecisionSubmissionRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// ISBODYPARAM
	if o.RecallDecisionSubmissionCreation != nil {
		if err := r.SetBodyParam(o.RecallDecisionSubmissionCreation); err != nil {
			return err
		}
	}

	// path param decisionId
	if err := r.SetPathParam("decisionId", o.DecisionID.String()); err != nil {
		return err
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
