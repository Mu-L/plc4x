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

// SequenceHeader is the corresponding interface of SequenceHeader
type SequenceHeader interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetSequenceNumber returns SequenceNumber (property field)
	GetSequenceNumber() int32
	// GetRequestId returns RequestId (property field)
	GetRequestId() int32
}

// SequenceHeaderExactly can be used when we want exactly this type and not a type which fulfills SequenceHeader.
// This is useful for switch cases.
type SequenceHeaderExactly interface {
	SequenceHeader
	isSequenceHeader() bool
}

// _SequenceHeader is the data-structure of this message
type _SequenceHeader struct {
	SequenceNumber int32
	RequestId      int32
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SequenceHeader) GetSequenceNumber() int32 {
	return m.SequenceNumber
}

func (m *_SequenceHeader) GetRequestId() int32 {
	return m.RequestId
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewSequenceHeader factory function for _SequenceHeader
func NewSequenceHeader(sequenceNumber int32, requestId int32) *_SequenceHeader {
	return &_SequenceHeader{SequenceNumber: sequenceNumber, RequestId: requestId}
}

// Deprecated: use the interface for direct cast
func CastSequenceHeader(structType any) SequenceHeader {
	if casted, ok := structType.(SequenceHeader); ok {
		return casted
	}
	if casted, ok := structType.(*SequenceHeader); ok {
		return *casted
	}
	return nil
}

func (m *_SequenceHeader) GetTypeName() string {
	return "SequenceHeader"
}

func (m *_SequenceHeader) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (sequenceNumber)
	lengthInBits += 32

	// Simple field (requestId)
	lengthInBits += 32

	return lengthInBits
}

func (m *_SequenceHeader) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func SequenceHeaderParse(ctx context.Context, theBytes []byte) (SequenceHeader, error) {
	return SequenceHeaderParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func SequenceHeaderParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (SequenceHeader, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (SequenceHeader, error) {
		return SequenceHeaderParseWithBuffer(ctx, readBuffer)
	}
}

func SequenceHeaderParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (SequenceHeader, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("SequenceHeader"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SequenceHeader")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	sequenceNumber, err := ReadSimpleField(ctx, "sequenceNumber", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'sequenceNumber' field"))
	}

	requestId, err := ReadSimpleField(ctx, "requestId", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestId' field"))
	}

	if closeErr := readBuffer.CloseContext("SequenceHeader"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SequenceHeader")
	}

	// Create the instance
	return &_SequenceHeader{
		SequenceNumber: sequenceNumber,
		RequestId:      requestId,
	}, nil
}

func (m *_SequenceHeader) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SequenceHeader) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("SequenceHeader"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for SequenceHeader")
	}

	// Simple Field (sequenceNumber)
	sequenceNumber := int32(m.GetSequenceNumber())
	_sequenceNumberErr := /*TODO: migrate me*/ writeBuffer.WriteInt32("sequenceNumber", 32, int32((sequenceNumber)))
	if _sequenceNumberErr != nil {
		return errors.Wrap(_sequenceNumberErr, "Error serializing 'sequenceNumber' field")
	}

	// Simple Field (requestId)
	requestId := int32(m.GetRequestId())
	_requestIdErr := /*TODO: migrate me*/ writeBuffer.WriteInt32("requestId", 32, int32((requestId)))
	if _requestIdErr != nil {
		return errors.Wrap(_requestIdErr, "Error serializing 'requestId' field")
	}

	if popErr := writeBuffer.PopContext("SequenceHeader"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for SequenceHeader")
	}
	return nil
}

func (m *_SequenceHeader) isSequenceHeader() bool {
	return true
}

func (m *_SequenceHeader) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
