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

// BACnetServiceAckAtomicReadFileStream is the corresponding interface of BACnetServiceAckAtomicReadFileStream
type BACnetServiceAckAtomicReadFileStream interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetServiceAckAtomicReadFileStreamOrRecord
	// GetFileStartPosition returns FileStartPosition (property field)
	GetFileStartPosition() BACnetApplicationTagSignedInteger
	// GetFileData returns FileData (property field)
	GetFileData() BACnetApplicationTagOctetString
}

// BACnetServiceAckAtomicReadFileStreamExactly can be used when we want exactly this type and not a type which fulfills BACnetServiceAckAtomicReadFileStream.
// This is useful for switch cases.
type BACnetServiceAckAtomicReadFileStreamExactly interface {
	BACnetServiceAckAtomicReadFileStream
	isBACnetServiceAckAtomicReadFileStream() bool
}

// _BACnetServiceAckAtomicReadFileStream is the data-structure of this message
type _BACnetServiceAckAtomicReadFileStream struct {
	BACnetServiceAckAtomicReadFileStreamOrRecordContract
	FileStartPosition BACnetApplicationTagSignedInteger
	FileData          BACnetApplicationTagOctetString
}

var _ BACnetServiceAckAtomicReadFileStream = (*_BACnetServiceAckAtomicReadFileStream)(nil)
var _ BACnetServiceAckAtomicReadFileStreamOrRecordRequirements = (*_BACnetServiceAckAtomicReadFileStream)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetServiceAckAtomicReadFileStream) GetParent() BACnetServiceAckAtomicReadFileStreamOrRecordContract {
	return m.BACnetServiceAckAtomicReadFileStreamOrRecordContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetServiceAckAtomicReadFileStream) GetFileStartPosition() BACnetApplicationTagSignedInteger {
	return m.FileStartPosition
}

func (m *_BACnetServiceAckAtomicReadFileStream) GetFileData() BACnetApplicationTagOctetString {
	return m.FileData
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetServiceAckAtomicReadFileStream factory function for _BACnetServiceAckAtomicReadFileStream
func NewBACnetServiceAckAtomicReadFileStream(fileStartPosition BACnetApplicationTagSignedInteger, fileData BACnetApplicationTagOctetString, peekedTagHeader BACnetTagHeader, openingTag BACnetOpeningTag, closingTag BACnetClosingTag) *_BACnetServiceAckAtomicReadFileStream {
	_result := &_BACnetServiceAckAtomicReadFileStream{
		BACnetServiceAckAtomicReadFileStreamOrRecordContract: NewBACnetServiceAckAtomicReadFileStreamOrRecord(peekedTagHeader, openingTag, closingTag),
		FileStartPosition: fileStartPosition,
		FileData:          fileData,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetServiceAckAtomicReadFileStream(structType any) BACnetServiceAckAtomicReadFileStream {
	if casted, ok := structType.(BACnetServiceAckAtomicReadFileStream); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetServiceAckAtomicReadFileStream); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetServiceAckAtomicReadFileStream) GetTypeName() string {
	return "BACnetServiceAckAtomicReadFileStream"
}

func (m *_BACnetServiceAckAtomicReadFileStream) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetServiceAckAtomicReadFileStreamOrRecordContract.(*_BACnetServiceAckAtomicReadFileStreamOrRecord).getLengthInBits(ctx))

	// Simple field (fileStartPosition)
	lengthInBits += m.FileStartPosition.GetLengthInBits(ctx)

	// Simple field (fileData)
	lengthInBits += m.FileData.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetServiceAckAtomicReadFileStream) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetServiceAckAtomicReadFileStream) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetServiceAckAtomicReadFileStreamOrRecord) (__bACnetServiceAckAtomicReadFileStream BACnetServiceAckAtomicReadFileStream, err error) {
	m.BACnetServiceAckAtomicReadFileStreamOrRecordContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetServiceAckAtomicReadFileStream"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetServiceAckAtomicReadFileStream")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	fileStartPosition, err := ReadSimpleField[BACnetApplicationTagSignedInteger](ctx, "fileStartPosition", ReadComplex[BACnetApplicationTagSignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagSignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'fileStartPosition' field"))
	}
	m.FileStartPosition = fileStartPosition

	fileData, err := ReadSimpleField[BACnetApplicationTagOctetString](ctx, "fileData", ReadComplex[BACnetApplicationTagOctetString](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagOctetString](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'fileData' field"))
	}
	m.FileData = fileData

	if closeErr := readBuffer.CloseContext("BACnetServiceAckAtomicReadFileStream"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetServiceAckAtomicReadFileStream")
	}

	return m, nil
}

func (m *_BACnetServiceAckAtomicReadFileStream) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetServiceAckAtomicReadFileStream) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetServiceAckAtomicReadFileStream"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetServiceAckAtomicReadFileStream")
		}

		if err := WriteSimpleField[BACnetApplicationTagSignedInteger](ctx, "fileStartPosition", m.GetFileStartPosition(), WriteComplex[BACnetApplicationTagSignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'fileStartPosition' field")
		}

		if err := WriteSimpleField[BACnetApplicationTagOctetString](ctx, "fileData", m.GetFileData(), WriteComplex[BACnetApplicationTagOctetString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'fileData' field")
		}

		if popErr := writeBuffer.PopContext("BACnetServiceAckAtomicReadFileStream"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetServiceAckAtomicReadFileStream")
		}
		return nil
	}
	return m.BACnetServiceAckAtomicReadFileStreamOrRecordContract.(*_BACnetServiceAckAtomicReadFileStreamOrRecord).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetServiceAckAtomicReadFileStream) isBACnetServiceAckAtomicReadFileStream() bool {
	return true
}

func (m *_BACnetServiceAckAtomicReadFileStream) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
