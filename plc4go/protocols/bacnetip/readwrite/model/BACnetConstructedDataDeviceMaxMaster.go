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

// BACnetConstructedDataDeviceMaxMaster is the data-structure of this message
type BACnetConstructedDataDeviceMaxMaster struct {
	*BACnetConstructedData
	MaxMaster *BACnetApplicationTagUnsignedInteger

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataDeviceMaxMaster is the corresponding interface of BACnetConstructedDataDeviceMaxMaster
type IBACnetConstructedDataDeviceMaxMaster interface {
	IBACnetConstructedData
	// GetMaxMaster returns MaxMaster (property field)
	GetMaxMaster() *BACnetApplicationTagUnsignedInteger
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *BACnetConstructedDataDeviceMaxMaster) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_DEVICE
}

func (m *BACnetConstructedDataDeviceMaxMaster) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_MAX_MASTER
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataDeviceMaxMaster) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataDeviceMaxMaster) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataDeviceMaxMaster) GetMaxMaster() *BACnetApplicationTagUnsignedInteger {
	return m.MaxMaster
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataDeviceMaxMaster factory function for BACnetConstructedDataDeviceMaxMaster
func NewBACnetConstructedDataDeviceMaxMaster(maxMaster *BACnetApplicationTagUnsignedInteger, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataDeviceMaxMaster {
	_result := &BACnetConstructedDataDeviceMaxMaster{
		MaxMaster:             maxMaster,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataDeviceMaxMaster(structType interface{}) *BACnetConstructedDataDeviceMaxMaster {
	if casted, ok := structType.(BACnetConstructedDataDeviceMaxMaster); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataDeviceMaxMaster); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataDeviceMaxMaster(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataDeviceMaxMaster(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataDeviceMaxMaster) GetTypeName() string {
	return "BACnetConstructedDataDeviceMaxMaster"
}

func (m *BACnetConstructedDataDeviceMaxMaster) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataDeviceMaxMaster) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (maxMaster)
	lengthInBits += m.MaxMaster.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataDeviceMaxMaster) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataDeviceMaxMasterParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataDeviceMaxMaster, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataDeviceMaxMaster"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataDeviceMaxMaster")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (maxMaster)
	if pullErr := readBuffer.PullContext("maxMaster"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for maxMaster")
	}
	_maxMaster, _maxMasterErr := BACnetApplicationTagParse(readBuffer)
	if _maxMasterErr != nil {
		return nil, errors.Wrap(_maxMasterErr, "Error parsing 'maxMaster' field")
	}
	maxMaster := CastBACnetApplicationTagUnsignedInteger(_maxMaster)
	if closeErr := readBuffer.CloseContext("maxMaster"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for maxMaster")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataDeviceMaxMaster"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataDeviceMaxMaster")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataDeviceMaxMaster{
		MaxMaster:             CastBACnetApplicationTagUnsignedInteger(maxMaster),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataDeviceMaxMaster) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataDeviceMaxMaster"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataDeviceMaxMaster")
		}

		// Simple Field (maxMaster)
		if pushErr := writeBuffer.PushContext("maxMaster"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for maxMaster")
		}
		_maxMasterErr := m.MaxMaster.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("maxMaster"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for maxMaster")
		}
		if _maxMasterErr != nil {
			return errors.Wrap(_maxMasterErr, "Error serializing 'maxMaster' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataDeviceMaxMaster"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataDeviceMaxMaster")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataDeviceMaxMaster) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}