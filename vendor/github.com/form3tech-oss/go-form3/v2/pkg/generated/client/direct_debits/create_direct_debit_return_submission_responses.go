// Code generated by go-swagger; DO NOT EDIT.

package direct_debits

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v2/pkg/generated/models"
)

// CreateDirectDebitReturnSubmissionReader is a Reader for the CreateDirectDebitReturnSubmission structure.
type CreateDirectDebitReturnSubmissionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateDirectDebitReturnSubmissionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreateDirectDebitReturnSubmissionCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateDirectDebitReturnSubmissionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateDirectDebitReturnSubmissionCreated creates a CreateDirectDebitReturnSubmissionCreated with default headers values
func NewCreateDirectDebitReturnSubmissionCreated() *CreateDirectDebitReturnSubmissionCreated {
	return &CreateDirectDebitReturnSubmissionCreated{}
}

/*CreateDirectDebitReturnSubmissionCreated handles this case with default header values.

Return submission creation response
*/
type CreateDirectDebitReturnSubmissionCreated struct {

	//Payload

	// isStream: false
	*models.DirectDebitReturnSubmissionCreationResponse
}

func (o *CreateDirectDebitReturnSubmissionCreated) Error() string {
	return fmt.Sprintf("[POST /transaction/directdebits/{id}/returns/{returnId}/submissions][%d] createDirectDebitReturnSubmissionCreated", 201)
}

func (o *CreateDirectDebitReturnSubmissionCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.DirectDebitReturnSubmissionCreationResponse = new(models.DirectDebitReturnSubmissionCreationResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.DirectDebitReturnSubmissionCreationResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDirectDebitReturnSubmissionBadRequest creates a CreateDirectDebitReturnSubmissionBadRequest with default headers values
func NewCreateDirectDebitReturnSubmissionBadRequest() *CreateDirectDebitReturnSubmissionBadRequest {
	return &CreateDirectDebitReturnSubmissionBadRequest{}
}

/*CreateDirectDebitReturnSubmissionBadRequest handles this case with default header values.

Return submission creation error
*/
type CreateDirectDebitReturnSubmissionBadRequest struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *CreateDirectDebitReturnSubmissionBadRequest) Error() string {
	return fmt.Sprintf("[POST /transaction/directdebits/{id}/returns/{returnId}/submissions][%d] createDirectDebitReturnSubmissionBadRequest", 400)
}

func (o *CreateDirectDebitReturnSubmissionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}