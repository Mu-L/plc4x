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

// BACnetConstructedDataTransition is the data-structure of this message
type BACnetConstructedDataTransition struct {
	*BACnetConstructedData
	Transition *BACnetLightingTransitionTagged

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataTransition is the corresponding interface of BACnetConstructedDataTransition
type IBACnetConstructedDataTransition interface {
	IBACnetConstructedData
	// GetTransition returns Transition (property field)
	GetTransition() *BACnetLightingTransitionTagged
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

func (m *BACnetConstructedDataTransition) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataTransition) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_TRANSITION
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataTransition) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataTransition) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataTransition) GetTransition() *BACnetLightingTransitionTagged {
	return m.Transition
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataTransition factory function for BACnetConstructedDataTransition
func NewBACnetConstructedDataTransition(transition *BACnetLightingTransitionTagged, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataTransition {
	_result := &BACnetConstructedDataTransition{
		Transition:            transition,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataTransition(structType interface{}) *BACnetConstructedDataTransition {
	if casted, ok := structType.(BACnetConstructedDataTransition); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataTransition); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataTransition(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataTransition(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataTransition) GetTypeName() string {
	return "BACnetConstructedDataTransition"
}

func (m *BACnetConstructedDataTransition) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataTransition) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (transition)
	lengthInBits += m.Transition.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataTransition) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataTransitionParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataTransition, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataTransition"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataTransition")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (transition)
	if pullErr := readBuffer.PullContext("transition"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for transition")
	}
	_transition, _transitionErr := BACnetLightingTransitionTaggedParse(readBuffer, uint8(uint8(0)), TagClass(TagClass_APPLICATION_TAGS))
	if _transitionErr != nil {
		return nil, errors.Wrap(_transitionErr, "Error parsing 'transition' field")
	}
	transition := CastBACnetLightingTransitionTagged(_transition)
	if closeErr := readBuffer.CloseContext("transition"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for transition")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataTransition"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataTransition")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataTransition{
		Transition:            CastBACnetLightingTransitionTagged(transition),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataTransition) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataTransition"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataTransition")
		}

		// Simple Field (transition)
		if pushErr := writeBuffer.PushContext("transition"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for transition")
		}
		_transitionErr := m.Transition.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("transition"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for transition")
		}
		if _transitionErr != nil {
			return errors.Wrap(_transitionErr, "Error serializing 'transition' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataTransition"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataTransition")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataTransition) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}