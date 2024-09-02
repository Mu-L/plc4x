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

// BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger is the corresponding interface of BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger
type BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetFaultParameterFaultOutOfRangeMinNormalValue
	// GetIntegerValue returns IntegerValue (property field)
	GetIntegerValue() BACnetApplicationTagSignedInteger
}

// BACnetFaultParameterFaultOutOfRangeMinNormalValueIntegerExactly can be used when we want exactly this type and not a type which fulfills BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger.
// This is useful for switch cases.
type BACnetFaultParameterFaultOutOfRangeMinNormalValueIntegerExactly interface {
	BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger
	isBACnetFaultParameterFaultOutOfRangeMinNormalValueInteger() bool
}

// _BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger is the data-structure of this message
type _BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger struct {
	BACnetFaultParameterFaultOutOfRangeMinNormalValueContract
	IntegerValue BACnetApplicationTagSignedInteger
}

var _ BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger = (*_BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger)(nil)
var _ BACnetFaultParameterFaultOutOfRangeMinNormalValueRequirements = (*_BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger) GetParent() BACnetFaultParameterFaultOutOfRangeMinNormalValueContract {
	return m.BACnetFaultParameterFaultOutOfRangeMinNormalValueContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger) GetIntegerValue() BACnetApplicationTagSignedInteger {
	return m.IntegerValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetFaultParameterFaultOutOfRangeMinNormalValueInteger factory function for _BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger
func NewBACnetFaultParameterFaultOutOfRangeMinNormalValueInteger(integerValue BACnetApplicationTagSignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger {
	_result := &_BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger{
		BACnetFaultParameterFaultOutOfRangeMinNormalValueContract: NewBACnetFaultParameterFaultOutOfRangeMinNormalValue(openingTag, peekedTagHeader, closingTag, tagNumber),
		IntegerValue: integerValue,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetFaultParameterFaultOutOfRangeMinNormalValueInteger(structType any) BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger {
	if casted, ok := structType.(BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger) GetTypeName() string {
	return "BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger"
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetFaultParameterFaultOutOfRangeMinNormalValueContract.(*_BACnetFaultParameterFaultOutOfRangeMinNormalValue).getLengthInBits(ctx))

	// Simple field (integerValue)
	lengthInBits += m.IntegerValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetFaultParameterFaultOutOfRangeMinNormalValue, tagNumber uint8) (__bACnetFaultParameterFaultOutOfRangeMinNormalValueInteger BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger, err error) {
	m.BACnetFaultParameterFaultOutOfRangeMinNormalValueContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	integerValue, err := ReadSimpleField[BACnetApplicationTagSignedInteger](ctx, "integerValue", ReadComplex[BACnetApplicationTagSignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagSignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'integerValue' field"))
	}
	m.IntegerValue = integerValue

	if closeErr := readBuffer.CloseContext("BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger")
	}

	return m, nil
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger")
		}

		if err := WriteSimpleField[BACnetApplicationTagSignedInteger](ctx, "integerValue", m.GetIntegerValue(), WriteComplex[BACnetApplicationTagSignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'integerValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger")
		}
		return nil
	}
	return m.BACnetFaultParameterFaultOutOfRangeMinNormalValueContract.(*_BACnetFaultParameterFaultOutOfRangeMinNormalValue).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger) isBACnetFaultParameterFaultOutOfRangeMinNormalValueInteger() bool {
	return true
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
