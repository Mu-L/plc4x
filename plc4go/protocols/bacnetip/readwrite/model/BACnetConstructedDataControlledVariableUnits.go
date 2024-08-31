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

// BACnetConstructedDataControlledVariableUnits is the corresponding interface of BACnetConstructedDataControlledVariableUnits
type BACnetConstructedDataControlledVariableUnits interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetUnits returns Units (property field)
	GetUnits() BACnetEngineeringUnitsTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetEngineeringUnitsTagged
}

// BACnetConstructedDataControlledVariableUnitsExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataControlledVariableUnits.
// This is useful for switch cases.
type BACnetConstructedDataControlledVariableUnitsExactly interface {
	BACnetConstructedDataControlledVariableUnits
	isBACnetConstructedDataControlledVariableUnits() bool
}

// _BACnetConstructedDataControlledVariableUnits is the data-structure of this message
type _BACnetConstructedDataControlledVariableUnits struct {
	*_BACnetConstructedData
	Units BACnetEngineeringUnitsTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataControlledVariableUnits) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataControlledVariableUnits) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_CONTROLLED_VARIABLE_UNITS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataControlledVariableUnits) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataControlledVariableUnits) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataControlledVariableUnits) GetUnits() BACnetEngineeringUnitsTagged {
	return m.Units
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataControlledVariableUnits) GetActualValue() BACnetEngineeringUnitsTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetEngineeringUnitsTagged(m.GetUnits())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataControlledVariableUnits factory function for _BACnetConstructedDataControlledVariableUnits
func NewBACnetConstructedDataControlledVariableUnits(units BACnetEngineeringUnitsTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataControlledVariableUnits {
	_result := &_BACnetConstructedDataControlledVariableUnits{
		Units:                  units,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataControlledVariableUnits(structType any) BACnetConstructedDataControlledVariableUnits {
	if casted, ok := structType.(BACnetConstructedDataControlledVariableUnits); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataControlledVariableUnits); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataControlledVariableUnits) GetTypeName() string {
	return "BACnetConstructedDataControlledVariableUnits"
}

func (m *_BACnetConstructedDataControlledVariableUnits) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (units)
	lengthInBits += m.Units.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataControlledVariableUnits) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataControlledVariableUnitsParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataControlledVariableUnits, error) {
	return BACnetConstructedDataControlledVariableUnitsParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataControlledVariableUnitsParseWithBufferProducer(tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetConstructedDataControlledVariableUnits, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetConstructedDataControlledVariableUnits, error) {
		return BACnetConstructedDataControlledVariableUnitsParseWithBuffer(ctx, readBuffer, tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
	}
}

func BACnetConstructedDataControlledVariableUnitsParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataControlledVariableUnits, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetConstructedDataControlledVariableUnits"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataControlledVariableUnits")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	units, err := ReadSimpleField[BACnetEngineeringUnitsTagged](ctx, "units", ReadComplex[BACnetEngineeringUnitsTagged](BACnetEngineeringUnitsTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'units' field"))
	}

	// Virtual field
	_actualValue := units
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataControlledVariableUnits"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataControlledVariableUnits")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataControlledVariableUnits{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		Units: units,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataControlledVariableUnits) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataControlledVariableUnits) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataControlledVariableUnits"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataControlledVariableUnits")
		}

		// Simple Field (units)
		if pushErr := writeBuffer.PushContext("units"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for units")
		}
		_unitsErr := writeBuffer.WriteSerializable(ctx, m.GetUnits())
		if popErr := writeBuffer.PopContext("units"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for units")
		}
		if _unitsErr != nil {
			return errors.Wrap(_unitsErr, "Error serializing 'units' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataControlledVariableUnits"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataControlledVariableUnits")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataControlledVariableUnits) isBACnetConstructedDataControlledVariableUnits() bool {
	return true
}

func (m *_BACnetConstructedDataControlledVariableUnits) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
