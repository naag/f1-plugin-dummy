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

// GetPaymentsIDRecallsRecallIDReversalsReversalIDAdmissionsAdmissionIDReader is a Reader for the GetPaymentsIDRecallsRecallIDReversalsReversalIDAdmissionsAdmissionID structure.
type GetPaymentsIDRecallsRecallIDReversalsReversalIDAdmissionsAdmissionIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPaymentsIDRecallsRecallIDReversalsReversalIDAdmissionsAdmissionIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPaymentsIDRecallsRecallIDReversalsReversalIDAdmissionsAdmissionIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPaymentsIDRecallsRecallIDReversalsReversalIDAdmissionsAdmissionIDOK creates a GetPaymentsIDRecallsRecallIDReversalsReversalIDAdmissionsAdmissionIDOK with default headers values
func NewGetPaymentsIDRecallsRecallIDReversalsReversalIDAdmissionsAdmissionIDOK() *GetPaymentsIDRecallsRecallIDReversalsReversalIDAdmissionsAdmissionIDOK {
	return &GetPaymentsIDRecallsRecallIDReversalsReversalIDAdmissionsAdmissionIDOK{}
}

/*GetPaymentsIDRecallsRecallIDReversalsReversalIDAdmissionsAdmissionIDOK handles this case with default header values.

Reversal admission details
*/
type GetPaymentsIDRecallsRecallIDReversalsReversalIDAdmissionsAdmissionIDOK struct {
	Payload *models.RecallReversalAdmissionDetailsResponse
}

func (o *GetPaymentsIDRecallsRecallIDReversalsReversalIDAdmissionsAdmissionIDOK) Error() string {
	return fmt.Sprintf("[GET /payments/{id}/recalls/{recallId}/reversals/{reversalId}/admissions/{admissionId}][%d] getPaymentsIdRecallsRecallIdReversalsReversalIdAdmissionsAdmissionIdOK  %+v", 200, o.Payload)
}

func (o *GetPaymentsIDRecallsRecallIDReversalsReversalIDAdmissionsAdmissionIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RecallReversalAdmissionDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
