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

// PostPaymentsIDRecallsRecallIDDecisionsDecisionIDAdmissionsReader is a Reader for the PostPaymentsIDRecallsRecallIDDecisionsDecisionIDAdmissions structure.
type PostPaymentsIDRecallsRecallIDDecisionsDecisionIDAdmissionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostPaymentsIDRecallsRecallIDDecisionsDecisionIDAdmissionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostPaymentsIDRecallsRecallIDDecisionsDecisionIDAdmissionsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostPaymentsIDRecallsRecallIDDecisionsDecisionIDAdmissionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostPaymentsIDRecallsRecallIDDecisionsDecisionIDAdmissionsCreated creates a PostPaymentsIDRecallsRecallIDDecisionsDecisionIDAdmissionsCreated with default headers values
func NewPostPaymentsIDRecallsRecallIDDecisionsDecisionIDAdmissionsCreated() *PostPaymentsIDRecallsRecallIDDecisionsDecisionIDAdmissionsCreated {
	return &PostPaymentsIDRecallsRecallIDDecisionsDecisionIDAdmissionsCreated{}
}

/*PostPaymentsIDRecallsRecallIDDecisionsDecisionIDAdmissionsCreated handles this case with default header values.

Recall decision admission creation response
*/
type PostPaymentsIDRecallsRecallIDDecisionsDecisionIDAdmissionsCreated struct {
	Payload *models.RecallDecisionAdmissionCreationResponse
}

func (o *PostPaymentsIDRecallsRecallIDDecisionsDecisionIDAdmissionsCreated) Error() string {
	return fmt.Sprintf("[POST /payments/{id}/recalls/{recallId}/decisions/{decisionId}/admissions][%d] postPaymentsIdRecallsRecallIdDecisionsDecisionIdAdmissionsCreated  %+v", 201, o.Payload)
}

func (o *PostPaymentsIDRecallsRecallIDDecisionsDecisionIDAdmissionsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RecallDecisionAdmissionCreationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPaymentsIDRecallsRecallIDDecisionsDecisionIDAdmissionsBadRequest creates a PostPaymentsIDRecallsRecallIDDecisionsDecisionIDAdmissionsBadRequest with default headers values
func NewPostPaymentsIDRecallsRecallIDDecisionsDecisionIDAdmissionsBadRequest() *PostPaymentsIDRecallsRecallIDDecisionsDecisionIDAdmissionsBadRequest {
	return &PostPaymentsIDRecallsRecallIDDecisionsDecisionIDAdmissionsBadRequest{}
}

/*PostPaymentsIDRecallsRecallIDDecisionsDecisionIDAdmissionsBadRequest handles this case with default header values.

Recall decision admission creation error
*/
type PostPaymentsIDRecallsRecallIDDecisionsDecisionIDAdmissionsBadRequest struct {
	Payload *models.APIError
}

func (o *PostPaymentsIDRecallsRecallIDDecisionsDecisionIDAdmissionsBadRequest) Error() string {
	return fmt.Sprintf("[POST /payments/{id}/recalls/{recallId}/decisions/{decisionId}/admissions][%d] postPaymentsIdRecallsRecallIdDecisionsDecisionIdAdmissionsBadRequest  %+v", 400, o.Payload)
}

func (o *PostPaymentsIDRecallsRecallIDDecisionsDecisionIDAdmissionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
