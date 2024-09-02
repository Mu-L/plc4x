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

// BACnetVendorIdTagged is the corresponding interface of BACnetVendorIdTagged
type BACnetVendorIdTagged interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetValue returns Value (property field)
	GetValue() BACnetVendorId
	// GetUnknownId returns UnknownId (property field)
	GetUnknownId() uint32
	// GetIsUnknownId returns IsUnknownId (virtual field)
	GetIsUnknownId() bool
	// IsBACnetVendorIdTagged is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetVendorIdTagged()
}

// _BACnetVendorIdTagged is the data-structure of this message
type _BACnetVendorIdTagged struct {
	Header    BACnetTagHeader
	Value     BACnetVendorId
	UnknownId uint32

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

var _ BACnetVendorIdTagged = (*_BACnetVendorIdTagged)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetVendorIdTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_BACnetVendorIdTagged) GetValue() BACnetVendorId {
	return m.Value
}

func (m *_BACnetVendorIdTagged) GetUnknownId() uint32 {
	return m.UnknownId
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetVendorIdTagged) GetIsUnknownId() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetValue()) == (BACnetVendorId_UNKNOWN_VENDOR)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetVendorIdTagged factory function for _BACnetVendorIdTagged
func NewBACnetVendorIdTagged(header BACnetTagHeader, value BACnetVendorId, unknownId uint32, tagNumber uint8, tagClass TagClass) *_BACnetVendorIdTagged {
	return &_BACnetVendorIdTagged{Header: header, Value: value, UnknownId: unknownId, TagNumber: tagNumber, TagClass: tagClass}
}

// Deprecated: use the interface for direct cast
func CastBACnetVendorIdTagged(structType any) BACnetVendorIdTagged {
	if casted, ok := structType.(BACnetVendorIdTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetVendorIdTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetVendorIdTagged) GetTypeName() string {
	return "BACnetVendorIdTagged"
}

func (m *_BACnetVendorIdTagged) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits(ctx)

	// Manual Field (value)
	lengthInBits += uint16(utils.InlineIf(m.GetIsUnknownId(), func() any { return int32(int32(0)) }, func() any { return int32((int32(m.GetHeader().GetActualLength()) * int32(int32(8)))) }).(int32))

	// A virtual field doesn't have any in- or output.

	// Manual Field (unknownId)
	lengthInBits += uint16(utils.InlineIf(m.GetIsUnknownId(), func() any { return int32((int32(m.GetHeader().GetActualLength()) * int32(int32(8)))) }, func() any { return int32(int32(0)) }).(int32))

	return lengthInBits
}

func (m *_BACnetVendorIdTagged) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetVendorIdTaggedParse(ctx context.Context, theBytes []byte, tagNumber uint8, tagClass TagClass) (BACnetVendorIdTagged, error) {
	return BACnetVendorIdTaggedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, tagClass)
}

func BACnetVendorIdTaggedParseWithBufferProducer(tagNumber uint8, tagClass TagClass) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetVendorIdTagged, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetVendorIdTagged, error) {
		return BACnetVendorIdTaggedParseWithBuffer(ctx, readBuffer, tagNumber, tagClass)
	}
}

func BACnetVendorIdTaggedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (BACnetVendorIdTagged, error) {
	v, err := (&_BACnetVendorIdTagged{TagNumber: tagNumber, TagClass: tagClass}).parse(ctx, readBuffer, tagNumber, tagClass)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *_BACnetVendorIdTagged) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (__bACnetVendorIdTagged BACnetVendorIdTagged, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetVendorIdTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetVendorIdTagged")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	header, err := ReadSimpleField[BACnetTagHeader](ctx, "header", ReadComplex[BACnetTagHeader](BACnetTagHeaderParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'header' field"))
	}
	m.Header = header

	// Validation
	if !(bool((header.GetTagClass()) == (tagClass))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "tag class doesn't match"})
	}

	// Validation
	if !(bool((bool(bool((header.GetTagClass()) == (TagClass_APPLICATION_TAGS))) && bool(bool((header.GetActualTagNumber()) == (2))))) || bool((bool((header.GetActualTagNumber()) == (tagNumber))))) {
		return nil, errors.WithStack(utils.ParseAssertError{Message: "tagnumber doesn't match"})
	}

	value, err := ReadManualField[BACnetVendorId](ctx, "value", readBuffer, EnsureType[BACnetVendorId](ReadEnumGeneric(ctx, readBuffer, header.GetActualLength(), BACnetVendorId_UNKNOWN_VENDOR)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	isUnknownId, err := ReadVirtualField[bool](ctx, "isUnknownId", (*bool)(nil), bool((value) == (BACnetVendorId_UNKNOWN_VENDOR)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isUnknownId' field"))
	}
	_ = isUnknownId

	unknownId, err := ReadManualField[uint32](ctx, "unknownId", readBuffer, EnsureType[uint32](ReadProprietaryEnumGeneric(ctx, readBuffer, header.GetActualLength(), isUnknownId)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'unknownId' field"))
	}
	m.UnknownId = unknownId

	if closeErr := readBuffer.CloseContext("BACnetVendorIdTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetVendorIdTagged")
	}

	return m, nil
}

func (m *_BACnetVendorIdTagged) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetVendorIdTagged) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetVendorIdTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetVendorIdTagged")
	}

	if err := WriteSimpleField[BACnetTagHeader](ctx, "header", m.GetHeader(), WriteComplex[BACnetTagHeader](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'header' field")
	}

	if err := WriteManualField[BACnetVendorId](ctx, "value", func(ctx context.Context) error { return WriteEnumGeneric(ctx, writeBuffer, m.GetValue()) }, writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'value' field")
	}
	// Virtual field
	isUnknownId := m.GetIsUnknownId()
	_ = isUnknownId
	if _isUnknownIdErr := writeBuffer.WriteVirtual(ctx, "isUnknownId", m.GetIsUnknownId()); _isUnknownIdErr != nil {
		return errors.Wrap(_isUnknownIdErr, "Error serializing 'isUnknownId' field")
	}

	if err := WriteManualField[uint32](ctx, "unknownId", func(ctx context.Context) error {
		return WriteProprietaryEnumGeneric(ctx, writeBuffer, m.GetUnknownId(), m.GetIsUnknownId())
	}, writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'unknownId' field")
	}

	if popErr := writeBuffer.PopContext("BACnetVendorIdTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetVendorIdTagged")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetVendorIdTagged) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_BACnetVendorIdTagged) GetTagClass() TagClass {
	return m.TagClass
}

//
////

func (m *_BACnetVendorIdTagged) IsBACnetVendorIdTagged() {}

func (m *_BACnetVendorIdTagged) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
