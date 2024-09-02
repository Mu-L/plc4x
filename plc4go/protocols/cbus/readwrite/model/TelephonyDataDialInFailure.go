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

// TelephonyDataDialInFailure is the corresponding interface of TelephonyDataDialInFailure
type TelephonyDataDialInFailure interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	TelephonyData
	// GetReason returns Reason (property field)
	GetReason() DialInFailureReason
	// IsTelephonyDataDialInFailure is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsTelephonyDataDialInFailure()
}

// _TelephonyDataDialInFailure is the data-structure of this message
type _TelephonyDataDialInFailure struct {
	TelephonyDataContract
	Reason DialInFailureReason
}

var _ TelephonyDataDialInFailure = (*_TelephonyDataDialInFailure)(nil)
var _ TelephonyDataRequirements = (*_TelephonyDataDialInFailure)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_TelephonyDataDialInFailure) GetParent() TelephonyDataContract {
	return m.TelephonyDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_TelephonyDataDialInFailure) GetReason() DialInFailureReason {
	return m.Reason
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewTelephonyDataDialInFailure factory function for _TelephonyDataDialInFailure
func NewTelephonyDataDialInFailure(reason DialInFailureReason, commandTypeContainer TelephonyCommandTypeContainer, argument byte) *_TelephonyDataDialInFailure {
	_result := &_TelephonyDataDialInFailure{
		TelephonyDataContract: NewTelephonyData(commandTypeContainer, argument),
		Reason:                reason,
	}
	_result.TelephonyDataContract.(*_TelephonyData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastTelephonyDataDialInFailure(structType any) TelephonyDataDialInFailure {
	if casted, ok := structType.(TelephonyDataDialInFailure); ok {
		return casted
	}
	if casted, ok := structType.(*TelephonyDataDialInFailure); ok {
		return *casted
	}
	return nil
}

func (m *_TelephonyDataDialInFailure) GetTypeName() string {
	return "TelephonyDataDialInFailure"
}

func (m *_TelephonyDataDialInFailure) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.TelephonyDataContract.(*_TelephonyData).getLengthInBits(ctx))

	// Simple field (reason)
	lengthInBits += 8

	return lengthInBits
}

func (m *_TelephonyDataDialInFailure) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_TelephonyDataDialInFailure) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_TelephonyData) (__telephonyDataDialInFailure TelephonyDataDialInFailure, err error) {
	m.TelephonyDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("TelephonyDataDialInFailure"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for TelephonyDataDialInFailure")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	reason, err := ReadEnumField[DialInFailureReason](ctx, "reason", "DialInFailureReason", ReadEnum(DialInFailureReasonByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'reason' field"))
	}
	m.Reason = reason

	if closeErr := readBuffer.CloseContext("TelephonyDataDialInFailure"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for TelephonyDataDialInFailure")
	}

	return m, nil
}

func (m *_TelephonyDataDialInFailure) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_TelephonyDataDialInFailure) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("TelephonyDataDialInFailure"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for TelephonyDataDialInFailure")
		}

		if err := WriteSimpleEnumField[DialInFailureReason](ctx, "reason", "DialInFailureReason", m.GetReason(), WriteEnum[DialInFailureReason, uint8](DialInFailureReason.GetValue, DialInFailureReason.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
			return errors.Wrap(err, "Error serializing 'reason' field")
		}

		if popErr := writeBuffer.PopContext("TelephonyDataDialInFailure"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for TelephonyDataDialInFailure")
		}
		return nil
	}
	return m.TelephonyDataContract.(*_TelephonyData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_TelephonyDataDialInFailure) IsTelephonyDataDialInFailure() {}

func (m *_TelephonyDataDialInFailure) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
