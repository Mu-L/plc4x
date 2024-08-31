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

// MediaTransportControlDataStatusRequest is the corresponding interface of MediaTransportControlDataStatusRequest
type MediaTransportControlDataStatusRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	MediaTransportControlData
}

// MediaTransportControlDataStatusRequestExactly can be used when we want exactly this type and not a type which fulfills MediaTransportControlDataStatusRequest.
// This is useful for switch cases.
type MediaTransportControlDataStatusRequestExactly interface {
	MediaTransportControlDataStatusRequest
	isMediaTransportControlDataStatusRequest() bool
}

// _MediaTransportControlDataStatusRequest is the data-structure of this message
type _MediaTransportControlDataStatusRequest struct {
	*_MediaTransportControlData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_MediaTransportControlDataStatusRequest) InitializeParent(parent MediaTransportControlData, commandTypeContainer MediaTransportControlCommandTypeContainer, mediaLinkGroup byte) {
	m.CommandTypeContainer = commandTypeContainer
	m.MediaLinkGroup = mediaLinkGroup
}

func (m *_MediaTransportControlDataStatusRequest) GetParent() MediaTransportControlData {
	return m._MediaTransportControlData
}

// NewMediaTransportControlDataStatusRequest factory function for _MediaTransportControlDataStatusRequest
func NewMediaTransportControlDataStatusRequest(commandTypeContainer MediaTransportControlCommandTypeContainer, mediaLinkGroup byte) *_MediaTransportControlDataStatusRequest {
	_result := &_MediaTransportControlDataStatusRequest{
		_MediaTransportControlData: NewMediaTransportControlData(commandTypeContainer, mediaLinkGroup),
	}
	_result._MediaTransportControlData._MediaTransportControlDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastMediaTransportControlDataStatusRequest(structType any) MediaTransportControlDataStatusRequest {
	if casted, ok := structType.(MediaTransportControlDataStatusRequest); ok {
		return casted
	}
	if casted, ok := structType.(*MediaTransportControlDataStatusRequest); ok {
		return *casted
	}
	return nil
}

func (m *_MediaTransportControlDataStatusRequest) GetTypeName() string {
	return "MediaTransportControlDataStatusRequest"
}

func (m *_MediaTransportControlDataStatusRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	return lengthInBits
}

func (m *_MediaTransportControlDataStatusRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func MediaTransportControlDataStatusRequestParse(ctx context.Context, theBytes []byte) (MediaTransportControlDataStatusRequest, error) {
	return MediaTransportControlDataStatusRequestParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func MediaTransportControlDataStatusRequestParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (MediaTransportControlDataStatusRequest, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (MediaTransportControlDataStatusRequest, error) {
		return MediaTransportControlDataStatusRequestParseWithBuffer(ctx, readBuffer)
	}
}

func MediaTransportControlDataStatusRequestParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (MediaTransportControlDataStatusRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("MediaTransportControlDataStatusRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MediaTransportControlDataStatusRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("MediaTransportControlDataStatusRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MediaTransportControlDataStatusRequest")
	}

	// Create a partially initialized instance
	_child := &_MediaTransportControlDataStatusRequest{
		_MediaTransportControlData: &_MediaTransportControlData{},
	}
	_child._MediaTransportControlData._MediaTransportControlDataChildRequirements = _child
	return _child, nil
}

func (m *_MediaTransportControlDataStatusRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MediaTransportControlDataStatusRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MediaTransportControlDataStatusRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MediaTransportControlDataStatusRequest")
		}

		if popErr := writeBuffer.PopContext("MediaTransportControlDataStatusRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MediaTransportControlDataStatusRequest")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_MediaTransportControlDataStatusRequest) isMediaTransportControlDataStatusRequest() bool {
	return true
}

func (m *_MediaTransportControlDataStatusRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
