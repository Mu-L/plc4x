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

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// BVLCReadBroadcastDistributionTable is the corresponding interface of BVLCReadBroadcastDistributionTable
type BVLCReadBroadcastDistributionTable interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BVLC
}

// BVLCReadBroadcastDistributionTableExactly can be used when we want exactly this type and not a type which fulfills BVLCReadBroadcastDistributionTable.
// This is useful for switch cases.
type BVLCReadBroadcastDistributionTableExactly interface {
	BVLCReadBroadcastDistributionTable
	isBVLCReadBroadcastDistributionTable() bool
}

// _BVLCReadBroadcastDistributionTable is the data-structure of this message
type _BVLCReadBroadcastDistributionTable struct {
	BVLCContract
}

var _ BVLCReadBroadcastDistributionTable = (*_BVLCReadBroadcastDistributionTable)(nil)
var _ BVLCRequirements = (*_BVLCReadBroadcastDistributionTable)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BVLCReadBroadcastDistributionTable) GetBvlcFunction() uint8 {
	return 0x02
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BVLCReadBroadcastDistributionTable) GetParent() BVLCContract {
	return m.BVLCContract
}

// NewBVLCReadBroadcastDistributionTable factory function for _BVLCReadBroadcastDistributionTable
func NewBVLCReadBroadcastDistributionTable() *_BVLCReadBroadcastDistributionTable {
	_result := &_BVLCReadBroadcastDistributionTable{
		BVLCContract: NewBVLC(),
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastBVLCReadBroadcastDistributionTable(structType any) BVLCReadBroadcastDistributionTable {
	if casted, ok := structType.(BVLCReadBroadcastDistributionTable); ok {
		return casted
	}
	if casted, ok := structType.(*BVLCReadBroadcastDistributionTable); ok {
		return *casted
	}
	return nil
}

func (m *_BVLCReadBroadcastDistributionTable) GetTypeName() string {
	return "BVLCReadBroadcastDistributionTable"
}

func (m *_BVLCReadBroadcastDistributionTable) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BVLCContract.(*_BVLC).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_BVLCReadBroadcastDistributionTable) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BVLCReadBroadcastDistributionTable) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BVLC) (__bVLCReadBroadcastDistributionTable BVLCReadBroadcastDistributionTable, err error) {
	m.BVLCContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BVLCReadBroadcastDistributionTable"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BVLCReadBroadcastDistributionTable")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("BVLCReadBroadcastDistributionTable"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BVLCReadBroadcastDistributionTable")
	}

	return m, nil
}

func (m *_BVLCReadBroadcastDistributionTable) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))), utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BVLCReadBroadcastDistributionTable) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BVLCReadBroadcastDistributionTable"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BVLCReadBroadcastDistributionTable")
		}

		if popErr := writeBuffer.PopContext("BVLCReadBroadcastDistributionTable"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BVLCReadBroadcastDistributionTable")
		}
		return nil
	}
	return m.BVLCContract.(*_BVLC).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BVLCReadBroadcastDistributionTable) isBVLCReadBroadcastDistributionTable() bool {
	return true
}

func (m *_BVLCReadBroadcastDistributionTable) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
