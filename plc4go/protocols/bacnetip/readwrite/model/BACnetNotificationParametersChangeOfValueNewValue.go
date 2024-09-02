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

// BACnetNotificationParametersChangeOfValueNewValue is the corresponding interface of BACnetNotificationParametersChangeOfValueNewValue
type BACnetNotificationParametersChangeOfValueNewValue interface {
	BACnetNotificationParametersChangeOfValueNewValueContract
	BACnetNotificationParametersChangeOfValueNewValueRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

// BACnetNotificationParametersChangeOfValueNewValueContract provides a set of functions which can be overwritten by a sub struct
type BACnetNotificationParametersChangeOfValueNewValueContract interface {
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetPeekedTagHeader returns PeekedTagHeader (property field)
	GetPeekedTagHeader() BACnetTagHeader
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
	// GetPeekedTagNumber returns PeekedTagNumber (virtual field)
	GetPeekedTagNumber() uint8
	// GetTagNumber() returns a parser argument
	GetTagNumber() uint8
}

// BACnetNotificationParametersChangeOfValueNewValueRequirements provides a set of functions which need to be implemented by a sub struct
type BACnetNotificationParametersChangeOfValueNewValueRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetPeekedTagNumber returns PeekedTagNumber (discriminator field)
	GetPeekedTagNumber() uint8
}

// BACnetNotificationParametersChangeOfValueNewValueExactly can be used when we want exactly this type and not a type which fulfills BACnetNotificationParametersChangeOfValueNewValue.
// This is useful for switch cases.
type BACnetNotificationParametersChangeOfValueNewValueExactly interface {
	BACnetNotificationParametersChangeOfValueNewValue
	isBACnetNotificationParametersChangeOfValueNewValue() bool
}

// _BACnetNotificationParametersChangeOfValueNewValue is the data-structure of this message
type _BACnetNotificationParametersChangeOfValueNewValue struct {
	_SubType        BACnetNotificationParametersChangeOfValueNewValue
	OpeningTag      BACnetOpeningTag
	PeekedTagHeader BACnetTagHeader
	ClosingTag      BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

var _ BACnetNotificationParametersChangeOfValueNewValueContract = (*_BACnetNotificationParametersChangeOfValueNewValue)(nil)

type BACnetNotificationParametersChangeOfValueNewValueChild interface {
	utils.Serializable

	GetParent() *BACnetNotificationParametersChangeOfValueNewValue

	GetTypeName() string
	BACnetNotificationParametersChangeOfValueNewValue
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetNotificationParametersChangeOfValueNewValue) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetNotificationParametersChangeOfValueNewValue) GetPeekedTagHeader() BACnetTagHeader {
	return m.PeekedTagHeader
}

func (m *_BACnetNotificationParametersChangeOfValueNewValue) GetClosingTag() BACnetClosingTag {
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

func (pm *_BACnetNotificationParametersChangeOfValueNewValue) GetPeekedTagNumber() uint8 {
	m := pm._SubType
	ctx := context.Background()
	_ = ctx
	return uint8(m.GetPeekedTagHeader().GetActualTagNumber())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetNotificationParametersChangeOfValueNewValue factory function for _BACnetNotificationParametersChangeOfValueNewValue
func NewBACnetNotificationParametersChangeOfValueNewValue(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetNotificationParametersChangeOfValueNewValue {
	return &_BACnetNotificationParametersChangeOfValueNewValue{OpeningTag: openingTag, PeekedTagHeader: peekedTagHeader, ClosingTag: closingTag, TagNumber: tagNumber}
}

// Deprecated: use the interface for direct cast
func CastBACnetNotificationParametersChangeOfValueNewValue(structType any) BACnetNotificationParametersChangeOfValueNewValue {
	if casted, ok := structType.(BACnetNotificationParametersChangeOfValueNewValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetNotificationParametersChangeOfValueNewValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetNotificationParametersChangeOfValueNewValue) GetTypeName() string {
	return "BACnetNotificationParametersChangeOfValueNewValue"
}

func (m *_BACnetNotificationParametersChangeOfValueNewValue) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetNotificationParametersChangeOfValueNewValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func BACnetNotificationParametersChangeOfValueNewValueParse[T BACnetNotificationParametersChangeOfValueNewValue](ctx context.Context, theBytes []byte, tagNumber uint8) (T, error) {
	return BACnetNotificationParametersChangeOfValueNewValueParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetNotificationParametersChangeOfValueNewValueParseWithBufferProducer[T BACnetNotificationParametersChangeOfValueNewValue](tagNumber uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := BACnetNotificationParametersChangeOfValueNewValueParseWithBuffer[T](ctx, readBuffer, tagNumber)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, err
	}
}

func BACnetNotificationParametersChangeOfValueNewValueParseWithBuffer[T BACnetNotificationParametersChangeOfValueNewValue](ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (T, error) {
	v, err := (&_BACnetNotificationParametersChangeOfValueNewValue{}).parse(ctx, readBuffer, tagNumber)
	if err != nil {
		var zero T
		return zero, err
	}
	return v.(T), err
}

func (m *_BACnetNotificationParametersChangeOfValueNewValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (__bACnetNotificationParametersChangeOfValueNewValue BACnetNotificationParametersChangeOfValueNewValue, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetNotificationParametersChangeOfValueNewValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetNotificationParametersChangeOfValueNewValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}
	m.OpeningTag = openingTag

	peekedTagHeader, err := ReadPeekField[BACnetTagHeader](ctx, "peekedTagHeader", ReadComplex[BACnetTagHeader](BACnetTagHeaderParseWithBuffer, readBuffer), 0)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'peekedTagHeader' field"))
	}
	m.PeekedTagHeader = peekedTagHeader

	peekedTagNumber, err := ReadVirtualField[uint8](ctx, "peekedTagNumber", (*uint8)(nil), peekedTagHeader.GetActualTagNumber())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'peekedTagNumber' field"))
	}
	_ = peekedTagNumber

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child BACnetNotificationParametersChangeOfValueNewValue
	switch {
	case peekedTagNumber == uint8(0): // BACnetNotificationParametersChangeOfValueNewValueChangedBits
		if _child, err = (&_BACnetNotificationParametersChangeOfValueNewValueChangedBits{}).parse(ctx, readBuffer, m, peekedTagNumber, tagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetNotificationParametersChangeOfValueNewValueChangedBits for type-switch of BACnetNotificationParametersChangeOfValueNewValue")
		}
	case peekedTagNumber == uint8(1): // BACnetNotificationParametersChangeOfValueNewValueChangedValue
		if _child, err = (&_BACnetNotificationParametersChangeOfValueNewValueChangedValue{}).parse(ctx, readBuffer, m, peekedTagNumber, tagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetNotificationParametersChangeOfValueNewValueChangedValue for type-switch of BACnetNotificationParametersChangeOfValueNewValue")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [peekedTagNumber=%v]", peekedTagNumber)
	}

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	m.ClosingTag = closingTag

	if closeErr := readBuffer.CloseContext("BACnetNotificationParametersChangeOfValueNewValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetNotificationParametersChangeOfValueNewValue")
	}

	return _child, nil
}

func (pm *_BACnetNotificationParametersChangeOfValueNewValue) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetNotificationParametersChangeOfValueNewValue, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetNotificationParametersChangeOfValueNewValue"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetNotificationParametersChangeOfValueNewValue")
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

	if popErr := writeBuffer.PopContext("BACnetNotificationParametersChangeOfValueNewValue"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetNotificationParametersChangeOfValueNewValue")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetNotificationParametersChangeOfValueNewValue) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_BACnetNotificationParametersChangeOfValueNewValue) isBACnetNotificationParametersChangeOfValueNewValue() bool {
	return true
}
