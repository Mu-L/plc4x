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

// BACnetConstructedDataDescriptionOfHalt is the corresponding interface of BACnetConstructedDataDescriptionOfHalt
type BACnetConstructedDataDescriptionOfHalt interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetDescriptionForHalt returns DescriptionForHalt (property field)
	GetDescriptionForHalt() BACnetApplicationTagCharacterString
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagCharacterString
}

// BACnetConstructedDataDescriptionOfHaltExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataDescriptionOfHalt.
// This is useful for switch cases.
type BACnetConstructedDataDescriptionOfHaltExactly interface {
	BACnetConstructedDataDescriptionOfHalt
	isBACnetConstructedDataDescriptionOfHalt() bool
}

// _BACnetConstructedDataDescriptionOfHalt is the data-structure of this message
type _BACnetConstructedDataDescriptionOfHalt struct {
	*_BACnetConstructedData
	DescriptionForHalt BACnetApplicationTagCharacterString
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataDescriptionOfHalt) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataDescriptionOfHalt) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_DESCRIPTION_OF_HALT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataDescriptionOfHalt) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataDescriptionOfHalt) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataDescriptionOfHalt) GetDescriptionForHalt() BACnetApplicationTagCharacterString {
	return m.DescriptionForHalt
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataDescriptionOfHalt) GetActualValue() BACnetApplicationTagCharacterString {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagCharacterString(m.GetDescriptionForHalt())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataDescriptionOfHalt factory function for _BACnetConstructedDataDescriptionOfHalt
func NewBACnetConstructedDataDescriptionOfHalt(descriptionForHalt BACnetApplicationTagCharacterString, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataDescriptionOfHalt {
	_result := &_BACnetConstructedDataDescriptionOfHalt{
		DescriptionForHalt:     descriptionForHalt,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataDescriptionOfHalt(structType any) BACnetConstructedDataDescriptionOfHalt {
	if casted, ok := structType.(BACnetConstructedDataDescriptionOfHalt); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataDescriptionOfHalt); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataDescriptionOfHalt) GetTypeName() string {
	return "BACnetConstructedDataDescriptionOfHalt"
}

func (m *_BACnetConstructedDataDescriptionOfHalt) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (descriptionForHalt)
	lengthInBits += m.DescriptionForHalt.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataDescriptionOfHalt) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataDescriptionOfHaltParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataDescriptionOfHalt, error) {
	return BACnetConstructedDataDescriptionOfHaltParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataDescriptionOfHaltParseWithBufferProducer(tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetConstructedDataDescriptionOfHalt, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetConstructedDataDescriptionOfHalt, error) {
		return BACnetConstructedDataDescriptionOfHaltParseWithBuffer(ctx, readBuffer, tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
	}
}

func BACnetConstructedDataDescriptionOfHaltParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataDescriptionOfHalt, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetConstructedDataDescriptionOfHalt"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataDescriptionOfHalt")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	descriptionForHalt, err := ReadSimpleField[BACnetApplicationTagCharacterString](ctx, "descriptionForHalt", ReadComplex[BACnetApplicationTagCharacterString](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagCharacterString](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'descriptionForHalt' field"))
	}

	// Virtual field
	_actualValue := descriptionForHalt
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataDescriptionOfHalt"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataDescriptionOfHalt")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataDescriptionOfHalt{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		DescriptionForHalt: descriptionForHalt,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataDescriptionOfHalt) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataDescriptionOfHalt) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataDescriptionOfHalt"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataDescriptionOfHalt")
		}

		// Simple Field (descriptionForHalt)
		if pushErr := writeBuffer.PushContext("descriptionForHalt"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for descriptionForHalt")
		}
		_descriptionForHaltErr := writeBuffer.WriteSerializable(ctx, m.GetDescriptionForHalt())
		if popErr := writeBuffer.PopContext("descriptionForHalt"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for descriptionForHalt")
		}
		if _descriptionForHaltErr != nil {
			return errors.Wrap(_descriptionForHaltErr, "Error serializing 'descriptionForHalt' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataDescriptionOfHalt"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataDescriptionOfHalt")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataDescriptionOfHalt) isBACnetConstructedDataDescriptionOfHalt() bool {
	return true
}

func (m *_BACnetConstructedDataDescriptionOfHalt) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
