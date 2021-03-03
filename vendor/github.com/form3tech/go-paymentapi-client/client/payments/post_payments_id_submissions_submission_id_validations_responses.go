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

// PostPaymentsIDSubmissionsSubmissionIDValidationsReader is a Reader for the PostPaymentsIDSubmissionsSubmissionIDValidations structure.
type PostPaymentsIDSubmissionsSubmissionIDValidationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostPaymentsIDSubmissionsSubmissionIDValidationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostPaymentsIDSubmissionsSubmissionIDValidationsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostPaymentsIDSubmissionsSubmissionIDValidationsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostPaymentsIDSubmissionsSubmissionIDValidationsCreated creates a PostPaymentsIDSubmissionsSubmissionIDValidationsCreated with default headers values
func NewPostPaymentsIDSubmissionsSubmissionIDValidationsCreated() *PostPaymentsIDSubmissionsSubmissionIDValidationsCreated {
	return &PostPaymentsIDSubmissionsSubmissionIDValidationsCreated{}
}

/*PostPaymentsIDSubmissionsSubmissionIDValidationsCreated handles this case with default header values.

Payment submission validation creation response
*/
type PostPaymentsIDSubmissionsSubmissionIDValidationsCreated struct {
	Payload *models.PaymentSubmissionValidationCreationResponse
}

func (o *PostPaymentsIDSubmissionsSubmissionIDValidationsCreated) Error() string {
	return fmt.Sprintf("[POST /payments/{id}/submissions/{submissionId}/validations][%d] postPaymentsIdSubmissionsSubmissionIdValidationsCreated  %+v", 201, o.Payload)
}

func (o *PostPaymentsIDSubmissionsSubmissionIDValidationsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PaymentSubmissionValidationCreationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPaymentsIDSubmissionsSubmissionIDValidationsBadRequest creates a PostPaymentsIDSubmissionsSubmissionIDValidationsBadRequest with default headers values
func NewPostPaymentsIDSubmissionsSubmissionIDValidationsBadRequest() *PostPaymentsIDSubmissionsSubmissionIDValidationsBadRequest {
	return &PostPaymentsIDSubmissionsSubmissionIDValidationsBadRequest{}
}

/*PostPaymentsIDSubmissionsSubmissionIDValidationsBadRequest handles this case with default header values.

Payment submission validation creation error
*/
type PostPaymentsIDSubmissionsSubmissionIDValidationsBadRequest struct {
	Payload *models.APIError
}

func (o *PostPaymentsIDSubmissionsSubmissionIDValidationsBadRequest) Error() string {
	return fmt.Sprintf("[POST /payments/{id}/submissions/{submissionId}/validations][%d] postPaymentsIdSubmissionsSubmissionIdValidationsBadRequest  %+v", 400, o.Payload)
}

func (o *PostPaymentsIDSubmissionsSubmissionIDValidationsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
