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

// SecurityDataFireAlarmRaised is the corresponding interface of SecurityDataFireAlarmRaised
type SecurityDataFireAlarmRaised interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	SecurityData
}

// SecurityDataFireAlarmRaisedExactly can be used when we want exactly this type and not a type which fulfills SecurityDataFireAlarmRaised.
// This is useful for switch cases.
type SecurityDataFireAlarmRaisedExactly interface {
	SecurityDataFireAlarmRaised
	isSecurityDataFireAlarmRaised() bool
}

// _SecurityDataFireAlarmRaised is the data-structure of this message
type _SecurityDataFireAlarmRaised struct {
	SecurityDataContract
}

var _ SecurityDataFireAlarmRaised = (*_SecurityDataFireAlarmRaised)(nil)
var _ SecurityDataRequirements = (*_SecurityDataFireAlarmRaised)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SecurityDataFireAlarmRaised) GetParent() SecurityDataContract {
	return m.SecurityDataContract
}

// NewSecurityDataFireAlarmRaised factory function for _SecurityDataFireAlarmRaised
func NewSecurityDataFireAlarmRaised(commandTypeContainer SecurityCommandTypeContainer, argument byte) *_SecurityDataFireAlarmRaised {
	_result := &_SecurityDataFireAlarmRaised{
		SecurityDataContract: NewSecurityData(commandTypeContainer, argument),
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastSecurityDataFireAlarmRaised(structType any) SecurityDataFireAlarmRaised {
	if casted, ok := structType.(SecurityDataFireAlarmRaised); ok {
		return casted
	}
	if casted, ok := structType.(*SecurityDataFireAlarmRaised); ok {
		return *casted
	}
	return nil
}

func (m *_SecurityDataFireAlarmRaised) GetTypeName() string {
	return "SecurityDataFireAlarmRaised"
}

func (m *_SecurityDataFireAlarmRaised) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.SecurityDataContract.(*_SecurityData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_SecurityDataFireAlarmRaised) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SecurityDataFireAlarmRaised) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_SecurityData) (__securityDataFireAlarmRaised SecurityDataFireAlarmRaised, err error) {
	m.SecurityDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SecurityDataFireAlarmRaised"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SecurityDataFireAlarmRaised")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("SecurityDataFireAlarmRaised"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SecurityDataFireAlarmRaised")
	}

	return m, nil
}

func (m *_SecurityDataFireAlarmRaised) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SecurityDataFireAlarmRaised) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SecurityDataFireAlarmRaised"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SecurityDataFireAlarmRaised")
		}

		if popErr := writeBuffer.PopContext("SecurityDataFireAlarmRaised"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SecurityDataFireAlarmRaised")
		}
		return nil
	}
	return m.SecurityDataContract.(*_SecurityData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SecurityDataFireAlarmRaised) isSecurityDataFireAlarmRaised() bool {
	return true
}

func (m *_SecurityDataFireAlarmRaised) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
