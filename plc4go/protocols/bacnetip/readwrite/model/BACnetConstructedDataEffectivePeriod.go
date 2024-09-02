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

// BACnetConstructedDataEffectivePeriod is the corresponding interface of BACnetConstructedDataEffectivePeriod
type BACnetConstructedDataEffectivePeriod interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetDateRange returns DateRange (property field)
	GetDateRange() BACnetDateRange
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetDateRange
}

// BACnetConstructedDataEffectivePeriodExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataEffectivePeriod.
// This is useful for switch cases.
type BACnetConstructedDataEffectivePeriodExactly interface {
	BACnetConstructedDataEffectivePeriod
	isBACnetConstructedDataEffectivePeriod() bool
}

// _BACnetConstructedDataEffectivePeriod is the data-structure of this message
type _BACnetConstructedDataEffectivePeriod struct {
	BACnetConstructedDataContract
	DateRange BACnetDateRange
}

var _ BACnetConstructedDataEffectivePeriod = (*_BACnetConstructedDataEffectivePeriod)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataEffectivePeriod)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataEffectivePeriod) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataEffectivePeriod) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_EFFECTIVE_PERIOD
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataEffectivePeriod) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataEffectivePeriod) GetDateRange() BACnetDateRange {
	return m.DateRange
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataEffectivePeriod) GetActualValue() BACnetDateRange {
	ctx := context.Background()
	_ = ctx
	return CastBACnetDateRange(m.GetDateRange())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataEffectivePeriod factory function for _BACnetConstructedDataEffectivePeriod
func NewBACnetConstructedDataEffectivePeriod(dateRange BACnetDateRange, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataEffectivePeriod {
	_result := &_BACnetConstructedDataEffectivePeriod{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		DateRange:                     dateRange,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataEffectivePeriod(structType any) BACnetConstructedDataEffectivePeriod {
	if casted, ok := structType.(BACnetConstructedDataEffectivePeriod); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataEffectivePeriod); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataEffectivePeriod) GetTypeName() string {
	return "BACnetConstructedDataEffectivePeriod"
}

func (m *_BACnetConstructedDataEffectivePeriod) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (dateRange)
	lengthInBits += m.DateRange.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataEffectivePeriod) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataEffectivePeriod) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataEffectivePeriod BACnetConstructedDataEffectivePeriod, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataEffectivePeriod"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataEffectivePeriod")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	dateRange, err := ReadSimpleField[BACnetDateRange](ctx, "dateRange", ReadComplex[BACnetDateRange](BACnetDateRangeParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dateRange' field"))
	}
	m.DateRange = dateRange

	actualValue, err := ReadVirtualField[BACnetDateRange](ctx, "actualValue", (*BACnetDateRange)(nil), dateRange)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataEffectivePeriod"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataEffectivePeriod")
	}

	return m, nil
}

func (m *_BACnetConstructedDataEffectivePeriod) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataEffectivePeriod) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataEffectivePeriod"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataEffectivePeriod")
		}

		if err := WriteSimpleField[BACnetDateRange](ctx, "dateRange", m.GetDateRange(), WriteComplex[BACnetDateRange](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'dateRange' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataEffectivePeriod"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataEffectivePeriod")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataEffectivePeriod) isBACnetConstructedDataEffectivePeriod() bool {
	return true
}

func (m *_BACnetConstructedDataEffectivePeriod) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
