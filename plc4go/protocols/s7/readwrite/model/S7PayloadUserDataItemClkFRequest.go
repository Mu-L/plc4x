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

// S7PayloadUserDataItemClkFRequest is the corresponding interface of S7PayloadUserDataItemClkFRequest
type S7PayloadUserDataItemClkFRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	S7PayloadUserDataItem
}

// S7PayloadUserDataItemClkFRequestExactly can be used when we want exactly this type and not a type which fulfills S7PayloadUserDataItemClkFRequest.
// This is useful for switch cases.
type S7PayloadUserDataItemClkFRequestExactly interface {
	S7PayloadUserDataItemClkFRequest
	isS7PayloadUserDataItemClkFRequest() bool
}

// _S7PayloadUserDataItemClkFRequest is the data-structure of this message
type _S7PayloadUserDataItemClkFRequest struct {
	*_S7PayloadUserDataItem
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_S7PayloadUserDataItemClkFRequest) GetCpuFunctionGroup() uint8 {
	return 0x07
}

func (m *_S7PayloadUserDataItemClkFRequest) GetCpuFunctionType() uint8 {
	return 0x04
}

func (m *_S7PayloadUserDataItemClkFRequest) GetCpuSubfunction() uint8 {
	return 0x03
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_S7PayloadUserDataItemClkFRequest) InitializeParent(parent S7PayloadUserDataItem, returnCode DataTransportErrorCode, transportSize DataTransportSize, dataLength uint16) {
	m.ReturnCode = returnCode
	m.TransportSize = transportSize
	m.DataLength = dataLength
}

func (m *_S7PayloadUserDataItemClkFRequest) GetParent() S7PayloadUserDataItem {
	return m._S7PayloadUserDataItem
}

// NewS7PayloadUserDataItemClkFRequest factory function for _S7PayloadUserDataItemClkFRequest
func NewS7PayloadUserDataItemClkFRequest(returnCode DataTransportErrorCode, transportSize DataTransportSize, dataLength uint16) *_S7PayloadUserDataItemClkFRequest {
	_result := &_S7PayloadUserDataItemClkFRequest{
		_S7PayloadUserDataItem: NewS7PayloadUserDataItem(returnCode, transportSize, dataLength),
	}
	_result._S7PayloadUserDataItem._S7PayloadUserDataItemChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastS7PayloadUserDataItemClkFRequest(structType any) S7PayloadUserDataItemClkFRequest {
	if casted, ok := structType.(S7PayloadUserDataItemClkFRequest); ok {
		return casted
	}
	if casted, ok := structType.(*S7PayloadUserDataItemClkFRequest); ok {
		return *casted
	}
	return nil
}

func (m *_S7PayloadUserDataItemClkFRequest) GetTypeName() string {
	return "S7PayloadUserDataItemClkFRequest"
}

func (m *_S7PayloadUserDataItemClkFRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	return lengthInBits
}

func (m *_S7PayloadUserDataItemClkFRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func S7PayloadUserDataItemClkFRequestParse(ctx context.Context, theBytes []byte, cpuFunctionGroup uint8, cpuFunctionType uint8, cpuSubfunction uint8) (S7PayloadUserDataItemClkFRequest, error) {
	return S7PayloadUserDataItemClkFRequestParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), cpuFunctionGroup, cpuFunctionType, cpuSubfunction)
}

func S7PayloadUserDataItemClkFRequestParseWithBufferProducer(cpuFunctionGroup uint8, cpuFunctionType uint8, cpuSubfunction uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (S7PayloadUserDataItemClkFRequest, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (S7PayloadUserDataItemClkFRequest, error) {
		return S7PayloadUserDataItemClkFRequestParseWithBuffer(ctx, readBuffer, cpuFunctionGroup, cpuFunctionType, cpuSubfunction)
	}
}

func S7PayloadUserDataItemClkFRequestParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, cpuFunctionGroup uint8, cpuFunctionType uint8, cpuSubfunction uint8) (S7PayloadUserDataItemClkFRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("S7PayloadUserDataItemClkFRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7PayloadUserDataItemClkFRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("S7PayloadUserDataItemClkFRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7PayloadUserDataItemClkFRequest")
	}

	// Create a partially initialized instance
	_child := &_S7PayloadUserDataItemClkFRequest{
		_S7PayloadUserDataItem: &_S7PayloadUserDataItem{},
	}
	_child._S7PayloadUserDataItem._S7PayloadUserDataItemChildRequirements = _child
	return _child, nil
}

func (m *_S7PayloadUserDataItemClkFRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_S7PayloadUserDataItemClkFRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7PayloadUserDataItemClkFRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for S7PayloadUserDataItemClkFRequest")
		}

		if popErr := writeBuffer.PopContext("S7PayloadUserDataItemClkFRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for S7PayloadUserDataItemClkFRequest")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_S7PayloadUserDataItemClkFRequest) isS7PayloadUserDataItemClkFRequest() bool {
	return true
}

func (m *_S7PayloadUserDataItemClkFRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
