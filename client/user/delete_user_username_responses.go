// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteUserUsernameReader is a Reader for the DeleteUserUsername structure.
type DeleteUserUsernameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteUserUsernameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteUserUsernameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteUserUsernameBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteUserUsernameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteUserUsernameOK creates a DeleteUserUsernameOK with default headers values
func NewDeleteUserUsernameOK() *DeleteUserUsernameOK {
	return &DeleteUserUsernameOK{}
}

/*DeleteUserUsernameOK handles this case with default header values.

Successful
*/
type DeleteUserUsernameOK struct {
}

func (o *DeleteUserUsernameOK) Error() string {
	return fmt.Sprintf("[DELETE /user/{username}][%d] deleteUserUsernameOK ", 200)
}

func (o *DeleteUserUsernameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUserUsernameBadRequest creates a DeleteUserUsernameBadRequest with default headers values
func NewDeleteUserUsernameBadRequest() *DeleteUserUsernameBadRequest {
	return &DeleteUserUsernameBadRequest{}
}

/*DeleteUserUsernameBadRequest handles this case with default header values.

Invalid username
*/
type DeleteUserUsernameBadRequest struct {
}

func (o *DeleteUserUsernameBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /user/{username}][%d] deleteUserUsernameBadRequest ", 400)
}

func (o *DeleteUserUsernameBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUserUsernameNotFound creates a DeleteUserUsernameNotFound with default headers values
func NewDeleteUserUsernameNotFound() *DeleteUserUsernameNotFound {
	return &DeleteUserUsernameNotFound{}
}

/*DeleteUserUsernameNotFound handles this case with default header values.

User not found
*/
type DeleteUserUsernameNotFound struct {
}

func (o *DeleteUserUsernameNotFound) Error() string {
	return fmt.Sprintf("[DELETE /user/{username}][%d] deleteUserUsernameNotFound ", 404)
}

func (o *DeleteUserUsernameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
