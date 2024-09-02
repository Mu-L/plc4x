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

// NLMInitializeRoutingTable is the corresponding interface of NLMInitializeRoutingTable
type NLMInitializeRoutingTable interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	NLM
	// GetNumberOfPorts returns NumberOfPorts (property field)
	GetNumberOfPorts() uint8
	// GetPortMappings returns PortMappings (property field)
	GetPortMappings() []NLMInitializeRoutingTablePortMapping
}

// NLMInitializeRoutingTableExactly can be used when we want exactly this type and not a type which fulfills NLMInitializeRoutingTable.
// This is useful for switch cases.
type NLMInitializeRoutingTableExactly interface {
	NLMInitializeRoutingTable
	isNLMInitializeRoutingTable() bool
}

// _NLMInitializeRoutingTable is the data-structure of this message
type _NLMInitializeRoutingTable struct {
	NLMContract
	NumberOfPorts uint8
	PortMappings  []NLMInitializeRoutingTablePortMapping
}

var _ NLMInitializeRoutingTable = (*_NLMInitializeRoutingTable)(nil)
var _ NLMRequirements = (*_NLMInitializeRoutingTable)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_NLMInitializeRoutingTable) GetMessageType() uint8 {
	return 0x06
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_NLMInitializeRoutingTable) GetParent() NLMContract {
	return m.NLMContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_NLMInitializeRoutingTable) GetNumberOfPorts() uint8 {
	return m.NumberOfPorts
}

func (m *_NLMInitializeRoutingTable) GetPortMappings() []NLMInitializeRoutingTablePortMapping {
	return m.PortMappings
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewNLMInitializeRoutingTable factory function for _NLMInitializeRoutingTable
func NewNLMInitializeRoutingTable(numberOfPorts uint8, portMappings []NLMInitializeRoutingTablePortMapping, apduLength uint16) *_NLMInitializeRoutingTable {
	_result := &_NLMInitializeRoutingTable{
		NLMContract:   NewNLM(apduLength),
		NumberOfPorts: numberOfPorts,
		PortMappings:  portMappings,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastNLMInitializeRoutingTable(structType any) NLMInitializeRoutingTable {
	if casted, ok := structType.(NLMInitializeRoutingTable); ok {
		return casted
	}
	if casted, ok := structType.(*NLMInitializeRoutingTable); ok {
		return *casted
	}
	return nil
}

func (m *_NLMInitializeRoutingTable) GetTypeName() string {
	return "NLMInitializeRoutingTable"
}

func (m *_NLMInitializeRoutingTable) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.NLMContract.(*_NLM).getLengthInBits(ctx))

	// Simple field (numberOfPorts)
	lengthInBits += 8

	// Array field
	if len(m.PortMappings) > 0 {
		for _curItem, element := range m.PortMappings {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.PortMappings), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_NLMInitializeRoutingTable) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_NLMInitializeRoutingTable) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_NLM, apduLength uint16) (__nLMInitializeRoutingTable NLMInitializeRoutingTable, err error) {
	m.NLMContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("NLMInitializeRoutingTable"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NLMInitializeRoutingTable")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	numberOfPorts, err := ReadSimpleField(ctx, "numberOfPorts", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'numberOfPorts' field"))
	}
	m.NumberOfPorts = numberOfPorts

	portMappings, err := ReadCountArrayField[NLMInitializeRoutingTablePortMapping](ctx, "portMappings", ReadComplex[NLMInitializeRoutingTablePortMapping](NLMInitializeRoutingTablePortMappingParseWithBuffer, readBuffer), uint64(numberOfPorts))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'portMappings' field"))
	}
	m.PortMappings = portMappings

	if closeErr := readBuffer.CloseContext("NLMInitializeRoutingTable"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NLMInitializeRoutingTable")
	}

	return m, nil
}

func (m *_NLMInitializeRoutingTable) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_NLMInitializeRoutingTable) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("NLMInitializeRoutingTable"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for NLMInitializeRoutingTable")
		}

		if err := WriteSimpleField[uint8](ctx, "numberOfPorts", m.GetNumberOfPorts(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'numberOfPorts' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "portMappings", m.GetPortMappings(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'portMappings' field")
		}

		if popErr := writeBuffer.PopContext("NLMInitializeRoutingTable"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for NLMInitializeRoutingTable")
		}
		return nil
	}
	return m.NLMContract.(*_NLM).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_NLMInitializeRoutingTable) isNLMInitializeRoutingTable() bool {
	return true
}

func (m *_NLMInitializeRoutingTable) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
