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
	"encoding/binary"
	"fmt"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	"github.com/apache/plc4x/plc4go/spi/codegen"
	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// Constant values.
const DF1Symbol_MESSAGESTART uint8 = 0x10

// DF1Symbol is the corresponding interface of DF1Symbol
type DF1Symbol interface {
	DF1SymbolContract
	DF1SymbolRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

// DF1SymbolContract provides a set of functions which can be overwritten by a sub struct
type DF1SymbolContract interface {
}

// DF1SymbolRequirements provides a set of functions which need to be implemented by a sub struct
type DF1SymbolRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetSymbolType returns SymbolType (discriminator field)
	GetSymbolType() uint8
}

// DF1SymbolExactly can be used when we want exactly this type and not a type which fulfills DF1Symbol.
// This is useful for switch cases.
type DF1SymbolExactly interface {
	DF1Symbol
	isDF1Symbol() bool
}

// _DF1Symbol is the data-structure of this message
type _DF1Symbol struct {
	_SubType DF1Symbol
}

var _ DF1SymbolContract = (*_DF1Symbol)(nil)

type DF1SymbolChild interface {
	utils.Serializable

	GetParent() *DF1Symbol

	GetTypeName() string
	DF1Symbol
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_DF1Symbol) GetMessageStart() uint8 {
	return DF1Symbol_MESSAGESTART
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewDF1Symbol factory function for _DF1Symbol
func NewDF1Symbol() *_DF1Symbol {
	return &_DF1Symbol{}
}

// Deprecated: use the interface for direct cast
func CastDF1Symbol(structType any) DF1Symbol {
	if casted, ok := structType.(DF1Symbol); ok {
		return casted
	}
	if casted, ok := structType.(*DF1Symbol); ok {
		return *casted
	}
	return nil
}

func (m *_DF1Symbol) GetTypeName() string {
	return "DF1Symbol"
}

func (m *_DF1Symbol) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Const Field (messageStart)
	lengthInBits += 8
	// Discriminator Field (symbolType)
	lengthInBits += 8

	return lengthInBits
}

func (m *_DF1Symbol) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func DF1SymbolParse[T DF1Symbol](ctx context.Context, theBytes []byte) (T, error) {
	return DF1SymbolParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian)))
}

func DF1SymbolParseWithBufferProducer[T DF1Symbol]() func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := DF1SymbolParseWithBuffer[T](ctx, readBuffer)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, err
	}
}

func DF1SymbolParseWithBuffer[T DF1Symbol](ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	v, err := (&_DF1Symbol{}).parse(ctx, readBuffer)
	if err != nil {
		var zero T
		return zero, err
	}
	return v.(T), err
}

func (m *_DF1Symbol) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__dF1Symbol DF1Symbol, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("DF1Symbol"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DF1Symbol")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	messageStart, err := ReadConstField[uint8](ctx, "messageStart", ReadUnsignedByte(readBuffer, uint8(8)), DF1Symbol_MESSAGESTART, codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'messageStart' field"))
	}
	_ = messageStart

	symbolType, err := ReadDiscriminatorField[uint8](ctx, "symbolType", ReadUnsignedByte(readBuffer, uint8(8)), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'symbolType' field"))
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child DF1Symbol
	switch {
	case symbolType == 0x02: // DF1SymbolMessageFrame
		if _child, err = (&_DF1SymbolMessageFrame{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type DF1SymbolMessageFrame for type-switch of DF1Symbol")
		}
	case symbolType == 0x06: // DF1SymbolMessageFrameACK
		if _child, err = (&_DF1SymbolMessageFrameACK{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type DF1SymbolMessageFrameACK for type-switch of DF1Symbol")
		}
	case symbolType == 0x15: // DF1SymbolMessageFrameNAK
		if _child, err = (&_DF1SymbolMessageFrameNAK{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type DF1SymbolMessageFrameNAK for type-switch of DF1Symbol")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [symbolType=%v]", symbolType)
	}

	if closeErr := readBuffer.CloseContext("DF1Symbol"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DF1Symbol")
	}

	return _child, nil
}

func (pm *_DF1Symbol) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child DF1Symbol, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("DF1Symbol"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for DF1Symbol")
	}

	if err := WriteConstField(ctx, "messageStart", DF1Symbol_MESSAGESTART, WriteUnsignedByte(writeBuffer, 8), codegen.WithByteOrder(binary.BigEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'messageStart' field")
	}

	if err := WriteDiscriminatorField(ctx, "symbolType", m.GetSymbolType(), WriteUnsignedByte(writeBuffer, 8), codegen.WithByteOrder(binary.BigEndian)); err != nil {
		return errors.Wrap(err, "Error serializing 'symbolType' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("DF1Symbol"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for DF1Symbol")
	}
	return nil
}

func (m *_DF1Symbol) isDF1Symbol() bool {
	return true
}
