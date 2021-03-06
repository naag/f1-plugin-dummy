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

// PatchPaymentsIDRecallsRecallIDSubmissionsSubmissionIDReader is a Reader for the PatchPaymentsIDRecallsRecallIDSubmissionsSubmissionID structure.
type PatchPaymentsIDRecallsRecallIDSubmissionsSubmissionIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchPaymentsIDRecallsRecallIDSubmissionsSubmissionIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPatchPaymentsIDRecallsRecallIDSubmissionsSubmissionIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPatchPaymentsIDRecallsRecallIDSubmissionsSubmissionIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchPaymentsIDRecallsRecallIDSubmissionsSubmissionIDOK creates a PatchPaymentsIDRecallsRecallIDSubmissionsSubmissionIDOK with default headers values
func NewPatchPaymentsIDRecallsRecallIDSubmissionsSubmissionIDOK() *PatchPaymentsIDRecallsRecallIDSubmissionsSubmissionIDOK {
	return &PatchPaymentsIDRecallsRecallIDSubmissionsSubmissionIDOK{}
}

/*PatchPaymentsIDRecallsRecallIDSubmissionsSubmissionIDOK handles this case with default header values.

Recall submission update response
*/
type PatchPaymentsIDRecallsRecallIDSubmissionsSubmissionIDOK struct {
	Payload *models.RecallSubmissionDetailsResponse
}

func (o *PatchPaymentsIDRecallsRecallIDSubmissionsSubmissionIDOK) Error() string {
	return fmt.Sprintf("[PATCH /payments/{id}/recalls/{recallId}/submissions/{submissionId}][%d] patchPaymentsIdRecallsRecallIdSubmissionsSubmissionIdOK  %+v", 200, o.Payload)
}

func (o *PatchPaymentsIDRecallsRecallIDSubmissionsSubmissionIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RecallSubmissionDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchPaymentsIDRecallsRecallIDSubmissionsSubmissionIDBadRequest creates a PatchPaymentsIDRecallsRecallIDSubmissionsSubmissionIDBadRequest with default headers values
func NewPatchPaymentsIDRecallsRecallIDSubmissionsSubmissionIDBadRequest() *PatchPaymentsIDRecallsRecallIDSubmissionsSubmissionIDBadRequest {
	return &PatchPaymentsIDRecallsRecallIDSubmissionsSubmissionIDBadRequest{}
}

/*PatchPaymentsIDRecallsRecallIDSubmissionsSubmissionIDBadRequest handles this case with default header values.

Error
*/
type PatchPaymentsIDRecallsRecallIDSubmissionsSubmissionIDBadRequest struct {
	Payload *models.APIError
}

func (o *PatchPaymentsIDRecallsRecallIDSubmissionsSubmissionIDBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /payments/{id}/recalls/{recallId}/submissions/{submissionId}][%d] patchPaymentsIdRecallsRecallIdSubmissionsSubmissionIdBadRequest  %+v", 400, o.Payload)
}

func (o *PatchPaymentsIDRecallsRecallIDSubmissionsSubmissionIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
