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

// ModbusPDUWriteSingleRegisterRequest is the corresponding interface of ModbusPDUWriteSingleRegisterRequest
type ModbusPDUWriteSingleRegisterRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ModbusPDU
	// GetAddress returns Address (property field)
	GetAddress() uint16
	// GetValue returns Value (property field)
	GetValue() uint16
}

// ModbusPDUWriteSingleRegisterRequestExactly can be used when we want exactly this type and not a type which fulfills ModbusPDUWriteSingleRegisterRequest.
// This is useful for switch cases.
type ModbusPDUWriteSingleRegisterRequestExactly interface {
	ModbusPDUWriteSingleRegisterRequest
	isModbusPDUWriteSingleRegisterRequest() bool
}

// _ModbusPDUWriteSingleRegisterRequest is the data-structure of this message
type _ModbusPDUWriteSingleRegisterRequest struct {
	ModbusPDUContract
	Address uint16
	Value   uint16
}

var _ ModbusPDUWriteSingleRegisterRequest = (*_ModbusPDUWriteSingleRegisterRequest)(nil)
var _ ModbusPDURequirements = (*_ModbusPDUWriteSingleRegisterRequest)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ModbusPDUWriteSingleRegisterRequest) GetErrorFlag() bool {
	return bool(false)
}

func (m *_ModbusPDUWriteSingleRegisterRequest) GetFunctionFlag() uint8 {
	return 0x06
}

func (m *_ModbusPDUWriteSingleRegisterRequest) GetResponse() bool {
	return bool(false)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ModbusPDUWriteSingleRegisterRequest) GetParent() ModbusPDUContract {
	return m.ModbusPDUContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ModbusPDUWriteSingleRegisterRequest) GetAddress() uint16 {
	return m.Address
}

func (m *_ModbusPDUWriteSingleRegisterRequest) GetValue() uint16 {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewModbusPDUWriteSingleRegisterRequest factory function for _ModbusPDUWriteSingleRegisterRequest
func NewModbusPDUWriteSingleRegisterRequest(address uint16, value uint16) *_ModbusPDUWriteSingleRegisterRequest {
	_result := &_ModbusPDUWriteSingleRegisterRequest{
		ModbusPDUContract: NewModbusPDU(),
		Address:           address,
		Value:             value,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastModbusPDUWriteSingleRegisterRequest(structType any) ModbusPDUWriteSingleRegisterRequest {
	if casted, ok := structType.(ModbusPDUWriteSingleRegisterRequest); ok {
		return casted
	}
	if casted, ok := structType.(*ModbusPDUWriteSingleRegisterRequest); ok {
		return *casted
	}
	return nil
}

func (m *_ModbusPDUWriteSingleRegisterRequest) GetTypeName() string {
	return "ModbusPDUWriteSingleRegisterRequest"
}

func (m *_ModbusPDUWriteSingleRegisterRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ModbusPDUContract.(*_ModbusPDU).getLengthInBits(ctx))

	// Simple field (address)
	lengthInBits += 16

	// Simple field (value)
	lengthInBits += 16

	return lengthInBits
}

func (m *_ModbusPDUWriteSingleRegisterRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ModbusPDUWriteSingleRegisterRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ModbusPDU, response bool) (__modbusPDUWriteSingleRegisterRequest ModbusPDUWriteSingleRegisterRequest, err error) {
	m.ModbusPDUContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModbusPDUWriteSingleRegisterRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusPDUWriteSingleRegisterRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	address, err := ReadSimpleField(ctx, "address", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'address' field"))
	}
	m.Address = address

	value, err := ReadSimpleField(ctx, "value", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	if closeErr := readBuffer.CloseContext("ModbusPDUWriteSingleRegisterRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusPDUWriteSingleRegisterRequest")
	}

	return m, nil
}

func (m *_ModbusPDUWriteSingleRegisterRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ModbusPDUWriteSingleRegisterRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUWriteSingleRegisterRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ModbusPDUWriteSingleRegisterRequest")
		}

		if err := WriteSimpleField[uint16](ctx, "address", m.GetAddress(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'address' field")
		}

		if err := WriteSimpleField[uint16](ctx, "value", m.GetValue(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'value' field")
		}

		if popErr := writeBuffer.PopContext("ModbusPDUWriteSingleRegisterRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ModbusPDUWriteSingleRegisterRequest")
		}
		return nil
	}
	return m.ModbusPDUContract.(*_ModbusPDU).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ModbusPDUWriteSingleRegisterRequest) isModbusPDUWriteSingleRegisterRequest() bool {
	return true
}

func (m *_ModbusPDUWriteSingleRegisterRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
