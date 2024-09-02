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

// LRawReq is the corresponding interface of LRawReq
type LRawReq interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	CEMI
}

// LRawReqExactly can be used when we want exactly this type and not a type which fulfills LRawReq.
// This is useful for switch cases.
type LRawReqExactly interface {
	LRawReq
	isLRawReq() bool
}

// _LRawReq is the data-structure of this message
type _LRawReq struct {
	CEMIContract
}

var _ LRawReq = (*_LRawReq)(nil)
var _ CEMIRequirements = (*_LRawReq)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_LRawReq) GetMessageCode() uint8 {
	return 0x10
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_LRawReq) GetParent() CEMIContract {
	return m.CEMIContract
}

// NewLRawReq factory function for _LRawReq
func NewLRawReq(size uint16) *_LRawReq {
	_result := &_LRawReq{
		CEMIContract: NewCEMI(size),
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastLRawReq(structType any) LRawReq {
	if casted, ok := structType.(LRawReq); ok {
		return casted
	}
	if casted, ok := structType.(*LRawReq); ok {
		return *casted
	}
	return nil
}

func (m *_LRawReq) GetTypeName() string {
	return "LRawReq"
}

func (m *_LRawReq) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.CEMIContract.(*_CEMI).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_LRawReq) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_LRawReq) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_CEMI, size uint16) (__lRawReq LRawReq, err error) {
	m.CEMIContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("LRawReq"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for LRawReq")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("LRawReq"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for LRawReq")
	}

	return m, nil
}

func (m *_LRawReq) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_LRawReq) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("LRawReq"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for LRawReq")
		}

		if popErr := writeBuffer.PopContext("LRawReq"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for LRawReq")
		}
		return nil
	}
	return m.CEMIContract.(*_CEMI).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_LRawReq) isLRawReq() bool {
	return true
}

func (m *_LRawReq) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
