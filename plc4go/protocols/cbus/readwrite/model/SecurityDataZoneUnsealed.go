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

// SecurityDataZoneUnsealed is the corresponding interface of SecurityDataZoneUnsealed
type SecurityDataZoneUnsealed interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	SecurityData
	// GetZoneNumber returns ZoneNumber (property field)
	GetZoneNumber() uint8
}

// SecurityDataZoneUnsealedExactly can be used when we want exactly this type and not a type which fulfills SecurityDataZoneUnsealed.
// This is useful for switch cases.
type SecurityDataZoneUnsealedExactly interface {
	SecurityDataZoneUnsealed
	isSecurityDataZoneUnsealed() bool
}

// _SecurityDataZoneUnsealed is the data-structure of this message
type _SecurityDataZoneUnsealed struct {
	SecurityDataContract
	ZoneNumber uint8
}

var _ SecurityDataZoneUnsealed = (*_SecurityDataZoneUnsealed)(nil)
var _ SecurityDataRequirements = (*_SecurityDataZoneUnsealed)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SecurityDataZoneUnsealed) GetParent() SecurityDataContract {
	return m.SecurityDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SecurityDataZoneUnsealed) GetZoneNumber() uint8 {
	return m.ZoneNumber
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewSecurityDataZoneUnsealed factory function for _SecurityDataZoneUnsealed
func NewSecurityDataZoneUnsealed(zoneNumber uint8, commandTypeContainer SecurityCommandTypeContainer, argument byte) *_SecurityDataZoneUnsealed {
	_result := &_SecurityDataZoneUnsealed{
		SecurityDataContract: NewSecurityData(commandTypeContainer, argument),
		ZoneNumber:           zoneNumber,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastSecurityDataZoneUnsealed(structType any) SecurityDataZoneUnsealed {
	if casted, ok := structType.(SecurityDataZoneUnsealed); ok {
		return casted
	}
	if casted, ok := structType.(*SecurityDataZoneUnsealed); ok {
		return *casted
	}
	return nil
}

func (m *_SecurityDataZoneUnsealed) GetTypeName() string {
	return "SecurityDataZoneUnsealed"
}

func (m *_SecurityDataZoneUnsealed) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.SecurityDataContract.(*_SecurityData).getLengthInBits(ctx))

	// Simple field (zoneNumber)
	lengthInBits += 8

	return lengthInBits
}

func (m *_SecurityDataZoneUnsealed) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SecurityDataZoneUnsealed) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_SecurityData) (__securityDataZoneUnsealed SecurityDataZoneUnsealed, err error) {
	m.SecurityDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SecurityDataZoneUnsealed"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SecurityDataZoneUnsealed")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	zoneNumber, err := ReadSimpleField(ctx, "zoneNumber", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'zoneNumber' field"))
	}
	m.ZoneNumber = zoneNumber

	if closeErr := readBuffer.CloseContext("SecurityDataZoneUnsealed"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SecurityDataZoneUnsealed")
	}

	return m, nil
}

func (m *_SecurityDataZoneUnsealed) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SecurityDataZoneUnsealed) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SecurityDataZoneUnsealed"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SecurityDataZoneUnsealed")
		}

		if err := WriteSimpleField[uint8](ctx, "zoneNumber", m.GetZoneNumber(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'zoneNumber' field")
		}

		if popErr := writeBuffer.PopContext("SecurityDataZoneUnsealed"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SecurityDataZoneUnsealed")
		}
		return nil
	}
	return m.SecurityDataContract.(*_SecurityData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SecurityDataZoneUnsealed) isSecurityDataZoneUnsealed() bool {
	return true
}

func (m *_SecurityDataZoneUnsealed) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
