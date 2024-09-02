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

// BACnetContextTagEnumerated is the corresponding interface of BACnetContextTagEnumerated
type BACnetContextTagEnumerated interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetContextTag
	// GetPayload returns Payload (property field)
	GetPayload() BACnetTagPayloadEnumerated
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() uint32
}

// BACnetContextTagEnumeratedExactly can be used when we want exactly this type and not a type which fulfills BACnetContextTagEnumerated.
// This is useful for switch cases.
type BACnetContextTagEnumeratedExactly interface {
	BACnetContextTagEnumerated
	isBACnetContextTagEnumerated() bool
}

// _BACnetContextTagEnumerated is the data-structure of this message
type _BACnetContextTagEnumerated struct {
	BACnetContextTagContract
	Payload BACnetTagPayloadEnumerated
}

var _ BACnetContextTagEnumerated = (*_BACnetContextTagEnumerated)(nil)
var _ BACnetContextTagRequirements = (*_BACnetContextTagEnumerated)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetContextTagEnumerated) GetDataType() BACnetDataType {
	return BACnetDataType_ENUMERATED
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetContextTagEnumerated) GetParent() BACnetContextTagContract {
	return m.BACnetContextTagContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetContextTagEnumerated) GetPayload() BACnetTagPayloadEnumerated {
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

func (m *_BACnetContextTagEnumerated) GetActualValue() uint32 {
	ctx := context.Background()
	_ = ctx
	return uint32(m.GetPayload().GetActualValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetContextTagEnumerated factory function for _BACnetContextTagEnumerated
func NewBACnetContextTagEnumerated(payload BACnetTagPayloadEnumerated, header BACnetTagHeader, tagNumberArgument uint8) *_BACnetContextTagEnumerated {
	_result := &_BACnetContextTagEnumerated{
		BACnetContextTagContract: NewBACnetContextTag(header, tagNumberArgument),
		Payload:                  payload,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetContextTagEnumerated(structType any) BACnetContextTagEnumerated {
	if casted, ok := structType.(BACnetContextTagEnumerated); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetContextTagEnumerated); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetContextTagEnumerated) GetTypeName() string {
	return "BACnetContextTagEnumerated"
}

func (m *_BACnetContextTagEnumerated) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetContextTagContract.(*_BACnetContextTag).getLengthInBits(ctx))

	// Simple field (payload)
	lengthInBits += m.Payload.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetContextTagEnumerated) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetContextTagEnumerated) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetContextTag, header BACnetTagHeader, tagNumberArgument uint8, dataType BACnetDataType) (__bACnetContextTagEnumerated BACnetContextTagEnumerated, err error) {
	m.BACnetContextTagContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetContextTagEnumerated"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetContextTagEnumerated")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	payload, err := ReadSimpleField[BACnetTagPayloadEnumerated](ctx, "payload", ReadComplex[BACnetTagPayloadEnumerated](BACnetTagPayloadEnumeratedParseWithBufferProducer((uint32)(header.GetActualLength())), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'payload' field"))
	}
	m.Payload = payload

	actualValue, err := ReadVirtualField[uint32](ctx, "actualValue", (*uint32)(nil), payload.GetActualValue())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetContextTagEnumerated"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetContextTagEnumerated")
	}

	return m, nil
}

func (m *_BACnetContextTagEnumerated) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetContextTagEnumerated) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetContextTagEnumerated"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetContextTagEnumerated")
		}

		if err := WriteSimpleField[BACnetTagPayloadEnumerated](ctx, "payload", m.GetPayload(), WriteComplex[BACnetTagPayloadEnumerated](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'payload' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetContextTagEnumerated"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetContextTagEnumerated")
		}
		return nil
	}
	return m.BACnetContextTagContract.(*_BACnetContextTag).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetContextTagEnumerated) isBACnetContextTagEnumerated() bool {
	return true
}

func (m *_BACnetContextTagEnumerated) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
