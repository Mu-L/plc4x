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

// BACnetTimeStampSequence is the corresponding interface of BACnetTimeStampSequence
type BACnetTimeStampSequence interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetTimeStamp
	// GetSequenceNumber returns SequenceNumber (property field)
	GetSequenceNumber() BACnetContextTagUnsignedInteger
	// IsBACnetTimeStampSequence is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetTimeStampSequence()
}

// _BACnetTimeStampSequence is the data-structure of this message
type _BACnetTimeStampSequence struct {
	BACnetTimeStampContract
	SequenceNumber BACnetContextTagUnsignedInteger
}

var _ BACnetTimeStampSequence = (*_BACnetTimeStampSequence)(nil)
var _ BACnetTimeStampRequirements = (*_BACnetTimeStampSequence)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetTimeStampSequence) GetParent() BACnetTimeStampContract {
	return m.BACnetTimeStampContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetTimeStampSequence) GetSequenceNumber() BACnetContextTagUnsignedInteger {
	return m.SequenceNumber
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetTimeStampSequence factory function for _BACnetTimeStampSequence
func NewBACnetTimeStampSequence(sequenceNumber BACnetContextTagUnsignedInteger, peekedTagHeader BACnetTagHeader) *_BACnetTimeStampSequence {
	_result := &_BACnetTimeStampSequence{
		BACnetTimeStampContract: NewBACnetTimeStamp(peekedTagHeader),
		SequenceNumber:          sequenceNumber,
	}
	_result.BACnetTimeStampContract.(*_BACnetTimeStamp)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetTimeStampSequence(structType any) BACnetTimeStampSequence {
	if casted, ok := structType.(BACnetTimeStampSequence); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetTimeStampSequence); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetTimeStampSequence) GetTypeName() string {
	return "BACnetTimeStampSequence"
}

func (m *_BACnetTimeStampSequence) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetTimeStampContract.(*_BACnetTimeStamp).getLengthInBits(ctx))

	// Simple field (sequenceNumber)
	lengthInBits += m.SequenceNumber.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetTimeStampSequence) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetTimeStampSequence) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetTimeStamp) (__bACnetTimeStampSequence BACnetTimeStampSequence, err error) {
	m.BACnetTimeStampContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetTimeStampSequence"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetTimeStampSequence")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	sequenceNumber, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "sequenceNumber", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(1)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'sequenceNumber' field"))
	}
	m.SequenceNumber = sequenceNumber

	if closeErr := readBuffer.CloseContext("BACnetTimeStampSequence"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetTimeStampSequence")
	}

	return m, nil
}

func (m *_BACnetTimeStampSequence) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetTimeStampSequence) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetTimeStampSequence"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetTimeStampSequence")
		}

		if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "sequenceNumber", m.GetSequenceNumber(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'sequenceNumber' field")
		}

		if popErr := writeBuffer.PopContext("BACnetTimeStampSequence"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetTimeStampSequence")
		}
		return nil
	}
	return m.BACnetTimeStampContract.(*_BACnetTimeStamp).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetTimeStampSequence) IsBACnetTimeStampSequence() {}

func (m *_BACnetTimeStampSequence) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
