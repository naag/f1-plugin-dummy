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

// PostPaymentsIDRecallsRecallIDReversalsReversalIDAdmissionsReader is a Reader for the PostPaymentsIDRecallsRecallIDReversalsReversalIDAdmissions structure.
type PostPaymentsIDRecallsRecallIDReversalsReversalIDAdmissionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostPaymentsIDRecallsRecallIDReversalsReversalIDAdmissionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostPaymentsIDRecallsRecallIDReversalsReversalIDAdmissionsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostPaymentsIDRecallsRecallIDReversalsReversalIDAdmissionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostPaymentsIDRecallsRecallIDReversalsReversalIDAdmissionsCreated creates a PostPaymentsIDRecallsRecallIDReversalsReversalIDAdmissionsCreated with default headers values
func NewPostPaymentsIDRecallsRecallIDReversalsReversalIDAdmissionsCreated() *PostPaymentsIDRecallsRecallIDReversalsReversalIDAdmissionsCreated {
	return &PostPaymentsIDRecallsRecallIDReversalsReversalIDAdmissionsCreated{}
}

/*PostPaymentsIDRecallsRecallIDReversalsReversalIDAdmissionsCreated handles this case with default header values.

Recall reversal admission creation response
*/
type PostPaymentsIDRecallsRecallIDReversalsReversalIDAdmissionsCreated struct {
	Payload *models.RecallReversalAdmissionCreationResponse
}

func (o *PostPaymentsIDRecallsRecallIDReversalsReversalIDAdmissionsCreated) Error() string {
	return fmt.Sprintf("[POST /payments/{id}/recalls/{recallId}/reversals/{reversalId}/admissions][%d] postPaymentsIdRecallsRecallIdReversalsReversalIdAdmissionsCreated  %+v", 201, o.Payload)
}

func (o *PostPaymentsIDRecallsRecallIDReversalsReversalIDAdmissionsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RecallReversalAdmissionCreationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPaymentsIDRecallsRecallIDReversalsReversalIDAdmissionsBadRequest creates a PostPaymentsIDRecallsRecallIDReversalsReversalIDAdmissionsBadRequest with default headers values
func NewPostPaymentsIDRecallsRecallIDReversalsReversalIDAdmissionsBadRequest() *PostPaymentsIDRecallsRecallIDReversalsReversalIDAdmissionsBadRequest {
	return &PostPaymentsIDRecallsRecallIDReversalsReversalIDAdmissionsBadRequest{}
}

/*PostPaymentsIDRecallsRecallIDReversalsReversalIDAdmissionsBadRequest handles this case with default header values.

Reversal admission creation error
*/
type PostPaymentsIDRecallsRecallIDReversalsReversalIDAdmissionsBadRequest struct {
	Payload *models.APIError
}

func (o *PostPaymentsIDRecallsRecallIDReversalsReversalIDAdmissionsBadRequest) Error() string {
	return fmt.Sprintf("[POST /payments/{id}/recalls/{recallId}/reversals/{reversalId}/admissions][%d] postPaymentsIdRecallsRecallIdReversalsReversalIdAdmissionsBadRequest  %+v", 400, o.Payload)
}

func (o *PostPaymentsIDRecallsRecallIDReversalsReversalIDAdmissionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
