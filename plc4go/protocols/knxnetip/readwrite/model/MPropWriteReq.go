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

// MPropWriteReq is the corresponding interface of MPropWriteReq
type MPropWriteReq interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	CEMI
}

// MPropWriteReqExactly can be used when we want exactly this type and not a type which fulfills MPropWriteReq.
// This is useful for switch cases.
type MPropWriteReqExactly interface {
	MPropWriteReq
	isMPropWriteReq() bool
}

// _MPropWriteReq is the data-structure of this message
type _MPropWriteReq struct {
	CEMIContract
}

var _ MPropWriteReq = (*_MPropWriteReq)(nil)
var _ CEMIRequirements = (*_MPropWriteReq)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_MPropWriteReq) GetMessageCode() uint8 {
	return 0xF6
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_MPropWriteReq) GetParent() CEMIContract {
	return m.CEMIContract
}

// NewMPropWriteReq factory function for _MPropWriteReq
func NewMPropWriteReq(size uint16) *_MPropWriteReq {
	_result := &_MPropWriteReq{
		CEMIContract: NewCEMI(size),
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastMPropWriteReq(structType any) MPropWriteReq {
	if casted, ok := structType.(MPropWriteReq); ok {
		return casted
	}
	if casted, ok := structType.(*MPropWriteReq); ok {
		return *casted
	}
	return nil
}

func (m *_MPropWriteReq) GetTypeName() string {
	return "MPropWriteReq"
}

func (m *_MPropWriteReq) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.CEMIContract.(*_CEMI).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_MPropWriteReq) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_MPropWriteReq) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_CEMI, size uint16) (__mPropWriteReq MPropWriteReq, err error) {
	m.CEMIContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MPropWriteReq"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MPropWriteReq")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("MPropWriteReq"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MPropWriteReq")
	}

	return m, nil
}

func (m *_MPropWriteReq) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MPropWriteReq) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MPropWriteReq"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MPropWriteReq")
		}

		if popErr := writeBuffer.PopContext("MPropWriteReq"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MPropWriteReq")
		}
		return nil
	}
	return m.CEMIContract.(*_CEMI).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_MPropWriteReq) isMPropWriteReq() bool {
	return true
}

func (m *_MPropWriteReq) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
