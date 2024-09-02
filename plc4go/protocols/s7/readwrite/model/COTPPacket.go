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

// COTPPacket is the corresponding interface of COTPPacket
type COTPPacket interface {
	COTPPacketContract
	COTPPacketRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

// COTPPacketContract provides a set of functions which can be overwritten by a sub struct
type COTPPacketContract interface {
	// GetParameters returns Parameters (property field)
	GetParameters() []COTPParameter
	// GetPayload returns Payload (property field)
	GetPayload() S7Message
	// GetCotpLen() returns a parser argument
	GetCotpLen() uint16
}

// COTPPacketRequirements provides a set of functions which need to be implemented by a sub struct
type COTPPacketRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetTpduCode returns TpduCode (discriminator field)
	GetTpduCode() uint8
}

// COTPPacketExactly can be used when we want exactly this type and not a type which fulfills COTPPacket.
// This is useful for switch cases.
type COTPPacketExactly interface {
	COTPPacket
	isCOTPPacket() bool
}

// _COTPPacket is the data-structure of this message
type _COTPPacket struct {
	_SubType   COTPPacket
	Parameters []COTPParameter
	Payload    S7Message

	// Arguments.
	CotpLen uint16
}

var _ COTPPacketContract = (*_COTPPacket)(nil)

type COTPPacketChild interface {
	utils.Serializable

	GetParent() *COTPPacket

	GetTypeName() string
	COTPPacket
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_COTPPacket) GetParameters() []COTPParameter {
	return m.Parameters
}

func (m *_COTPPacket) GetPayload() S7Message {
	return m.Payload
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCOTPPacket factory function for _COTPPacket
func NewCOTPPacket(parameters []COTPParameter, payload S7Message, cotpLen uint16) *_COTPPacket {
	return &_COTPPacket{Parameters: parameters, Payload: payload, CotpLen: cotpLen}
}

// Deprecated: use the interface for direct cast
func CastCOTPPacket(structType any) COTPPacket {
	if casted, ok := structType.(COTPPacket); ok {
		return casted
	}
	if casted, ok := structType.(*COTPPacket); ok {
		return *casted
	}
	return nil
}

func (m *_COTPPacket) GetTypeName() string {
	return "COTPPacket"
}

func (m *_COTPPacket) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Implicit Field (headerLength)
	lengthInBits += 8
	// Discriminator Field (tpduCode)
	lengthInBits += 8

	// Array field
	if len(m.Parameters) > 0 {
		for _, element := range m.Parameters {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	// Optional Field (payload)
	if m.Payload != nil {
		lengthInBits += m.Payload.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_COTPPacket) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func COTPPacketParse[T COTPPacket](ctx context.Context, theBytes []byte, cotpLen uint16) (T, error) {
	return COTPPacketParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes), cotpLen)
}

func COTPPacketParseWithBufferProducer[T COTPPacket](cotpLen uint16) func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := COTPPacketParseWithBuffer[T](ctx, readBuffer, cotpLen)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, err
	}
}

func COTPPacketParseWithBuffer[T COTPPacket](ctx context.Context, readBuffer utils.ReadBuffer, cotpLen uint16) (T, error) {
	v, err := (&_COTPPacket{}).parse(ctx, readBuffer, cotpLen)
	if err != nil {
		var zero T
		return zero, err
	}
	return v.(T), err
}

func (m *_COTPPacket) parse(ctx context.Context, readBuffer utils.ReadBuffer, cotpLen uint16) (__cOTPPacket COTPPacket, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("COTPPacket"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for COTPPacket")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos
	var startPos = positionAware.GetPos()
	_ = startPos

	headerLength, err := ReadImplicitField[uint8](ctx, "headerLength", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'headerLength' field"))
	}
	_ = headerLength

	tpduCode, err := ReadDiscriminatorField[uint8](ctx, "tpduCode", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'tpduCode' field"))
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child COTPPacket
	switch {
	case tpduCode == 0xF0: // COTPPacketData
		if _child, err = (&_COTPPacketData{}).parse(ctx, readBuffer, m, cotpLen); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type COTPPacketData for type-switch of COTPPacket")
		}
	case tpduCode == 0xE0: // COTPPacketConnectionRequest
		if _child, err = (&_COTPPacketConnectionRequest{}).parse(ctx, readBuffer, m, cotpLen); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type COTPPacketConnectionRequest for type-switch of COTPPacket")
		}
	case tpduCode == 0xD0: // COTPPacketConnectionResponse
		if _child, err = (&_COTPPacketConnectionResponse{}).parse(ctx, readBuffer, m, cotpLen); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type COTPPacketConnectionResponse for type-switch of COTPPacket")
		}
	case tpduCode == 0x80: // COTPPacketDisconnectRequest
		if _child, err = (&_COTPPacketDisconnectRequest{}).parse(ctx, readBuffer, m, cotpLen); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type COTPPacketDisconnectRequest for type-switch of COTPPacket")
		}
	case tpduCode == 0xC0: // COTPPacketDisconnectResponse
		if _child, err = (&_COTPPacketDisconnectResponse{}).parse(ctx, readBuffer, m, cotpLen); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type COTPPacketDisconnectResponse for type-switch of COTPPacket")
		}
	case tpduCode == 0x70: // COTPPacketTpduError
		if _child, err = (&_COTPPacketTpduError{}).parse(ctx, readBuffer, m, cotpLen); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type COTPPacketTpduError for type-switch of COTPPacket")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [tpduCode=%v]", tpduCode)
	}

	parameters, err := ReadLengthArrayField[COTPParameter](ctx, "parameters", ReadComplex[COTPParameter](COTPParameterParseWithBufferProducer[COTPParameter]((uint8)(uint8((uint8(headerLength)+uint8(uint8(1))))-uint8((positionAware.GetPos()-startPos)))), readBuffer), int(int32((int32(headerLength)+int32(int32(1))))-int32((positionAware.GetPos()-startPos))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'parameters' field"))
	}
	m.Parameters = parameters

	var payload S7Message
	_payload, err := ReadOptionalField[S7Message](ctx, "payload", ReadComplex[S7Message](S7MessageParseWithBuffer, readBuffer), bool((positionAware.GetPos()-startPos) < (cotpLen)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'payload' field"))
	}
	if _payload != nil {
		payload = *_payload
		m.Payload = payload
	}

	if closeErr := readBuffer.CloseContext("COTPPacket"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for COTPPacket")
	}

	return _child, nil
}

func (pm *_COTPPacket) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child COTPPacket, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("COTPPacket"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for COTPPacket")
	}
	headerLength := uint8(uint8(uint8(m.GetLengthInBytes(ctx))) - uint8((uint8((utils.InlineIf((bool((m.GetPayload()) != (nil))), func() any { return uint8((m.GetPayload()).GetLengthInBytes(ctx)) }, func() any { return uint8(uint8(0)) }).(uint8))) + uint8(uint8(1)))))
	if err := WriteImplicitField(ctx, "headerLength", headerLength, WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'headerLength' field")
	}

	if err := WriteDiscriminatorField(ctx, "tpduCode", m.GetTpduCode(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'tpduCode' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if err := WriteComplexTypeArrayField(ctx, "parameters", m.GetParameters(), writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'parameters' field")
	}

	if err := WriteOptionalField[S7Message](ctx, "payload", GetRef(m.GetPayload()), WriteComplex[S7Message](writeBuffer), true); err != nil {
		return errors.Wrap(err, "Error serializing 'payload' field")
	}

	if popErr := writeBuffer.PopContext("COTPPacket"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for COTPPacket")
	}
	return nil
}

////
// Arguments Getter

func (m *_COTPPacket) GetCotpLen() uint16 {
	return m.CotpLen
}

//
////

func (m *_COTPPacket) isCOTPPacket() bool {
	return true
}
