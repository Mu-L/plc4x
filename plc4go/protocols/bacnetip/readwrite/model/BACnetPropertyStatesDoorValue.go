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
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetPropertyStatesDoorValue is the corresponding interface of BACnetPropertyStatesDoorValue
type BACnetPropertyStatesDoorValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetPropertyStates
	// GetDoorValue returns DoorValue (property field)
	GetDoorValue() BACnetDoorValueTagged
}

// BACnetPropertyStatesDoorValueExactly can be used when we want exactly this type and not a type which fulfills BACnetPropertyStatesDoorValue.
// This is useful for switch cases.
type BACnetPropertyStatesDoorValueExactly interface {
	BACnetPropertyStatesDoorValue
	isBACnetPropertyStatesDoorValue() bool
}

// _BACnetPropertyStatesDoorValue is the data-structure of this message
type _BACnetPropertyStatesDoorValue struct {
	*_BACnetPropertyStates
	DoorValue BACnetDoorValueTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPropertyStatesDoorValue) InitializeParent(parent BACnetPropertyStates, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetPropertyStatesDoorValue) GetParent() BACnetPropertyStates {
	return m._BACnetPropertyStates
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesDoorValue) GetDoorValue() BACnetDoorValueTagged {
	return m.DoorValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStatesDoorValue factory function for _BACnetPropertyStatesDoorValue
func NewBACnetPropertyStatesDoorValue(doorValue BACnetDoorValueTagged, peekedTagHeader BACnetTagHeader) *_BACnetPropertyStatesDoorValue {
	_result := &_BACnetPropertyStatesDoorValue{
		DoorValue:             doorValue,
		_BACnetPropertyStates: NewBACnetPropertyStates(peekedTagHeader),
	}
	_result._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesDoorValue(structType any) BACnetPropertyStatesDoorValue {
	if casted, ok := structType.(BACnetPropertyStatesDoorValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesDoorValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesDoorValue) GetTypeName() string {
	return "BACnetPropertyStatesDoorValue"
}

func (m *_BACnetPropertyStatesDoorValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (doorValue)
	lengthInBits += m.DoorValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPropertyStatesDoorValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetPropertyStatesDoorValueParse(ctx context.Context, theBytes []byte, peekedTagNumber uint8) (BACnetPropertyStatesDoorValue, error) {
	return BACnetPropertyStatesDoorValueParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), peekedTagNumber)
}

func BACnetPropertyStatesDoorValueParseWithBufferProducer(peekedTagNumber uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetPropertyStatesDoorValue, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetPropertyStatesDoorValue, error) {
		return BACnetPropertyStatesDoorValueParseWithBuffer(ctx, readBuffer, peekedTagNumber)
	}
}

func BACnetPropertyStatesDoorValueParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, peekedTagNumber uint8) (BACnetPropertyStatesDoorValue, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesDoorValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesDoorValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	doorValue, err := ReadSimpleField[BACnetDoorValueTagged](ctx, "doorValue", ReadComplex[BACnetDoorValueTagged](BACnetDoorValueTaggedParseWithBufferProducer((uint8)(peekedTagNumber), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'doorValue' field"))
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesDoorValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesDoorValue")
	}

	// Create a partially initialized instance
	_child := &_BACnetPropertyStatesDoorValue{
		_BACnetPropertyStates: &_BACnetPropertyStates{},
		DoorValue:             doorValue,
	}
	_child._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _child
	return _child, nil
}

func (m *_BACnetPropertyStatesDoorValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesDoorValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesDoorValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesDoorValue")
		}

		// Simple Field (doorValue)
		if pushErr := writeBuffer.PushContext("doorValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for doorValue")
		}
		_doorValueErr := writeBuffer.WriteSerializable(ctx, m.GetDoorValue())
		if popErr := writeBuffer.PopContext("doorValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for doorValue")
		}
		if _doorValueErr != nil {
			return errors.Wrap(_doorValueErr, "Error serializing 'doorValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesDoorValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesDoorValue")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesDoorValue) isBACnetPropertyStatesDoorValue() bool {
	return true
}

func (m *_BACnetPropertyStatesDoorValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
