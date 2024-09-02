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

// BACnetOptionalREALNull is the corresponding interface of BACnetOptionalREALNull
type BACnetOptionalREALNull interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetOptionalREAL
	// GetNullValue returns NullValue (property field)
	GetNullValue() BACnetApplicationTagNull
}

// BACnetOptionalREALNullExactly can be used when we want exactly this type and not a type which fulfills BACnetOptionalREALNull.
// This is useful for switch cases.
type BACnetOptionalREALNullExactly interface {
	BACnetOptionalREALNull
	isBACnetOptionalREALNull() bool
}

// _BACnetOptionalREALNull is the data-structure of this message
type _BACnetOptionalREALNull struct {
	BACnetOptionalREALContract
	NullValue BACnetApplicationTagNull
}

var _ BACnetOptionalREALNull = (*_BACnetOptionalREALNull)(nil)
var _ BACnetOptionalREALRequirements = (*_BACnetOptionalREALNull)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetOptionalREALNull) GetParent() BACnetOptionalREALContract {
	return m.BACnetOptionalREALContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetOptionalREALNull) GetNullValue() BACnetApplicationTagNull {
	return m.NullValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetOptionalREALNull factory function for _BACnetOptionalREALNull
func NewBACnetOptionalREALNull(nullValue BACnetApplicationTagNull, peekedTagHeader BACnetTagHeader) *_BACnetOptionalREALNull {
	_result := &_BACnetOptionalREALNull{
		BACnetOptionalREALContract: NewBACnetOptionalREAL(peekedTagHeader),
		NullValue:                  nullValue,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetOptionalREALNull(structType any) BACnetOptionalREALNull {
	if casted, ok := structType.(BACnetOptionalREALNull); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetOptionalREALNull); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetOptionalREALNull) GetTypeName() string {
	return "BACnetOptionalREALNull"
}

func (m *_BACnetOptionalREALNull) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetOptionalREALContract.(*_BACnetOptionalREAL).getLengthInBits(ctx))

	// Simple field (nullValue)
	lengthInBits += m.NullValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetOptionalREALNull) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetOptionalREALNull) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetOptionalREAL) (__bACnetOptionalREALNull BACnetOptionalREALNull, err error) {
	m.BACnetOptionalREALContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetOptionalREALNull"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetOptionalREALNull")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	nullValue, err := ReadSimpleField[BACnetApplicationTagNull](ctx, "nullValue", ReadComplex[BACnetApplicationTagNull](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagNull](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'nullValue' field"))
	}
	m.NullValue = nullValue

	if closeErr := readBuffer.CloseContext("BACnetOptionalREALNull"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetOptionalREALNull")
	}

	return m, nil
}

func (m *_BACnetOptionalREALNull) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetOptionalREALNull) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetOptionalREALNull"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetOptionalREALNull")
		}

		if err := WriteSimpleField[BACnetApplicationTagNull](ctx, "nullValue", m.GetNullValue(), WriteComplex[BACnetApplicationTagNull](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'nullValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetOptionalREALNull"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetOptionalREALNull")
		}
		return nil
	}
	return m.BACnetOptionalREALContract.(*_BACnetOptionalREAL).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetOptionalREALNull) isBACnetOptionalREALNull() bool {
	return true
}

func (m *_BACnetOptionalREALNull) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
