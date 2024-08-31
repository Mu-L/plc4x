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

// BACnetCalendarEntryWeekNDay is the corresponding interface of BACnetCalendarEntryWeekNDay
type BACnetCalendarEntryWeekNDay interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetCalendarEntry
	// GetWeekNDay returns WeekNDay (property field)
	GetWeekNDay() BACnetWeekNDayTagged
}

// BACnetCalendarEntryWeekNDayExactly can be used when we want exactly this type and not a type which fulfills BACnetCalendarEntryWeekNDay.
// This is useful for switch cases.
type BACnetCalendarEntryWeekNDayExactly interface {
	BACnetCalendarEntryWeekNDay
	isBACnetCalendarEntryWeekNDay() bool
}

// _BACnetCalendarEntryWeekNDay is the data-structure of this message
type _BACnetCalendarEntryWeekNDay struct {
	*_BACnetCalendarEntry
	WeekNDay BACnetWeekNDayTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetCalendarEntryWeekNDay) InitializeParent(parent BACnetCalendarEntry, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetCalendarEntryWeekNDay) GetParent() BACnetCalendarEntry {
	return m._BACnetCalendarEntry
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetCalendarEntryWeekNDay) GetWeekNDay() BACnetWeekNDayTagged {
	return m.WeekNDay
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetCalendarEntryWeekNDay factory function for _BACnetCalendarEntryWeekNDay
func NewBACnetCalendarEntryWeekNDay(weekNDay BACnetWeekNDayTagged, peekedTagHeader BACnetTagHeader) *_BACnetCalendarEntryWeekNDay {
	_result := &_BACnetCalendarEntryWeekNDay{
		WeekNDay:             weekNDay,
		_BACnetCalendarEntry: NewBACnetCalendarEntry(peekedTagHeader),
	}
	_result._BACnetCalendarEntry._BACnetCalendarEntryChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetCalendarEntryWeekNDay(structType any) BACnetCalendarEntryWeekNDay {
	if casted, ok := structType.(BACnetCalendarEntryWeekNDay); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetCalendarEntryWeekNDay); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetCalendarEntryWeekNDay) GetTypeName() string {
	return "BACnetCalendarEntryWeekNDay"
}

func (m *_BACnetCalendarEntryWeekNDay) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (weekNDay)
	lengthInBits += m.WeekNDay.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetCalendarEntryWeekNDay) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetCalendarEntryWeekNDayParse(ctx context.Context, theBytes []byte) (BACnetCalendarEntryWeekNDay, error) {
	return BACnetCalendarEntryWeekNDayParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetCalendarEntryWeekNDayParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetCalendarEntryWeekNDay, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetCalendarEntryWeekNDay, error) {
		return BACnetCalendarEntryWeekNDayParseWithBuffer(ctx, readBuffer)
	}
}

func BACnetCalendarEntryWeekNDayParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetCalendarEntryWeekNDay, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetCalendarEntryWeekNDay"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetCalendarEntryWeekNDay")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	weekNDay, err := ReadSimpleField[BACnetWeekNDayTagged](ctx, "weekNDay", ReadComplex[BACnetWeekNDayTagged](BACnetWeekNDayTaggedParseWithBufferProducer((uint8)(uint8(2)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'weekNDay' field"))
	}

	if closeErr := readBuffer.CloseContext("BACnetCalendarEntryWeekNDay"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetCalendarEntryWeekNDay")
	}

	// Create a partially initialized instance
	_child := &_BACnetCalendarEntryWeekNDay{
		_BACnetCalendarEntry: &_BACnetCalendarEntry{},
		WeekNDay:             weekNDay,
	}
	_child._BACnetCalendarEntry._BACnetCalendarEntryChildRequirements = _child
	return _child, nil
}

func (m *_BACnetCalendarEntryWeekNDay) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetCalendarEntryWeekNDay) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetCalendarEntryWeekNDay"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetCalendarEntryWeekNDay")
		}

		// Simple Field (weekNDay)
		if pushErr := writeBuffer.PushContext("weekNDay"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for weekNDay")
		}
		_weekNDayErr := writeBuffer.WriteSerializable(ctx, m.GetWeekNDay())
		if popErr := writeBuffer.PopContext("weekNDay"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for weekNDay")
		}
		if _weekNDayErr != nil {
			return errors.Wrap(_weekNDayErr, "Error serializing 'weekNDay' field")
		}

		if popErr := writeBuffer.PopContext("BACnetCalendarEntryWeekNDay"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetCalendarEntryWeekNDay")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetCalendarEntryWeekNDay) isBACnetCalendarEntryWeekNDay() bool {
	return true
}

func (m *_BACnetCalendarEntryWeekNDay) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
