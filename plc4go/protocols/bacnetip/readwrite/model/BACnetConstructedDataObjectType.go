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

// BACnetConstructedDataObjectType is the corresponding interface of BACnetConstructedDataObjectType
type BACnetConstructedDataObjectType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetObjectType returns ObjectType (property field)
	GetObjectType() BACnetObjectTypeTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetObjectTypeTagged
}

// BACnetConstructedDataObjectTypeExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataObjectType.
// This is useful for switch cases.
type BACnetConstructedDataObjectTypeExactly interface {
	BACnetConstructedDataObjectType
	isBACnetConstructedDataObjectType() bool
}

// _BACnetConstructedDataObjectType is the data-structure of this message
type _BACnetConstructedDataObjectType struct {
	BACnetConstructedDataContract
	ObjectType BACnetObjectTypeTagged
}

var _ BACnetConstructedDataObjectType = (*_BACnetConstructedDataObjectType)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataObjectType)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataObjectType) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataObjectType) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_OBJECT_TYPE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataObjectType) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataObjectType) GetObjectType() BACnetObjectTypeTagged {
	return m.ObjectType
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataObjectType) GetActualValue() BACnetObjectTypeTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetObjectTypeTagged(m.GetObjectType())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataObjectType factory function for _BACnetConstructedDataObjectType
func NewBACnetConstructedDataObjectType(objectType BACnetObjectTypeTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataObjectType {
	_result := &_BACnetConstructedDataObjectType{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		ObjectType:                    objectType,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataObjectType(structType any) BACnetConstructedDataObjectType {
	if casted, ok := structType.(BACnetConstructedDataObjectType); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataObjectType); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataObjectType) GetTypeName() string {
	return "BACnetConstructedDataObjectType"
}

func (m *_BACnetConstructedDataObjectType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (objectType)
	lengthInBits += m.ObjectType.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataObjectType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataObjectType) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataObjectType BACnetConstructedDataObjectType, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataObjectType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataObjectType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	objectType, err := ReadSimpleField[BACnetObjectTypeTagged](ctx, "objectType", ReadComplex[BACnetObjectTypeTagged](BACnetObjectTypeTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'objectType' field"))
	}
	m.ObjectType = objectType

	actualValue, err := ReadVirtualField[BACnetObjectTypeTagged](ctx, "actualValue", (*BACnetObjectTypeTagged)(nil), objectType)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataObjectType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataObjectType")
	}

	return m, nil
}

func (m *_BACnetConstructedDataObjectType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataObjectType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataObjectType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataObjectType")
		}

		if err := WriteSimpleField[BACnetObjectTypeTagged](ctx, "objectType", m.GetObjectType(), WriteComplex[BACnetObjectTypeTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'objectType' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataObjectType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataObjectType")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataObjectType) isBACnetConstructedDataObjectType() bool {
	return true
}

func (m *_BACnetConstructedDataObjectType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
