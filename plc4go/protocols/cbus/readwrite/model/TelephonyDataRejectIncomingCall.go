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

// TelephonyDataRejectIncomingCall is the corresponding interface of TelephonyDataRejectIncomingCall
type TelephonyDataRejectIncomingCall interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	TelephonyData
	// IsTelephonyDataRejectIncomingCall is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsTelephonyDataRejectIncomingCall()
}

// _TelephonyDataRejectIncomingCall is the data-structure of this message
type _TelephonyDataRejectIncomingCall struct {
	TelephonyDataContract
}

var _ TelephonyDataRejectIncomingCall = (*_TelephonyDataRejectIncomingCall)(nil)
var _ TelephonyDataRequirements = (*_TelephonyDataRejectIncomingCall)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_TelephonyDataRejectIncomingCall) GetParent() TelephonyDataContract {
	return m.TelephonyDataContract
}

// NewTelephonyDataRejectIncomingCall factory function for _TelephonyDataRejectIncomingCall
func NewTelephonyDataRejectIncomingCall(commandTypeContainer TelephonyCommandTypeContainer, argument byte) *_TelephonyDataRejectIncomingCall {
	_result := &_TelephonyDataRejectIncomingCall{
		TelephonyDataContract: NewTelephonyData(commandTypeContainer, argument),
	}
	_result.TelephonyDataContract.(*_TelephonyData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastTelephonyDataRejectIncomingCall(structType any) TelephonyDataRejectIncomingCall {
	if casted, ok := structType.(TelephonyDataRejectIncomingCall); ok {
		return casted
	}
	if casted, ok := structType.(*TelephonyDataRejectIncomingCall); ok {
		return *casted
	}
	return nil
}

func (m *_TelephonyDataRejectIncomingCall) GetTypeName() string {
	return "TelephonyDataRejectIncomingCall"
}

func (m *_TelephonyDataRejectIncomingCall) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.TelephonyDataContract.(*_TelephonyData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_TelephonyDataRejectIncomingCall) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_TelephonyDataRejectIncomingCall) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_TelephonyData) (__telephonyDataRejectIncomingCall TelephonyDataRejectIncomingCall, err error) {
	m.TelephonyDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("TelephonyDataRejectIncomingCall"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for TelephonyDataRejectIncomingCall")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("TelephonyDataRejectIncomingCall"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for TelephonyDataRejectIncomingCall")
	}

	return m, nil
}

func (m *_TelephonyDataRejectIncomingCall) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_TelephonyDataRejectIncomingCall) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("TelephonyDataRejectIncomingCall"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for TelephonyDataRejectIncomingCall")
		}

		if popErr := writeBuffer.PopContext("TelephonyDataRejectIncomingCall"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for TelephonyDataRejectIncomingCall")
		}
		return nil
	}
	return m.TelephonyDataContract.(*_TelephonyData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_TelephonyDataRejectIncomingCall) IsTelephonyDataRejectIncomingCall() {}

func (m *_TelephonyDataRejectIncomingCall) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
