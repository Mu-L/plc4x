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

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// EipDisconnectRequest is the corresponding interface of EipDisconnectRequest
type EipDisconnectRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	EipPacket
}

// EipDisconnectRequestExactly can be used when we want exactly this type and not a type which fulfills EipDisconnectRequest.
// This is useful for switch cases.
type EipDisconnectRequestExactly interface {
	EipDisconnectRequest
	isEipDisconnectRequest() bool
}

// _EipDisconnectRequest is the data-structure of this message
type _EipDisconnectRequest struct {
	*_EipPacket
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_EipDisconnectRequest) GetCommand() uint16 {
	return 0x0066
}

func (m *_EipDisconnectRequest) GetResponse() bool {
	return false
}

func (m *_EipDisconnectRequest) GetPacketLength() uint16 {
	return 0
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_EipDisconnectRequest) InitializeParent(parent EipPacket, sessionHandle uint32, status uint32, senderContext []byte, options uint32) {
	m.SessionHandle = sessionHandle
	m.Status = status
	m.SenderContext = senderContext
	m.Options = options
}

func (m *_EipDisconnectRequest) GetParent() EipPacket {
	return m._EipPacket
}

// NewEipDisconnectRequest factory function for _EipDisconnectRequest
func NewEipDisconnectRequest(sessionHandle uint32, status uint32, senderContext []byte, options uint32) *_EipDisconnectRequest {
	_result := &_EipDisconnectRequest{
		_EipPacket: NewEipPacket(sessionHandle, status, senderContext, options),
	}
	_result._EipPacket._EipPacketChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastEipDisconnectRequest(structType any) EipDisconnectRequest {
	if casted, ok := structType.(EipDisconnectRequest); ok {
		return casted
	}
	if casted, ok := structType.(*EipDisconnectRequest); ok {
		return *casted
	}
	return nil
}

func (m *_EipDisconnectRequest) GetTypeName() string {
	return "EipDisconnectRequest"
}

func (m *_EipDisconnectRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	return lengthInBits
}

func (m *_EipDisconnectRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func EipDisconnectRequestParse(ctx context.Context, theBytes []byte, response bool) (EipDisconnectRequest, error) {
	return EipDisconnectRequestParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), response)
}

func EipDisconnectRequestParseWithBufferProducer(response bool) func(ctx context.Context, readBuffer utils.ReadBuffer) (EipDisconnectRequest, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (EipDisconnectRequest, error) {
		return EipDisconnectRequestParseWithBuffer(ctx, readBuffer, response)
	}
}

func EipDisconnectRequestParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, response bool) (EipDisconnectRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("EipDisconnectRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for EipDisconnectRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("EipDisconnectRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for EipDisconnectRequest")
	}

	// Create a partially initialized instance
	_child := &_EipDisconnectRequest{
		_EipPacket: &_EipPacket{},
	}
	_child._EipPacket._EipPacketChildRequirements = _child
	return _child, nil
}

func (m *_EipDisconnectRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_EipDisconnectRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("EipDisconnectRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for EipDisconnectRequest")
		}

		if popErr := writeBuffer.PopContext("EipDisconnectRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for EipDisconnectRequest")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_EipDisconnectRequest) isEipDisconnectRequest() bool {
	return true
}

func (m *_EipDisconnectRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
