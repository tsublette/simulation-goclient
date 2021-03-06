package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/3dsim/simulation-goclient/models"
)

// AddMaterialReader is a Reader for the AddMaterial structure.
type AddMaterialReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddMaterialReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAddMaterialOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewAddMaterialUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewAddMaterialForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewAddMaterialDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddMaterialOK creates a AddMaterialOK with default headers values
func NewAddMaterialOK() *AddMaterialOK {
	return &AddMaterialOK{}
}

/*AddMaterialOK handles this case with default header values.

material response
*/
type AddMaterialOK struct {
	Payload *models.Material
}

func (o *AddMaterialOK) Error() string {
	return fmt.Sprintf("[POST /materials][%d] addMaterialOK  %+v", 200, o.Payload)
}

func (o *AddMaterialOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Material)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddMaterialUnauthorized creates a AddMaterialUnauthorized with default headers values
func NewAddMaterialUnauthorized() *AddMaterialUnauthorized {
	return &AddMaterialUnauthorized{}
}

/*AddMaterialUnauthorized handles this case with default header values.

Not authorized
*/
type AddMaterialUnauthorized struct {
	Payload *models.Error
}

func (o *AddMaterialUnauthorized) Error() string {
	return fmt.Sprintf("[POST /materials][%d] addMaterialUnauthorized  %+v", 401, o.Payload)
}

func (o *AddMaterialUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddMaterialForbidden creates a AddMaterialForbidden with default headers values
func NewAddMaterialForbidden() *AddMaterialForbidden {
	return &AddMaterialForbidden{}
}

/*AddMaterialForbidden handles this case with default header values.

Forbidden
*/
type AddMaterialForbidden struct {
	Payload *models.Error
}

func (o *AddMaterialForbidden) Error() string {
	return fmt.Sprintf("[POST /materials][%d] addMaterialForbidden  %+v", 403, o.Payload)
}

func (o *AddMaterialForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddMaterialDefault creates a AddMaterialDefault with default headers values
func NewAddMaterialDefault(code int) *AddMaterialDefault {
	return &AddMaterialDefault{
		_statusCode: code,
	}
}

/*AddMaterialDefault handles this case with default header values.

unexpected error
*/
type AddMaterialDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the add material default response
func (o *AddMaterialDefault) Code() int {
	return o._statusCode
}

func (o *AddMaterialDefault) Error() string {
	return fmt.Sprintf("[POST /materials][%d] addMaterial default  %+v", o._statusCode, o.Payload)
}

func (o *AddMaterialDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
