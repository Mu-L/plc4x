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

// BACnetApplicationTagUnsignedInteger is the corresponding interface of BACnetApplicationTagUnsignedInteger
type BACnetApplicationTagUnsignedInteger interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetApplicationTag
	// GetPayload returns Payload (property field)
	GetPayload() BACnetTagPayloadUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() uint64
	// IsBACnetApplicationTagUnsignedInteger is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetApplicationTagUnsignedInteger()
}

// _BACnetApplicationTagUnsignedInteger is the data-structure of this message
type _BACnetApplicationTagUnsignedInteger struct {
	BACnetApplicationTagContract
	Payload BACnetTagPayloadUnsignedInteger
}

var _ BACnetApplicationTagUnsignedInteger = (*_BACnetApplicationTagUnsignedInteger)(nil)
var _ BACnetApplicationTagRequirements = (*_BACnetApplicationTagUnsignedInteger)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetApplicationTagUnsignedInteger) GetParent() BACnetApplicationTagContract {
	return m.BACnetApplicationTagContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetApplicationTagUnsignedInteger) GetPayload() BACnetTagPayloadUnsignedInteger {
	return m.Payload
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetApplicationTagUnsignedInteger) GetActualValue() uint64 {
	ctx := context.Background()
	_ = ctx
	return uint64(m.GetPayload().GetActualValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetApplicationTagUnsignedInteger factory function for _BACnetApplicationTagUnsignedInteger
func NewBACnetApplicationTagUnsignedInteger(payload BACnetTagPayloadUnsignedInteger, header BACnetTagHeader) *_BACnetApplicationTagUnsignedInteger {
	_result := &_BACnetApplicationTagUnsignedInteger{
		BACnetApplicationTagContract: NewBACnetApplicationTag(header),
		Payload:                      payload,
	}
	_result.BACnetApplicationTagContract.(*_BACnetApplicationTag)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetApplicationTagUnsignedInteger(structType any) BACnetApplicationTagUnsignedInteger {
	if casted, ok := structType.(BACnetApplicationTagUnsignedInteger); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetApplicationTagUnsignedInteger); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetApplicationTagUnsignedInteger) GetTypeName() string {
	return "BACnetApplicationTagUnsignedInteger"
}

func (m *_BACnetApplicationTagUnsignedInteger) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetApplicationTagContract.(*_BACnetApplicationTag).getLengthInBits(ctx))

	// Simple field (payload)
	lengthInBits += m.Payload.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetApplicationTagUnsignedInteger) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetApplicationTagUnsignedInteger) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetApplicationTag, header BACnetTagHeader) (__bACnetApplicationTagUnsignedInteger BACnetApplicationTagUnsignedInteger, err error) {
	m.BACnetApplicationTagContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetApplicationTagUnsignedInteger"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetApplicationTagUnsignedInteger")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	payload, err := ReadSimpleField[BACnetTagPayloadUnsignedInteger](ctx, "payload", ReadComplex[BACnetTagPayloadUnsignedInteger](BACnetTagPayloadUnsignedIntegerParseWithBufferProducer((uint32)(header.GetActualLength())), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'payload' field"))
	}
	m.Payload = payload

	actualValue, err := ReadVirtualField[uint64](ctx, "actualValue", (*uint64)(nil), payload.GetActualValue())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetApplicationTagUnsignedInteger"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetApplicationTagUnsignedInteger")
	}

	return m, nil
}

func (m *_BACnetApplicationTagUnsignedInteger) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetApplicationTagUnsignedInteger) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetApplicationTagUnsignedInteger"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetApplicationTagUnsignedInteger")
		}

		if err := WriteSimpleField[BACnetTagPayloadUnsignedInteger](ctx, "payload", m.GetPayload(), WriteComplex[BACnetTagPayloadUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'payload' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetApplicationTagUnsignedInteger"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetApplicationTagUnsignedInteger")
		}
		return nil
	}
	return m.BACnetApplicationTagContract.(*_BACnetApplicationTag).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetApplicationTagUnsignedInteger) IsBACnetApplicationTagUnsignedInteger() {}

func (m *_BACnetApplicationTagUnsignedInteger) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
