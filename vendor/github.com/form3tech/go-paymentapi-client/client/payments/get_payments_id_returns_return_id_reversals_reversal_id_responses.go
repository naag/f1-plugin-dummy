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

// GetPaymentsIDReturnsReturnIDReversalsReversalIDReader is a Reader for the GetPaymentsIDReturnsReturnIDReversalsReversalID structure.
type GetPaymentsIDReturnsReturnIDReversalsReversalIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPaymentsIDReturnsReturnIDReversalsReversalIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPaymentsIDReturnsReturnIDReversalsReversalIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPaymentsIDReturnsReturnIDReversalsReversalIDOK creates a GetPaymentsIDReturnsReturnIDReversalsReversalIDOK with default headers values
func NewGetPaymentsIDReturnsReturnIDReversalsReversalIDOK() *GetPaymentsIDReturnsReturnIDReversalsReversalIDOK {
	return &GetPaymentsIDReturnsReturnIDReversalsReversalIDOK{}
}

/*GetPaymentsIDReturnsReturnIDReversalsReversalIDOK handles this case with default header values.

Return reversal details
*/
type GetPaymentsIDReturnsReturnIDReversalsReversalIDOK struct {
	Payload *models.ReturnReversalDetailsResponse
}

func (o *GetPaymentsIDReturnsReturnIDReversalsReversalIDOK) Error() string {
	return fmt.Sprintf("[GET /payments/{id}/returns/{returnId}/reversals/{reversalId}][%d] getPaymentsIdReturnsReturnIdReversalsReversalIdOK  %+v", 200, o.Payload)
}

func (o *GetPaymentsIDReturnsReturnIDReversalsReversalIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ReturnReversalDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
