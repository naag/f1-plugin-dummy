// Code generated by go-swagger; DO NOT EDIT.

package payments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v2/pkg/generated/models"
)

// CreatePaymentRecallSubmissionReader is a Reader for the CreatePaymentRecallSubmission structure.
type CreatePaymentRecallSubmissionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreatePaymentRecallSubmissionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreatePaymentRecallSubmissionCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreatePaymentRecallSubmissionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreatePaymentRecallSubmissionCreated creates a CreatePaymentRecallSubmissionCreated with default headers values
func NewCreatePaymentRecallSubmissionCreated() *CreatePaymentRecallSubmissionCreated {
	return &CreatePaymentRecallSubmissionCreated{}
}

/*CreatePaymentRecallSubmissionCreated handles this case with default header values.

Recall submission creation response
*/
type CreatePaymentRecallSubmissionCreated struct {

	//Payload

	// isStream: false
	*models.RecallSubmissionCreationResponse
}

func (o *CreatePaymentRecallSubmissionCreated) Error() string {
	return fmt.Sprintf("[POST /transaction/payments/{id}/recalls/{recallId}/submissions][%d] createPaymentRecallSubmissionCreated", 201)
}

func (o *CreatePaymentRecallSubmissionCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.RecallSubmissionCreationResponse = new(models.RecallSubmissionCreationResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.RecallSubmissionCreationResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePaymentRecallSubmissionBadRequest creates a CreatePaymentRecallSubmissionBadRequest with default headers values
func NewCreatePaymentRecallSubmissionBadRequest() *CreatePaymentRecallSubmissionBadRequest {
	return &CreatePaymentRecallSubmissionBadRequest{}
}

/*CreatePaymentRecallSubmissionBadRequest handles this case with default header values.

Recall submission creation error
*/
type CreatePaymentRecallSubmissionBadRequest struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *CreatePaymentRecallSubmissionBadRequest) Error() string {
	return fmt.Sprintf("[POST /transaction/payments/{id}/recalls/{recallId}/submissions][%d] createPaymentRecallSubmissionBadRequest", 400)
}

func (o *CreatePaymentRecallSubmissionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}