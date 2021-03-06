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

// GetPaymentsIDRecallsRecallIDDecisionsDecisionIDReader is a Reader for the GetPaymentsIDRecallsRecallIDDecisionsDecisionID structure.
type GetPaymentsIDRecallsRecallIDDecisionsDecisionIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPaymentsIDRecallsRecallIDDecisionsDecisionIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPaymentsIDRecallsRecallIDDecisionsDecisionIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPaymentsIDRecallsRecallIDDecisionsDecisionIDOK creates a GetPaymentsIDRecallsRecallIDDecisionsDecisionIDOK with default headers values
func NewGetPaymentsIDRecallsRecallIDDecisionsDecisionIDOK() *GetPaymentsIDRecallsRecallIDDecisionsDecisionIDOK {
	return &GetPaymentsIDRecallsRecallIDDecisionsDecisionIDOK{}
}

/*GetPaymentsIDRecallsRecallIDDecisionsDecisionIDOK handles this case with default header values.

Recall decision details
*/
type GetPaymentsIDRecallsRecallIDDecisionsDecisionIDOK struct {
	Payload *models.RecallDecisionDetailsResponse
}

func (o *GetPaymentsIDRecallsRecallIDDecisionsDecisionIDOK) Error() string {
	return fmt.Sprintf("[GET /payments/{id}/recalls/{recallId}/decisions/{decisionId}][%d] getPaymentsIdRecallsRecallIdDecisionsDecisionIdOK  %+v", 200, o.Payload)
}

func (o *GetPaymentsIDRecallsRecallIDDecisionsDecisionIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RecallDecisionDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
