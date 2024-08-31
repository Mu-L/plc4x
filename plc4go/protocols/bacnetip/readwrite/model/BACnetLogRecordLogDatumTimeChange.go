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

// BACnetLogRecordLogDatumTimeChange is the corresponding interface of BACnetLogRecordLogDatumTimeChange
type BACnetLogRecordLogDatumTimeChange interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetLogRecordLogDatum
	// GetTimeChange returns TimeChange (property field)
	GetTimeChange() BACnetContextTagReal
}

// BACnetLogRecordLogDatumTimeChangeExactly can be used when we want exactly this type and not a type which fulfills BACnetLogRecordLogDatumTimeChange.
// This is useful for switch cases.
type BACnetLogRecordLogDatumTimeChangeExactly interface {
	BACnetLogRecordLogDatumTimeChange
	isBACnetLogRecordLogDatumTimeChange() bool
}

// _BACnetLogRecordLogDatumTimeChange is the data-structure of this message
type _BACnetLogRecordLogDatumTimeChange struct {
	*_BACnetLogRecordLogDatum
	TimeChange BACnetContextTagReal
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetLogRecordLogDatumTimeChange) InitializeParent(parent BACnetLogRecordLogDatum, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetLogRecordLogDatumTimeChange) GetParent() BACnetLogRecordLogDatum {
	return m._BACnetLogRecordLogDatum
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetLogRecordLogDatumTimeChange) GetTimeChange() BACnetContextTagReal {
	return m.TimeChange
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetLogRecordLogDatumTimeChange factory function for _BACnetLogRecordLogDatumTimeChange
func NewBACnetLogRecordLogDatumTimeChange(timeChange BACnetContextTagReal, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetLogRecordLogDatumTimeChange {
	_result := &_BACnetLogRecordLogDatumTimeChange{
		TimeChange:               timeChange,
		_BACnetLogRecordLogDatum: NewBACnetLogRecordLogDatum(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result._BACnetLogRecordLogDatum._BACnetLogRecordLogDatumChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetLogRecordLogDatumTimeChange(structType any) BACnetLogRecordLogDatumTimeChange {
	if casted, ok := structType.(BACnetLogRecordLogDatumTimeChange); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetLogRecordLogDatumTimeChange); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetLogRecordLogDatumTimeChange) GetTypeName() string {
	return "BACnetLogRecordLogDatumTimeChange"
}

func (m *_BACnetLogRecordLogDatumTimeChange) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (timeChange)
	lengthInBits += m.TimeChange.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetLogRecordLogDatumTimeChange) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetLogRecordLogDatumTimeChangeParse(ctx context.Context, theBytes []byte, tagNumber uint8) (BACnetLogRecordLogDatumTimeChange, error) {
	return BACnetLogRecordLogDatumTimeChangeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetLogRecordLogDatumTimeChangeParseWithBufferProducer(tagNumber uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetLogRecordLogDatumTimeChange, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetLogRecordLogDatumTimeChange, error) {
		return BACnetLogRecordLogDatumTimeChangeParseWithBuffer(ctx, readBuffer, tagNumber)
	}
}

func BACnetLogRecordLogDatumTimeChangeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetLogRecordLogDatumTimeChange, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetLogRecordLogDatumTimeChange"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetLogRecordLogDatumTimeChange")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	timeChange, err := ReadSimpleField[BACnetContextTagReal](ctx, "timeChange", ReadComplex[BACnetContextTagReal](BACnetContextTagParseWithBufferProducer[BACnetContextTagReal]((uint8)(uint8(9)), (BACnetDataType)(BACnetDataType_REAL)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timeChange' field"))
	}

	if closeErr := readBuffer.CloseContext("BACnetLogRecordLogDatumTimeChange"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetLogRecordLogDatumTimeChange")
	}

	// Create a partially initialized instance
	_child := &_BACnetLogRecordLogDatumTimeChange{
		_BACnetLogRecordLogDatum: &_BACnetLogRecordLogDatum{
			TagNumber: tagNumber,
		},
		TimeChange: timeChange,
	}
	_child._BACnetLogRecordLogDatum._BACnetLogRecordLogDatumChildRequirements = _child
	return _child, nil
}

func (m *_BACnetLogRecordLogDatumTimeChange) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetLogRecordLogDatumTimeChange) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetLogRecordLogDatumTimeChange"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetLogRecordLogDatumTimeChange")
		}

		// Simple Field (timeChange)
		if pushErr := writeBuffer.PushContext("timeChange"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for timeChange")
		}
		_timeChangeErr := writeBuffer.WriteSerializable(ctx, m.GetTimeChange())
		if popErr := writeBuffer.PopContext("timeChange"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for timeChange")
		}
		if _timeChangeErr != nil {
			return errors.Wrap(_timeChangeErr, "Error serializing 'timeChange' field")
		}

		if popErr := writeBuffer.PopContext("BACnetLogRecordLogDatumTimeChange"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetLogRecordLogDatumTimeChange")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetLogRecordLogDatumTimeChange) isBACnetLogRecordLogDatumTimeChange() bool {
	return true
}

func (m *_BACnetLogRecordLogDatumTimeChange) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
