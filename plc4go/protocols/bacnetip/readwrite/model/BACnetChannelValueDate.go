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

// BACnetChannelValueDate is the corresponding interface of BACnetChannelValueDate
type BACnetChannelValueDate interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetChannelValue
	// GetDateValue returns DateValue (property field)
	GetDateValue() BACnetApplicationTagDate
}

// BACnetChannelValueDateExactly can be used when we want exactly this type and not a type which fulfills BACnetChannelValueDate.
// This is useful for switch cases.
type BACnetChannelValueDateExactly interface {
	BACnetChannelValueDate
	isBACnetChannelValueDate() bool
}

// _BACnetChannelValueDate is the data-structure of this message
type _BACnetChannelValueDate struct {
	BACnetChannelValueContract
	DateValue BACnetApplicationTagDate
}

var _ BACnetChannelValueDate = (*_BACnetChannelValueDate)(nil)
var _ BACnetChannelValueRequirements = (*_BACnetChannelValueDate)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetChannelValueDate) GetParent() BACnetChannelValueContract {
	return m.BACnetChannelValueContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetChannelValueDate) GetDateValue() BACnetApplicationTagDate {
	return m.DateValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetChannelValueDate factory function for _BACnetChannelValueDate
func NewBACnetChannelValueDate(dateValue BACnetApplicationTagDate, peekedTagHeader BACnetTagHeader) *_BACnetChannelValueDate {
	_result := &_BACnetChannelValueDate{
		BACnetChannelValueContract: NewBACnetChannelValue(peekedTagHeader),
		DateValue:                  dateValue,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetChannelValueDate(structType any) BACnetChannelValueDate {
	if casted, ok := structType.(BACnetChannelValueDate); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetChannelValueDate); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetChannelValueDate) GetTypeName() string {
	return "BACnetChannelValueDate"
}

func (m *_BACnetChannelValueDate) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetChannelValueContract.(*_BACnetChannelValue).getLengthInBits(ctx))

	// Simple field (dateValue)
	lengthInBits += m.DateValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetChannelValueDate) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetChannelValueDate) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetChannelValue) (__bACnetChannelValueDate BACnetChannelValueDate, err error) {
	m.BACnetChannelValueContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetChannelValueDate"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetChannelValueDate")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	dateValue, err := ReadSimpleField[BACnetApplicationTagDate](ctx, "dateValue", ReadComplex[BACnetApplicationTagDate](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagDate](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dateValue' field"))
	}
	m.DateValue = dateValue

	if closeErr := readBuffer.CloseContext("BACnetChannelValueDate"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetChannelValueDate")
	}

	return m, nil
}

func (m *_BACnetChannelValueDate) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetChannelValueDate) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetChannelValueDate"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetChannelValueDate")
		}

		if err := WriteSimpleField[BACnetApplicationTagDate](ctx, "dateValue", m.GetDateValue(), WriteComplex[BACnetApplicationTagDate](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'dateValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetChannelValueDate"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetChannelValueDate")
		}
		return nil
	}
	return m.BACnetChannelValueContract.(*_BACnetChannelValue).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetChannelValueDate) isBACnetChannelValueDate() bool {
	return true
}

func (m *_BACnetChannelValueDate) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
