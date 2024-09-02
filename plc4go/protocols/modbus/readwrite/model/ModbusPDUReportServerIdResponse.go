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

// ModbusPDUReportServerIdResponse is the corresponding interface of ModbusPDUReportServerIdResponse
type ModbusPDUReportServerIdResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ModbusPDU
	// GetValue returns Value (property field)
	GetValue() []byte
	// IsModbusPDUReportServerIdResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsModbusPDUReportServerIdResponse()
}

// _ModbusPDUReportServerIdResponse is the data-structure of this message
type _ModbusPDUReportServerIdResponse struct {
	ModbusPDUContract
	Value []byte
}

var _ ModbusPDUReportServerIdResponse = (*_ModbusPDUReportServerIdResponse)(nil)
var _ ModbusPDURequirements = (*_ModbusPDUReportServerIdResponse)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ModbusPDUReportServerIdResponse) GetErrorFlag() bool {
	return bool(false)
}

func (m *_ModbusPDUReportServerIdResponse) GetFunctionFlag() uint8 {
	return 0x11
}

func (m *_ModbusPDUReportServerIdResponse) GetResponse() bool {
	return bool(true)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ModbusPDUReportServerIdResponse) GetParent() ModbusPDUContract {
	return m.ModbusPDUContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ModbusPDUReportServerIdResponse) GetValue() []byte {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewModbusPDUReportServerIdResponse factory function for _ModbusPDUReportServerIdResponse
func NewModbusPDUReportServerIdResponse(value []byte) *_ModbusPDUReportServerIdResponse {
	_result := &_ModbusPDUReportServerIdResponse{
		ModbusPDUContract: NewModbusPDU(),
		Value:             value,
	}
	_result.ModbusPDUContract.(*_ModbusPDU)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastModbusPDUReportServerIdResponse(structType any) ModbusPDUReportServerIdResponse {
	if casted, ok := structType.(ModbusPDUReportServerIdResponse); ok {
		return casted
	}
	if casted, ok := structType.(*ModbusPDUReportServerIdResponse); ok {
		return *casted
	}
	return nil
}

func (m *_ModbusPDUReportServerIdResponse) GetTypeName() string {
	return "ModbusPDUReportServerIdResponse"
}

func (m *_ModbusPDUReportServerIdResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ModbusPDUContract.(*_ModbusPDU).getLengthInBits(ctx))

	// Implicit Field (byteCount)
	lengthInBits += 8

	// Array field
	if len(m.Value) > 0 {
		lengthInBits += 8 * uint16(len(m.Value))
	}

	return lengthInBits
}

func (m *_ModbusPDUReportServerIdResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ModbusPDUReportServerIdResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ModbusPDU, response bool) (__modbusPDUReportServerIdResponse ModbusPDUReportServerIdResponse, err error) {
	m.ModbusPDUContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModbusPDUReportServerIdResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusPDUReportServerIdResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	byteCount, err := ReadImplicitField[uint8](ctx, "byteCount", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'byteCount' field"))
	}
	_ = byteCount

	value, err := readBuffer.ReadByteArray("value", int(byteCount))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	if closeErr := readBuffer.CloseContext("ModbusPDUReportServerIdResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusPDUReportServerIdResponse")
	}

	return m, nil
}

func (m *_ModbusPDUReportServerIdResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ModbusPDUReportServerIdResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUReportServerIdResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ModbusPDUReportServerIdResponse")
		}
		byteCount := uint8(uint8(len(m.GetValue())))
		if err := WriteImplicitField(ctx, "byteCount", byteCount, WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'byteCount' field")
		}

		if err := WriteByteArrayField(ctx, "value", m.GetValue(), WriteByteArray(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'value' field")
		}

		if popErr := writeBuffer.PopContext("ModbusPDUReportServerIdResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ModbusPDUReportServerIdResponse")
		}
		return nil
	}
	return m.ModbusPDUContract.(*_ModbusPDU).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ModbusPDUReportServerIdResponse) IsModbusPDUReportServerIdResponse() {}

func (m *_ModbusPDUReportServerIdResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
