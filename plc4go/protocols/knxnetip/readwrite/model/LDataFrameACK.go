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

// LDataFrameACK is the corresponding interface of LDataFrameACK
type LDataFrameACK interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	LDataFrame
}

// LDataFrameACKExactly can be used when we want exactly this type and not a type which fulfills LDataFrameACK.
// This is useful for switch cases.
type LDataFrameACKExactly interface {
	LDataFrameACK
	isLDataFrameACK() bool
}

// _LDataFrameACK is the data-structure of this message
type _LDataFrameACK struct {
	LDataFrameContract
}

var _ LDataFrameACK = (*_LDataFrameACK)(nil)
var _ LDataFrameRequirements = (*_LDataFrameACK)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_LDataFrameACK) GetNotAckFrame() bool {
	return bool(false)
}

func (m *_LDataFrameACK) GetPolling() bool {
	return false
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_LDataFrameACK) GetParent() LDataFrameContract {
	return m.LDataFrameContract
}

// NewLDataFrameACK factory function for _LDataFrameACK
func NewLDataFrameACK(frameType bool, notRepeated bool, priority CEMIPriority, acknowledgeRequested bool, errorFlag bool) *_LDataFrameACK {
	_result := &_LDataFrameACK{
		LDataFrameContract: NewLDataFrame(frameType, notRepeated, priority, acknowledgeRequested, errorFlag),
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastLDataFrameACK(structType any) LDataFrameACK {
	if casted, ok := structType.(LDataFrameACK); ok {
		return casted
	}
	if casted, ok := structType.(*LDataFrameACK); ok {
		return *casted
	}
	return nil
}

func (m *_LDataFrameACK) GetTypeName() string {
	return "LDataFrameACK"
}

func (m *_LDataFrameACK) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.LDataFrameContract.(*_LDataFrame).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_LDataFrameACK) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_LDataFrameACK) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_LDataFrame) (__lDataFrameACK LDataFrameACK, err error) {
	m.LDataFrameContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("LDataFrameACK"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for LDataFrameACK")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("LDataFrameACK"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for LDataFrameACK")
	}

	return m, nil
}

func (m *_LDataFrameACK) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_LDataFrameACK) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("LDataFrameACK"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for LDataFrameACK")
		}

		if popErr := writeBuffer.PopContext("LDataFrameACK"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for LDataFrameACK")
		}
		return nil
	}
	return m.LDataFrameContract.(*_LDataFrame).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_LDataFrameACK) isLDataFrameACK() bool {
	return true
}

func (m *_LDataFrameACK) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
