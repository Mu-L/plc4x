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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// Variant is the corresponding interface of Variant
type Variant interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetVariantType returns VariantType (discriminator field)
	GetVariantType() uint8
	// GetArrayLengthSpecified returns ArrayLengthSpecified (property field)
	GetArrayLengthSpecified() bool
	// GetArrayDimensionsSpecified returns ArrayDimensionsSpecified (property field)
	GetArrayDimensionsSpecified() bool
	// GetNoOfArrayDimensions returns NoOfArrayDimensions (property field)
	GetNoOfArrayDimensions() *int32
	// GetArrayDimensions returns ArrayDimensions (property field)
	GetArrayDimensions() []bool
}

// VariantExactly can be used when we want exactly this type and not a type which fulfills Variant.
// This is useful for switch cases.
type VariantExactly interface {
	Variant
	isVariant() bool
}

// _Variant is the data-structure of this message
type _Variant struct {
	_VariantChildRequirements
	ArrayLengthSpecified     bool
	ArrayDimensionsSpecified bool
	NoOfArrayDimensions      *int32
	ArrayDimensions          []bool
}

type _VariantChildRequirements interface {
	utils.Serializable
	GetLengthInBits(ctx context.Context) uint16
	GetVariantType() uint8
	GetArrayLengthSpecified() bool
}

type VariantParent interface {
	SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child Variant, serializeChildFunction func() error) error
	GetTypeName() string
}

type VariantChild interface {
	utils.Serializable
	InitializeParent(parent Variant, arrayLengthSpecified bool, arrayDimensionsSpecified bool, noOfArrayDimensions *int32, arrayDimensions []bool)
	GetParent() *Variant

	GetTypeName() string
	Variant
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_Variant) GetArrayLengthSpecified() bool {
	return m.ArrayLengthSpecified
}

func (m *_Variant) GetArrayDimensionsSpecified() bool {
	return m.ArrayDimensionsSpecified
}

func (m *_Variant) GetNoOfArrayDimensions() *int32 {
	return m.NoOfArrayDimensions
}

func (m *_Variant) GetArrayDimensions() []bool {
	return m.ArrayDimensions
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewVariant factory function for _Variant
func NewVariant(arrayLengthSpecified bool, arrayDimensionsSpecified bool, noOfArrayDimensions *int32, arrayDimensions []bool) *_Variant {
	return &_Variant{ArrayLengthSpecified: arrayLengthSpecified, ArrayDimensionsSpecified: arrayDimensionsSpecified, NoOfArrayDimensions: noOfArrayDimensions, ArrayDimensions: arrayDimensions}
}

// Deprecated: use the interface for direct cast
func CastVariant(structType any) Variant {
	if casted, ok := structType.(Variant); ok {
		return casted
	}
	if casted, ok := structType.(*Variant); ok {
		return *casted
	}
	return nil
}

func (m *_Variant) GetTypeName() string {
	return "Variant"
}

func (m *_Variant) GetParentLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (arrayLengthSpecified)
	lengthInBits += 1

	// Simple field (arrayDimensionsSpecified)
	lengthInBits += 1
	// Discriminator Field (VariantType)
	lengthInBits += 6

	// Optional Field (noOfArrayDimensions)
	if m.NoOfArrayDimensions != nil {
		lengthInBits += 32
	}

	// Array field
	if len(m.ArrayDimensions) > 0 {
		lengthInBits += 1 * uint16(len(m.ArrayDimensions))
	}

	return lengthInBits
}

func (m *_Variant) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func VariantParse(ctx context.Context, theBytes []byte) (Variant, error) {
	return VariantParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func VariantParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (Variant, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("Variant"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for Variant")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (arrayLengthSpecified)
	_arrayLengthSpecified, _arrayLengthSpecifiedErr := readBuffer.ReadBit("arrayLengthSpecified")
	if _arrayLengthSpecifiedErr != nil {
		return nil, errors.Wrap(_arrayLengthSpecifiedErr, "Error parsing 'arrayLengthSpecified' field of Variant")
	}
	arrayLengthSpecified := _arrayLengthSpecified

	// Simple Field (arrayDimensionsSpecified)
	_arrayDimensionsSpecified, _arrayDimensionsSpecifiedErr := readBuffer.ReadBit("arrayDimensionsSpecified")
	if _arrayDimensionsSpecifiedErr != nil {
		return nil, errors.Wrap(_arrayDimensionsSpecifiedErr, "Error parsing 'arrayDimensionsSpecified' field of Variant")
	}
	arrayDimensionsSpecified := _arrayDimensionsSpecified

	// Discriminator Field (VariantType) (Used as input to a switch field)
	VariantType, _VariantTypeErr := readBuffer.ReadUint8("VariantType", 6)
	if _VariantTypeErr != nil {
		return nil, errors.Wrap(_VariantTypeErr, "Error parsing 'VariantType' field of Variant")
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type VariantChildSerializeRequirement interface {
		Variant
		InitializeParent(Variant, bool, bool, *int32, []bool)
		GetParent() Variant
	}
	var _childTemp any
	var _child VariantChildSerializeRequirement
	var typeSwitchError error
	switch {
	case VariantType == uint8(1): // VariantBoolean
		_childTemp, typeSwitchError = VariantBooleanParseWithBuffer(ctx, readBuffer, arrayLengthSpecified)
	case VariantType == uint8(2): // VariantSByte
		_childTemp, typeSwitchError = VariantSByteParseWithBuffer(ctx, readBuffer, arrayLengthSpecified)
	case VariantType == uint8(3): // VariantByte
		_childTemp, typeSwitchError = VariantByteParseWithBuffer(ctx, readBuffer, arrayLengthSpecified)
	case VariantType == uint8(4): // VariantInt16
		_childTemp, typeSwitchError = VariantInt16ParseWithBuffer(ctx, readBuffer, arrayLengthSpecified)
	case VariantType == uint8(5): // VariantUInt16
		_childTemp, typeSwitchError = VariantUInt16ParseWithBuffer(ctx, readBuffer, arrayLengthSpecified)
	case VariantType == uint8(6): // VariantInt32
		_childTemp, typeSwitchError = VariantInt32ParseWithBuffer(ctx, readBuffer, arrayLengthSpecified)
	case VariantType == uint8(7): // VariantUInt32
		_childTemp, typeSwitchError = VariantUInt32ParseWithBuffer(ctx, readBuffer, arrayLengthSpecified)
	case VariantType == uint8(8): // VariantInt64
		_childTemp, typeSwitchError = VariantInt64ParseWithBuffer(ctx, readBuffer, arrayLengthSpecified)
	case VariantType == uint8(9): // VariantUInt64
		_childTemp, typeSwitchError = VariantUInt64ParseWithBuffer(ctx, readBuffer, arrayLengthSpecified)
	case VariantType == uint8(10): // VariantFloat
		_childTemp, typeSwitchError = VariantFloatParseWithBuffer(ctx, readBuffer, arrayLengthSpecified)
	case VariantType == uint8(11): // VariantDouble
		_childTemp, typeSwitchError = VariantDoubleParseWithBuffer(ctx, readBuffer, arrayLengthSpecified)
	case VariantType == uint8(12): // VariantString
		_childTemp, typeSwitchError = VariantStringParseWithBuffer(ctx, readBuffer, arrayLengthSpecified)
	case VariantType == uint8(13): // VariantDateTime
		_childTemp, typeSwitchError = VariantDateTimeParseWithBuffer(ctx, readBuffer, arrayLengthSpecified)
	case VariantType == uint8(14): // VariantGuid
		_childTemp, typeSwitchError = VariantGuidParseWithBuffer(ctx, readBuffer, arrayLengthSpecified)
	case VariantType == uint8(15): // VariantByteString
		_childTemp, typeSwitchError = VariantByteStringParseWithBuffer(ctx, readBuffer, arrayLengthSpecified)
	case VariantType == uint8(16): // VariantXmlElement
		_childTemp, typeSwitchError = VariantXmlElementParseWithBuffer(ctx, readBuffer, arrayLengthSpecified)
	case VariantType == uint8(17): // VariantNodeId
		_childTemp, typeSwitchError = VariantNodeIdParseWithBuffer(ctx, readBuffer, arrayLengthSpecified)
	case VariantType == uint8(18): // VariantExpandedNodeId
		_childTemp, typeSwitchError = VariantExpandedNodeIdParseWithBuffer(ctx, readBuffer, arrayLengthSpecified)
	case VariantType == uint8(19): // VariantStatusCode
		_childTemp, typeSwitchError = VariantStatusCodeParseWithBuffer(ctx, readBuffer, arrayLengthSpecified)
	case VariantType == uint8(20): // VariantQualifiedName
		_childTemp, typeSwitchError = VariantQualifiedNameParseWithBuffer(ctx, readBuffer, arrayLengthSpecified)
	case VariantType == uint8(21): // VariantLocalizedText
		_childTemp, typeSwitchError = VariantLocalizedTextParseWithBuffer(ctx, readBuffer, arrayLengthSpecified)
	case VariantType == uint8(22): // VariantExtensionObject
		_childTemp, typeSwitchError = VariantExtensionObjectParseWithBuffer(ctx, readBuffer, arrayLengthSpecified)
	case VariantType == uint8(23): // VariantDataValue
		_childTemp, typeSwitchError = VariantDataValueParseWithBuffer(ctx, readBuffer, arrayLengthSpecified)
	case VariantType == uint8(24): // VariantVariant
		_childTemp, typeSwitchError = VariantVariantParseWithBuffer(ctx, readBuffer, arrayLengthSpecified)
	case VariantType == uint8(25): // VariantDiagnosticInfo
		_childTemp, typeSwitchError = VariantDiagnosticInfoParseWithBuffer(ctx, readBuffer, arrayLengthSpecified)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [VariantType=%v, arrayLengthSpecified=%v]", VariantType, arrayLengthSpecified)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of Variant")
	}
	_child = _childTemp.(VariantChildSerializeRequirement)

	// Optional Field (noOfArrayDimensions) (Can be skipped, if a given expression evaluates to false)
	var noOfArrayDimensions *int32 = nil
	if arrayDimensionsSpecified {
		_val, _err := readBuffer.ReadInt32("noOfArrayDimensions", 32)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'noOfArrayDimensions' field of Variant")
		}
		noOfArrayDimensions = &_val
	}

	// Array field (arrayDimensions)
	if pullErr := readBuffer.PullContext("arrayDimensions", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for arrayDimensions")
	}
	// Count array
	arrayDimensions := make([]bool, utils.InlineIf(bool((noOfArrayDimensions) == (nil)), func() any { return uint16(uint16(0)) }, func() any { return uint16((*noOfArrayDimensions)) }).(uint16))
	// This happens when the size is set conditional to 0
	if len(arrayDimensions) == 0 {
		arrayDimensions = nil
	}
	{
		_numItems := uint16(utils.InlineIf(bool((noOfArrayDimensions) == (nil)), func() any { return uint16(uint16(0)) }, func() any { return uint16((*noOfArrayDimensions)) }).(uint16))
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
			_item, _err := readBuffer.ReadBit("")
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'arrayDimensions' field of Variant")
			}
			arrayDimensions[_curItem] = _item
		}
	}
	if closeErr := readBuffer.CloseContext("arrayDimensions", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for arrayDimensions")
	}

	if closeErr := readBuffer.CloseContext("Variant"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for Variant")
	}

	// Finish initializing
	_child.InitializeParent(_child, arrayLengthSpecified, arrayDimensionsSpecified, noOfArrayDimensions, arrayDimensions)
	return _child, nil
}

func (pm *_Variant) SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child Variant, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("Variant"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for Variant")
	}

	// Simple Field (arrayLengthSpecified)
	arrayLengthSpecified := bool(m.GetArrayLengthSpecified())
	_arrayLengthSpecifiedErr := writeBuffer.WriteBit("arrayLengthSpecified", (arrayLengthSpecified))
	if _arrayLengthSpecifiedErr != nil {
		return errors.Wrap(_arrayLengthSpecifiedErr, "Error serializing 'arrayLengthSpecified' field")
	}

	// Simple Field (arrayDimensionsSpecified)
	arrayDimensionsSpecified := bool(m.GetArrayDimensionsSpecified())
	_arrayDimensionsSpecifiedErr := writeBuffer.WriteBit("arrayDimensionsSpecified", (arrayDimensionsSpecified))
	if _arrayDimensionsSpecifiedErr != nil {
		return errors.Wrap(_arrayDimensionsSpecifiedErr, "Error serializing 'arrayDimensionsSpecified' field")
	}

	// Discriminator Field (VariantType) (Used as input to a switch field)
	VariantType := uint8(child.GetVariantType())
	_VariantTypeErr := writeBuffer.WriteUint8("VariantType", 6, (VariantType))

	if _VariantTypeErr != nil {
		return errors.Wrap(_VariantTypeErr, "Error serializing 'VariantType' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	// Optional Field (noOfArrayDimensions) (Can be skipped, if the value is null)
	var noOfArrayDimensions *int32 = nil
	if m.GetNoOfArrayDimensions() != nil {
		noOfArrayDimensions = m.GetNoOfArrayDimensions()
		_noOfArrayDimensionsErr := writeBuffer.WriteInt32("noOfArrayDimensions", 32, *(noOfArrayDimensions))
		if _noOfArrayDimensionsErr != nil {
			return errors.Wrap(_noOfArrayDimensionsErr, "Error serializing 'noOfArrayDimensions' field")
		}
	}

	// Array Field (arrayDimensions)
	if pushErr := writeBuffer.PushContext("arrayDimensions", utils.WithRenderAsList(true)); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for arrayDimensions")
	}
	for _curItem, _element := range m.GetArrayDimensions() {
		_ = _curItem
		_elementErr := writeBuffer.WriteBit("", _element)
		if _elementErr != nil {
			return errors.Wrap(_elementErr, "Error serializing 'arrayDimensions' field")
		}
	}
	if popErr := writeBuffer.PopContext("arrayDimensions", utils.WithRenderAsList(true)); popErr != nil {
		return errors.Wrap(popErr, "Error popping for arrayDimensions")
	}

	if popErr := writeBuffer.PopContext("Variant"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for Variant")
	}
	return nil
}

func (m *_Variant) isVariant() bool {
	return true
}

func (m *_Variant) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
