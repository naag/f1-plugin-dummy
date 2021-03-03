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

// GetPaymentsIDAdvicesAdviceIDSubmissionsSubmissionIDReader is a Reader for the GetPaymentsIDAdvicesAdviceIDSubmissionsSubmissionID structure.
type GetPaymentsIDAdvicesAdviceIDSubmissionsSubmissionIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPaymentsIDAdvicesAdviceIDSubmissionsSubmissionIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPaymentsIDAdvicesAdviceIDSubmissionsSubmissionIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPaymentsIDAdvicesAdviceIDSubmissionsSubmissionIDOK creates a GetPaymentsIDAdvicesAdviceIDSubmissionsSubmissionIDOK with default headers values
func NewGetPaymentsIDAdvicesAdviceIDSubmissionsSubmissionIDOK() *GetPaymentsIDAdvicesAdviceIDSubmissionsSubmissionIDOK {
	return &GetPaymentsIDAdvicesAdviceIDSubmissionsSubmissionIDOK{}
}

/*GetPaymentsIDAdvicesAdviceIDSubmissionsSubmissionIDOK handles this case with default header values.

Advice submission details
*/
type GetPaymentsIDAdvicesAdviceIDSubmissionsSubmissionIDOK struct {
	Payload *models.AdviceSubmissionDetailsResponse
}

func (o *GetPaymentsIDAdvicesAdviceIDSubmissionsSubmissionIDOK) Error() string {
	return fmt.Sprintf("[GET /payments/{id}/advices/{adviceId}/submissions/{submissionId}][%d] getPaymentsIdAdvicesAdviceIdSubmissionsSubmissionIdOK  %+v", 200, o.Payload)
}

func (o *GetPaymentsIDAdvicesAdviceIDSubmissionsSubmissionIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AdviceSubmissionDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
