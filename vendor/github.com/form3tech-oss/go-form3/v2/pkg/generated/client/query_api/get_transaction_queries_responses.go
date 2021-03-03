// Code generated by go-swagger; DO NOT EDIT.

package query_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v2/pkg/generated/models"
)

// GetTransactionQueriesReader is a Reader for the GetTransactionQueries structure.
type GetTransactionQueriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTransactionQueriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetTransactionQueriesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetTransactionQueriesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 502:
		result := NewGetTransactionQueriesBadGateway()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetTransactionQueriesOK creates a GetTransactionQueriesOK with default headers values
func NewGetTransactionQueriesOK() *GetTransactionQueriesOK {
	return &GetTransactionQueriesOK{}
}

/*GetTransactionQueriesOK handles this case with default header values.

Query response
*/
type GetTransactionQueriesOK struct {

	//Payload

	// isStream: false
	*models.QueryListResponse
}

func (o *GetTransactionQueriesOK) Error() string {
	return fmt.Sprintf("[GET /transaction/queries][%d] getTransactionQueriesOK", 200)
}

func (o *GetTransactionQueriesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.QueryListResponse = new(models.QueryListResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.QueryListResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTransactionQueriesBadRequest creates a GetTransactionQueriesBadRequest with default headers values
func NewGetTransactionQueriesBadRequest() *GetTransactionQueriesBadRequest {
	return &GetTransactionQueriesBadRequest{}
}

/*GetTransactionQueriesBadRequest handles this case with default header values.

Query bad request
*/
type GetTransactionQueriesBadRequest struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetTransactionQueriesBadRequest) Error() string {
	return fmt.Sprintf("[GET /transaction/queries][%d] getTransactionQueriesBadRequest", 400)
}

func (o *GetTransactionQueriesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTransactionQueriesBadGateway creates a GetTransactionQueriesBadGateway with default headers values
func NewGetTransactionQueriesBadGateway() *GetTransactionQueriesBadGateway {
	return &GetTransactionQueriesBadGateway{}
}

/*GetTransactionQueriesBadGateway handles this case with default header values.

Query gateway issue
*/
type GetTransactionQueriesBadGateway struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetTransactionQueriesBadGateway) Error() string {
	return fmt.Sprintf("[GET /transaction/queries][%d] getTransactionQueriesBadGateway", 502)
}

func (o *GetTransactionQueriesBadGateway) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
