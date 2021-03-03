// Code generated by go-swagger; DO NOT EDIT.

package payments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetBatchesIDFileReader is a Reader for the GetBatchesIDFile structure.
type GetBatchesIDFileReader struct {
	formats strfmt.Registry
	writer  io.Writer
}

// ReadResponse reads a server response into the received o.
func (o *GetBatchesIDFileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetBatchesIDFileOK(o.writer)
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetBatchesIDFileOK creates a GetBatchesIDFileOK with default headers values
func NewGetBatchesIDFileOK(writer io.Writer) *GetBatchesIDFileOK {
	return &GetBatchesIDFileOK{
		Payload: writer,
	}
}

/*GetBatchesIDFileOK handles this case with default header values.

Batch file
*/
type GetBatchesIDFileOK struct {
	Payload io.Writer
}

func (o *GetBatchesIDFileOK) Error() string {
	return fmt.Sprintf("[GET /batches/{id}/file][%d] getBatchesIdFileOK  %+v", 200, o.Payload)
}

func (o *GetBatchesIDFileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
