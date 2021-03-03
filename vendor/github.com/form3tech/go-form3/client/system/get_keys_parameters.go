// Code generated by go-swagger; DO NOT EDIT.

package system

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetKeysParams creates a new GetKeysParams object
// with the default values initialized.
func NewGetKeysParams() *GetKeysParams {

	return &GetKeysParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetKeysParamsWithTimeout creates a new GetKeysParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetKeysParamsWithTimeout(timeout time.Duration) *GetKeysParams {

	return &GetKeysParams{

		timeout: timeout,
	}
}

// NewGetKeysParamsWithContext creates a new GetKeysParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetKeysParamsWithContext(ctx context.Context) *GetKeysParams {

	return &GetKeysParams{

		Context: ctx,
	}
}

// NewGetKeysParamsWithHTTPClient creates a new GetKeysParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetKeysParamsWithHTTPClient(client *http.Client) *GetKeysParams {

	return &GetKeysParams{
		HTTPClient: client,
	}
}

/*GetKeysParams contains all the parameters to send to the API endpoint
for the get keys operation typically these are written to a http.Request
*/
type GetKeysParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get keys params
func (o *GetKeysParams) WithTimeout(timeout time.Duration) *GetKeysParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get keys params
func (o *GetKeysParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get keys params
func (o *GetKeysParams) WithContext(ctx context.Context) *GetKeysParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get keys params
func (o *GetKeysParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get keys params
func (o *GetKeysParams) WithHTTPClient(client *http.Client) *GetKeysParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get keys params
func (o *GetKeysParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetKeysParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}