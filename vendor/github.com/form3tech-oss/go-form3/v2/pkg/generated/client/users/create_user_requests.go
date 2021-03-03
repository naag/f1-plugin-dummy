// Code generated by go-swagger; DO NOT EDIT.

package users

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

// Client.CreateUser creates a new CreateUserRequest object
// with the default values initialized.
func (c *Client) CreateUser() *CreateUserRequest {
	var ()
	return &CreateUserRequest{

		UserCreation: models.UserCreationWithDefaults(c.Defaults),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type CreateUserRequest struct {

	/*UserCreationRequest*/

	*models.UserCreation

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *CreateUserRequest) FromJson(j string) *CreateUserRequest {

	var m models.UserCreation
	if err := json.Unmarshal([]byte(j), &m); err != nil {
		log.Fatal(err)
	}

	o.UserCreation = &m

	return o
}

func (o *CreateUserRequest) WithUserCreationRequest(userCreationRequest models.UserCreation) *CreateUserRequest {

	o.UserCreation = &userCreationRequest

	return o
}

func (o *CreateUserRequest) WithoutUserCreationRequest() *CreateUserRequest {

	o.UserCreation = &models.UserCreation{}

	return o
}

//////////////////
// WithContext adds the context to the create user Request
func (o *CreateUserRequest) WithContext(ctx context.Context) *CreateUserRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the create user Request
func (o *CreateUserRequest) WithHTTPClient(client *http.Client) *CreateUserRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *CreateUserRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// ISBODYPARAM
	if o.UserCreation != nil {
		if err := r.SetBodyParam(o.UserCreation); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}