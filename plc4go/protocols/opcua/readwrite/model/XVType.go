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

// XVType is the corresponding interface of XVType
type XVType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetX returns X (property field)
	GetX() float64
	// GetValue returns Value (property field)
	GetValue() float32
}

// XVTypeExactly can be used when we want exactly this type and not a type which fulfills XVType.
// This is useful for switch cases.
type XVTypeExactly interface {
	XVType
	isXVType() bool
}

// _XVType is the data-structure of this message
type _XVType struct {
	*_ExtensionObjectDefinition
	X     float64
	Value float32
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_XVType) GetIdentifier() string {
	return "12082"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_XVType) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_XVType) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_XVType) GetX() float64 {
	return m.X
}

func (m *_XVType) GetValue() float32 {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewXVType factory function for _XVType
func NewXVType(x float64, value float32) *_XVType {
	_result := &_XVType{
		X:                          x,
		Value:                      value,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastXVType(structType any) XVType {
	if casted, ok := structType.(XVType); ok {
		return casted
	}
	if casted, ok := structType.(*XVType); ok {
		return *casted
	}
	return nil
}

func (m *_XVType) GetTypeName() string {
	return "XVType"
}

func (m *_XVType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (x)
	lengthInBits += 64

	// Simple field (value)
	lengthInBits += 32

	return lengthInBits
}

func (m *_XVType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func XVTypeParse(ctx context.Context, theBytes []byte, identifier string) (XVType, error) {
	return XVTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func XVTypeParseWithBufferProducer(identifier string) func(ctx context.Context, readBuffer utils.ReadBuffer) (XVType, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (XVType, error) {
		return XVTypeParseWithBuffer(ctx, readBuffer, identifier)
	}
}

func XVTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (XVType, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("XVType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for XVType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	x, err := ReadSimpleField(ctx, "x", ReadDouble(readBuffer, uint8(64)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'x' field"))
	}

	value, err := ReadSimpleField(ctx, "value", ReadFloat(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}

	if closeErr := readBuffer.CloseContext("XVType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for XVType")
	}

	// Create a partially initialized instance
	_child := &_XVType{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		X:                          x,
		Value:                      value,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_XVType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_XVType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("XVType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for XVType")
		}

		// Simple Field (x)
		x := float64(m.GetX())
		_xErr := /*TODO: migrate me*/ writeBuffer.WriteFloat64("x", 64, (x))
		if _xErr != nil {
			return errors.Wrap(_xErr, "Error serializing 'x' field")
		}

		// Simple Field (value)
		value := float32(m.GetValue())
		_valueErr := /*TODO: migrate me*/ writeBuffer.WriteFloat32("value", 32, (value))
		if _valueErr != nil {
			return errors.Wrap(_valueErr, "Error serializing 'value' field")
		}

		if popErr := writeBuffer.PopContext("XVType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for XVType")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_XVType) isXVType() bool {
	return true
}

func (m *_XVType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
