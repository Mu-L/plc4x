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

// BACnetConstructedDataNextStoppingFloor is the corresponding interface of BACnetConstructedDataNextStoppingFloor
type BACnetConstructedDataNextStoppingFloor interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetNextStoppingFloor returns NextStoppingFloor (property field)
	GetNextStoppingFloor() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataNextStoppingFloorExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataNextStoppingFloor.
// This is useful for switch cases.
type BACnetConstructedDataNextStoppingFloorExactly interface {
	BACnetConstructedDataNextStoppingFloor
	isBACnetConstructedDataNextStoppingFloor() bool
}

// _BACnetConstructedDataNextStoppingFloor is the data-structure of this message
type _BACnetConstructedDataNextStoppingFloor struct {
	*_BACnetConstructedData
	NextStoppingFloor BACnetApplicationTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataNextStoppingFloor) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataNextStoppingFloor) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_NEXT_STOPPING_FLOOR
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataNextStoppingFloor) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataNextStoppingFloor) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataNextStoppingFloor) GetNextStoppingFloor() BACnetApplicationTagUnsignedInteger {
	return m.NextStoppingFloor
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataNextStoppingFloor) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetNextStoppingFloor())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataNextStoppingFloor factory function for _BACnetConstructedDataNextStoppingFloor
func NewBACnetConstructedDataNextStoppingFloor(nextStoppingFloor BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataNextStoppingFloor {
	_result := &_BACnetConstructedDataNextStoppingFloor{
		NextStoppingFloor:      nextStoppingFloor,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataNextStoppingFloor(structType any) BACnetConstructedDataNextStoppingFloor {
	if casted, ok := structType.(BACnetConstructedDataNextStoppingFloor); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataNextStoppingFloor); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataNextStoppingFloor) GetTypeName() string {
	return "BACnetConstructedDataNextStoppingFloor"
}

func (m *_BACnetConstructedDataNextStoppingFloor) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (nextStoppingFloor)
	lengthInBits += m.NextStoppingFloor.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataNextStoppingFloor) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataNextStoppingFloorParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataNextStoppingFloor, error) {
	return BACnetConstructedDataNextStoppingFloorParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataNextStoppingFloorParseWithBufferProducer(tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetConstructedDataNextStoppingFloor, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetConstructedDataNextStoppingFloor, error) {
		return BACnetConstructedDataNextStoppingFloorParseWithBuffer(ctx, readBuffer, tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
	}
}

func BACnetConstructedDataNextStoppingFloorParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataNextStoppingFloor, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetConstructedDataNextStoppingFloor"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataNextStoppingFloor")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	nextStoppingFloor, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "nextStoppingFloor", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'nextStoppingFloor' field"))
	}

	// Virtual field
	_actualValue := nextStoppingFloor
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataNextStoppingFloor"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataNextStoppingFloor")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataNextStoppingFloor{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		NextStoppingFloor: nextStoppingFloor,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataNextStoppingFloor) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataNextStoppingFloor) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataNextStoppingFloor"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataNextStoppingFloor")
		}

		// Simple Field (nextStoppingFloor)
		if pushErr := writeBuffer.PushContext("nextStoppingFloor"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for nextStoppingFloor")
		}
		_nextStoppingFloorErr := writeBuffer.WriteSerializable(ctx, m.GetNextStoppingFloor())
		if popErr := writeBuffer.PopContext("nextStoppingFloor"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for nextStoppingFloor")
		}
		if _nextStoppingFloorErr != nil {
			return errors.Wrap(_nextStoppingFloorErr, "Error serializing 'nextStoppingFloor' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataNextStoppingFloor"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataNextStoppingFloor")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataNextStoppingFloor) isBACnetConstructedDataNextStoppingFloor() bool {
	return true
}

func (m *_BACnetConstructedDataNextStoppingFloor) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
