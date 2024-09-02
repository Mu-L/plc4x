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

// UnConnectedDataItem is the corresponding interface of UnConnectedDataItem
type UnConnectedDataItem interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	TypeId
	// GetService returns Service (property field)
	GetService() CipService
	// IsUnConnectedDataItem is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsUnConnectedDataItem()
}

// _UnConnectedDataItem is the data-structure of this message
type _UnConnectedDataItem struct {
	TypeIdContract
	Service CipService
}

var _ UnConnectedDataItem = (*_UnConnectedDataItem)(nil)
var _ TypeIdRequirements = (*_UnConnectedDataItem)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_UnConnectedDataItem) GetId() uint16 {
	return 0x00B2
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_UnConnectedDataItem) GetParent() TypeIdContract {
	return m.TypeIdContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_UnConnectedDataItem) GetService() CipService {
	return m.Service
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewUnConnectedDataItem factory function for _UnConnectedDataItem
func NewUnConnectedDataItem(service CipService) *_UnConnectedDataItem {
	_result := &_UnConnectedDataItem{
		TypeIdContract: NewTypeId(),
		Service:        service,
	}
	_result.TypeIdContract.(*_TypeId)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastUnConnectedDataItem(structType any) UnConnectedDataItem {
	if casted, ok := structType.(UnConnectedDataItem); ok {
		return casted
	}
	if casted, ok := structType.(*UnConnectedDataItem); ok {
		return *casted
	}
	return nil
}

func (m *_UnConnectedDataItem) GetTypeName() string {
	return "UnConnectedDataItem"
}

func (m *_UnConnectedDataItem) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.TypeIdContract.(*_TypeId).getLengthInBits(ctx))

	// Implicit Field (packetSize)
	lengthInBits += 16

	// Simple field (service)
	lengthInBits += m.Service.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_UnConnectedDataItem) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_UnConnectedDataItem) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_TypeId) (__unConnectedDataItem UnConnectedDataItem, err error) {
	m.TypeIdContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("UnConnectedDataItem"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for UnConnectedDataItem")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	packetSize, err := ReadImplicitField[uint16](ctx, "packetSize", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'packetSize' field"))
	}
	_ = packetSize

	service, err := ReadSimpleField[CipService](ctx, "service", ReadComplex[CipService](CipServiceParseWithBufferProducer[CipService]((bool)(bool(false)), (uint16)(packetSize)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'service' field"))
	}
	m.Service = service

	if closeErr := readBuffer.CloseContext("UnConnectedDataItem"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for UnConnectedDataItem")
	}

	return m, nil
}

func (m *_UnConnectedDataItem) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_UnConnectedDataItem) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("UnConnectedDataItem"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for UnConnectedDataItem")
		}
		packetSize := uint16(m.GetService().GetLengthInBytes(ctx))
		if err := WriteImplicitField(ctx, "packetSize", packetSize, WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'packetSize' field")
		}

		if err := WriteSimpleField[CipService](ctx, "service", m.GetService(), WriteComplex[CipService](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'service' field")
		}

		if popErr := writeBuffer.PopContext("UnConnectedDataItem"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for UnConnectedDataItem")
		}
		return nil
	}
	return m.TypeIdContract.(*_TypeId).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_UnConnectedDataItem) IsUnConnectedDataItem() {}

func (m *_UnConnectedDataItem) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
