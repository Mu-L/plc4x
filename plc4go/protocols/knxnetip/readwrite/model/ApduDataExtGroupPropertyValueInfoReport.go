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

// ApduDataExtGroupPropertyValueInfoReport is the corresponding interface of ApduDataExtGroupPropertyValueInfoReport
type ApduDataExtGroupPropertyValueInfoReport interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ApduDataExt
	// IsApduDataExtGroupPropertyValueInfoReport is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsApduDataExtGroupPropertyValueInfoReport()
}

// _ApduDataExtGroupPropertyValueInfoReport is the data-structure of this message
type _ApduDataExtGroupPropertyValueInfoReport struct {
	ApduDataExtContract
}

var _ ApduDataExtGroupPropertyValueInfoReport = (*_ApduDataExtGroupPropertyValueInfoReport)(nil)
var _ ApduDataExtRequirements = (*_ApduDataExtGroupPropertyValueInfoReport)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataExtGroupPropertyValueInfoReport) GetExtApciType() uint8 {
	return 0x2B
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataExtGroupPropertyValueInfoReport) GetParent() ApduDataExtContract {
	return m.ApduDataExtContract
}

// NewApduDataExtGroupPropertyValueInfoReport factory function for _ApduDataExtGroupPropertyValueInfoReport
func NewApduDataExtGroupPropertyValueInfoReport(length uint8) *_ApduDataExtGroupPropertyValueInfoReport {
	_result := &_ApduDataExtGroupPropertyValueInfoReport{
		ApduDataExtContract: NewApduDataExt(length),
	}
	_result.ApduDataExtContract.(*_ApduDataExt)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastApduDataExtGroupPropertyValueInfoReport(structType any) ApduDataExtGroupPropertyValueInfoReport {
	if casted, ok := structType.(ApduDataExtGroupPropertyValueInfoReport); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataExtGroupPropertyValueInfoReport); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataExtGroupPropertyValueInfoReport) GetTypeName() string {
	return "ApduDataExtGroupPropertyValueInfoReport"
}

func (m *_ApduDataExtGroupPropertyValueInfoReport) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ApduDataExtContract.(*_ApduDataExt).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_ApduDataExtGroupPropertyValueInfoReport) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ApduDataExtGroupPropertyValueInfoReport) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ApduDataExt, length uint8) (__apduDataExtGroupPropertyValueInfoReport ApduDataExtGroupPropertyValueInfoReport, err error) {
	m.ApduDataExtContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataExtGroupPropertyValueInfoReport"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataExtGroupPropertyValueInfoReport")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduDataExtGroupPropertyValueInfoReport"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataExtGroupPropertyValueInfoReport")
	}

	return m, nil
}

func (m *_ApduDataExtGroupPropertyValueInfoReport) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ApduDataExtGroupPropertyValueInfoReport) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtGroupPropertyValueInfoReport"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataExtGroupPropertyValueInfoReport")
		}

		if popErr := writeBuffer.PopContext("ApduDataExtGroupPropertyValueInfoReport"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataExtGroupPropertyValueInfoReport")
		}
		return nil
	}
	return m.ApduDataExtContract.(*_ApduDataExt).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ApduDataExtGroupPropertyValueInfoReport) IsApduDataExtGroupPropertyValueInfoReport() {}

func (m *_ApduDataExtGroupPropertyValueInfoReport) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
