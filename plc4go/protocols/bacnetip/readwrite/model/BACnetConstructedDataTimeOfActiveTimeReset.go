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

// BACnetConstructedDataTimeOfActiveTimeReset is the corresponding interface of BACnetConstructedDataTimeOfActiveTimeReset
type BACnetConstructedDataTimeOfActiveTimeReset interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetTimeOfActiveTimeReset returns TimeOfActiveTimeReset (property field)
	GetTimeOfActiveTimeReset() BACnetDateTime
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetDateTime
	// IsBACnetConstructedDataTimeOfActiveTimeReset is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataTimeOfActiveTimeReset()
}

// _BACnetConstructedDataTimeOfActiveTimeReset is the data-structure of this message
type _BACnetConstructedDataTimeOfActiveTimeReset struct {
	BACnetConstructedDataContract
	TimeOfActiveTimeReset BACnetDateTime
}

var _ BACnetConstructedDataTimeOfActiveTimeReset = (*_BACnetConstructedDataTimeOfActiveTimeReset)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataTimeOfActiveTimeReset)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataTimeOfActiveTimeReset) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataTimeOfActiveTimeReset) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_TIME_OF_ACTIVE_TIME_RESET
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataTimeOfActiveTimeReset) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataTimeOfActiveTimeReset) GetTimeOfActiveTimeReset() BACnetDateTime {
	return m.TimeOfActiveTimeReset
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataTimeOfActiveTimeReset) GetActualValue() BACnetDateTime {
	ctx := context.Background()
	_ = ctx
	return CastBACnetDateTime(m.GetTimeOfActiveTimeReset())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataTimeOfActiveTimeReset factory function for _BACnetConstructedDataTimeOfActiveTimeReset
func NewBACnetConstructedDataTimeOfActiveTimeReset(timeOfActiveTimeReset BACnetDateTime, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataTimeOfActiveTimeReset {
	_result := &_BACnetConstructedDataTimeOfActiveTimeReset{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		TimeOfActiveTimeReset:         timeOfActiveTimeReset,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataTimeOfActiveTimeReset(structType any) BACnetConstructedDataTimeOfActiveTimeReset {
	if casted, ok := structType.(BACnetConstructedDataTimeOfActiveTimeReset); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataTimeOfActiveTimeReset); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataTimeOfActiveTimeReset) GetTypeName() string {
	return "BACnetConstructedDataTimeOfActiveTimeReset"
}

func (m *_BACnetConstructedDataTimeOfActiveTimeReset) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (timeOfActiveTimeReset)
	lengthInBits += m.TimeOfActiveTimeReset.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataTimeOfActiveTimeReset) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataTimeOfActiveTimeReset) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataTimeOfActiveTimeReset BACnetConstructedDataTimeOfActiveTimeReset, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataTimeOfActiveTimeReset"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataTimeOfActiveTimeReset")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	timeOfActiveTimeReset, err := ReadSimpleField[BACnetDateTime](ctx, "timeOfActiveTimeReset", ReadComplex[BACnetDateTime](BACnetDateTimeParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timeOfActiveTimeReset' field"))
	}
	m.TimeOfActiveTimeReset = timeOfActiveTimeReset

	actualValue, err := ReadVirtualField[BACnetDateTime](ctx, "actualValue", (*BACnetDateTime)(nil), timeOfActiveTimeReset)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataTimeOfActiveTimeReset"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataTimeOfActiveTimeReset")
	}

	return m, nil
}

func (m *_BACnetConstructedDataTimeOfActiveTimeReset) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataTimeOfActiveTimeReset) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataTimeOfActiveTimeReset"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataTimeOfActiveTimeReset")
		}

		if err := WriteSimpleField[BACnetDateTime](ctx, "timeOfActiveTimeReset", m.GetTimeOfActiveTimeReset(), WriteComplex[BACnetDateTime](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'timeOfActiveTimeReset' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataTimeOfActiveTimeReset"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataTimeOfActiveTimeReset")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataTimeOfActiveTimeReset) IsBACnetConstructedDataTimeOfActiveTimeReset() {
}

func (m *_BACnetConstructedDataTimeOfActiveTimeReset) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
