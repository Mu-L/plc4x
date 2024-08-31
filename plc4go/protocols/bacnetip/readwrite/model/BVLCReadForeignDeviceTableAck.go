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

// BVLCReadForeignDeviceTableAck is the corresponding interface of BVLCReadForeignDeviceTableAck
type BVLCReadForeignDeviceTableAck interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BVLC
	// GetTable returns Table (property field)
	GetTable() []BVLCForeignDeviceTableEntry
}

// BVLCReadForeignDeviceTableAckExactly can be used when we want exactly this type and not a type which fulfills BVLCReadForeignDeviceTableAck.
// This is useful for switch cases.
type BVLCReadForeignDeviceTableAckExactly interface {
	BVLCReadForeignDeviceTableAck
	isBVLCReadForeignDeviceTableAck() bool
}

// _BVLCReadForeignDeviceTableAck is the data-structure of this message
type _BVLCReadForeignDeviceTableAck struct {
	*_BVLC
	Table []BVLCForeignDeviceTableEntry

	// Arguments.
	BvlcPayloadLength uint16
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BVLCReadForeignDeviceTableAck) GetBvlcFunction() uint8 {
	return 0x07
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BVLCReadForeignDeviceTableAck) InitializeParent(parent BVLC) {}

func (m *_BVLCReadForeignDeviceTableAck) GetParent() BVLC {
	return m._BVLC
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BVLCReadForeignDeviceTableAck) GetTable() []BVLCForeignDeviceTableEntry {
	return m.Table
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBVLCReadForeignDeviceTableAck factory function for _BVLCReadForeignDeviceTableAck
func NewBVLCReadForeignDeviceTableAck(table []BVLCForeignDeviceTableEntry, bvlcPayloadLength uint16) *_BVLCReadForeignDeviceTableAck {
	_result := &_BVLCReadForeignDeviceTableAck{
		Table: table,
		_BVLC: NewBVLC(),
	}
	_result._BVLC._BVLCChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBVLCReadForeignDeviceTableAck(structType any) BVLCReadForeignDeviceTableAck {
	if casted, ok := structType.(BVLCReadForeignDeviceTableAck); ok {
		return casted
	}
	if casted, ok := structType.(*BVLCReadForeignDeviceTableAck); ok {
		return *casted
	}
	return nil
}

func (m *_BVLCReadForeignDeviceTableAck) GetTypeName() string {
	return "BVLCReadForeignDeviceTableAck"
}

func (m *_BVLCReadForeignDeviceTableAck) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Array field
	if len(m.Table) > 0 {
		for _, element := range m.Table {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BVLCReadForeignDeviceTableAck) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BVLCReadForeignDeviceTableAckParse(ctx context.Context, theBytes []byte, bvlcPayloadLength uint16) (BVLCReadForeignDeviceTableAck, error) {
	return BVLCReadForeignDeviceTableAckParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian)), bvlcPayloadLength)
}

func BVLCReadForeignDeviceTableAckParseWithBufferProducer(bvlcPayloadLength uint16) func(ctx context.Context, readBuffer utils.ReadBuffer) (BVLCReadForeignDeviceTableAck, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BVLCReadForeignDeviceTableAck, error) {
		return BVLCReadForeignDeviceTableAckParseWithBuffer(ctx, readBuffer, bvlcPayloadLength)
	}
}

func BVLCReadForeignDeviceTableAckParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, bvlcPayloadLength uint16) (BVLCReadForeignDeviceTableAck, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BVLCReadForeignDeviceTableAck"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BVLCReadForeignDeviceTableAck")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	table, err := ReadLengthArrayField[BVLCForeignDeviceTableEntry](ctx, "table", ReadComplex[BVLCForeignDeviceTableEntry](BVLCForeignDeviceTableEntryParseWithBuffer, readBuffer), int(bvlcPayloadLength), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'table' field"))
	}

	if closeErr := readBuffer.CloseContext("BVLCReadForeignDeviceTableAck"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BVLCReadForeignDeviceTableAck")
	}

	// Create a partially initialized instance
	_child := &_BVLCReadForeignDeviceTableAck{
		_BVLC: &_BVLC{},
		Table: table,
	}
	_child._BVLC._BVLCChildRequirements = _child
	return _child, nil
}

func (m *_BVLCReadForeignDeviceTableAck) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))), utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BVLCReadForeignDeviceTableAck) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BVLCReadForeignDeviceTableAck"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BVLCReadForeignDeviceTableAck")
		}

		// Array Field (table)
		if pushErr := writeBuffer.PushContext("table", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for table")
		}
		for _curItem, _element := range m.GetTable() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetTable()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'table' field")
			}
		}
		if popErr := writeBuffer.PopContext("table", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for table")
		}

		if popErr := writeBuffer.PopContext("BVLCReadForeignDeviceTableAck"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BVLCReadForeignDeviceTableAck")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

////
// Arguments Getter

func (m *_BVLCReadForeignDeviceTableAck) GetBvlcPayloadLength() uint16 {
	return m.BvlcPayloadLength
}

//
////

func (m *_BVLCReadForeignDeviceTableAck) isBVLCReadForeignDeviceTableAck() bool {
	return true
}

func (m *_BVLCReadForeignDeviceTableAck) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
