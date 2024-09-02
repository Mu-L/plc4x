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

// CipRRData is the corresponding interface of CipRRData
type CipRRData interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	EipPacket
	// GetInterfaceHandle returns InterfaceHandle (property field)
	GetInterfaceHandle() uint32
	// GetTimeout returns Timeout (property field)
	GetTimeout() uint16
	// GetTypeIds returns TypeIds (property field)
	GetTypeIds() []TypeId
}

// CipRRDataExactly can be used when we want exactly this type and not a type which fulfills CipRRData.
// This is useful for switch cases.
type CipRRDataExactly interface {
	CipRRData
	isCipRRData() bool
}

// _CipRRData is the data-structure of this message
type _CipRRData struct {
	EipPacketContract
	InterfaceHandle uint32
	Timeout         uint16
	TypeIds         []TypeId
}

var _ CipRRData = (*_CipRRData)(nil)
var _ EipPacketRequirements = (*_CipRRData)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_CipRRData) GetCommand() uint16 {
	return 0x006F
}

func (m *_CipRRData) GetResponse() bool {
	return false
}

func (m *_CipRRData) GetPacketLength() uint16 {
	return 0
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CipRRData) GetParent() EipPacketContract {
	return m.EipPacketContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CipRRData) GetInterfaceHandle() uint32 {
	return m.InterfaceHandle
}

func (m *_CipRRData) GetTimeout() uint16 {
	return m.Timeout
}

func (m *_CipRRData) GetTypeIds() []TypeId {
	return m.TypeIds
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCipRRData factory function for _CipRRData
func NewCipRRData(interfaceHandle uint32, timeout uint16, typeIds []TypeId, sessionHandle uint32, status uint32, senderContext []byte, options uint32) *_CipRRData {
	_result := &_CipRRData{
		EipPacketContract: NewEipPacket(sessionHandle, status, senderContext, options),
		InterfaceHandle:   interfaceHandle,
		Timeout:           timeout,
		TypeIds:           typeIds,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastCipRRData(structType any) CipRRData {
	if casted, ok := structType.(CipRRData); ok {
		return casted
	}
	if casted, ok := structType.(*CipRRData); ok {
		return *casted
	}
	return nil
}

func (m *_CipRRData) GetTypeName() string {
	return "CipRRData"
}

func (m *_CipRRData) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.EipPacketContract.(*_EipPacket).getLengthInBits(ctx))

	// Simple field (interfaceHandle)
	lengthInBits += 32

	// Simple field (timeout)
	lengthInBits += 16

	// Implicit Field (typeIdCount)
	lengthInBits += 16

	// Array field
	if len(m.TypeIds) > 0 {
		for _curItem, element := range m.TypeIds {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.TypeIds), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_CipRRData) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_CipRRData) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_EipPacket, response bool) (__cipRRData CipRRData, err error) {
	m.EipPacketContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CipRRData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CipRRData")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	interfaceHandle, err := ReadSimpleField(ctx, "interfaceHandle", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'interfaceHandle' field"))
	}
	m.InterfaceHandle = interfaceHandle

	timeout, err := ReadSimpleField(ctx, "timeout", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timeout' field"))
	}
	m.Timeout = timeout

	typeIdCount, err := ReadImplicitField[uint16](ctx, "typeIdCount", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'typeIdCount' field"))
	}
	_ = typeIdCount

	typeIds, err := ReadCountArrayField[TypeId](ctx, "typeIds", ReadComplex[TypeId](TypeIdParseWithBuffer, readBuffer), uint64(typeIdCount))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'typeIds' field"))
	}
	m.TypeIds = typeIds

	if closeErr := readBuffer.CloseContext("CipRRData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CipRRData")
	}

	return m, nil
}

func (m *_CipRRData) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CipRRData) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CipRRData"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CipRRData")
		}

		if err := WriteSimpleField[uint32](ctx, "interfaceHandle", m.GetInterfaceHandle(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'interfaceHandle' field")
		}

		if err := WriteSimpleField[uint16](ctx, "timeout", m.GetTimeout(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'timeout' field")
		}
		typeIdCount := uint16(uint16(len(m.GetTypeIds())))
		if err := WriteImplicitField(ctx, "typeIdCount", typeIdCount, WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'typeIdCount' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "typeIds", m.GetTypeIds(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'typeIds' field")
		}

		if popErr := writeBuffer.PopContext("CipRRData"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CipRRData")
		}
		return nil
	}
	return m.EipPacketContract.(*_EipPacket).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CipRRData) isCipRRData() bool {
	return true
}

func (m *_CipRRData) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
