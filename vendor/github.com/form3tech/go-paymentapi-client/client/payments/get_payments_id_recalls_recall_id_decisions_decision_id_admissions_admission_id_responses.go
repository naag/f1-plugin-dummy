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

// GetPaymentsIDRecallsRecallIDDecisionsDecisionIDAdmissionsAdmissionIDReader is a Reader for the GetPaymentsIDRecallsRecallIDDecisionsDecisionIDAdmissionsAdmissionID structure.
type GetPaymentsIDRecallsRecallIDDecisionsDecisionIDAdmissionsAdmissionIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPaymentsIDRecallsRecallIDDecisionsDecisionIDAdmissionsAdmissionIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPaymentsIDRecallsRecallIDDecisionsDecisionIDAdmissionsAdmissionIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPaymentsIDRecallsRecallIDDecisionsDecisionIDAdmissionsAdmissionIDOK creates a GetPaymentsIDRecallsRecallIDDecisionsDecisionIDAdmissionsAdmissionIDOK with default headers values
func NewGetPaymentsIDRecallsRecallIDDecisionsDecisionIDAdmissionsAdmissionIDOK() *GetPaymentsIDRecallsRecallIDDecisionsDecisionIDAdmissionsAdmissionIDOK {
	return &GetPaymentsIDRecallsRecallIDDecisionsDecisionIDAdmissionsAdmissionIDOK{}
}

/*GetPaymentsIDRecallsRecallIDDecisionsDecisionIDAdmissionsAdmissionIDOK handles this case with default header values.

Recall decision admission details
*/
type GetPaymentsIDRecallsRecallIDDecisionsDecisionIDAdmissionsAdmissionIDOK struct {
	Payload *models.RecallDecisionAdmissionDetailsResponse
}

func (o *GetPaymentsIDRecallsRecallIDDecisionsDecisionIDAdmissionsAdmissionIDOK) Error() string {
	return fmt.Sprintf("[GET /payments/{id}/recalls/{recallId}/decisions/{decisionId}/admissions/{admissionId}][%d] getPaymentsIdRecallsRecallIdDecisionsDecisionIdAdmissionsAdmissionIdOK  %+v", 200, o.Payload)
}

func (o *GetPaymentsIDRecallsRecallIDDecisionsDecisionIDAdmissionsAdmissionIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RecallDecisionAdmissionDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
