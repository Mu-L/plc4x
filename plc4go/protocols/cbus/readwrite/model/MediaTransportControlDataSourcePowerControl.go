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

// MediaTransportControlDataSourcePowerControl is the corresponding interface of MediaTransportControlDataSourcePowerControl
type MediaTransportControlDataSourcePowerControl interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	MediaTransportControlData
	// GetState returns State (property field)
	GetState() byte
	// GetIsShouldPowerOn returns IsShouldPowerOn (virtual field)
	GetIsShouldPowerOn() bool
	// GetIsShouldPowerOff returns IsShouldPowerOff (virtual field)
	GetIsShouldPowerOff() bool
}

// MediaTransportControlDataSourcePowerControlExactly can be used when we want exactly this type and not a type which fulfills MediaTransportControlDataSourcePowerControl.
// This is useful for switch cases.
type MediaTransportControlDataSourcePowerControlExactly interface {
	MediaTransportControlDataSourcePowerControl
	isMediaTransportControlDataSourcePowerControl() bool
}

// _MediaTransportControlDataSourcePowerControl is the data-structure of this message
type _MediaTransportControlDataSourcePowerControl struct {
	MediaTransportControlDataContract
	State byte
}

var _ MediaTransportControlDataSourcePowerControl = (*_MediaTransportControlDataSourcePowerControl)(nil)
var _ MediaTransportControlDataRequirements = (*_MediaTransportControlDataSourcePowerControl)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_MediaTransportControlDataSourcePowerControl) GetParent() MediaTransportControlDataContract {
	return m.MediaTransportControlDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_MediaTransportControlDataSourcePowerControl) GetState() byte {
	return m.State
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_MediaTransportControlDataSourcePowerControl) GetIsShouldPowerOn() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetState()) == (0x00)))
}

func (m *_MediaTransportControlDataSourcePowerControl) GetIsShouldPowerOff() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetState()) != (0x00)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewMediaTransportControlDataSourcePowerControl factory function for _MediaTransportControlDataSourcePowerControl
func NewMediaTransportControlDataSourcePowerControl(state byte, commandTypeContainer MediaTransportControlCommandTypeContainer, mediaLinkGroup byte) *_MediaTransportControlDataSourcePowerControl {
	_result := &_MediaTransportControlDataSourcePowerControl{
		MediaTransportControlDataContract: NewMediaTransportControlData(commandTypeContainer, mediaLinkGroup),
		State:                             state,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastMediaTransportControlDataSourcePowerControl(structType any) MediaTransportControlDataSourcePowerControl {
	if casted, ok := structType.(MediaTransportControlDataSourcePowerControl); ok {
		return casted
	}
	if casted, ok := structType.(*MediaTransportControlDataSourcePowerControl); ok {
		return *casted
	}
	return nil
}

func (m *_MediaTransportControlDataSourcePowerControl) GetTypeName() string {
	return "MediaTransportControlDataSourcePowerControl"
}

func (m *_MediaTransportControlDataSourcePowerControl) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.MediaTransportControlDataContract.(*_MediaTransportControlData).getLengthInBits(ctx))

	// Simple field (state)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_MediaTransportControlDataSourcePowerControl) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_MediaTransportControlDataSourcePowerControl) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_MediaTransportControlData) (__mediaTransportControlDataSourcePowerControl MediaTransportControlDataSourcePowerControl, err error) {
	m.MediaTransportControlDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MediaTransportControlDataSourcePowerControl"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MediaTransportControlDataSourcePowerControl")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	state, err := ReadSimpleField(ctx, "state", ReadByte(readBuffer, 8))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'state' field"))
	}
	m.State = state

	isShouldPowerOn, err := ReadVirtualField[bool](ctx, "isShouldPowerOn", (*bool)(nil), bool((state) == (0x00)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isShouldPowerOn' field"))
	}
	_ = isShouldPowerOn

	isShouldPowerOff, err := ReadVirtualField[bool](ctx, "isShouldPowerOff", (*bool)(nil), bool((state) != (0x00)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isShouldPowerOff' field"))
	}
	_ = isShouldPowerOff

	if closeErr := readBuffer.CloseContext("MediaTransportControlDataSourcePowerControl"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MediaTransportControlDataSourcePowerControl")
	}

	return m, nil
}

func (m *_MediaTransportControlDataSourcePowerControl) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MediaTransportControlDataSourcePowerControl) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MediaTransportControlDataSourcePowerControl"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MediaTransportControlDataSourcePowerControl")
		}

		if err := WriteSimpleField[byte](ctx, "state", m.GetState(), WriteByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'state' field")
		}
		// Virtual field
		isShouldPowerOn := m.GetIsShouldPowerOn()
		_ = isShouldPowerOn
		if _isShouldPowerOnErr := writeBuffer.WriteVirtual(ctx, "isShouldPowerOn", m.GetIsShouldPowerOn()); _isShouldPowerOnErr != nil {
			return errors.Wrap(_isShouldPowerOnErr, "Error serializing 'isShouldPowerOn' field")
		}
		// Virtual field
		isShouldPowerOff := m.GetIsShouldPowerOff()
		_ = isShouldPowerOff
		if _isShouldPowerOffErr := writeBuffer.WriteVirtual(ctx, "isShouldPowerOff", m.GetIsShouldPowerOff()); _isShouldPowerOffErr != nil {
			return errors.Wrap(_isShouldPowerOffErr, "Error serializing 'isShouldPowerOff' field")
		}

		if popErr := writeBuffer.PopContext("MediaTransportControlDataSourcePowerControl"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MediaTransportControlDataSourcePowerControl")
		}
		return nil
	}
	return m.MediaTransportControlDataContract.(*_MediaTransportControlData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_MediaTransportControlDataSourcePowerControl) isMediaTransportControlDataSourcePowerControl() bool {
	return true
}

func (m *_MediaTransportControlDataSourcePowerControl) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
