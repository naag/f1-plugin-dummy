// Code generated by go-swagger; DO NOT EDIT.

package system

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/form3tech/go-form3/models"
)

// GetKeysKeyIDCertificatesReader is a Reader for the GetKeysKeyIDCertificates structure.
type GetKeysKeyIDCertificatesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetKeysKeyIDCertificatesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetKeysKeyIDCertificatesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetKeysKeyIDCertificatesOK creates a GetKeysKeyIDCertificatesOK with default headers values
func NewGetKeysKeyIDCertificatesOK() *GetKeysKeyIDCertificatesOK {
	return &GetKeysKeyIDCertificatesOK{}
}

/*GetKeysKeyIDCertificatesOK handles this case with default header values.

List of certificates
*/
type GetKeysKeyIDCertificatesOK struct {
	Payload *models.CertificateDetailsListResponse
}

func (o *GetKeysKeyIDCertificatesOK) Error() string {
	return fmt.Sprintf("[GET /keys/{key_id}/certificates][%d] getKeysKeyIdCertificatesOK  %+v", 200, o.Payload)
}

func (o *GetKeysKeyIDCertificatesOK) GetPayload() *models.CertificateDetailsListResponse {
	return o.Payload
}

func (o *GetKeysKeyIDCertificatesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CertificateDetailsListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
