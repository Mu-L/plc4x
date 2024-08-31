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

// VariantInt32 is the corresponding interface of VariantInt32
type VariantInt32 interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	Variant
	// GetArrayLength returns ArrayLength (property field)
	GetArrayLength() *int32
	// GetValue returns Value (property field)
	GetValue() []int32
}

// VariantInt32Exactly can be used when we want exactly this type and not a type which fulfills VariantInt32.
// This is useful for switch cases.
type VariantInt32Exactly interface {
	VariantInt32
	isVariantInt32() bool
}

// _VariantInt32 is the data-structure of this message
type _VariantInt32 struct {
	*_Variant
	ArrayLength *int32
	Value       []int32
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_VariantInt32) GetVariantType() uint8 {
	return uint8(6)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_VariantInt32) InitializeParent(parent Variant, arrayLengthSpecified bool, arrayDimensionsSpecified bool, noOfArrayDimensions *int32, arrayDimensions []bool) {
	m.ArrayLengthSpecified = arrayLengthSpecified
	m.ArrayDimensionsSpecified = arrayDimensionsSpecified
	m.NoOfArrayDimensions = noOfArrayDimensions
	m.ArrayDimensions = arrayDimensions
}

func (m *_VariantInt32) GetParent() Variant {
	return m._Variant
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_VariantInt32) GetArrayLength() *int32 {
	return m.ArrayLength
}

func (m *_VariantInt32) GetValue() []int32 {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewVariantInt32 factory function for _VariantInt32
func NewVariantInt32(arrayLength *int32, value []int32, arrayLengthSpecified bool, arrayDimensionsSpecified bool, noOfArrayDimensions *int32, arrayDimensions []bool) *_VariantInt32 {
	_result := &_VariantInt32{
		ArrayLength: arrayLength,
		Value:       value,
		_Variant:    NewVariant(arrayLengthSpecified, arrayDimensionsSpecified, noOfArrayDimensions, arrayDimensions),
	}
	_result._Variant._VariantChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastVariantInt32(structType any) VariantInt32 {
	if casted, ok := structType.(VariantInt32); ok {
		return casted
	}
	if casted, ok := structType.(*VariantInt32); ok {
		return *casted
	}
	return nil
}

func (m *_VariantInt32) GetTypeName() string {
	return "VariantInt32"
}

func (m *_VariantInt32) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Optional Field (arrayLength)
	if m.ArrayLength != nil {
		lengthInBits += 32
	}

	// Array field
	if len(m.Value) > 0 {
		lengthInBits += 32 * uint16(len(m.Value))
	}

	return lengthInBits
}

func (m *_VariantInt32) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func VariantInt32Parse(ctx context.Context, theBytes []byte, arrayLengthSpecified bool) (VariantInt32, error) {
	return VariantInt32ParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), arrayLengthSpecified)
}

func VariantInt32ParseWithBufferProducer(arrayLengthSpecified bool) func(ctx context.Context, readBuffer utils.ReadBuffer) (VariantInt32, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (VariantInt32, error) {
		return VariantInt32ParseWithBuffer(ctx, readBuffer, arrayLengthSpecified)
	}
}

func VariantInt32ParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, arrayLengthSpecified bool) (VariantInt32, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("VariantInt32"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for VariantInt32")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	arrayLength, err := ReadOptionalField[int32](ctx, "arrayLength", ReadSignedInt(readBuffer, uint8(32)), arrayLengthSpecified)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'arrayLength' field"))
	}

	value, err := ReadCountArrayField[int32](ctx, "value", ReadSignedInt(readBuffer, uint8(32)), uint64(utils.InlineIf(bool((arrayLength) == (nil)), func() any { return int32(int32(1)) }, func() any { return int32((*arrayLength)) }).(int32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}

	if closeErr := readBuffer.CloseContext("VariantInt32"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for VariantInt32")
	}

	// Create a partially initialized instance
	_child := &_VariantInt32{
		_Variant:    &_Variant{},
		ArrayLength: arrayLength,
		Value:       value,
	}
	_child._Variant._VariantChildRequirements = _child
	return _child, nil
}

func (m *_VariantInt32) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_VariantInt32) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("VariantInt32"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for VariantInt32")
		}

		// Optional Field (arrayLength) (Can be skipped, if the value is null)
		var arrayLength *int32 = nil
		if m.GetArrayLength() != nil {
			arrayLength = m.GetArrayLength()
			_arrayLengthErr := /*TODO: migrate me*/ writeBuffer.WriteInt32("arrayLength", 32, int32(*(arrayLength)))
			if _arrayLengthErr != nil {
				return errors.Wrap(_arrayLengthErr, "Error serializing 'arrayLength' field")
			}
		}

		// Array Field (value)
		if pushErr := writeBuffer.PushContext("value", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for value")
		}
		for _curItem, _element := range m.GetValue() {
			_ = _curItem
			_elementErr := /*TODO: migrate me*/ writeBuffer.WriteInt32("", 32, int32(_element))
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'value' field")
			}
		}
		if popErr := writeBuffer.PopContext("value", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for value")
		}

		if popErr := writeBuffer.PopContext("VariantInt32"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for VariantInt32")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_VariantInt32) isVariantInt32() bool {
	return true
}

func (m *_VariantInt32) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
