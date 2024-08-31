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

// AdsInvalidRequest is the corresponding interface of AdsInvalidRequest
type AdsInvalidRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	AmsPacket
}

// AdsInvalidRequestExactly can be used when we want exactly this type and not a type which fulfills AdsInvalidRequest.
// This is useful for switch cases.
type AdsInvalidRequestExactly interface {
	AdsInvalidRequest
	isAdsInvalidRequest() bool
}

// _AdsInvalidRequest is the data-structure of this message
type _AdsInvalidRequest struct {
	*_AmsPacket
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_AdsInvalidRequest) GetCommandId() CommandId {
	return CommandId_INVALID
}

func (m *_AdsInvalidRequest) GetResponse() bool {
	return bool(false)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AdsInvalidRequest) InitializeParent(parent AmsPacket, targetAmsNetId AmsNetId, targetAmsPort uint16, sourceAmsNetId AmsNetId, sourceAmsPort uint16, errorCode uint32, invokeId uint32) {
	m.TargetAmsNetId = targetAmsNetId
	m.TargetAmsPort = targetAmsPort
	m.SourceAmsNetId = sourceAmsNetId
	m.SourceAmsPort = sourceAmsPort
	m.ErrorCode = errorCode
	m.InvokeId = invokeId
}

func (m *_AdsInvalidRequest) GetParent() AmsPacket {
	return m._AmsPacket
}

// NewAdsInvalidRequest factory function for _AdsInvalidRequest
func NewAdsInvalidRequest(targetAmsNetId AmsNetId, targetAmsPort uint16, sourceAmsNetId AmsNetId, sourceAmsPort uint16, errorCode uint32, invokeId uint32) *_AdsInvalidRequest {
	_result := &_AdsInvalidRequest{
		_AmsPacket: NewAmsPacket(targetAmsNetId, targetAmsPort, sourceAmsNetId, sourceAmsPort, errorCode, invokeId),
	}
	_result._AmsPacket._AmsPacketChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastAdsInvalidRequest(structType any) AdsInvalidRequest {
	if casted, ok := structType.(AdsInvalidRequest); ok {
		return casted
	}
	if casted, ok := structType.(*AdsInvalidRequest); ok {
		return *casted
	}
	return nil
}

func (m *_AdsInvalidRequest) GetTypeName() string {
	return "AdsInvalidRequest"
}

func (m *_AdsInvalidRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	return lengthInBits
}

func (m *_AdsInvalidRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func AdsInvalidRequestParse(ctx context.Context, theBytes []byte) (AdsInvalidRequest, error) {
	return AdsInvalidRequestParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func AdsInvalidRequestParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (AdsInvalidRequest, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (AdsInvalidRequest, error) {
		return AdsInvalidRequestParseWithBuffer(ctx, readBuffer)
	}
}

func AdsInvalidRequestParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (AdsInvalidRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("AdsInvalidRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AdsInvalidRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("AdsInvalidRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AdsInvalidRequest")
	}

	// Create a partially initialized instance
	_child := &_AdsInvalidRequest{
		_AmsPacket: &_AmsPacket{},
	}
	_child._AmsPacket._AmsPacketChildRequirements = _child
	return _child, nil
}

func (m *_AdsInvalidRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AdsInvalidRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AdsInvalidRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AdsInvalidRequest")
		}

		if popErr := writeBuffer.PopContext("AdsInvalidRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AdsInvalidRequest")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AdsInvalidRequest) isAdsInvalidRequest() bool {
	return true
}

func (m *_AdsInvalidRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
