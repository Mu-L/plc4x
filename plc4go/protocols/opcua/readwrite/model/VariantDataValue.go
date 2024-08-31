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

// VariantDataValue is the corresponding interface of VariantDataValue
type VariantDataValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	Variant
	// GetArrayLength returns ArrayLength (property field)
	GetArrayLength() *int32
	// GetValue returns Value (property field)
	GetValue() []DataValue
}

// VariantDataValueExactly can be used when we want exactly this type and not a type which fulfills VariantDataValue.
// This is useful for switch cases.
type VariantDataValueExactly interface {
	VariantDataValue
	isVariantDataValue() bool
}

// _VariantDataValue is the data-structure of this message
type _VariantDataValue struct {
	*_Variant
	ArrayLength *int32
	Value       []DataValue
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_VariantDataValue) GetVariantType() uint8 {
	return uint8(23)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_VariantDataValue) InitializeParent(parent Variant, arrayLengthSpecified bool, arrayDimensionsSpecified bool, noOfArrayDimensions *int32, arrayDimensions []bool) {
	m.ArrayLengthSpecified = arrayLengthSpecified
	m.ArrayDimensionsSpecified = arrayDimensionsSpecified
	m.NoOfArrayDimensions = noOfArrayDimensions
	m.ArrayDimensions = arrayDimensions
}

func (m *_VariantDataValue) GetParent() Variant {
	return m._Variant
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_VariantDataValue) GetArrayLength() *int32 {
	return m.ArrayLength
}

func (m *_VariantDataValue) GetValue() []DataValue {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewVariantDataValue factory function for _VariantDataValue
func NewVariantDataValue(arrayLength *int32, value []DataValue, arrayLengthSpecified bool, arrayDimensionsSpecified bool, noOfArrayDimensions *int32, arrayDimensions []bool) *_VariantDataValue {
	_result := &_VariantDataValue{
		ArrayLength: arrayLength,
		Value:       value,
		_Variant:    NewVariant(arrayLengthSpecified, arrayDimensionsSpecified, noOfArrayDimensions, arrayDimensions),
	}
	_result._Variant._VariantChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastVariantDataValue(structType any) VariantDataValue {
	if casted, ok := structType.(VariantDataValue); ok {
		return casted
	}
	if casted, ok := structType.(*VariantDataValue); ok {
		return *casted
	}
	return nil
}

func (m *_VariantDataValue) GetTypeName() string {
	return "VariantDataValue"
}

func (m *_VariantDataValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Optional Field (arrayLength)
	if m.ArrayLength != nil {
		lengthInBits += 32
	}

	// Array field
	if len(m.Value) > 0 {
		for _curItem, element := range m.Value {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.Value), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_VariantDataValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func VariantDataValueParse(ctx context.Context, theBytes []byte, arrayLengthSpecified bool) (VariantDataValue, error) {
	return VariantDataValueParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), arrayLengthSpecified)
}

func VariantDataValueParseWithBufferProducer(arrayLengthSpecified bool) func(ctx context.Context, readBuffer utils.ReadBuffer) (VariantDataValue, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (VariantDataValue, error) {
		return VariantDataValueParseWithBuffer(ctx, readBuffer, arrayLengthSpecified)
	}
}

func VariantDataValueParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, arrayLengthSpecified bool) (VariantDataValue, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("VariantDataValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for VariantDataValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	arrayLength, err := ReadOptionalField[int32](ctx, "arrayLength", ReadSignedInt(readBuffer, uint8(32)), arrayLengthSpecified)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'arrayLength' field"))
	}

	value, err := ReadCountArrayField[DataValue](ctx, "value", ReadComplex[DataValue](DataValueParseWithBuffer, readBuffer), uint64(utils.InlineIf(bool((arrayLength) == (nil)), func() any { return int32(int32(1)) }, func() any { return int32((*arrayLength)) }).(int32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}

	if closeErr := readBuffer.CloseContext("VariantDataValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for VariantDataValue")
	}

	// Create a partially initialized instance
	_child := &_VariantDataValue{
		_Variant:    &_Variant{},
		ArrayLength: arrayLength,
		Value:       value,
	}
	_child._Variant._VariantChildRequirements = _child
	return _child, nil
}

func (m *_VariantDataValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_VariantDataValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("VariantDataValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for VariantDataValue")
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
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetValue()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'value' field")
			}
		}
		if popErr := writeBuffer.PopContext("value", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for value")
		}

		if popErr := writeBuffer.PopContext("VariantDataValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for VariantDataValue")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_VariantDataValue) isVariantDataValue() bool {
	return true
}

func (m *_VariantDataValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
