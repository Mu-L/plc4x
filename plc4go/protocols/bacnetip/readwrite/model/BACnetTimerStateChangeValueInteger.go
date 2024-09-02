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

// BACnetTimerStateChangeValueInteger is the corresponding interface of BACnetTimerStateChangeValueInteger
type BACnetTimerStateChangeValueInteger interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetTimerStateChangeValue
	// GetIntegerValue returns IntegerValue (property field)
	GetIntegerValue() BACnetApplicationTagSignedInteger
	// IsBACnetTimerStateChangeValueInteger is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetTimerStateChangeValueInteger()
}

// _BACnetTimerStateChangeValueInteger is the data-structure of this message
type _BACnetTimerStateChangeValueInteger struct {
	BACnetTimerStateChangeValueContract
	IntegerValue BACnetApplicationTagSignedInteger
}

var _ BACnetTimerStateChangeValueInteger = (*_BACnetTimerStateChangeValueInteger)(nil)
var _ BACnetTimerStateChangeValueRequirements = (*_BACnetTimerStateChangeValueInteger)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetTimerStateChangeValueInteger) GetParent() BACnetTimerStateChangeValueContract {
	return m.BACnetTimerStateChangeValueContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetTimerStateChangeValueInteger) GetIntegerValue() BACnetApplicationTagSignedInteger {
	return m.IntegerValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetTimerStateChangeValueInteger factory function for _BACnetTimerStateChangeValueInteger
func NewBACnetTimerStateChangeValueInteger(integerValue BACnetApplicationTagSignedInteger, peekedTagHeader BACnetTagHeader, objectTypeArgument BACnetObjectType) *_BACnetTimerStateChangeValueInteger {
	_result := &_BACnetTimerStateChangeValueInteger{
		BACnetTimerStateChangeValueContract: NewBACnetTimerStateChangeValue(peekedTagHeader, objectTypeArgument),
		IntegerValue:                        integerValue,
	}
	_result.BACnetTimerStateChangeValueContract.(*_BACnetTimerStateChangeValue)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetTimerStateChangeValueInteger(structType any) BACnetTimerStateChangeValueInteger {
	if casted, ok := structType.(BACnetTimerStateChangeValueInteger); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetTimerStateChangeValueInteger); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetTimerStateChangeValueInteger) GetTypeName() string {
	return "BACnetTimerStateChangeValueInteger"
}

func (m *_BACnetTimerStateChangeValueInteger) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetTimerStateChangeValueContract.(*_BACnetTimerStateChangeValue).getLengthInBits(ctx))

	// Simple field (integerValue)
	lengthInBits += m.IntegerValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetTimerStateChangeValueInteger) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetTimerStateChangeValueInteger) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetTimerStateChangeValue, objectTypeArgument BACnetObjectType) (__bACnetTimerStateChangeValueInteger BACnetTimerStateChangeValueInteger, err error) {
	m.BACnetTimerStateChangeValueContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetTimerStateChangeValueInteger"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetTimerStateChangeValueInteger")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	integerValue, err := ReadSimpleField[BACnetApplicationTagSignedInteger](ctx, "integerValue", ReadComplex[BACnetApplicationTagSignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagSignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'integerValue' field"))
	}
	m.IntegerValue = integerValue

	if closeErr := readBuffer.CloseContext("BACnetTimerStateChangeValueInteger"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetTimerStateChangeValueInteger")
	}

	return m, nil
}

func (m *_BACnetTimerStateChangeValueInteger) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetTimerStateChangeValueInteger) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetTimerStateChangeValueInteger"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetTimerStateChangeValueInteger")
		}

		if err := WriteSimpleField[BACnetApplicationTagSignedInteger](ctx, "integerValue", m.GetIntegerValue(), WriteComplex[BACnetApplicationTagSignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'integerValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetTimerStateChangeValueInteger"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetTimerStateChangeValueInteger")
		}
		return nil
	}
	return m.BACnetTimerStateChangeValueContract.(*_BACnetTimerStateChangeValue).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetTimerStateChangeValueInteger) IsBACnetTimerStateChangeValueInteger() {}

func (m *_BACnetTimerStateChangeValueInteger) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
