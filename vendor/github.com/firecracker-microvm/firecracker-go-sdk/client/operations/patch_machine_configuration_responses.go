// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2018 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
// 	http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/firecracker-microvm/firecracker-go-sdk/client/models"
)

// PatchMachineConfigurationReader is a Reader for the PatchMachineConfiguration structure.
type PatchMachineConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchMachineConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPatchMachineConfigurationNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchMachineConfigurationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPatchMachineConfigurationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchMachineConfigurationNoContent creates a PatchMachineConfigurationNoContent with default headers values
func NewPatchMachineConfigurationNoContent() *PatchMachineConfigurationNoContent {
	return &PatchMachineConfigurationNoContent{}
}

/*PatchMachineConfigurationNoContent handles this case with default header values.

Machine Configuration created/updated
*/
type PatchMachineConfigurationNoContent struct {
}

func (o *PatchMachineConfigurationNoContent) Error() string {
	return fmt.Sprintf("[PATCH /machine-config][%d] patchMachineConfigurationNoContent ", 204)
}

func (o *PatchMachineConfigurationNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchMachineConfigurationBadRequest creates a PatchMachineConfigurationBadRequest with default headers values
func NewPatchMachineConfigurationBadRequest() *PatchMachineConfigurationBadRequest {
	return &PatchMachineConfigurationBadRequest{}
}

/*PatchMachineConfigurationBadRequest handles this case with default header values.

Machine Configuration cannot be updated due to bad input
*/
type PatchMachineConfigurationBadRequest struct {
	Payload *models.Error
}

func (o *PatchMachineConfigurationBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /machine-config][%d] patchMachineConfigurationBadRequest  %+v", 400, o.Payload)
}

func (o *PatchMachineConfigurationBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchMachineConfigurationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchMachineConfigurationDefault creates a PatchMachineConfigurationDefault with default headers values
func NewPatchMachineConfigurationDefault(code int) *PatchMachineConfigurationDefault {
	return &PatchMachineConfigurationDefault{
		_statusCode: code,
	}
}

/*PatchMachineConfigurationDefault handles this case with default header values.

Internal server error
*/
type PatchMachineConfigurationDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the patch machine configuration default response
func (o *PatchMachineConfigurationDefault) Code() int {
	return o._statusCode
}

func (o *PatchMachineConfigurationDefault) Error() string {
	return fmt.Sprintf("[PATCH /machine-config][%d] patchMachineConfiguration default  %+v", o._statusCode, o.Payload)
}

func (o *PatchMachineConfigurationDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchMachineConfigurationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
