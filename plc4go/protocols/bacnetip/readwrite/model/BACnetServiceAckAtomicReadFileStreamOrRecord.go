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

// BACnetServiceAckAtomicReadFileStreamOrRecord is the corresponding interface of BACnetServiceAckAtomicReadFileStreamOrRecord
type BACnetServiceAckAtomicReadFileStreamOrRecord interface {
	BACnetServiceAckAtomicReadFileStreamOrRecordContract
	BACnetServiceAckAtomicReadFileStreamOrRecordRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// IsBACnetServiceAckAtomicReadFileStreamOrRecord is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetServiceAckAtomicReadFileStreamOrRecord()
}

// BACnetServiceAckAtomicReadFileStreamOrRecordContract provides a set of functions which can be overwritten by a sub struct
type BACnetServiceAckAtomicReadFileStreamOrRecordContract interface {
	// GetPeekedTagHeader returns PeekedTagHeader (property field)
	GetPeekedTagHeader() BACnetTagHeader
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
	// GetPeekedTagNumber returns PeekedTagNumber (virtual field)
	GetPeekedTagNumber() uint8
	// IsBACnetServiceAckAtomicReadFileStreamOrRecord is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetServiceAckAtomicReadFileStreamOrRecord()
}

// BACnetServiceAckAtomicReadFileStreamOrRecordRequirements provides a set of functions which need to be implemented by a sub struct
type BACnetServiceAckAtomicReadFileStreamOrRecordRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetPeekedTagNumber returns PeekedTagNumber (discriminator field)
	GetPeekedTagNumber() uint8
}

// _BACnetServiceAckAtomicReadFileStreamOrRecord is the data-structure of this message
type _BACnetServiceAckAtomicReadFileStreamOrRecord struct {
	_SubType        BACnetServiceAckAtomicReadFileStreamOrRecord
	PeekedTagHeader BACnetTagHeader
	OpeningTag      BACnetOpeningTag
	ClosingTag      BACnetClosingTag
}

var _ BACnetServiceAckAtomicReadFileStreamOrRecordContract = (*_BACnetServiceAckAtomicReadFileStreamOrRecord)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetServiceAckAtomicReadFileStreamOrRecord) GetPeekedTagHeader() BACnetTagHeader {
	return m.PeekedTagHeader
}

func (m *_BACnetServiceAckAtomicReadFileStreamOrRecord) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetServiceAckAtomicReadFileStreamOrRecord) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (pm *_BACnetServiceAckAtomicReadFileStreamOrRecord) GetPeekedTagNumber() uint8 {
	m := pm._SubType
	ctx := context.Background()
	_ = ctx
	return uint8(m.GetPeekedTagHeader().GetActualTagNumber())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetServiceAckAtomicReadFileStreamOrRecord factory function for _BACnetServiceAckAtomicReadFileStreamOrRecord
func NewBACnetServiceAckAtomicReadFileStreamOrRecord(peekedTagHeader BACnetTagHeader, openingTag BACnetOpeningTag, closingTag BACnetClosingTag) *_BACnetServiceAckAtomicReadFileStreamOrRecord {
	return &_BACnetServiceAckAtomicReadFileStreamOrRecord{PeekedTagHeader: peekedTagHeader, OpeningTag: openingTag, ClosingTag: closingTag}
}

// Deprecated: use the interface for direct cast
func CastBACnetServiceAckAtomicReadFileStreamOrRecord(structType any) BACnetServiceAckAtomicReadFileStreamOrRecord {
	if casted, ok := structType.(BACnetServiceAckAtomicReadFileStreamOrRecord); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetServiceAckAtomicReadFileStreamOrRecord); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetServiceAckAtomicReadFileStreamOrRecord) GetTypeName() string {
	return "BACnetServiceAckAtomicReadFileStreamOrRecord"
}

func (m *_BACnetServiceAckAtomicReadFileStreamOrRecord) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetServiceAckAtomicReadFileStreamOrRecord) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func BACnetServiceAckAtomicReadFileStreamOrRecordParse[T BACnetServiceAckAtomicReadFileStreamOrRecord](ctx context.Context, theBytes []byte) (T, error) {
	return BACnetServiceAckAtomicReadFileStreamOrRecordParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetServiceAckAtomicReadFileStreamOrRecordParseWithBufferProducer[T BACnetServiceAckAtomicReadFileStreamOrRecord]() func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := BACnetServiceAckAtomicReadFileStreamOrRecordParseWithBuffer[T](ctx, readBuffer)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, err
	}
}

func BACnetServiceAckAtomicReadFileStreamOrRecordParseWithBuffer[T BACnetServiceAckAtomicReadFileStreamOrRecord](ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	v, err := (&_BACnetServiceAckAtomicReadFileStreamOrRecord{}).parse(ctx, readBuffer)
	if err != nil {
		var zero T
		return zero, err
	}
	return v.(T), err
}

func (m *_BACnetServiceAckAtomicReadFileStreamOrRecord) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__bACnetServiceAckAtomicReadFileStreamOrRecord BACnetServiceAckAtomicReadFileStreamOrRecord, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetServiceAckAtomicReadFileStreamOrRecord"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetServiceAckAtomicReadFileStreamOrRecord")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	peekedTagHeader, err := ReadPeekField[BACnetTagHeader](ctx, "peekedTagHeader", ReadComplex[BACnetTagHeader](BACnetTagHeaderParseWithBuffer, readBuffer), 0)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'peekedTagHeader' field"))
	}
	m.PeekedTagHeader = peekedTagHeader

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(peekedTagHeader.GetActualTagNumber())), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}
	m.OpeningTag = openingTag

	peekedTagNumber, err := ReadVirtualField[uint8](ctx, "peekedTagNumber", (*uint8)(nil), peekedTagHeader.GetActualTagNumber())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'peekedTagNumber' field"))
	}
	_ = peekedTagNumber

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child BACnetServiceAckAtomicReadFileStreamOrRecord
	switch {
	case peekedTagNumber == 0x0: // BACnetServiceAckAtomicReadFileStream
		if _child, err = (&_BACnetServiceAckAtomicReadFileStream{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetServiceAckAtomicReadFileStream for type-switch of BACnetServiceAckAtomicReadFileStreamOrRecord")
		}
	case peekedTagNumber == 0x1: // BACnetServiceAckAtomicReadFileRecord
		if _child, err = (&_BACnetServiceAckAtomicReadFileRecord{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetServiceAckAtomicReadFileRecord for type-switch of BACnetServiceAckAtomicReadFileStreamOrRecord")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [peekedTagNumber=%v]", peekedTagNumber)
	}

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(peekedTagHeader.GetActualTagNumber())), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	m.ClosingTag = closingTag

	if closeErr := readBuffer.CloseContext("BACnetServiceAckAtomicReadFileStreamOrRecord"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetServiceAckAtomicReadFileStreamOrRecord")
	}

	return _child, nil
}

func (pm *_BACnetServiceAckAtomicReadFileStreamOrRecord) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetServiceAckAtomicReadFileStreamOrRecord, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetServiceAckAtomicReadFileStreamOrRecord"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetServiceAckAtomicReadFileStreamOrRecord")
	}

	if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'openingTag' field")
	}
	// Virtual field
	peekedTagNumber := m.GetPeekedTagNumber()
	_ = peekedTagNumber
	if _peekedTagNumberErr := writeBuffer.WriteVirtual(ctx, "peekedTagNumber", m.GetPeekedTagNumber()); _peekedTagNumberErr != nil {
		return errors.Wrap(_peekedTagNumberErr, "Error serializing 'peekedTagNumber' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetServiceAckAtomicReadFileStreamOrRecord"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetServiceAckAtomicReadFileStreamOrRecord")
	}
	return nil
}

func (m *_BACnetServiceAckAtomicReadFileStreamOrRecord) IsBACnetServiceAckAtomicReadFileStreamOrRecord() {
}
