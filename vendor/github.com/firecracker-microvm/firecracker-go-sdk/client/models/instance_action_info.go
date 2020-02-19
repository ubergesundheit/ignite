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

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// InstanceActionInfo Variant wrapper containing the real action.
// swagger:model InstanceActionInfo
type InstanceActionInfo struct {

	// Enumeration indicating what type of action is contained in the payload
	// Required: true
	// Enum: [BlockDeviceRescan FlushMetrics InstanceStart SendCtrlAltDel]
	ActionType *string `json:"action_type"`

	// payload
	Payload string `json:"payload,omitempty"`
}

// Validate validates this instance action info
func (m *InstanceActionInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActionType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var instanceActionInfoTypeActionTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["BlockDeviceRescan","FlushMetrics","InstanceStart","SendCtrlAltDel"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		instanceActionInfoTypeActionTypePropEnum = append(instanceActionInfoTypeActionTypePropEnum, v)
	}
}

const (

	// InstanceActionInfoActionTypeBlockDeviceRescan captures enum value "BlockDeviceRescan"
	InstanceActionInfoActionTypeBlockDeviceRescan string = "BlockDeviceRescan"

	// InstanceActionInfoActionTypeFlushMetrics captures enum value "FlushMetrics"
	InstanceActionInfoActionTypeFlushMetrics string = "FlushMetrics"

	// InstanceActionInfoActionTypeInstanceStart captures enum value "InstanceStart"
	InstanceActionInfoActionTypeInstanceStart string = "InstanceStart"

	// InstanceActionInfoActionTypeSendCtrlAltDel captures enum value "SendCtrlAltDel"
	InstanceActionInfoActionTypeSendCtrlAltDel string = "SendCtrlAltDel"
)

// prop value enum
func (m *InstanceActionInfo) validateActionTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, instanceActionInfoTypeActionTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *InstanceActionInfo) validateActionType(formats strfmt.Registry) error {

	if err := validate.Required("action_type", "body", m.ActionType); err != nil {
		return err
	}

	// value enum
	if err := m.validateActionTypeEnum("action_type", "body", *m.ActionType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InstanceActionInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InstanceActionInfo) UnmarshalBinary(b []byte) error {
	var res InstanceActionInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
