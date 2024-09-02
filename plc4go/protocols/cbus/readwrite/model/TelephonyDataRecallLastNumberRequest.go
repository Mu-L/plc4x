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

// TelephonyDataRecallLastNumberRequest is the corresponding interface of TelephonyDataRecallLastNumberRequest
type TelephonyDataRecallLastNumberRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	TelephonyData
	// GetRecallLastNumberType returns RecallLastNumberType (property field)
	GetRecallLastNumberType() byte
	// GetIsNumberOfLastOutgoingCall returns IsNumberOfLastOutgoingCall (virtual field)
	GetIsNumberOfLastOutgoingCall() bool
	// GetIsNumberOfLastIncomingCall returns IsNumberOfLastIncomingCall (virtual field)
	GetIsNumberOfLastIncomingCall() bool
}

// TelephonyDataRecallLastNumberRequestExactly can be used when we want exactly this type and not a type which fulfills TelephonyDataRecallLastNumberRequest.
// This is useful for switch cases.
type TelephonyDataRecallLastNumberRequestExactly interface {
	TelephonyDataRecallLastNumberRequest
	isTelephonyDataRecallLastNumberRequest() bool
}

// _TelephonyDataRecallLastNumberRequest is the data-structure of this message
type _TelephonyDataRecallLastNumberRequest struct {
	TelephonyDataContract
	RecallLastNumberType byte
}

var _ TelephonyDataRecallLastNumberRequest = (*_TelephonyDataRecallLastNumberRequest)(nil)
var _ TelephonyDataRequirements = (*_TelephonyDataRecallLastNumberRequest)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_TelephonyDataRecallLastNumberRequest) GetParent() TelephonyDataContract {
	return m.TelephonyDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_TelephonyDataRecallLastNumberRequest) GetRecallLastNumberType() byte {
	return m.RecallLastNumberType
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_TelephonyDataRecallLastNumberRequest) GetIsNumberOfLastOutgoingCall() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetRecallLastNumberType()) == (0x01)))
}

func (m *_TelephonyDataRecallLastNumberRequest) GetIsNumberOfLastIncomingCall() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetRecallLastNumberType()) == (0x02)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewTelephonyDataRecallLastNumberRequest factory function for _TelephonyDataRecallLastNumberRequest
func NewTelephonyDataRecallLastNumberRequest(recallLastNumberType byte, commandTypeContainer TelephonyCommandTypeContainer, argument byte) *_TelephonyDataRecallLastNumberRequest {
	_result := &_TelephonyDataRecallLastNumberRequest{
		TelephonyDataContract: NewTelephonyData(commandTypeContainer, argument),
		RecallLastNumberType:  recallLastNumberType,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastTelephonyDataRecallLastNumberRequest(structType any) TelephonyDataRecallLastNumberRequest {
	if casted, ok := structType.(TelephonyDataRecallLastNumberRequest); ok {
		return casted
	}
	if casted, ok := structType.(*TelephonyDataRecallLastNumberRequest); ok {
		return *casted
	}
	return nil
}

func (m *_TelephonyDataRecallLastNumberRequest) GetTypeName() string {
	return "TelephonyDataRecallLastNumberRequest"
}

func (m *_TelephonyDataRecallLastNumberRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.TelephonyDataContract.(*_TelephonyData).getLengthInBits(ctx))

	// Simple field (recallLastNumberType)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_TelephonyDataRecallLastNumberRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_TelephonyDataRecallLastNumberRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_TelephonyData) (__telephonyDataRecallLastNumberRequest TelephonyDataRecallLastNumberRequest, err error) {
	m.TelephonyDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("TelephonyDataRecallLastNumberRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for TelephonyDataRecallLastNumberRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	recallLastNumberType, err := ReadSimpleField(ctx, "recallLastNumberType", ReadByte(readBuffer, 8))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'recallLastNumberType' field"))
	}
	m.RecallLastNumberType = recallLastNumberType

	isNumberOfLastOutgoingCall, err := ReadVirtualField[bool](ctx, "isNumberOfLastOutgoingCall", (*bool)(nil), bool((recallLastNumberType) == (0x01)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isNumberOfLastOutgoingCall' field"))
	}
	_ = isNumberOfLastOutgoingCall

	isNumberOfLastIncomingCall, err := ReadVirtualField[bool](ctx, "isNumberOfLastIncomingCall", (*bool)(nil), bool((recallLastNumberType) == (0x02)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isNumberOfLastIncomingCall' field"))
	}
	_ = isNumberOfLastIncomingCall

	if closeErr := readBuffer.CloseContext("TelephonyDataRecallLastNumberRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for TelephonyDataRecallLastNumberRequest")
	}

	return m, nil
}

func (m *_TelephonyDataRecallLastNumberRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_TelephonyDataRecallLastNumberRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("TelephonyDataRecallLastNumberRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for TelephonyDataRecallLastNumberRequest")
		}

		if err := WriteSimpleField[byte](ctx, "recallLastNumberType", m.GetRecallLastNumberType(), WriteByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'recallLastNumberType' field")
		}
		// Virtual field
		isNumberOfLastOutgoingCall := m.GetIsNumberOfLastOutgoingCall()
		_ = isNumberOfLastOutgoingCall
		if _isNumberOfLastOutgoingCallErr := writeBuffer.WriteVirtual(ctx, "isNumberOfLastOutgoingCall", m.GetIsNumberOfLastOutgoingCall()); _isNumberOfLastOutgoingCallErr != nil {
			return errors.Wrap(_isNumberOfLastOutgoingCallErr, "Error serializing 'isNumberOfLastOutgoingCall' field")
		}
		// Virtual field
		isNumberOfLastIncomingCall := m.GetIsNumberOfLastIncomingCall()
		_ = isNumberOfLastIncomingCall
		if _isNumberOfLastIncomingCallErr := writeBuffer.WriteVirtual(ctx, "isNumberOfLastIncomingCall", m.GetIsNumberOfLastIncomingCall()); _isNumberOfLastIncomingCallErr != nil {
			return errors.Wrap(_isNumberOfLastIncomingCallErr, "Error serializing 'isNumberOfLastIncomingCall' field")
		}

		if popErr := writeBuffer.PopContext("TelephonyDataRecallLastNumberRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for TelephonyDataRecallLastNumberRequest")
		}
		return nil
	}
	return m.TelephonyDataContract.(*_TelephonyData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_TelephonyDataRecallLastNumberRequest) isTelephonyDataRecallLastNumberRequest() bool {
	return true
}

func (m *_TelephonyDataRecallLastNumberRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
