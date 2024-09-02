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

// AccessControlDataLockAccessPoint is the corresponding interface of AccessControlDataLockAccessPoint
type AccessControlDataLockAccessPoint interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	AccessControlData
}

// AccessControlDataLockAccessPointExactly can be used when we want exactly this type and not a type which fulfills AccessControlDataLockAccessPoint.
// This is useful for switch cases.
type AccessControlDataLockAccessPointExactly interface {
	AccessControlDataLockAccessPoint
	isAccessControlDataLockAccessPoint() bool
}

// _AccessControlDataLockAccessPoint is the data-structure of this message
type _AccessControlDataLockAccessPoint struct {
	AccessControlDataContract
}

var _ AccessControlDataLockAccessPoint = (*_AccessControlDataLockAccessPoint)(nil)
var _ AccessControlDataRequirements = (*_AccessControlDataLockAccessPoint)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AccessControlDataLockAccessPoint) GetParent() AccessControlDataContract {
	return m.AccessControlDataContract
}

// NewAccessControlDataLockAccessPoint factory function for _AccessControlDataLockAccessPoint
func NewAccessControlDataLockAccessPoint(commandTypeContainer AccessControlCommandTypeContainer, networkId byte, accessPointId byte) *_AccessControlDataLockAccessPoint {
	_result := &_AccessControlDataLockAccessPoint{
		AccessControlDataContract: NewAccessControlData(commandTypeContainer, networkId, accessPointId),
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastAccessControlDataLockAccessPoint(structType any) AccessControlDataLockAccessPoint {
	if casted, ok := structType.(AccessControlDataLockAccessPoint); ok {
		return casted
	}
	if casted, ok := structType.(*AccessControlDataLockAccessPoint); ok {
		return *casted
	}
	return nil
}

func (m *_AccessControlDataLockAccessPoint) GetTypeName() string {
	return "AccessControlDataLockAccessPoint"
}

func (m *_AccessControlDataLockAccessPoint) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.AccessControlDataContract.(*_AccessControlData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_AccessControlDataLockAccessPoint) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_AccessControlDataLockAccessPoint) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_AccessControlData) (__accessControlDataLockAccessPoint AccessControlDataLockAccessPoint, err error) {
	m.AccessControlDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AccessControlDataLockAccessPoint"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AccessControlDataLockAccessPoint")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("AccessControlDataLockAccessPoint"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AccessControlDataLockAccessPoint")
	}

	return m, nil
}

func (m *_AccessControlDataLockAccessPoint) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AccessControlDataLockAccessPoint) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AccessControlDataLockAccessPoint"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AccessControlDataLockAccessPoint")
		}

		if popErr := writeBuffer.PopContext("AccessControlDataLockAccessPoint"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AccessControlDataLockAccessPoint")
		}
		return nil
	}
	return m.AccessControlDataContract.(*_AccessControlData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AccessControlDataLockAccessPoint) isAccessControlDataLockAccessPoint() bool {
	return true
}

func (m *_AccessControlDataLockAccessPoint) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
