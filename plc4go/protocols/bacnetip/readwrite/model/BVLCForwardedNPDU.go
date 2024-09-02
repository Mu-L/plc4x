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
	"encoding/binary"
	"fmt"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	"github.com/apache/plc4x/plc4go/spi/codegen"
	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// BVLCForwardedNPDU is the corresponding interface of BVLCForwardedNPDU
type BVLCForwardedNPDU interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BVLC
	// GetIp returns Ip (property field)
	GetIp() []uint8
	// GetPort returns Port (property field)
	GetPort() uint16
	// GetNpdu returns Npdu (property field)
	GetNpdu() NPDU
}

// BVLCForwardedNPDUExactly can be used when we want exactly this type and not a type which fulfills BVLCForwardedNPDU.
// This is useful for switch cases.
type BVLCForwardedNPDUExactly interface {
	BVLCForwardedNPDU
	isBVLCForwardedNPDU() bool
}

// _BVLCForwardedNPDU is the data-structure of this message
type _BVLCForwardedNPDU struct {
	BVLCContract
	Ip   []uint8
	Port uint16
	Npdu NPDU

	// Arguments.
	BvlcPayloadLength uint16
}

var _ BVLCForwardedNPDU = (*_BVLCForwardedNPDU)(nil)
var _ BVLCRequirements = (*_BVLCForwardedNPDU)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BVLCForwardedNPDU) GetBvlcFunction() uint8 {
	return 0x04
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BVLCForwardedNPDU) GetParent() BVLCContract {
	return m.BVLCContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BVLCForwardedNPDU) GetIp() []uint8 {
	return m.Ip
}

func (m *_BVLCForwardedNPDU) GetPort() uint16 {
	return m.Port
}

func (m *_BVLCForwardedNPDU) GetNpdu() NPDU {
	return m.Npdu
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBVLCForwardedNPDU factory function for _BVLCForwardedNPDU
func NewBVLCForwardedNPDU(ip []uint8, port uint16, npdu NPDU, bvlcPayloadLength uint16) *_BVLCForwardedNPDU {
	_result := &_BVLCForwardedNPDU{
		BVLCContract: NewBVLC(),
		Ip:           ip,
		Port:         port,
		Npdu:         npdu,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastBVLCForwardedNPDU(structType any) BVLCForwardedNPDU {
	if casted, ok := structType.(BVLCForwardedNPDU); ok {
		return casted
	}
	if casted, ok := structType.(*BVLCForwardedNPDU); ok {
		return *casted
	}
	return nil
}

func (m *_BVLCForwardedNPDU) GetTypeName() string {
	return "BVLCForwardedNPDU"
}

func (m *_BVLCForwardedNPDU) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BVLCContract.(*_BVLC).getLengthInBits(ctx))

	// Array field
	if len(m.Ip) > 0 {
		lengthInBits += 8 * uint16(len(m.Ip))
	}

	// Simple field (port)
	lengthInBits += 16

	// Simple field (npdu)
	lengthInBits += m.Npdu.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BVLCForwardedNPDU) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BVLCForwardedNPDU) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BVLC, bvlcPayloadLength uint16) (__bVLCForwardedNPDU BVLCForwardedNPDU, err error) {
	m.BVLCContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BVLCForwardedNPDU"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BVLCForwardedNPDU")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	ip, err := ReadCountArrayField[uint8](ctx, "ip", ReadUnsignedByte(readBuffer, uint8(8)), uint64(int32(4)), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'ip' field"))
	}
	m.Ip = ip

	port, err := ReadSimpleField(ctx, "port", ReadUnsignedShort(readBuffer, uint8(16)), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'port' field"))
	}
	m.Port = port

	npdu, err := ReadSimpleField[NPDU](ctx, "npdu", ReadComplex[NPDU](NPDUParseWithBufferProducer((uint16)(uint16(bvlcPayloadLength)-uint16(uint16(6)))), readBuffer), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'npdu' field"))
	}
	m.Npdu = npdu

	if closeErr := readBuffer.CloseContext("BVLCForwardedNPDU"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BVLCForwardedNPDU")
	}

	return m, nil
}

func (m *_BVLCForwardedNPDU) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))), utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BVLCForwardedNPDU) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BVLCForwardedNPDU"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BVLCForwardedNPDU")
		}

		if err := WriteSimpleTypeArrayField(ctx, "ip", m.GetIp(), WriteUnsignedByte(writeBuffer, 8), codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'ip' field")
		}

		if err := WriteSimpleField[uint16](ctx, "port", m.GetPort(), WriteUnsignedShort(writeBuffer, 16), codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'port' field")
		}

		if err := WriteSimpleField[NPDU](ctx, "npdu", m.GetNpdu(), WriteComplex[NPDU](writeBuffer), codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'npdu' field")
		}

		if popErr := writeBuffer.PopContext("BVLCForwardedNPDU"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BVLCForwardedNPDU")
		}
		return nil
	}
	return m.BVLCContract.(*_BVLC).serializeParent(ctx, writeBuffer, m, ser)
}

////
// Arguments Getter

func (m *_BVLCForwardedNPDU) GetBvlcPayloadLength() uint16 {
	return m.BvlcPayloadLength
}

//
////

func (m *_BVLCForwardedNPDU) isBVLCForwardedNPDU() bool {
	return true
}

func (m *_BVLCForwardedNPDU) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
