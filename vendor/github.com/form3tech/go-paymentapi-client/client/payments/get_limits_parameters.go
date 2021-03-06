// Code generated by go-swagger; DO NOT EDIT.

package payments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetLimitsParams creates a new GetLimitsParams object
// with the default values initialized.
func NewGetLimitsParams() *GetLimitsParams {
	var ()
	return &GetLimitsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLimitsParamsWithTimeout creates a new GetLimitsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLimitsParamsWithTimeout(timeout time.Duration) *GetLimitsParams {
	var ()
	return &GetLimitsParams{

		timeout: timeout,
	}
}

// NewGetLimitsParamsWithContext creates a new GetLimitsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLimitsParamsWithContext(ctx context.Context) *GetLimitsParams {
	var ()
	return &GetLimitsParams{

		Context: ctx,
	}
}

// NewGetLimitsParamsWithHTTPClient creates a new GetLimitsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLimitsParamsWithHTTPClient(client *http.Client) *GetLimitsParams {
	var ()
	return &GetLimitsParams{
		HTTPClient: client,
	}
}

/*GetLimitsParams contains all the parameters to send to the API endpoint
for the get limits operation typically these are written to a http.Request
*/
type GetLimitsParams struct {

	/*XForm3TransactionTime*/
	XForm3TransactionTime *string
	/*XForm3CorrelationID*/
	XForm3CorrelationID *string
	/*FilterOrganisationID
	  Filter by organisation id

	*/
	FilterOrganisationID []strfmt.UUID
	/*PageNumber
	  Which page to select

	*/
	PageNumber *string
	/*PageSize
	  Number of items to select

	*/
	PageSize *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get limits params
func (o *GetLimitsParams) WithTimeout(timeout time.Duration) *GetLimitsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get limits params
func (o *GetLimitsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get limits params
func (o *GetLimitsParams) WithContext(ctx context.Context) *GetLimitsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get limits params
func (o *GetLimitsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get limits params
func (o *GetLimitsParams) WithHTTPClient(client *http.Client) *GetLimitsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get limits params
func (o *GetLimitsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXForm3TransactionTime adds the xForm3TransactionTime to the get limits params
func (o *GetLimitsParams) WithXForm3TransactionTime(xForm3TransactionTime *string) *GetLimitsParams {
	o.SetXForm3TransactionTime(xForm3TransactionTime)
	return o
}

// SetXForm3TransactionTime adds the xForm3TransactionTime to the get limits params
func (o *GetLimitsParams) SetXForm3TransactionTime(xForm3TransactionTime *string) {
	o.XForm3TransactionTime = xForm3TransactionTime
}

// WithXForm3CorrelationID adds the xForm3CorrelationID to the get limits params
func (o *GetLimitsParams) WithXForm3CorrelationID(xForm3CorrelationID *string) *GetLimitsParams {
	o.SetXForm3CorrelationID(xForm3CorrelationID)
	return o
}

// SetXForm3CorrelationID adds the xForm3CorrelationId to the get limits params
func (o *GetLimitsParams) SetXForm3CorrelationID(xForm3CorrelationID *string) {
	o.XForm3CorrelationID = xForm3CorrelationID
}

// WithFilterOrganisationID adds the filterOrganisationID to the get limits params
func (o *GetLimitsParams) WithFilterOrganisationID(filterOrganisationID []strfmt.UUID) *GetLimitsParams {
	o.SetFilterOrganisationID(filterOrganisationID)
	return o
}

// SetFilterOrganisationID adds the filterOrganisationId to the get limits params
func (o *GetLimitsParams) SetFilterOrganisationID(filterOrganisationID []strfmt.UUID) {
	o.FilterOrganisationID = filterOrganisationID
}

// WithPageNumber adds the pageNumber to the get limits params
func (o *GetLimitsParams) WithPageNumber(pageNumber *string) *GetLimitsParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get limits params
func (o *GetLimitsParams) SetPageNumber(pageNumber *string) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get limits params
func (o *GetLimitsParams) WithPageSize(pageSize *int64) *GetLimitsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get limits params
func (o *GetLimitsParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WriteToRequest writes these params to a swagger request
func (o *GetLimitsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XForm3TransactionTime != nil {

		// header param X-Form3-Transaction-Time
		if err := r.SetHeaderParam("X-Form3-Transaction-Time", *o.XForm3TransactionTime); err != nil {
			return err
		}

	}

	if o.XForm3CorrelationID != nil {

		// header param X-Form3-correlation-id
		if err := r.SetHeaderParam("X-Form3-correlation-id", *o.XForm3CorrelationID); err != nil {
			return err
		}

	}

	var valuesFilterOrganisationID []string
	for _, v := range o.FilterOrganisationID {
		valuesFilterOrganisationID = append(valuesFilterOrganisationID, v.String())
	}

	joinedFilterOrganisationID := swag.JoinByFormat(valuesFilterOrganisationID, "")
	// query array param filter[organisation_id]
	if err := r.SetQueryParam("filter[organisation_id]", joinedFilterOrganisationID...); err != nil {
		return err
	}

	if o.PageNumber != nil {

		// query param page[number]
		var qrPageNumber string
		if o.PageNumber != nil {
			qrPageNumber = *o.PageNumber
		}
		qPageNumber := qrPageNumber
		if qPageNumber != "" {
			if err := r.SetQueryParam("page[number]", qPageNumber); err != nil {
				return err
			}
		}

	}

	if o.PageSize != nil {

		// query param page[size]
		var qrPageSize int64
		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatInt64(qrPageSize)
		if qPageSize != "" {
			if err := r.SetQueryParam("page[size]", qPageSize); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
