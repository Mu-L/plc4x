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

// S7PayloadUserDataItemCpuFunctionAlarmQueryResponse is the corresponding interface of S7PayloadUserDataItemCpuFunctionAlarmQueryResponse
type S7PayloadUserDataItemCpuFunctionAlarmQueryResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	S7PayloadUserDataItem
	// GetItems returns Items (property field)
	GetItems() []byte
}

// S7PayloadUserDataItemCpuFunctionAlarmQueryResponseExactly can be used when we want exactly this type and not a type which fulfills S7PayloadUserDataItemCpuFunctionAlarmQueryResponse.
// This is useful for switch cases.
type S7PayloadUserDataItemCpuFunctionAlarmQueryResponseExactly interface {
	S7PayloadUserDataItemCpuFunctionAlarmQueryResponse
	isS7PayloadUserDataItemCpuFunctionAlarmQueryResponse() bool
}

// _S7PayloadUserDataItemCpuFunctionAlarmQueryResponse is the data-structure of this message
type _S7PayloadUserDataItemCpuFunctionAlarmQueryResponse struct {
	S7PayloadUserDataItemContract
	Items []byte
}

var _ S7PayloadUserDataItemCpuFunctionAlarmQueryResponse = (*_S7PayloadUserDataItemCpuFunctionAlarmQueryResponse)(nil)
var _ S7PayloadUserDataItemRequirements = (*_S7PayloadUserDataItemCpuFunctionAlarmQueryResponse)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_S7PayloadUserDataItemCpuFunctionAlarmQueryResponse) GetCpuFunctionGroup() uint8 {
	return 0x04
}

func (m *_S7PayloadUserDataItemCpuFunctionAlarmQueryResponse) GetCpuFunctionType() uint8 {
	return 0x08
}

func (m *_S7PayloadUserDataItemCpuFunctionAlarmQueryResponse) GetCpuSubfunction() uint8 {
	return 0x13
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_S7PayloadUserDataItemCpuFunctionAlarmQueryResponse) GetParent() S7PayloadUserDataItemContract {
	return m.S7PayloadUserDataItemContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_S7PayloadUserDataItemCpuFunctionAlarmQueryResponse) GetItems() []byte {
	return m.Items
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewS7PayloadUserDataItemCpuFunctionAlarmQueryResponse factory function for _S7PayloadUserDataItemCpuFunctionAlarmQueryResponse
func NewS7PayloadUserDataItemCpuFunctionAlarmQueryResponse(items []byte, returnCode DataTransportErrorCode, transportSize DataTransportSize, dataLength uint16) *_S7PayloadUserDataItemCpuFunctionAlarmQueryResponse {
	_result := &_S7PayloadUserDataItemCpuFunctionAlarmQueryResponse{
		S7PayloadUserDataItemContract: NewS7PayloadUserDataItem(returnCode, transportSize, dataLength),
		Items:                         items,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastS7PayloadUserDataItemCpuFunctionAlarmQueryResponse(structType any) S7PayloadUserDataItemCpuFunctionAlarmQueryResponse {
	if casted, ok := structType.(S7PayloadUserDataItemCpuFunctionAlarmQueryResponse); ok {
		return casted
	}
	if casted, ok := structType.(*S7PayloadUserDataItemCpuFunctionAlarmQueryResponse); ok {
		return *casted
	}
	return nil
}

func (m *_S7PayloadUserDataItemCpuFunctionAlarmQueryResponse) GetTypeName() string {
	return "S7PayloadUserDataItemCpuFunctionAlarmQueryResponse"
}

func (m *_S7PayloadUserDataItemCpuFunctionAlarmQueryResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem).getLengthInBits(ctx))

	// Array field
	if len(m.Items) > 0 {
		lengthInBits += 8 * uint16(len(m.Items))
	}

	return lengthInBits
}

func (m *_S7PayloadUserDataItemCpuFunctionAlarmQueryResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_S7PayloadUserDataItemCpuFunctionAlarmQueryResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_S7PayloadUserDataItem, dataLength uint16, cpuFunctionGroup uint8, cpuFunctionType uint8, cpuSubfunction uint8) (__s7PayloadUserDataItemCpuFunctionAlarmQueryResponse S7PayloadUserDataItemCpuFunctionAlarmQueryResponse, err error) {
	m.S7PayloadUserDataItemContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("S7PayloadUserDataItemCpuFunctionAlarmQueryResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7PayloadUserDataItemCpuFunctionAlarmQueryResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	items, err := readBuffer.ReadByteArray("items", int(dataLength))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'items' field"))
	}
	m.Items = items

	if closeErr := readBuffer.CloseContext("S7PayloadUserDataItemCpuFunctionAlarmQueryResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7PayloadUserDataItemCpuFunctionAlarmQueryResponse")
	}

	return m, nil
}

func (m *_S7PayloadUserDataItemCpuFunctionAlarmQueryResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_S7PayloadUserDataItemCpuFunctionAlarmQueryResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7PayloadUserDataItemCpuFunctionAlarmQueryResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for S7PayloadUserDataItemCpuFunctionAlarmQueryResponse")
		}

		if err := WriteByteArrayField(ctx, "items", m.GetItems(), WriteByteArray(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'items' field")
		}

		if popErr := writeBuffer.PopContext("S7PayloadUserDataItemCpuFunctionAlarmQueryResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for S7PayloadUserDataItemCpuFunctionAlarmQueryResponse")
		}
		return nil
	}
	return m.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_S7PayloadUserDataItemCpuFunctionAlarmQueryResponse) isS7PayloadUserDataItemCpuFunctionAlarmQueryResponse() bool {
	return true
}

func (m *_S7PayloadUserDataItemCpuFunctionAlarmQueryResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
