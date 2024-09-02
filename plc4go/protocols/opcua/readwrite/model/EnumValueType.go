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

// EnumValueType is the corresponding interface of EnumValueType
type EnumValueType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetValue returns Value (property field)
	GetValue() int64
	// GetDisplayName returns DisplayName (property field)
	GetDisplayName() LocalizedText
	// GetDescription returns Description (property field)
	GetDescription() LocalizedText
}

// EnumValueTypeExactly can be used when we want exactly this type and not a type which fulfills EnumValueType.
// This is useful for switch cases.
type EnumValueTypeExactly interface {
	EnumValueType
	isEnumValueType() bool
}

// _EnumValueType is the data-structure of this message
type _EnumValueType struct {
	ExtensionObjectDefinitionContract
	Value       int64
	DisplayName LocalizedText
	Description LocalizedText
}

var _ EnumValueType = (*_EnumValueType)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_EnumValueType)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_EnumValueType) GetIdentifier() string {
	return "7596"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_EnumValueType) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_EnumValueType) GetValue() int64 {
	return m.Value
}

func (m *_EnumValueType) GetDisplayName() LocalizedText {
	return m.DisplayName
}

func (m *_EnumValueType) GetDescription() LocalizedText {
	return m.Description
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewEnumValueType factory function for _EnumValueType
func NewEnumValueType(value int64, displayName LocalizedText, description LocalizedText) *_EnumValueType {
	_result := &_EnumValueType{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		Value:                             value,
		DisplayName:                       displayName,
		Description:                       description,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastEnumValueType(structType any) EnumValueType {
	if casted, ok := structType.(EnumValueType); ok {
		return casted
	}
	if casted, ok := structType.(*EnumValueType); ok {
		return *casted
	}
	return nil
}

func (m *_EnumValueType) GetTypeName() string {
	return "EnumValueType"
}

func (m *_EnumValueType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (value)
	lengthInBits += 64

	// Simple field (displayName)
	lengthInBits += m.DisplayName.GetLengthInBits(ctx)

	// Simple field (description)
	lengthInBits += m.Description.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_EnumValueType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_EnumValueType) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__enumValueType EnumValueType, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("EnumValueType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for EnumValueType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	value, err := ReadSimpleField(ctx, "value", ReadSignedLong(readBuffer, uint8(64)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	displayName, err := ReadSimpleField[LocalizedText](ctx, "displayName", ReadComplex[LocalizedText](LocalizedTextParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'displayName' field"))
	}
	m.DisplayName = displayName

	description, err := ReadSimpleField[LocalizedText](ctx, "description", ReadComplex[LocalizedText](LocalizedTextParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'description' field"))
	}
	m.Description = description

	if closeErr := readBuffer.CloseContext("EnumValueType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for EnumValueType")
	}

	return m, nil
}

func (m *_EnumValueType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_EnumValueType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("EnumValueType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for EnumValueType")
		}

		if err := WriteSimpleField[int64](ctx, "value", m.GetValue(), WriteSignedLong(writeBuffer, 64)); err != nil {
			return errors.Wrap(err, "Error serializing 'value' field")
		}

		if err := WriteSimpleField[LocalizedText](ctx, "displayName", m.GetDisplayName(), WriteComplex[LocalizedText](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'displayName' field")
		}

		if err := WriteSimpleField[LocalizedText](ctx, "description", m.GetDescription(), WriteComplex[LocalizedText](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'description' field")
		}

		if popErr := writeBuffer.PopContext("EnumValueType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for EnumValueType")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_EnumValueType) isEnumValueType() bool {
	return true
}

func (m *_EnumValueType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
