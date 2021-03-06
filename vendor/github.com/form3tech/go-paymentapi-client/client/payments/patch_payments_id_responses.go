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

// PatchPaymentsIDReader is a Reader for the PatchPaymentsID structure.
type PatchPaymentsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchPaymentsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPatchPaymentsIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchPaymentsIDOK creates a PatchPaymentsIDOK with default headers values
func NewPatchPaymentsIDOK() *PatchPaymentsIDOK {
	return &PatchPaymentsIDOK{}
}

/*PatchPaymentsIDOK handles this case with default header values.

Payment updated
*/
type PatchPaymentsIDOK struct {
	Payload *models.PaymentDetailsResponse
}

func (o *PatchPaymentsIDOK) Error() string {
	return fmt.Sprintf("[PATCH /payments/{id}][%d] patchPaymentsIdOK  %+v", 200, o.Payload)
}

func (o *PatchPaymentsIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PaymentDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
