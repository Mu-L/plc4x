/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package model

import (
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// TelephonyDataDialOutFailure is the corresponding interface of TelephonyDataDialOutFailure
type TelephonyDataDialOutFailure interface {
	utils.LengthAware
	utils.Serializable
	TelephonyData
	// GetReason returns Reason (property field)
	GetReason() DialInFailureReason
}

// TelephonyDataDialOutFailureExactly can be used when we want exactly this type and not a type which fulfills TelephonyDataDialOutFailure.
// This is useful for switch cases.
type TelephonyDataDialOutFailureExactly interface {
	TelephonyDataDialOutFailure
	isTelephonyDataDialOutFailure() bool
}

// _TelephonyDataDialOutFailure is the data-structure of this message
type _TelephonyDataDialOutFailure struct {
	*_TelephonyData
	Reason DialInFailureReason
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_TelephonyDataDialOutFailure) InitializeParent(parent TelephonyData, commandTypeContainer TelephonyCommandTypeContainer, argument byte) {
	m.CommandTypeContainer = commandTypeContainer
	m.Argument = argument
}

func (m *_TelephonyDataDialOutFailure) GetParent() TelephonyData {
	return m._TelephonyData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_TelephonyDataDialOutFailure) GetReason() DialInFailureReason {
	return m.Reason
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewTelephonyDataDialOutFailure factory function for _TelephonyDataDialOutFailure
func NewTelephonyDataDialOutFailure(reason DialInFailureReason, commandTypeContainer TelephonyCommandTypeContainer, argument byte) *_TelephonyDataDialOutFailure {
	_result := &_TelephonyDataDialOutFailure{
		Reason:         reason,
		_TelephonyData: NewTelephonyData(commandTypeContainer, argument),
	}
	_result._TelephonyData._TelephonyDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastTelephonyDataDialOutFailure(structType interface{}) TelephonyDataDialOutFailure {
	if casted, ok := structType.(TelephonyDataDialOutFailure); ok {
		return casted
	}
	if casted, ok := structType.(*TelephonyDataDialOutFailure); ok {
		return *casted
	}
	return nil
}

func (m *_TelephonyDataDialOutFailure) GetTypeName() string {
	return "TelephonyDataDialOutFailure"
}

func (m *_TelephonyDataDialOutFailure) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_TelephonyDataDialOutFailure) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (reason)
	lengthInBits += 8

	return lengthInBits
}

func (m *_TelephonyDataDialOutFailure) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func TelephonyDataDialOutFailureParse(readBuffer utils.ReadBuffer) (TelephonyDataDialOutFailure, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("TelephonyDataDialOutFailure"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for TelephonyDataDialOutFailure")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (reason)
	if pullErr := readBuffer.PullContext("reason"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for reason")
	}
	_reason, _reasonErr := DialInFailureReasonParse(readBuffer)
	if _reasonErr != nil {
		return nil, errors.Wrap(_reasonErr, "Error parsing 'reason' field of TelephonyDataDialOutFailure")
	}
	reason := _reason
	if closeErr := readBuffer.CloseContext("reason"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for reason")
	}

	if closeErr := readBuffer.CloseContext("TelephonyDataDialOutFailure"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for TelephonyDataDialOutFailure")
	}

	// Create a partially initialized instance
	_child := &_TelephonyDataDialOutFailure{
		Reason:         reason,
		_TelephonyData: &_TelephonyData{},
	}
	_child._TelephonyData._TelephonyDataChildRequirements = _child
	return _child, nil
}

func (m *_TelephonyDataDialOutFailure) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("TelephonyDataDialOutFailure"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for TelephonyDataDialOutFailure")
		}

		// Simple Field (reason)
		if pushErr := writeBuffer.PushContext("reason"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for reason")
		}
		_reasonErr := writeBuffer.WriteSerializable(m.GetReason())
		if popErr := writeBuffer.PopContext("reason"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for reason")
		}
		if _reasonErr != nil {
			return errors.Wrap(_reasonErr, "Error serializing 'reason' field")
		}

		if popErr := writeBuffer.PopContext("TelephonyDataDialOutFailure"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for TelephonyDataDialOutFailure")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_TelephonyDataDialOutFailure) isTelephonyDataDialOutFailure() bool {
	return true
}

func (m *_TelephonyDataDialOutFailure) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}