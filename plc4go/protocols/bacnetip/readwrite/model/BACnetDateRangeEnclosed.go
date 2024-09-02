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

// BACnetDateRangeEnclosed is the corresponding interface of BACnetDateRangeEnclosed
type BACnetDateRangeEnclosed interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetDateRange returns DateRange (property field)
	GetDateRange() BACnetDateRange
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
	// IsBACnetDateRangeEnclosed is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetDateRangeEnclosed()
}

// _BACnetDateRangeEnclosed is the data-structure of this message
type _BACnetDateRangeEnclosed struct {
	OpeningTag BACnetOpeningTag
	DateRange  BACnetDateRange
	ClosingTag BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

var _ BACnetDateRangeEnclosed = (*_BACnetDateRangeEnclosed)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetDateRangeEnclosed) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetDateRangeEnclosed) GetDateRange() BACnetDateRange {
	return m.DateRange
}

func (m *_BACnetDateRangeEnclosed) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetDateRangeEnclosed factory function for _BACnetDateRangeEnclosed
func NewBACnetDateRangeEnclosed(openingTag BACnetOpeningTag, dateRange BACnetDateRange, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetDateRangeEnclosed {
	return &_BACnetDateRangeEnclosed{OpeningTag: openingTag, DateRange: dateRange, ClosingTag: closingTag, TagNumber: tagNumber}
}

// Deprecated: use the interface for direct cast
func CastBACnetDateRangeEnclosed(structType any) BACnetDateRangeEnclosed {
	if casted, ok := structType.(BACnetDateRangeEnclosed); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetDateRangeEnclosed); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetDateRangeEnclosed) GetTypeName() string {
	return "BACnetDateRangeEnclosed"
}

func (m *_BACnetDateRangeEnclosed) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Simple field (dateRange)
	lengthInBits += m.DateRange.GetLengthInBits(ctx)

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetDateRangeEnclosed) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetDateRangeEnclosedParse(ctx context.Context, theBytes []byte, tagNumber uint8) (BACnetDateRangeEnclosed, error) {
	return BACnetDateRangeEnclosedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetDateRangeEnclosedParseWithBufferProducer(tagNumber uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetDateRangeEnclosed, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetDateRangeEnclosed, error) {
		return BACnetDateRangeEnclosedParseWithBuffer(ctx, readBuffer, tagNumber)
	}
}

func BACnetDateRangeEnclosedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetDateRangeEnclosed, error) {
	v, err := (&_BACnetDateRangeEnclosed{TagNumber: tagNumber}).parse(ctx, readBuffer, tagNumber)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *_BACnetDateRangeEnclosed) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (__bACnetDateRangeEnclosed BACnetDateRangeEnclosed, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetDateRangeEnclosed"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetDateRangeEnclosed")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}
	m.OpeningTag = openingTag

	dateRange, err := ReadSimpleField[BACnetDateRange](ctx, "dateRange", ReadComplex[BACnetDateRange](BACnetDateRangeParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dateRange' field"))
	}
	m.DateRange = dateRange

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	m.ClosingTag = closingTag

	if closeErr := readBuffer.CloseContext("BACnetDateRangeEnclosed"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetDateRangeEnclosed")
	}

	return m, nil
}

func (m *_BACnetDateRangeEnclosed) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetDateRangeEnclosed) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetDateRangeEnclosed"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetDateRangeEnclosed")
	}

	if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'openingTag' field")
	}

	if err := WriteSimpleField[BACnetDateRange](ctx, "dateRange", m.GetDateRange(), WriteComplex[BACnetDateRange](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'dateRange' field")
	}

	if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetDateRangeEnclosed"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetDateRangeEnclosed")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetDateRangeEnclosed) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_BACnetDateRangeEnclosed) IsBACnetDateRangeEnclosed() {}

func (m *_BACnetDateRangeEnclosed) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
