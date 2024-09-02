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

// BACnetNotificationParameters is the corresponding interface of BACnetNotificationParameters
type BACnetNotificationParameters interface {
	BACnetNotificationParametersContract
	BACnetNotificationParametersRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

// BACnetNotificationParametersContract provides a set of functions which can be overwritten by a sub struct
type BACnetNotificationParametersContract interface {
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
	// GetObjectTypeArgument() returns a parser argument
	GetObjectTypeArgument() BACnetObjectType
}

// BACnetNotificationParametersRequirements provides a set of functions which need to be implemented by a sub struct
type BACnetNotificationParametersRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetPeekedTagNumber returns PeekedTagNumber (discriminator field)
	GetPeekedTagNumber() uint8
}

// BACnetNotificationParametersExactly can be used when we want exactly this type and not a type which fulfills BACnetNotificationParameters.
// This is useful for switch cases.
type BACnetNotificationParametersExactly interface {
	BACnetNotificationParameters
	isBACnetNotificationParameters() bool
}

// _BACnetNotificationParameters is the data-structure of this message
type _BACnetNotificationParameters struct {
	_SubType        BACnetNotificationParameters
	OpeningTag      BACnetOpeningTag
	PeekedTagHeader BACnetTagHeader
	ClosingTag      BACnetClosingTag

	// Arguments.
	TagNumber          uint8
	ObjectTypeArgument BACnetObjectType
}

var _ BACnetNotificationParametersContract = (*_BACnetNotificationParameters)(nil)

type BACnetNotificationParametersChild interface {
	utils.Serializable

	GetParent() *BACnetNotificationParameters

	GetTypeName() string
	BACnetNotificationParameters
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetNotificationParameters) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetNotificationParameters) GetPeekedTagHeader() BACnetTagHeader {
	return m.PeekedTagHeader
}

func (m *_BACnetNotificationParameters) GetClosingTag() BACnetClosingTag {
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

func (pm *_BACnetNotificationParameters) GetPeekedTagNumber() uint8 {
	m := pm._SubType
	ctx := context.Background()
	_ = ctx
	return uint8(m.GetPeekedTagHeader().GetActualTagNumber())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetNotificationParameters factory function for _BACnetNotificationParameters
func NewBACnetNotificationParameters(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, objectTypeArgument BACnetObjectType) *_BACnetNotificationParameters {
	return &_BACnetNotificationParameters{OpeningTag: openingTag, PeekedTagHeader: peekedTagHeader, ClosingTag: closingTag, TagNumber: tagNumber, ObjectTypeArgument: objectTypeArgument}
}

// Deprecated: use the interface for direct cast
func CastBACnetNotificationParameters(structType any) BACnetNotificationParameters {
	if casted, ok := structType.(BACnetNotificationParameters); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetNotificationParameters); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetNotificationParameters) GetTypeName() string {
	return "BACnetNotificationParameters"
}

func (m *_BACnetNotificationParameters) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetNotificationParameters) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func BACnetNotificationParametersParse[T BACnetNotificationParameters](ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType) (T, error) {
	return BACnetNotificationParametersParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument)
}

func BACnetNotificationParametersParseWithBufferProducer[T BACnetNotificationParameters](tagNumber uint8, objectTypeArgument BACnetObjectType) func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := BACnetNotificationParametersParseWithBuffer[T](ctx, readBuffer, tagNumber, objectTypeArgument)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, err
	}
}

func BACnetNotificationParametersParseWithBuffer[T BACnetNotificationParameters](ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType) (T, error) {
	v, err := (&_BACnetNotificationParameters{}).parse(ctx, readBuffer, tagNumber, objectTypeArgument)
	if err != nil {
		var zero T
		return zero, err
	}
	return v.(T), err
}

func (m *_BACnetNotificationParameters) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType) (__bACnetNotificationParameters BACnetNotificationParameters, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetNotificationParameters"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetNotificationParameters")
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
	var _child BACnetNotificationParameters
	switch {
	case peekedTagNumber == uint8(0): // BACnetNotificationParametersChangeOfBitString
		if _child, err = (&_BACnetNotificationParametersChangeOfBitString{}).parse(ctx, readBuffer, m, peekedTagNumber, tagNumber, objectTypeArgument); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetNotificationParametersChangeOfBitString for type-switch of BACnetNotificationParameters")
		}
	case peekedTagNumber == uint8(1): // BACnetNotificationParametersChangeOfState
		if _child, err = (&_BACnetNotificationParametersChangeOfState{}).parse(ctx, readBuffer, m, peekedTagNumber, tagNumber, objectTypeArgument); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetNotificationParametersChangeOfState for type-switch of BACnetNotificationParameters")
		}
	case peekedTagNumber == uint8(2): // BACnetNotificationParametersChangeOfValue
		if _child, err = (&_BACnetNotificationParametersChangeOfValue{}).parse(ctx, readBuffer, m, peekedTagNumber, tagNumber, objectTypeArgument); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetNotificationParametersChangeOfValue for type-switch of BACnetNotificationParameters")
		}
	case peekedTagNumber == uint8(3): // BACnetNotificationParametersCommandFailure
		if _child, err = (&_BACnetNotificationParametersCommandFailure{}).parse(ctx, readBuffer, m, peekedTagNumber, tagNumber, objectTypeArgument); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetNotificationParametersCommandFailure for type-switch of BACnetNotificationParameters")
		}
	case peekedTagNumber == uint8(4): // BACnetNotificationParametersFloatingLimit
		if _child, err = (&_BACnetNotificationParametersFloatingLimit{}).parse(ctx, readBuffer, m, peekedTagNumber, tagNumber, objectTypeArgument); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetNotificationParametersFloatingLimit for type-switch of BACnetNotificationParameters")
		}
	case peekedTagNumber == uint8(5): // BACnetNotificationParametersOutOfRange
		if _child, err = (&_BACnetNotificationParametersOutOfRange{}).parse(ctx, readBuffer, m, peekedTagNumber, tagNumber, objectTypeArgument); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetNotificationParametersOutOfRange for type-switch of BACnetNotificationParameters")
		}
	case peekedTagNumber == uint8(6): // BACnetNotificationParametersComplexEventType
		if _child, err = (&_BACnetNotificationParametersComplexEventType{}).parse(ctx, readBuffer, m, peekedTagNumber, tagNumber, objectTypeArgument); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetNotificationParametersComplexEventType for type-switch of BACnetNotificationParameters")
		}
	case peekedTagNumber == uint8(8): // BACnetNotificationParametersChangeOfLifeSafety
		if _child, err = (&_BACnetNotificationParametersChangeOfLifeSafety{}).parse(ctx, readBuffer, m, peekedTagNumber, tagNumber, objectTypeArgument); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetNotificationParametersChangeOfLifeSafety for type-switch of BACnetNotificationParameters")
		}
	case peekedTagNumber == uint8(9): // BACnetNotificationParametersExtended
		if _child, err = (&_BACnetNotificationParametersExtended{}).parse(ctx, readBuffer, m, peekedTagNumber, tagNumber, objectTypeArgument); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetNotificationParametersExtended for type-switch of BACnetNotificationParameters")
		}
	case peekedTagNumber == uint8(10): // BACnetNotificationParametersBufferReady
		if _child, err = (&_BACnetNotificationParametersBufferReady{}).parse(ctx, readBuffer, m, peekedTagNumber, tagNumber, objectTypeArgument); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetNotificationParametersBufferReady for type-switch of BACnetNotificationParameters")
		}
	case peekedTagNumber == uint8(11): // BACnetNotificationParametersUnsignedRange
		if _child, err = (&_BACnetNotificationParametersUnsignedRange{}).parse(ctx, readBuffer, m, peekedTagNumber, tagNumber, objectTypeArgument); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetNotificationParametersUnsignedRange for type-switch of BACnetNotificationParameters")
		}
	case peekedTagNumber == uint8(13): // BACnetNotificationParametersAccessEvent
		if _child, err = (&_BACnetNotificationParametersAccessEvent{}).parse(ctx, readBuffer, m, peekedTagNumber, tagNumber, objectTypeArgument); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetNotificationParametersAccessEvent for type-switch of BACnetNotificationParameters")
		}
	case peekedTagNumber == uint8(14): // BACnetNotificationParametersDoubleOutOfRange
		if _child, err = (&_BACnetNotificationParametersDoubleOutOfRange{}).parse(ctx, readBuffer, m, peekedTagNumber, tagNumber, objectTypeArgument); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetNotificationParametersDoubleOutOfRange for type-switch of BACnetNotificationParameters")
		}
	case peekedTagNumber == uint8(15): // BACnetNotificationParametersSignedOutOfRange
		if _child, err = (&_BACnetNotificationParametersSignedOutOfRange{}).parse(ctx, readBuffer, m, peekedTagNumber, tagNumber, objectTypeArgument); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetNotificationParametersSignedOutOfRange for type-switch of BACnetNotificationParameters")
		}
	case peekedTagNumber == uint8(16): // BACnetNotificationParametersUnsignedOutOfRange
		if _child, err = (&_BACnetNotificationParametersUnsignedOutOfRange{}).parse(ctx, readBuffer, m, peekedTagNumber, tagNumber, objectTypeArgument); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetNotificationParametersUnsignedOutOfRange for type-switch of BACnetNotificationParameters")
		}
	case peekedTagNumber == uint8(17): // BACnetNotificationParametersChangeOfCharacterString
		if _child, err = (&_BACnetNotificationParametersChangeOfCharacterString{}).parse(ctx, readBuffer, m, peekedTagNumber, tagNumber, objectTypeArgument); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetNotificationParametersChangeOfCharacterString for type-switch of BACnetNotificationParameters")
		}
	case peekedTagNumber == uint8(18): // BACnetNotificationParametersChangeOfStatusFlags
		if _child, err = (&_BACnetNotificationParametersChangeOfStatusFlags{}).parse(ctx, readBuffer, m, peekedTagNumber, tagNumber, objectTypeArgument); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetNotificationParametersChangeOfStatusFlags for type-switch of BACnetNotificationParameters")
		}
	case peekedTagNumber == uint8(19): // BACnetNotificationParametersChangeOfReliability
		if _child, err = (&_BACnetNotificationParametersChangeOfReliability{}).parse(ctx, readBuffer, m, peekedTagNumber, tagNumber, objectTypeArgument); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetNotificationParametersChangeOfReliability for type-switch of BACnetNotificationParameters")
		}
	case peekedTagNumber == uint8(21): // BACnetNotificationParametersChangeOfDiscreteValue
		if _child, err = (&_BACnetNotificationParametersChangeOfDiscreteValue{}).parse(ctx, readBuffer, m, peekedTagNumber, tagNumber, objectTypeArgument); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetNotificationParametersChangeOfDiscreteValue for type-switch of BACnetNotificationParameters")
		}
	case peekedTagNumber == uint8(22): // BACnetNotificationParametersChangeOfTimer
		if _child, err = (&_BACnetNotificationParametersChangeOfTimer{}).parse(ctx, readBuffer, m, peekedTagNumber, tagNumber, objectTypeArgument); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetNotificationParametersChangeOfTimer for type-switch of BACnetNotificationParameters")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [peekedTagNumber=%v]", peekedTagNumber)
	}

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	m.ClosingTag = closingTag

	if closeErr := readBuffer.CloseContext("BACnetNotificationParameters"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetNotificationParameters")
	}

	return _child, nil
}

func (pm *_BACnetNotificationParameters) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetNotificationParameters, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetNotificationParameters"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetNotificationParameters")
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

	if popErr := writeBuffer.PopContext("BACnetNotificationParameters"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetNotificationParameters")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetNotificationParameters) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_BACnetNotificationParameters) GetObjectTypeArgument() BACnetObjectType {
	return m.ObjectTypeArgument
}

//
////

func (m *_BACnetNotificationParameters) isBACnetNotificationParameters() bool {
	return true
}
