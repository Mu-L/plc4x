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

// BACnetEventTimestamps is the corresponding interface of BACnetEventTimestamps
type BACnetEventTimestamps interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetToOffnormal returns ToOffnormal (property field)
	GetToOffnormal() BACnetTimeStamp
	// GetToFault returns ToFault (property field)
	GetToFault() BACnetTimeStamp
	// GetToNormal returns ToNormal (property field)
	GetToNormal() BACnetTimeStamp
	// IsBACnetEventTimestamps is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetEventTimestamps()
}

// _BACnetEventTimestamps is the data-structure of this message
type _BACnetEventTimestamps struct {
	ToOffnormal BACnetTimeStamp
	ToFault     BACnetTimeStamp
	ToNormal    BACnetTimeStamp
}

var _ BACnetEventTimestamps = (*_BACnetEventTimestamps)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetEventTimestamps) GetToOffnormal() BACnetTimeStamp {
	return m.ToOffnormal
}

func (m *_BACnetEventTimestamps) GetToFault() BACnetTimeStamp {
	return m.ToFault
}

func (m *_BACnetEventTimestamps) GetToNormal() BACnetTimeStamp {
	return m.ToNormal
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetEventTimestamps factory function for _BACnetEventTimestamps
func NewBACnetEventTimestamps(toOffnormal BACnetTimeStamp, toFault BACnetTimeStamp, toNormal BACnetTimeStamp) *_BACnetEventTimestamps {
	return &_BACnetEventTimestamps{ToOffnormal: toOffnormal, ToFault: toFault, ToNormal: toNormal}
}

// Deprecated: use the interface for direct cast
func CastBACnetEventTimestamps(structType any) BACnetEventTimestamps {
	if casted, ok := structType.(BACnetEventTimestamps); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetEventTimestamps); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetEventTimestamps) GetTypeName() string {
	return "BACnetEventTimestamps"
}

func (m *_BACnetEventTimestamps) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (toOffnormal)
	lengthInBits += m.ToOffnormal.GetLengthInBits(ctx)

	// Simple field (toFault)
	lengthInBits += m.ToFault.GetLengthInBits(ctx)

	// Simple field (toNormal)
	lengthInBits += m.ToNormal.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetEventTimestamps) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetEventTimestampsParse(ctx context.Context, theBytes []byte) (BACnetEventTimestamps, error) {
	return BACnetEventTimestampsParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetEventTimestampsParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetEventTimestamps, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetEventTimestamps, error) {
		return BACnetEventTimestampsParseWithBuffer(ctx, readBuffer)
	}
}

func BACnetEventTimestampsParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetEventTimestamps, error) {
	v, err := (&_BACnetEventTimestamps{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *_BACnetEventTimestamps) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__bACnetEventTimestamps BACnetEventTimestamps, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetEventTimestamps"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetEventTimestamps")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	toOffnormal, err := ReadSimpleField[BACnetTimeStamp](ctx, "toOffnormal", ReadComplex[BACnetTimeStamp](BACnetTimeStampParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'toOffnormal' field"))
	}
	m.ToOffnormal = toOffnormal

	toFault, err := ReadSimpleField[BACnetTimeStamp](ctx, "toFault", ReadComplex[BACnetTimeStamp](BACnetTimeStampParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'toFault' field"))
	}
	m.ToFault = toFault

	toNormal, err := ReadSimpleField[BACnetTimeStamp](ctx, "toNormal", ReadComplex[BACnetTimeStamp](BACnetTimeStampParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'toNormal' field"))
	}
	m.ToNormal = toNormal

	if closeErr := readBuffer.CloseContext("BACnetEventTimestamps"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetEventTimestamps")
	}

	return m, nil
}

func (m *_BACnetEventTimestamps) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetEventTimestamps) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetEventTimestamps"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetEventTimestamps")
	}

	if err := WriteSimpleField[BACnetTimeStamp](ctx, "toOffnormal", m.GetToOffnormal(), WriteComplex[BACnetTimeStamp](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'toOffnormal' field")
	}

	if err := WriteSimpleField[BACnetTimeStamp](ctx, "toFault", m.GetToFault(), WriteComplex[BACnetTimeStamp](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'toFault' field")
	}

	if err := WriteSimpleField[BACnetTimeStamp](ctx, "toNormal", m.GetToNormal(), WriteComplex[BACnetTimeStamp](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'toNormal' field")
	}

	if popErr := writeBuffer.PopContext("BACnetEventTimestamps"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetEventTimestamps")
	}
	return nil
}

func (m *_BACnetEventTimestamps) IsBACnetEventTimestamps() {}

func (m *_BACnetEventTimestamps) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
