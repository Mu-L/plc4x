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

// OpcuaAcknowledgeResponse is the corresponding interface of OpcuaAcknowledgeResponse
type OpcuaAcknowledgeResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	MessagePDU
	// GetVersion returns Version (property field)
	GetVersion() uint32
	// GetLimits returns Limits (property field)
	GetLimits() OpcuaProtocolLimits
}

// OpcuaAcknowledgeResponseExactly can be used when we want exactly this type and not a type which fulfills OpcuaAcknowledgeResponse.
// This is useful for switch cases.
type OpcuaAcknowledgeResponseExactly interface {
	OpcuaAcknowledgeResponse
	isOpcuaAcknowledgeResponse() bool
}

// _OpcuaAcknowledgeResponse is the data-structure of this message
type _OpcuaAcknowledgeResponse struct {
	MessagePDUContract
	Version uint32
	Limits  OpcuaProtocolLimits
}

var _ OpcuaAcknowledgeResponse = (*_OpcuaAcknowledgeResponse)(nil)
var _ MessagePDURequirements = (*_OpcuaAcknowledgeResponse)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_OpcuaAcknowledgeResponse) GetMessageType() string {
	return "ACK"
}

func (m *_OpcuaAcknowledgeResponse) GetResponse() bool {
	return bool(true)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_OpcuaAcknowledgeResponse) GetParent() MessagePDUContract {
	return m.MessagePDUContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_OpcuaAcknowledgeResponse) GetVersion() uint32 {
	return m.Version
}

func (m *_OpcuaAcknowledgeResponse) GetLimits() OpcuaProtocolLimits {
	return m.Limits
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewOpcuaAcknowledgeResponse factory function for _OpcuaAcknowledgeResponse
func NewOpcuaAcknowledgeResponse(version uint32, limits OpcuaProtocolLimits, chunk ChunkType) *_OpcuaAcknowledgeResponse {
	_result := &_OpcuaAcknowledgeResponse{
		MessagePDUContract: NewMessagePDU(chunk),
		Version:            version,
		Limits:             limits,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastOpcuaAcknowledgeResponse(structType any) OpcuaAcknowledgeResponse {
	if casted, ok := structType.(OpcuaAcknowledgeResponse); ok {
		return casted
	}
	if casted, ok := structType.(*OpcuaAcknowledgeResponse); ok {
		return *casted
	}
	return nil
}

func (m *_OpcuaAcknowledgeResponse) GetTypeName() string {
	return "OpcuaAcknowledgeResponse"
}

func (m *_OpcuaAcknowledgeResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.MessagePDUContract.(*_MessagePDU).getLengthInBits(ctx))

	// Simple field (version)
	lengthInBits += 32

	// Simple field (limits)
	lengthInBits += m.Limits.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_OpcuaAcknowledgeResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_OpcuaAcknowledgeResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_MessagePDU, response bool) (__opcuaAcknowledgeResponse OpcuaAcknowledgeResponse, err error) {
	m.MessagePDUContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("OpcuaAcknowledgeResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for OpcuaAcknowledgeResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	version, err := ReadSimpleField(ctx, "version", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'version' field"))
	}
	m.Version = version

	limits, err := ReadSimpleField[OpcuaProtocolLimits](ctx, "limits", ReadComplex[OpcuaProtocolLimits](OpcuaProtocolLimitsParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'limits' field"))
	}
	m.Limits = limits

	if closeErr := readBuffer.CloseContext("OpcuaAcknowledgeResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for OpcuaAcknowledgeResponse")
	}

	return m, nil
}

func (m *_OpcuaAcknowledgeResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_OpcuaAcknowledgeResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("OpcuaAcknowledgeResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for OpcuaAcknowledgeResponse")
		}

		if err := WriteSimpleField[uint32](ctx, "version", m.GetVersion(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'version' field")
		}

		if err := WriteSimpleField[OpcuaProtocolLimits](ctx, "limits", m.GetLimits(), WriteComplex[OpcuaProtocolLimits](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'limits' field")
		}

		if popErr := writeBuffer.PopContext("OpcuaAcknowledgeResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for OpcuaAcknowledgeResponse")
		}
		return nil
	}
	return m.MessagePDUContract.(*_MessagePDU).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_OpcuaAcknowledgeResponse) isOpcuaAcknowledgeResponse() bool {
	return true
}

func (m *_OpcuaAcknowledgeResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
