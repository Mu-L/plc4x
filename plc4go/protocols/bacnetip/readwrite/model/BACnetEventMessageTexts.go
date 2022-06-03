/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
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

// BACnetEventMessageTexts is the data-structure of this message
type BACnetEventMessageTexts struct {
	ToOffnormalText *BACnetApplicationTagCharacterString
	ToFaultText     *BACnetApplicationTagCharacterString
	ToNormalText    *BACnetApplicationTagCharacterString
}

// IBACnetEventMessageTexts is the corresponding interface of BACnetEventMessageTexts
type IBACnetEventMessageTexts interface {
	// GetToOffnormalText returns ToOffnormalText (property field)
	GetToOffnormalText() *BACnetApplicationTagCharacterString
	// GetToFaultText returns ToFaultText (property field)
	GetToFaultText() *BACnetApplicationTagCharacterString
	// GetToNormalText returns ToNormalText (property field)
	GetToNormalText() *BACnetApplicationTagCharacterString
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetEventMessageTexts) GetToOffnormalText() *BACnetApplicationTagCharacterString {
	return m.ToOffnormalText
}

func (m *BACnetEventMessageTexts) GetToFaultText() *BACnetApplicationTagCharacterString {
	return m.ToFaultText
}

func (m *BACnetEventMessageTexts) GetToNormalText() *BACnetApplicationTagCharacterString {
	return m.ToNormalText
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetEventMessageTexts factory function for BACnetEventMessageTexts
func NewBACnetEventMessageTexts(toOffnormalText *BACnetApplicationTagCharacterString, toFaultText *BACnetApplicationTagCharacterString, toNormalText *BACnetApplicationTagCharacterString) *BACnetEventMessageTexts {
	return &BACnetEventMessageTexts{ToOffnormalText: toOffnormalText, ToFaultText: toFaultText, ToNormalText: toNormalText}
}

func CastBACnetEventMessageTexts(structType interface{}) *BACnetEventMessageTexts {
	if casted, ok := structType.(BACnetEventMessageTexts); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetEventMessageTexts); ok {
		return casted
	}
	return nil
}

func (m *BACnetEventMessageTexts) GetTypeName() string {
	return "BACnetEventMessageTexts"
}

func (m *BACnetEventMessageTexts) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetEventMessageTexts) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (toOffnormalText)
	lengthInBits += m.ToOffnormalText.GetLengthInBits()

	// Simple field (toFaultText)
	lengthInBits += m.ToFaultText.GetLengthInBits()

	// Simple field (toNormalText)
	lengthInBits += m.ToNormalText.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetEventMessageTexts) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetEventMessageTextsParse(readBuffer utils.ReadBuffer) (*BACnetEventMessageTexts, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetEventMessageTexts"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (toOffnormalText)
	if pullErr := readBuffer.PullContext("toOffnormalText"); pullErr != nil {
		return nil, pullErr
	}
	_toOffnormalText, _toOffnormalTextErr := BACnetApplicationTagParse(readBuffer)
	if _toOffnormalTextErr != nil {
		return nil, errors.Wrap(_toOffnormalTextErr, "Error parsing 'toOffnormalText' field")
	}
	toOffnormalText := CastBACnetApplicationTagCharacterString(_toOffnormalText)
	if closeErr := readBuffer.CloseContext("toOffnormalText"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (toFaultText)
	if pullErr := readBuffer.PullContext("toFaultText"); pullErr != nil {
		return nil, pullErr
	}
	_toFaultText, _toFaultTextErr := BACnetApplicationTagParse(readBuffer)
	if _toFaultTextErr != nil {
		return nil, errors.Wrap(_toFaultTextErr, "Error parsing 'toFaultText' field")
	}
	toFaultText := CastBACnetApplicationTagCharacterString(_toFaultText)
	if closeErr := readBuffer.CloseContext("toFaultText"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (toNormalText)
	if pullErr := readBuffer.PullContext("toNormalText"); pullErr != nil {
		return nil, pullErr
	}
	_toNormalText, _toNormalTextErr := BACnetApplicationTagParse(readBuffer)
	if _toNormalTextErr != nil {
		return nil, errors.Wrap(_toNormalTextErr, "Error parsing 'toNormalText' field")
	}
	toNormalText := CastBACnetApplicationTagCharacterString(_toNormalText)
	if closeErr := readBuffer.CloseContext("toNormalText"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetEventMessageTexts"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewBACnetEventMessageTexts(toOffnormalText, toFaultText, toNormalText), nil
}

func (m *BACnetEventMessageTexts) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetEventMessageTexts"); pushErr != nil {
		return pushErr
	}

	// Simple Field (toOffnormalText)
	if pushErr := writeBuffer.PushContext("toOffnormalText"); pushErr != nil {
		return pushErr
	}
	_toOffnormalTextErr := m.ToOffnormalText.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("toOffnormalText"); popErr != nil {
		return popErr
	}
	if _toOffnormalTextErr != nil {
		return errors.Wrap(_toOffnormalTextErr, "Error serializing 'toOffnormalText' field")
	}

	// Simple Field (toFaultText)
	if pushErr := writeBuffer.PushContext("toFaultText"); pushErr != nil {
		return pushErr
	}
	_toFaultTextErr := m.ToFaultText.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("toFaultText"); popErr != nil {
		return popErr
	}
	if _toFaultTextErr != nil {
		return errors.Wrap(_toFaultTextErr, "Error serializing 'toFaultText' field")
	}

	// Simple Field (toNormalText)
	if pushErr := writeBuffer.PushContext("toNormalText"); pushErr != nil {
		return pushErr
	}
	_toNormalTextErr := m.ToNormalText.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("toNormalText"); popErr != nil {
		return popErr
	}
	if _toNormalTextErr != nil {
		return errors.Wrap(_toNormalTextErr, "Error serializing 'toNormalText' field")
	}

	if popErr := writeBuffer.PopContext("BACnetEventMessageTexts"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *BACnetEventMessageTexts) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}