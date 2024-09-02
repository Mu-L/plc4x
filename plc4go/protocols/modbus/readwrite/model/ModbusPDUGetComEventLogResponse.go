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

// ModbusPDUGetComEventLogResponse is the corresponding interface of ModbusPDUGetComEventLogResponse
type ModbusPDUGetComEventLogResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ModbusPDU
	// GetStatus returns Status (property field)
	GetStatus() uint16
	// GetEventCount returns EventCount (property field)
	GetEventCount() uint16
	// GetMessageCount returns MessageCount (property field)
	GetMessageCount() uint16
	// GetEvents returns Events (property field)
	GetEvents() []byte
}

// ModbusPDUGetComEventLogResponseExactly can be used when we want exactly this type and not a type which fulfills ModbusPDUGetComEventLogResponse.
// This is useful for switch cases.
type ModbusPDUGetComEventLogResponseExactly interface {
	ModbusPDUGetComEventLogResponse
	isModbusPDUGetComEventLogResponse() bool
}

// _ModbusPDUGetComEventLogResponse is the data-structure of this message
type _ModbusPDUGetComEventLogResponse struct {
	ModbusPDUContract
	Status       uint16
	EventCount   uint16
	MessageCount uint16
	Events       []byte
}

var _ ModbusPDUGetComEventLogResponse = (*_ModbusPDUGetComEventLogResponse)(nil)
var _ ModbusPDURequirements = (*_ModbusPDUGetComEventLogResponse)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ModbusPDUGetComEventLogResponse) GetErrorFlag() bool {
	return bool(false)
}

func (m *_ModbusPDUGetComEventLogResponse) GetFunctionFlag() uint8 {
	return 0x0C
}

func (m *_ModbusPDUGetComEventLogResponse) GetResponse() bool {
	return bool(true)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ModbusPDUGetComEventLogResponse) GetParent() ModbusPDUContract {
	return m.ModbusPDUContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ModbusPDUGetComEventLogResponse) GetStatus() uint16 {
	return m.Status
}

func (m *_ModbusPDUGetComEventLogResponse) GetEventCount() uint16 {
	return m.EventCount
}

func (m *_ModbusPDUGetComEventLogResponse) GetMessageCount() uint16 {
	return m.MessageCount
}

func (m *_ModbusPDUGetComEventLogResponse) GetEvents() []byte {
	return m.Events
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewModbusPDUGetComEventLogResponse factory function for _ModbusPDUGetComEventLogResponse
func NewModbusPDUGetComEventLogResponse(status uint16, eventCount uint16, messageCount uint16, events []byte) *_ModbusPDUGetComEventLogResponse {
	_result := &_ModbusPDUGetComEventLogResponse{
		ModbusPDUContract: NewModbusPDU(),
		Status:            status,
		EventCount:        eventCount,
		MessageCount:      messageCount,
		Events:            events,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastModbusPDUGetComEventLogResponse(structType any) ModbusPDUGetComEventLogResponse {
	if casted, ok := structType.(ModbusPDUGetComEventLogResponse); ok {
		return casted
	}
	if casted, ok := structType.(*ModbusPDUGetComEventLogResponse); ok {
		return *casted
	}
	return nil
}

func (m *_ModbusPDUGetComEventLogResponse) GetTypeName() string {
	return "ModbusPDUGetComEventLogResponse"
}

func (m *_ModbusPDUGetComEventLogResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ModbusPDUContract.(*_ModbusPDU).getLengthInBits(ctx))

	// Implicit Field (byteCount)
	lengthInBits += 8

	// Simple field (status)
	lengthInBits += 16

	// Simple field (eventCount)
	lengthInBits += 16

	// Simple field (messageCount)
	lengthInBits += 16

	// Array field
	if len(m.Events) > 0 {
		lengthInBits += 8 * uint16(len(m.Events))
	}

	return lengthInBits
}

func (m *_ModbusPDUGetComEventLogResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ModbusPDUGetComEventLogResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ModbusPDU, response bool) (__modbusPDUGetComEventLogResponse ModbusPDUGetComEventLogResponse, err error) {
	m.ModbusPDUContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModbusPDUGetComEventLogResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusPDUGetComEventLogResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	byteCount, err := ReadImplicitField[uint8](ctx, "byteCount", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'byteCount' field"))
	}
	_ = byteCount

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

	messageCount, err := ReadSimpleField(ctx, "messageCount", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'messageCount' field"))
	}
	m.MessageCount = messageCount

	events, err := readBuffer.ReadByteArray("events", int(int32(byteCount)-int32(int32(6))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'events' field"))
	}
	m.Events = events

	if closeErr := readBuffer.CloseContext("ModbusPDUGetComEventLogResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusPDUGetComEventLogResponse")
	}

	return m, nil
}

func (m *_ModbusPDUGetComEventLogResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ModbusPDUGetComEventLogResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUGetComEventLogResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ModbusPDUGetComEventLogResponse")
		}
		byteCount := uint8(uint8(uint8(len(m.GetEvents()))) + uint8(uint8(6)))
		if err := WriteImplicitField(ctx, "byteCount", byteCount, WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'byteCount' field")
		}

		if err := WriteSimpleField[uint16](ctx, "status", m.GetStatus(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'status' field")
		}

		if err := WriteSimpleField[uint16](ctx, "eventCount", m.GetEventCount(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'eventCount' field")
		}

		if err := WriteSimpleField[uint16](ctx, "messageCount", m.GetMessageCount(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'messageCount' field")
		}

		if err := WriteByteArrayField(ctx, "events", m.GetEvents(), WriteByteArray(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'events' field")
		}

		if popErr := writeBuffer.PopContext("ModbusPDUGetComEventLogResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ModbusPDUGetComEventLogResponse")
		}
		return nil
	}
	return m.ModbusPDUContract.(*_ModbusPDU).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ModbusPDUGetComEventLogResponse) isModbusPDUGetComEventLogResponse() bool {
	return true
}

func (m *_ModbusPDUGetComEventLogResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
