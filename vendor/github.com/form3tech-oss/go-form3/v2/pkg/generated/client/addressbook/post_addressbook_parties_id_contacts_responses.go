// Code generated by go-swagger; DO NOT EDIT.

package addressbook

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v2/pkg/generated/models"
)

// PostAddressbookPartiesIDContactsReader is a Reader for the PostAddressbookPartiesIDContacts structure.
type PostAddressbookPartiesIDContactsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAddressbookPartiesIDContactsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostAddressbookPartiesIDContactsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostAddressbookPartiesIDContactsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPostAddressbookPartiesIDContactsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPostAddressbookPartiesIDContactsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewPostAddressbookPartiesIDContactsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostAddressbookPartiesIDContactsCreated creates a PostAddressbookPartiesIDContactsCreated with default headers values
func NewPostAddressbookPartiesIDContactsCreated() *PostAddressbookPartiesIDContactsCreated {
	return &PostAddressbookPartiesIDContactsCreated{}
}

/*PostAddressbookPartiesIDContactsCreated handles this case with default header values.

creation response
*/
type PostAddressbookPartiesIDContactsCreated struct {

	//Payload

	// isStream: false
	*models.ContactCreationResponse
}

func (o *PostAddressbookPartiesIDContactsCreated) Error() string {
	return fmt.Sprintf("[POST /addressbook/parties/{id}/contacts][%d] postAddressbookPartiesIdContactsCreated", 201)
}

func (o *PostAddressbookPartiesIDContactsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.ContactCreationResponse = new(models.ContactCreationResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.ContactCreationResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAddressbookPartiesIDContactsBadRequest creates a PostAddressbookPartiesIDContactsBadRequest with default headers values
func NewPostAddressbookPartiesIDContactsBadRequest() *PostAddressbookPartiesIDContactsBadRequest {
	return &PostAddressbookPartiesIDContactsBadRequest{}
}

/*PostAddressbookPartiesIDContactsBadRequest handles this case with default header values.

Bad Request
*/
type PostAddressbookPartiesIDContactsBadRequest struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *PostAddressbookPartiesIDContactsBadRequest) Error() string {
	return fmt.Sprintf("[POST /addressbook/parties/{id}/contacts][%d] postAddressbookPartiesIdContactsBadRequest", 400)
}

func (o *PostAddressbookPartiesIDContactsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAddressbookPartiesIDContactsForbidden creates a PostAddressbookPartiesIDContactsForbidden with default headers values
func NewPostAddressbookPartiesIDContactsForbidden() *PostAddressbookPartiesIDContactsForbidden {
	return &PostAddressbookPartiesIDContactsForbidden{}
}

/*PostAddressbookPartiesIDContactsForbidden handles this case with default header values.

Forbidden
*/
type PostAddressbookPartiesIDContactsForbidden struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *PostAddressbookPartiesIDContactsForbidden) Error() string {
	return fmt.Sprintf("[POST /addressbook/parties/{id}/contacts][%d] postAddressbookPartiesIdContactsForbidden", 403)
}

func (o *PostAddressbookPartiesIDContactsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAddressbookPartiesIDContactsNotFound creates a PostAddressbookPartiesIDContactsNotFound with default headers values
func NewPostAddressbookPartiesIDContactsNotFound() *PostAddressbookPartiesIDContactsNotFound {
	return &PostAddressbookPartiesIDContactsNotFound{}
}

/*PostAddressbookPartiesIDContactsNotFound handles this case with default header values.

Not Found
*/
type PostAddressbookPartiesIDContactsNotFound struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *PostAddressbookPartiesIDContactsNotFound) Error() string {
	return fmt.Sprintf("[POST /addressbook/parties/{id}/contacts][%d] postAddressbookPartiesIdContactsNotFound", 404)
}

func (o *PostAddressbookPartiesIDContactsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAddressbookPartiesIDContactsConflict creates a PostAddressbookPartiesIDContactsConflict with default headers values
func NewPostAddressbookPartiesIDContactsConflict() *PostAddressbookPartiesIDContactsConflict {
	return &PostAddressbookPartiesIDContactsConflict{}
}

/*PostAddressbookPartiesIDContactsConflict handles this case with default header values.

Conflict
*/
type PostAddressbookPartiesIDContactsConflict struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *PostAddressbookPartiesIDContactsConflict) Error() string {
	return fmt.Sprintf("[POST /addressbook/parties/{id}/contacts][%d] postAddressbookPartiesIdContactsConflict", 409)
}

func (o *PostAddressbookPartiesIDContactsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
