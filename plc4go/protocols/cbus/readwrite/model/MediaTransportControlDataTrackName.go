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

// MediaTransportControlDataTrackName is the corresponding interface of MediaTransportControlDataTrackName
type MediaTransportControlDataTrackName interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	MediaTransportControlData
	// GetTrackName returns TrackName (property field)
	GetTrackName() string
}

// MediaTransportControlDataTrackNameExactly can be used when we want exactly this type and not a type which fulfills MediaTransportControlDataTrackName.
// This is useful for switch cases.
type MediaTransportControlDataTrackNameExactly interface {
	MediaTransportControlDataTrackName
	isMediaTransportControlDataTrackName() bool
}

// _MediaTransportControlDataTrackName is the data-structure of this message
type _MediaTransportControlDataTrackName struct {
	MediaTransportControlDataContract
	TrackName string
}

var _ MediaTransportControlDataTrackName = (*_MediaTransportControlDataTrackName)(nil)
var _ MediaTransportControlDataRequirements = (*_MediaTransportControlDataTrackName)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_MediaTransportControlDataTrackName) GetParent() MediaTransportControlDataContract {
	return m.MediaTransportControlDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_MediaTransportControlDataTrackName) GetTrackName() string {
	return m.TrackName
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewMediaTransportControlDataTrackName factory function for _MediaTransportControlDataTrackName
func NewMediaTransportControlDataTrackName(trackName string, commandTypeContainer MediaTransportControlCommandTypeContainer, mediaLinkGroup byte) *_MediaTransportControlDataTrackName {
	_result := &_MediaTransportControlDataTrackName{
		MediaTransportControlDataContract: NewMediaTransportControlData(commandTypeContainer, mediaLinkGroup),
		TrackName:                         trackName,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastMediaTransportControlDataTrackName(structType any) MediaTransportControlDataTrackName {
	if casted, ok := structType.(MediaTransportControlDataTrackName); ok {
		return casted
	}
	if casted, ok := structType.(*MediaTransportControlDataTrackName); ok {
		return *casted
	}
	return nil
}

func (m *_MediaTransportControlDataTrackName) GetTypeName() string {
	return "MediaTransportControlDataTrackName"
}

func (m *_MediaTransportControlDataTrackName) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.MediaTransportControlDataContract.(*_MediaTransportControlData).getLengthInBits(ctx))

	// Simple field (trackName)
	lengthInBits += uint16(int32((int32(m.GetCommandTypeContainer().NumBytes()) - int32(int32(1)))) * int32(int32(8)))

	return lengthInBits
}

func (m *_MediaTransportControlDataTrackName) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_MediaTransportControlDataTrackName) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_MediaTransportControlData, commandTypeContainer MediaTransportControlCommandTypeContainer) (__mediaTransportControlDataTrackName MediaTransportControlDataTrackName, err error) {
	m.MediaTransportControlDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MediaTransportControlDataTrackName"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MediaTransportControlDataTrackName")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	trackName, err := ReadSimpleField(ctx, "trackName", ReadString(readBuffer, uint32(int32((int32(commandTypeContainer.NumBytes())-int32(int32(1))))*int32(int32(8)))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'trackName' field"))
	}
	m.TrackName = trackName

	if closeErr := readBuffer.CloseContext("MediaTransportControlDataTrackName"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MediaTransportControlDataTrackName")
	}

	return m, nil
}

func (m *_MediaTransportControlDataTrackName) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MediaTransportControlDataTrackName) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MediaTransportControlDataTrackName"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MediaTransportControlDataTrackName")
		}

		if err := WriteSimpleField[string](ctx, "trackName", m.GetTrackName(), WriteString(writeBuffer, int32(int32((int32(m.GetCommandTypeContainer().NumBytes())-int32(int32(1))))*int32(int32(8))))); err != nil {
			return errors.Wrap(err, "Error serializing 'trackName' field")
		}

		if popErr := writeBuffer.PopContext("MediaTransportControlDataTrackName"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MediaTransportControlDataTrackName")
		}
		return nil
	}
	return m.MediaTransportControlDataContract.(*_MediaTransportControlData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_MediaTransportControlDataTrackName) isMediaTransportControlDataTrackName() bool {
	return true
}

func (m *_MediaTransportControlDataTrackName) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
