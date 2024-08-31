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

// Constant values.
const EipConnectionResponse_PROTOCOLVERSION uint16 = 0x01
const EipConnectionResponse_FLAGS uint16 = 0x00

// EipConnectionResponse is the corresponding interface of EipConnectionResponse
type EipConnectionResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	EipPacket
}

// EipConnectionResponseExactly can be used when we want exactly this type and not a type which fulfills EipConnectionResponse.
// This is useful for switch cases.
type EipConnectionResponseExactly interface {
	EipConnectionResponse
	isEipConnectionResponse() bool
}

// _EipConnectionResponse is the data-structure of this message
type _EipConnectionResponse struct {
	*_EipPacket
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_EipConnectionResponse) GetCommand() uint16 {
	return 0x0065
}

func (m *_EipConnectionResponse) GetResponse() bool {
	return bool(true)
}

func (m *_EipConnectionResponse) GetPacketLength() uint16 {
	return 0
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_EipConnectionResponse) InitializeParent(parent EipPacket, sessionHandle uint32, status uint32, senderContext []byte, options uint32) {
	m.SessionHandle = sessionHandle
	m.Status = status
	m.SenderContext = senderContext
	m.Options = options
}

func (m *_EipConnectionResponse) GetParent() EipPacket {
	return m._EipPacket
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_EipConnectionResponse) GetProtocolVersion() uint16 {
	return EipConnectionResponse_PROTOCOLVERSION
}

func (m *_EipConnectionResponse) GetFlags() uint16 {
	return EipConnectionResponse_FLAGS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewEipConnectionResponse factory function for _EipConnectionResponse
func NewEipConnectionResponse(sessionHandle uint32, status uint32, senderContext []byte, options uint32) *_EipConnectionResponse {
	_result := &_EipConnectionResponse{
		_EipPacket: NewEipPacket(sessionHandle, status, senderContext, options),
	}
	_result._EipPacket._EipPacketChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastEipConnectionResponse(structType any) EipConnectionResponse {
	if casted, ok := structType.(EipConnectionResponse); ok {
		return casted
	}
	if casted, ok := structType.(*EipConnectionResponse); ok {
		return *casted
	}
	return nil
}

func (m *_EipConnectionResponse) GetTypeName() string {
	return "EipConnectionResponse"
}

func (m *_EipConnectionResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Const Field (protocolVersion)
	lengthInBits += 16

	// Const Field (flags)
	lengthInBits += 16

	return lengthInBits
}

func (m *_EipConnectionResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func EipConnectionResponseParse(ctx context.Context, theBytes []byte, response bool) (EipConnectionResponse, error) {
	return EipConnectionResponseParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), response)
}

func EipConnectionResponseParseWithBufferProducer(response bool) func(ctx context.Context, readBuffer utils.ReadBuffer) (EipConnectionResponse, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (EipConnectionResponse, error) {
		return EipConnectionResponseParseWithBuffer(ctx, readBuffer, response)
	}
}

func EipConnectionResponseParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, response bool) (EipConnectionResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("EipConnectionResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for EipConnectionResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	protocolVersion, err := ReadConstField[uint16](ctx, "protocolVersion", ReadUnsignedShort(readBuffer, uint8(16)), EipConnectionResponse_PROTOCOLVERSION)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'protocolVersion' field"))
	}
	_ = protocolVersion

	flags, err := ReadConstField[uint16](ctx, "flags", ReadUnsignedShort(readBuffer, uint8(16)), EipConnectionResponse_FLAGS)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'flags' field"))
	}
	_ = flags

	if closeErr := readBuffer.CloseContext("EipConnectionResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for EipConnectionResponse")
	}

	// Create a partially initialized instance
	_child := &_EipConnectionResponse{
		_EipPacket: &_EipPacket{},
	}
	_child._EipPacket._EipPacketChildRequirements = _child
	return _child, nil
}

func (m *_EipConnectionResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_EipConnectionResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("EipConnectionResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for EipConnectionResponse")
		}

		// Const Field (protocolVersion)
		_protocolVersionErr := /*TODO: migrate me*/ /*TODO: migrate me*/ writeBuffer.WriteUint16("protocolVersion", 16, uint16(0x01))
		if _protocolVersionErr != nil {
			return errors.Wrap(_protocolVersionErr, "Error serializing 'protocolVersion' field")
		}

		// Const Field (flags)
		_flagsErr := /*TODO: migrate me*/ /*TODO: migrate me*/ writeBuffer.WriteUint16("flags", 16, uint16(0x00))
		if _flagsErr != nil {
			return errors.Wrap(_flagsErr, "Error serializing 'flags' field")
		}

		if popErr := writeBuffer.PopContext("EipConnectionResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for EipConnectionResponse")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_EipConnectionResponse) isEipConnectionResponse() bool {
	return true
}

func (m *_EipConnectionResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
