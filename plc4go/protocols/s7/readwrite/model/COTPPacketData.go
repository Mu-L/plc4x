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

// COTPPacketData is the corresponding interface of COTPPacketData
type COTPPacketData interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	COTPPacket
	// GetEot returns Eot (property field)
	GetEot() bool
	// GetTpduRef returns TpduRef (property field)
	GetTpduRef() uint8
}

// COTPPacketDataExactly can be used when we want exactly this type and not a type which fulfills COTPPacketData.
// This is useful for switch cases.
type COTPPacketDataExactly interface {
	COTPPacketData
	isCOTPPacketData() bool
}

// _COTPPacketData is the data-structure of this message
type _COTPPacketData struct {
	COTPPacketContract
	Eot     bool
	TpduRef uint8
}

var _ COTPPacketData = (*_COTPPacketData)(nil)
var _ COTPPacketRequirements = (*_COTPPacketData)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_COTPPacketData) GetTpduCode() uint8 {
	return 0xF0
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_COTPPacketData) GetParent() COTPPacketContract {
	return m.COTPPacketContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_COTPPacketData) GetEot() bool {
	return m.Eot
}

func (m *_COTPPacketData) GetTpduRef() uint8 {
	return m.TpduRef
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCOTPPacketData factory function for _COTPPacketData
func NewCOTPPacketData(eot bool, tpduRef uint8, parameters []COTPParameter, payload S7Message, cotpLen uint16) *_COTPPacketData {
	_result := &_COTPPacketData{
		COTPPacketContract: NewCOTPPacket(parameters, payload, cotpLen),
		Eot:                eot,
		TpduRef:            tpduRef,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastCOTPPacketData(structType any) COTPPacketData {
	if casted, ok := structType.(COTPPacketData); ok {
		return casted
	}
	if casted, ok := structType.(*COTPPacketData); ok {
		return *casted
	}
	return nil
}

func (m *_COTPPacketData) GetTypeName() string {
	return "COTPPacketData"
}

func (m *_COTPPacketData) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.COTPPacketContract.(*_COTPPacket).getLengthInBits(ctx))

	// Simple field (eot)
	lengthInBits += 1

	// Simple field (tpduRef)
	lengthInBits += 7

	return lengthInBits
}

func (m *_COTPPacketData) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_COTPPacketData) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_COTPPacket, cotpLen uint16) (__cOTPPacketData COTPPacketData, err error) {
	m.COTPPacketContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("COTPPacketData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for COTPPacketData")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	eot, err := ReadSimpleField(ctx, "eot", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'eot' field"))
	}
	m.Eot = eot

	tpduRef, err := ReadSimpleField(ctx, "tpduRef", ReadUnsignedByte(readBuffer, uint8(7)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'tpduRef' field"))
	}
	m.TpduRef = tpduRef

	if closeErr := readBuffer.CloseContext("COTPPacketData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for COTPPacketData")
	}

	return m, nil
}

func (m *_COTPPacketData) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_COTPPacketData) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("COTPPacketData"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for COTPPacketData")
		}

		if err := WriteSimpleField[bool](ctx, "eot", m.GetEot(), WriteBoolean(writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'eot' field")
		}

		if err := WriteSimpleField[uint8](ctx, "tpduRef", m.GetTpduRef(), WriteUnsignedByte(writeBuffer, 7)); err != nil {
			return errors.Wrap(err, "Error serializing 'tpduRef' field")
		}

		if popErr := writeBuffer.PopContext("COTPPacketData"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for COTPPacketData")
		}
		return nil
	}
	return m.COTPPacketContract.(*_COTPPacket).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_COTPPacketData) isCOTPPacketData() bool {
	return true
}

func (m *_COTPPacketData) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
