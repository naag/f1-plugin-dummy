// Code generated by go-swagger; DO NOT EDIT.

package payments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech/go-paymentapi-client/models"
)

// PostPaymentsIDReturnsReturnIDAdmissionsReader is a Reader for the PostPaymentsIDReturnsReturnIDAdmissions structure.
type PostPaymentsIDReturnsReturnIDAdmissionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostPaymentsIDReturnsReturnIDAdmissionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostPaymentsIDReturnsReturnIDAdmissionsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostPaymentsIDReturnsReturnIDAdmissionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostPaymentsIDReturnsReturnIDAdmissionsCreated creates a PostPaymentsIDReturnsReturnIDAdmissionsCreated with default headers values
func NewPostPaymentsIDReturnsReturnIDAdmissionsCreated() *PostPaymentsIDReturnsReturnIDAdmissionsCreated {
	return &PostPaymentsIDReturnsReturnIDAdmissionsCreated{}
}

/*PostPaymentsIDReturnsReturnIDAdmissionsCreated handles this case with default header values.

Return admission creation response
*/
type PostPaymentsIDReturnsReturnIDAdmissionsCreated struct {
	Payload *models.ReturnAdmissionCreationResponse
}

func (o *PostPaymentsIDReturnsReturnIDAdmissionsCreated) Error() string {
	return fmt.Sprintf("[POST /payments/{id}/returns/{returnId}/admissions][%d] postPaymentsIdReturnsReturnIdAdmissionsCreated  %+v", 201, o.Payload)
}

func (o *PostPaymentsIDReturnsReturnIDAdmissionsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ReturnAdmissionCreationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPaymentsIDReturnsReturnIDAdmissionsBadRequest creates a PostPaymentsIDReturnsReturnIDAdmissionsBadRequest with default headers values
func NewPostPaymentsIDReturnsReturnIDAdmissionsBadRequest() *PostPaymentsIDReturnsReturnIDAdmissionsBadRequest {
	return &PostPaymentsIDReturnsReturnIDAdmissionsBadRequest{}
}

/*PostPaymentsIDReturnsReturnIDAdmissionsBadRequest handles this case with default header values.

Return admission creation error
*/
type PostPaymentsIDReturnsReturnIDAdmissionsBadRequest struct {
	Payload *models.APIError
}

func (o *PostPaymentsIDReturnsReturnIDAdmissionsBadRequest) Error() string {
	return fmt.Sprintf("[POST /payments/{id}/returns/{returnId}/admissions][%d] postPaymentsIdReturnsReturnIdAdmissionsBadRequest  %+v", 400, o.Payload)
}

func (o *PostPaymentsIDReturnsReturnIDAdmissionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
