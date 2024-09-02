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

// TelephonyDataLineOffHook is the corresponding interface of TelephonyDataLineOffHook
type TelephonyDataLineOffHook interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	TelephonyData
	// GetReason returns Reason (property field)
	GetReason() LineOffHookReason
	// GetNumber returns Number (property field)
	GetNumber() string
}

// TelephonyDataLineOffHookExactly can be used when we want exactly this type and not a type which fulfills TelephonyDataLineOffHook.
// This is useful for switch cases.
type TelephonyDataLineOffHookExactly interface {
	TelephonyDataLineOffHook
	isTelephonyDataLineOffHook() bool
}

// _TelephonyDataLineOffHook is the data-structure of this message
type _TelephonyDataLineOffHook struct {
	TelephonyDataContract
	Reason LineOffHookReason
	Number string
}

var _ TelephonyDataLineOffHook = (*_TelephonyDataLineOffHook)(nil)
var _ TelephonyDataRequirements = (*_TelephonyDataLineOffHook)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_TelephonyDataLineOffHook) GetParent() TelephonyDataContract {
	return m.TelephonyDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_TelephonyDataLineOffHook) GetReason() LineOffHookReason {
	return m.Reason
}

func (m *_TelephonyDataLineOffHook) GetNumber() string {
	return m.Number
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewTelephonyDataLineOffHook factory function for _TelephonyDataLineOffHook
func NewTelephonyDataLineOffHook(reason LineOffHookReason, number string, commandTypeContainer TelephonyCommandTypeContainer, argument byte) *_TelephonyDataLineOffHook {
	_result := &_TelephonyDataLineOffHook{
		TelephonyDataContract: NewTelephonyData(commandTypeContainer, argument),
		Reason:                reason,
		Number:                number,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastTelephonyDataLineOffHook(structType any) TelephonyDataLineOffHook {
	if casted, ok := structType.(TelephonyDataLineOffHook); ok {
		return casted
	}
	if casted, ok := structType.(*TelephonyDataLineOffHook); ok {
		return *casted
	}
	return nil
}

func (m *_TelephonyDataLineOffHook) GetTypeName() string {
	return "TelephonyDataLineOffHook"
}

func (m *_TelephonyDataLineOffHook) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.TelephonyDataContract.(*_TelephonyData).getLengthInBits(ctx))

	// Simple field (reason)
	lengthInBits += 8

	// Simple field (number)
	lengthInBits += uint16(int32((int32(m.GetCommandTypeContainer().NumBytes()) - int32(int32(2)))) * int32(int32(8)))

	return lengthInBits
}

func (m *_TelephonyDataLineOffHook) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_TelephonyDataLineOffHook) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_TelephonyData, commandTypeContainer TelephonyCommandTypeContainer) (__telephonyDataLineOffHook TelephonyDataLineOffHook, err error) {
	m.TelephonyDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("TelephonyDataLineOffHook"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for TelephonyDataLineOffHook")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	reason, err := ReadEnumField[LineOffHookReason](ctx, "reason", "LineOffHookReason", ReadEnum(LineOffHookReasonByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'reason' field"))
	}
	m.Reason = reason

	number, err := ReadSimpleField(ctx, "number", ReadString(readBuffer, uint32(int32((int32(commandTypeContainer.NumBytes())-int32(int32(2))))*int32(int32(8)))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'number' field"))
	}
	m.Number = number

	if closeErr := readBuffer.CloseContext("TelephonyDataLineOffHook"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for TelephonyDataLineOffHook")
	}

	return m, nil
}

func (m *_TelephonyDataLineOffHook) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_TelephonyDataLineOffHook) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("TelephonyDataLineOffHook"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for TelephonyDataLineOffHook")
		}

		if err := WriteSimpleEnumField[LineOffHookReason](ctx, "reason", "LineOffHookReason", m.GetReason(), WriteEnum[LineOffHookReason, uint8](LineOffHookReason.GetValue, LineOffHookReason.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
			return errors.Wrap(err, "Error serializing 'reason' field")
		}

		if err := WriteSimpleField[string](ctx, "number", m.GetNumber(), WriteString(writeBuffer, int32(int32((int32(m.GetCommandTypeContainer().NumBytes())-int32(int32(2))))*int32(int32(8))))); err != nil {
			return errors.Wrap(err, "Error serializing 'number' field")
		}

		if popErr := writeBuffer.PopContext("TelephonyDataLineOffHook"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for TelephonyDataLineOffHook")
		}
		return nil
	}
	return m.TelephonyDataContract.(*_TelephonyData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_TelephonyDataLineOffHook) isTelephonyDataLineOffHook() bool {
	return true
}

func (m *_TelephonyDataLineOffHook) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
