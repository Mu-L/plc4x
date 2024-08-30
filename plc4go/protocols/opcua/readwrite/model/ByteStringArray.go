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

// ByteStringArray is the corresponding interface of ByteStringArray
type ByteStringArray interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetArrayLength returns ArrayLength (property field)
	GetArrayLength() int32
	// GetValue returns Value (property field)
	GetValue() []uint8
}

// ByteStringArrayExactly can be used when we want exactly this type and not a type which fulfills ByteStringArray.
// This is useful for switch cases.
type ByteStringArrayExactly interface {
	ByteStringArray
	isByteStringArray() bool
}

// _ByteStringArray is the data-structure of this message
type _ByteStringArray struct {
	ArrayLength int32
	Value       []uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ByteStringArray) GetArrayLength() int32 {
	return m.ArrayLength
}

func (m *_ByteStringArray) GetValue() []uint8 {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewByteStringArray factory function for _ByteStringArray
func NewByteStringArray(arrayLength int32, value []uint8) *_ByteStringArray {
	return &_ByteStringArray{ArrayLength: arrayLength, Value: value}
}

// Deprecated: use the interface for direct cast
func CastByteStringArray(structType any) ByteStringArray {
	if casted, ok := structType.(ByteStringArray); ok {
		return casted
	}
	if casted, ok := structType.(*ByteStringArray); ok {
		return *casted
	}
	return nil
}

func (m *_ByteStringArray) GetTypeName() string {
	return "ByteStringArray"
}

func (m *_ByteStringArray) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (arrayLength)
	lengthInBits += 32

	// Array field
	if len(m.Value) > 0 {
		lengthInBits += 8 * uint16(len(m.Value))
	}

	return lengthInBits
}

func (m *_ByteStringArray) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ByteStringArrayParse(ctx context.Context, theBytes []byte) (ByteStringArray, error) {
	return ByteStringArrayParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func ByteStringArrayParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (ByteStringArray, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("ByteStringArray"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ByteStringArray")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (arrayLength)
	_arrayLength, _arrayLengthErr := /*TODO: migrate me*/ readBuffer.ReadInt32("arrayLength", 32)
	if _arrayLengthErr != nil {
		return nil, errors.Wrap(_arrayLengthErr, "Error parsing 'arrayLength' field of ByteStringArray")
	}
	arrayLength := _arrayLength

	value, err := ReadCountArrayField[uint8](ctx, "value", ReadUnsignedByte(readBuffer, 8), uint64(arrayLength))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}

	if closeErr := readBuffer.CloseContext("ByteStringArray"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ByteStringArray")
	}

	// Create the instance
	return &_ByteStringArray{
		ArrayLength: arrayLength,
		Value:       value,
	}, nil
}

func (m *_ByteStringArray) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ByteStringArray) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("ByteStringArray"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ByteStringArray")
	}

	// Simple Field (arrayLength)
	arrayLength := int32(m.GetArrayLength())
	_arrayLengthErr := /*TODO: migrate me*/ writeBuffer.WriteInt32("arrayLength", 32, int32((arrayLength)))
	if _arrayLengthErr != nil {
		return errors.Wrap(_arrayLengthErr, "Error serializing 'arrayLength' field")
	}

	// Array Field (value)
	if pushErr := writeBuffer.PushContext("value", utils.WithRenderAsList(true)); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for value")
	}
	for _curItem, _element := range m.GetValue() {
		_ = _curItem
		_elementErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("", 8, uint8(_element))
		if _elementErr != nil {
			return errors.Wrap(_elementErr, "Error serializing 'value' field")
		}
	}
	if popErr := writeBuffer.PopContext("value", utils.WithRenderAsList(true)); popErr != nil {
		return errors.Wrap(popErr, "Error popping for value")
	}

	if popErr := writeBuffer.PopContext("ByteStringArray"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ByteStringArray")
	}
	return nil
}

func (m *_ByteStringArray) isByteStringArray() bool {
	return true
}

func (m *_ByteStringArray) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
