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

// ComObjectTableRealisationType1 is the corresponding interface of ComObjectTableRealisationType1
type ComObjectTableRealisationType1 interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ComObjectTable
	// GetNumEntries returns NumEntries (property field)
	GetNumEntries() uint8
	// GetRamFlagsTablePointer returns RamFlagsTablePointer (property field)
	GetRamFlagsTablePointer() uint8
	// GetComObjectDescriptors returns ComObjectDescriptors (property field)
	GetComObjectDescriptors() []GroupObjectDescriptorRealisationType1
}

// ComObjectTableRealisationType1Exactly can be used when we want exactly this type and not a type which fulfills ComObjectTableRealisationType1.
// This is useful for switch cases.
type ComObjectTableRealisationType1Exactly interface {
	ComObjectTableRealisationType1
	isComObjectTableRealisationType1() bool
}

// _ComObjectTableRealisationType1 is the data-structure of this message
type _ComObjectTableRealisationType1 struct {
	*_ComObjectTable
	NumEntries           uint8
	RamFlagsTablePointer uint8
	ComObjectDescriptors []GroupObjectDescriptorRealisationType1
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ComObjectTableRealisationType1) GetFirmwareType() FirmwareType {
	return FirmwareType_SYSTEM_1
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ComObjectTableRealisationType1) InitializeParent(parent ComObjectTable) {}

func (m *_ComObjectTableRealisationType1) GetParent() ComObjectTable {
	return m._ComObjectTable
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ComObjectTableRealisationType1) GetNumEntries() uint8 {
	return m.NumEntries
}

func (m *_ComObjectTableRealisationType1) GetRamFlagsTablePointer() uint8 {
	return m.RamFlagsTablePointer
}

func (m *_ComObjectTableRealisationType1) GetComObjectDescriptors() []GroupObjectDescriptorRealisationType1 {
	return m.ComObjectDescriptors
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewComObjectTableRealisationType1 factory function for _ComObjectTableRealisationType1
func NewComObjectTableRealisationType1(numEntries uint8, ramFlagsTablePointer uint8, comObjectDescriptors []GroupObjectDescriptorRealisationType1) *_ComObjectTableRealisationType1 {
	_result := &_ComObjectTableRealisationType1{
		NumEntries:           numEntries,
		RamFlagsTablePointer: ramFlagsTablePointer,
		ComObjectDescriptors: comObjectDescriptors,
		_ComObjectTable:      NewComObjectTable(),
	}
	_result._ComObjectTable._ComObjectTableChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastComObjectTableRealisationType1(structType any) ComObjectTableRealisationType1 {
	if casted, ok := structType.(ComObjectTableRealisationType1); ok {
		return casted
	}
	if casted, ok := structType.(*ComObjectTableRealisationType1); ok {
		return *casted
	}
	return nil
}

func (m *_ComObjectTableRealisationType1) GetTypeName() string {
	return "ComObjectTableRealisationType1"
}

func (m *_ComObjectTableRealisationType1) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (numEntries)
	lengthInBits += 8

	// Simple field (ramFlagsTablePointer)
	lengthInBits += 8

	// Array field
	if len(m.ComObjectDescriptors) > 0 {
		for _curItem, element := range m.ComObjectDescriptors {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.ComObjectDescriptors), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_ComObjectTableRealisationType1) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ComObjectTableRealisationType1Parse(ctx context.Context, theBytes []byte, firmwareType FirmwareType) (ComObjectTableRealisationType1, error) {
	return ComObjectTableRealisationType1ParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), firmwareType)
}

func ComObjectTableRealisationType1ParseWithBufferProducer(firmwareType FirmwareType) func(ctx context.Context, readBuffer utils.ReadBuffer) (ComObjectTableRealisationType1, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (ComObjectTableRealisationType1, error) {
		return ComObjectTableRealisationType1ParseWithBuffer(ctx, readBuffer, firmwareType)
	}
}

func ComObjectTableRealisationType1ParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, firmwareType FirmwareType) (ComObjectTableRealisationType1, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("ComObjectTableRealisationType1"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ComObjectTableRealisationType1")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	numEntries, err := ReadSimpleField(ctx, "numEntries", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'numEntries' field"))
	}

	ramFlagsTablePointer, err := ReadSimpleField(ctx, "ramFlagsTablePointer", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'ramFlagsTablePointer' field"))
	}

	comObjectDescriptors, err := ReadCountArrayField[GroupObjectDescriptorRealisationType1](ctx, "comObjectDescriptors", ReadComplex[GroupObjectDescriptorRealisationType1](GroupObjectDescriptorRealisationType1ParseWithBuffer, readBuffer), uint64(numEntries))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'comObjectDescriptors' field"))
	}

	if closeErr := readBuffer.CloseContext("ComObjectTableRealisationType1"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ComObjectTableRealisationType1")
	}

	// Create a partially initialized instance
	_child := &_ComObjectTableRealisationType1{
		_ComObjectTable:      &_ComObjectTable{},
		NumEntries:           numEntries,
		RamFlagsTablePointer: ramFlagsTablePointer,
		ComObjectDescriptors: comObjectDescriptors,
	}
	_child._ComObjectTable._ComObjectTableChildRequirements = _child
	return _child, nil
}

func (m *_ComObjectTableRealisationType1) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ComObjectTableRealisationType1) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ComObjectTableRealisationType1"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ComObjectTableRealisationType1")
		}

		// Simple Field (numEntries)
		numEntries := uint8(m.GetNumEntries())
		_numEntriesErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("numEntries", 8, uint8((numEntries)))
		if _numEntriesErr != nil {
			return errors.Wrap(_numEntriesErr, "Error serializing 'numEntries' field")
		}

		// Simple Field (ramFlagsTablePointer)
		ramFlagsTablePointer := uint8(m.GetRamFlagsTablePointer())
		_ramFlagsTablePointerErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("ramFlagsTablePointer", 8, uint8((ramFlagsTablePointer)))
		if _ramFlagsTablePointerErr != nil {
			return errors.Wrap(_ramFlagsTablePointerErr, "Error serializing 'ramFlagsTablePointer' field")
		}

		// Array Field (comObjectDescriptors)
		if pushErr := writeBuffer.PushContext("comObjectDescriptors", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for comObjectDescriptors")
		}
		for _curItem, _element := range m.GetComObjectDescriptors() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetComObjectDescriptors()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'comObjectDescriptors' field")
			}
		}
		if popErr := writeBuffer.PopContext("comObjectDescriptors", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for comObjectDescriptors")
		}

		if popErr := writeBuffer.PopContext("ComObjectTableRealisationType1"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ComObjectTableRealisationType1")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ComObjectTableRealisationType1) isComObjectTableRealisationType1() bool {
	return true
}

func (m *_ComObjectTableRealisationType1) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
