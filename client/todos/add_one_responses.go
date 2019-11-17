// Code generated by go-swagger; DO NOT EDIT.

package todos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/spkane/todo-api-example/models"
)

// AddOneReader is a Reader for the AddOne structure.
type AddOneReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddOneReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewAddOneCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAddOneDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddOneCreated creates a AddOneCreated with default headers values
func NewAddOneCreated() *AddOneCreated {
	return &AddOneCreated{}
}

/*AddOneCreated handles this case with default header values.

Created
*/
type AddOneCreated struct {
	Payload *models.Item
}

func (o *AddOneCreated) Error() string {
	return fmt.Sprintf("[POST /][%d] addOneCreated  %+v", 201, o.Payload)
}

func (o *AddOneCreated) GetPayload() *models.Item {
	return o.Payload
}

func (o *AddOneCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Item)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddOneDefault creates a AddOneDefault with default headers values
func NewAddOneDefault(code int) *AddOneDefault {
	return &AddOneDefault{
		_statusCode: code,
	}
}

/*AddOneDefault handles this case with default header values.

error
*/
type AddOneDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the add one default response
func (o *AddOneDefault) Code() int {
	return o._statusCode
}

func (o *AddOneDefault) Error() string {
	return fmt.Sprintf("[POST /][%d] addOne default  %+v", o._statusCode, o.Payload)
}

func (o *AddOneDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *AddOneDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
