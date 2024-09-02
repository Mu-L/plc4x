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

// SecurityDataGasAlarmRaised is the corresponding interface of SecurityDataGasAlarmRaised
type SecurityDataGasAlarmRaised interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	SecurityData
	// IsSecurityDataGasAlarmRaised is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsSecurityDataGasAlarmRaised()
}

// _SecurityDataGasAlarmRaised is the data-structure of this message
type _SecurityDataGasAlarmRaised struct {
	SecurityDataContract
}

var _ SecurityDataGasAlarmRaised = (*_SecurityDataGasAlarmRaised)(nil)
var _ SecurityDataRequirements = (*_SecurityDataGasAlarmRaised)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SecurityDataGasAlarmRaised) GetParent() SecurityDataContract {
	return m.SecurityDataContract
}

// NewSecurityDataGasAlarmRaised factory function for _SecurityDataGasAlarmRaised
func NewSecurityDataGasAlarmRaised(commandTypeContainer SecurityCommandTypeContainer, argument byte) *_SecurityDataGasAlarmRaised {
	_result := &_SecurityDataGasAlarmRaised{
		SecurityDataContract: NewSecurityData(commandTypeContainer, argument),
	}
	_result.SecurityDataContract.(*_SecurityData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSecurityDataGasAlarmRaised(structType any) SecurityDataGasAlarmRaised {
	if casted, ok := structType.(SecurityDataGasAlarmRaised); ok {
		return casted
	}
	if casted, ok := structType.(*SecurityDataGasAlarmRaised); ok {
		return *casted
	}
	return nil
}

func (m *_SecurityDataGasAlarmRaised) GetTypeName() string {
	return "SecurityDataGasAlarmRaised"
}

func (m *_SecurityDataGasAlarmRaised) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.SecurityDataContract.(*_SecurityData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_SecurityDataGasAlarmRaised) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SecurityDataGasAlarmRaised) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_SecurityData) (__securityDataGasAlarmRaised SecurityDataGasAlarmRaised, err error) {
	m.SecurityDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SecurityDataGasAlarmRaised"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SecurityDataGasAlarmRaised")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("SecurityDataGasAlarmRaised"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SecurityDataGasAlarmRaised")
	}

	return m, nil
}

func (m *_SecurityDataGasAlarmRaised) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SecurityDataGasAlarmRaised) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SecurityDataGasAlarmRaised"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SecurityDataGasAlarmRaised")
		}

		if popErr := writeBuffer.PopContext("SecurityDataGasAlarmRaised"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SecurityDataGasAlarmRaised")
		}
		return nil
	}
	return m.SecurityDataContract.(*_SecurityData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SecurityDataGasAlarmRaised) IsSecurityDataGasAlarmRaised() {}

func (m *_SecurityDataGasAlarmRaised) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
