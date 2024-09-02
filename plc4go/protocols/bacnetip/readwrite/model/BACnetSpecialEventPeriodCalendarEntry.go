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

// BACnetSpecialEventPeriodCalendarEntry is the corresponding interface of BACnetSpecialEventPeriodCalendarEntry
type BACnetSpecialEventPeriodCalendarEntry interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetSpecialEventPeriod
	// GetCalendarEntry returns CalendarEntry (property field)
	GetCalendarEntry() BACnetCalendarEntryEnclosed
	// IsBACnetSpecialEventPeriodCalendarEntry is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetSpecialEventPeriodCalendarEntry()
}

// _BACnetSpecialEventPeriodCalendarEntry is the data-structure of this message
type _BACnetSpecialEventPeriodCalendarEntry struct {
	BACnetSpecialEventPeriodContract
	CalendarEntry BACnetCalendarEntryEnclosed
}

var _ BACnetSpecialEventPeriodCalendarEntry = (*_BACnetSpecialEventPeriodCalendarEntry)(nil)
var _ BACnetSpecialEventPeriodRequirements = (*_BACnetSpecialEventPeriodCalendarEntry)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetSpecialEventPeriodCalendarEntry) GetParent() BACnetSpecialEventPeriodContract {
	return m.BACnetSpecialEventPeriodContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetSpecialEventPeriodCalendarEntry) GetCalendarEntry() BACnetCalendarEntryEnclosed {
	return m.CalendarEntry
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetSpecialEventPeriodCalendarEntry factory function for _BACnetSpecialEventPeriodCalendarEntry
func NewBACnetSpecialEventPeriodCalendarEntry(calendarEntry BACnetCalendarEntryEnclosed, peekedTagHeader BACnetTagHeader) *_BACnetSpecialEventPeriodCalendarEntry {
	_result := &_BACnetSpecialEventPeriodCalendarEntry{
		BACnetSpecialEventPeriodContract: NewBACnetSpecialEventPeriod(peekedTagHeader),
		CalendarEntry:                    calendarEntry,
	}
	_result.BACnetSpecialEventPeriodContract.(*_BACnetSpecialEventPeriod)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetSpecialEventPeriodCalendarEntry(structType any) BACnetSpecialEventPeriodCalendarEntry {
	if casted, ok := structType.(BACnetSpecialEventPeriodCalendarEntry); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetSpecialEventPeriodCalendarEntry); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetSpecialEventPeriodCalendarEntry) GetTypeName() string {
	return "BACnetSpecialEventPeriodCalendarEntry"
}

func (m *_BACnetSpecialEventPeriodCalendarEntry) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetSpecialEventPeriodContract.(*_BACnetSpecialEventPeriod).getLengthInBits(ctx))

	// Simple field (calendarEntry)
	lengthInBits += m.CalendarEntry.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetSpecialEventPeriodCalendarEntry) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetSpecialEventPeriodCalendarEntry) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetSpecialEventPeriod) (__bACnetSpecialEventPeriodCalendarEntry BACnetSpecialEventPeriodCalendarEntry, err error) {
	m.BACnetSpecialEventPeriodContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetSpecialEventPeriodCalendarEntry"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetSpecialEventPeriodCalendarEntry")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	calendarEntry, err := ReadSimpleField[BACnetCalendarEntryEnclosed](ctx, "calendarEntry", ReadComplex[BACnetCalendarEntryEnclosed](BACnetCalendarEntryEnclosedParseWithBufferProducer((uint8)(uint8(0))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'calendarEntry' field"))
	}
	m.CalendarEntry = calendarEntry

	if closeErr := readBuffer.CloseContext("BACnetSpecialEventPeriodCalendarEntry"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetSpecialEventPeriodCalendarEntry")
	}

	return m, nil
}

func (m *_BACnetSpecialEventPeriodCalendarEntry) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetSpecialEventPeriodCalendarEntry) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetSpecialEventPeriodCalendarEntry"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetSpecialEventPeriodCalendarEntry")
		}

		if err := WriteSimpleField[BACnetCalendarEntryEnclosed](ctx, "calendarEntry", m.GetCalendarEntry(), WriteComplex[BACnetCalendarEntryEnclosed](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'calendarEntry' field")
		}

		if popErr := writeBuffer.PopContext("BACnetSpecialEventPeriodCalendarEntry"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetSpecialEventPeriodCalendarEntry")
		}
		return nil
	}
	return m.BACnetSpecialEventPeriodContract.(*_BACnetSpecialEventPeriod).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetSpecialEventPeriodCalendarEntry) IsBACnetSpecialEventPeriodCalendarEntry() {}

func (m *_BACnetSpecialEventPeriodCalendarEntry) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
