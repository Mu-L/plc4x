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

// ModbusPDUGetComEventCounterResponse is the corresponding interface of ModbusPDUGetComEventCounterResponse
type ModbusPDUGetComEventCounterResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ModbusPDU
	// GetStatus returns Status (property field)
	GetStatus() uint16
	// GetEventCount returns EventCount (property field)
	GetEventCount() uint16
}

// ModbusPDUGetComEventCounterResponseExactly can be used when we want exactly this type and not a type which fulfills ModbusPDUGetComEventCounterResponse.
// This is useful for switch cases.
type ModbusPDUGetComEventCounterResponseExactly interface {
	ModbusPDUGetComEventCounterResponse
	isModbusPDUGetComEventCounterResponse() bool
}

// _ModbusPDUGetComEventCounterResponse is the data-structure of this message
type _ModbusPDUGetComEventCounterResponse struct {
	ModbusPDUContract
	Status     uint16
	EventCount uint16
}

var _ ModbusPDUGetComEventCounterResponse = (*_ModbusPDUGetComEventCounterResponse)(nil)
var _ ModbusPDURequirements = (*_ModbusPDUGetComEventCounterResponse)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ModbusPDUGetComEventCounterResponse) GetErrorFlag() bool {
	return bool(false)
}

func (m *_ModbusPDUGetComEventCounterResponse) GetFunctionFlag() uint8 {
	return 0x0B
}

func (m *_ModbusPDUGetComEventCounterResponse) GetResponse() bool {
	return bool(true)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ModbusPDUGetComEventCounterResponse) GetParent() ModbusPDUContract {
	return m.ModbusPDUContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ModbusPDUGetComEventCounterResponse) GetStatus() uint16 {
	return m.Status
}

func (m *_ModbusPDUGetComEventCounterResponse) GetEventCount() uint16 {
	return m.EventCount
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewModbusPDUGetComEventCounterResponse factory function for _ModbusPDUGetComEventCounterResponse
func NewModbusPDUGetComEventCounterResponse(status uint16, eventCount uint16) *_ModbusPDUGetComEventCounterResponse {
	_result := &_ModbusPDUGetComEventCounterResponse{
		ModbusPDUContract: NewModbusPDU(),
		Status:            status,
		EventCount:        eventCount,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastModbusPDUGetComEventCounterResponse(structType any) ModbusPDUGetComEventCounterResponse {
	if casted, ok := structType.(ModbusPDUGetComEventCounterResponse); ok {
		return casted
	}
	if casted, ok := structType.(*ModbusPDUGetComEventCounterResponse); ok {
		return *casted
	}
	return nil
}

func (m *_ModbusPDUGetComEventCounterResponse) GetTypeName() string {
	return "ModbusPDUGetComEventCounterResponse"
}

func (m *_ModbusPDUGetComEventCounterResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ModbusPDUContract.(*_ModbusPDU).getLengthInBits(ctx))

	// Simple field (status)
	lengthInBits += 16

	// Simple field (eventCount)
	lengthInBits += 16

	return lengthInBits
}

func (m *_ModbusPDUGetComEventCounterResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ModbusPDUGetComEventCounterResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ModbusPDU, response bool) (__modbusPDUGetComEventCounterResponse ModbusPDUGetComEventCounterResponse, err error) {
	m.ModbusPDUContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModbusPDUGetComEventCounterResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusPDUGetComEventCounterResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	status, err := ReadSimpleField(ctx, "status", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'status' field"))
	}
	m.Status = status

	eventCount, err := ReadSimpleField(ctx, "eventCount", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'eventCount' field"))
	}
	m.EventCount = eventCount

	if closeErr := readBuffer.CloseContext("ModbusPDUGetComEventCounterResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusPDUGetComEventCounterResponse")
	}

	return m, nil
}

func (m *_ModbusPDUGetComEventCounterResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ModbusPDUGetComEventCounterResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUGetComEventCounterResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ModbusPDUGetComEventCounterResponse")
		}

		if err := WriteSimpleField[uint16](ctx, "status", m.GetStatus(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'status' field")
		}

		if err := WriteSimpleField[uint16](ctx, "eventCount", m.GetEventCount(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'eventCount' field")
		}

		if popErr := writeBuffer.PopContext("ModbusPDUGetComEventCounterResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ModbusPDUGetComEventCounterResponse")
		}
		return nil
	}
	return m.ModbusPDUContract.(*_ModbusPDU).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ModbusPDUGetComEventCounterResponse) isModbusPDUGetComEventCounterResponse() bool {
	return true
}

func (m *_ModbusPDUGetComEventCounterResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
