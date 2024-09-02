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

// BACnetConstructedDataDoorUnlockDelayTime is the corresponding interface of BACnetConstructedDataDoorUnlockDelayTime
type BACnetConstructedDataDoorUnlockDelayTime interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetDoorUnlockDelayTime returns DoorUnlockDelayTime (property field)
	GetDoorUnlockDelayTime() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetConstructedDataDoorUnlockDelayTime is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataDoorUnlockDelayTime()
}

// _BACnetConstructedDataDoorUnlockDelayTime is the data-structure of this message
type _BACnetConstructedDataDoorUnlockDelayTime struct {
	BACnetConstructedDataContract
	DoorUnlockDelayTime BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataDoorUnlockDelayTime = (*_BACnetConstructedDataDoorUnlockDelayTime)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataDoorUnlockDelayTime)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataDoorUnlockDelayTime) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataDoorUnlockDelayTime) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_DOOR_UNLOCK_DELAY_TIME
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataDoorUnlockDelayTime) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataDoorUnlockDelayTime) GetDoorUnlockDelayTime() BACnetApplicationTagUnsignedInteger {
	return m.DoorUnlockDelayTime
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataDoorUnlockDelayTime) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetDoorUnlockDelayTime())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataDoorUnlockDelayTime factory function for _BACnetConstructedDataDoorUnlockDelayTime
func NewBACnetConstructedDataDoorUnlockDelayTime(doorUnlockDelayTime BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataDoorUnlockDelayTime {
	_result := &_BACnetConstructedDataDoorUnlockDelayTime{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		DoorUnlockDelayTime:           doorUnlockDelayTime,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataDoorUnlockDelayTime(structType any) BACnetConstructedDataDoorUnlockDelayTime {
	if casted, ok := structType.(BACnetConstructedDataDoorUnlockDelayTime); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataDoorUnlockDelayTime); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataDoorUnlockDelayTime) GetTypeName() string {
	return "BACnetConstructedDataDoorUnlockDelayTime"
}

func (m *_BACnetConstructedDataDoorUnlockDelayTime) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (doorUnlockDelayTime)
	lengthInBits += m.DoorUnlockDelayTime.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataDoorUnlockDelayTime) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataDoorUnlockDelayTime) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataDoorUnlockDelayTime BACnetConstructedDataDoorUnlockDelayTime, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataDoorUnlockDelayTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataDoorUnlockDelayTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	doorUnlockDelayTime, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "doorUnlockDelayTime", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'doorUnlockDelayTime' field"))
	}
	m.DoorUnlockDelayTime = doorUnlockDelayTime

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), doorUnlockDelayTime)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataDoorUnlockDelayTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataDoorUnlockDelayTime")
	}

	return m, nil
}

func (m *_BACnetConstructedDataDoorUnlockDelayTime) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataDoorUnlockDelayTime) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataDoorUnlockDelayTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataDoorUnlockDelayTime")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "doorUnlockDelayTime", m.GetDoorUnlockDelayTime(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'doorUnlockDelayTime' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataDoorUnlockDelayTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataDoorUnlockDelayTime")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataDoorUnlockDelayTime) IsBACnetConstructedDataDoorUnlockDelayTime() {}

func (m *_BACnetConstructedDataDoorUnlockDelayTime) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
