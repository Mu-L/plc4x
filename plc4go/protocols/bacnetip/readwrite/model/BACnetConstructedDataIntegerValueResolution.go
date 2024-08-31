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

// BACnetConstructedDataIntegerValueResolution is the corresponding interface of BACnetConstructedDataIntegerValueResolution
type BACnetConstructedDataIntegerValueResolution interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetResolution returns Resolution (property field)
	GetResolution() BACnetApplicationTagSignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagSignedInteger
}

// BACnetConstructedDataIntegerValueResolutionExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataIntegerValueResolution.
// This is useful for switch cases.
type BACnetConstructedDataIntegerValueResolutionExactly interface {
	BACnetConstructedDataIntegerValueResolution
	isBACnetConstructedDataIntegerValueResolution() bool
}

// _BACnetConstructedDataIntegerValueResolution is the data-structure of this message
type _BACnetConstructedDataIntegerValueResolution struct {
	*_BACnetConstructedData
	Resolution BACnetApplicationTagSignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataIntegerValueResolution) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_INTEGER_VALUE
}

func (m *_BACnetConstructedDataIntegerValueResolution) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_RESOLUTION
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataIntegerValueResolution) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataIntegerValueResolution) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataIntegerValueResolution) GetResolution() BACnetApplicationTagSignedInteger {
	return m.Resolution
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataIntegerValueResolution) GetActualValue() BACnetApplicationTagSignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagSignedInteger(m.GetResolution())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataIntegerValueResolution factory function for _BACnetConstructedDataIntegerValueResolution
func NewBACnetConstructedDataIntegerValueResolution(resolution BACnetApplicationTagSignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataIntegerValueResolution {
	_result := &_BACnetConstructedDataIntegerValueResolution{
		Resolution:             resolution,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataIntegerValueResolution(structType any) BACnetConstructedDataIntegerValueResolution {
	if casted, ok := structType.(BACnetConstructedDataIntegerValueResolution); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataIntegerValueResolution); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataIntegerValueResolution) GetTypeName() string {
	return "BACnetConstructedDataIntegerValueResolution"
}

func (m *_BACnetConstructedDataIntegerValueResolution) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (resolution)
	lengthInBits += m.Resolution.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataIntegerValueResolution) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataIntegerValueResolutionParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataIntegerValueResolution, error) {
	return BACnetConstructedDataIntegerValueResolutionParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataIntegerValueResolutionParseWithBufferProducer(tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetConstructedDataIntegerValueResolution, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetConstructedDataIntegerValueResolution, error) {
		return BACnetConstructedDataIntegerValueResolutionParseWithBuffer(ctx, readBuffer, tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
	}
}

func BACnetConstructedDataIntegerValueResolutionParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataIntegerValueResolution, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetConstructedDataIntegerValueResolution"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataIntegerValueResolution")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	resolution, err := ReadSimpleField[BACnetApplicationTagSignedInteger](ctx, "resolution", ReadComplex[BACnetApplicationTagSignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagSignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'resolution' field"))
	}

	// Virtual field
	_actualValue := resolution
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataIntegerValueResolution"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataIntegerValueResolution")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataIntegerValueResolution{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		Resolution: resolution,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataIntegerValueResolution) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataIntegerValueResolution) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataIntegerValueResolution"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataIntegerValueResolution")
		}

		// Simple Field (resolution)
		if pushErr := writeBuffer.PushContext("resolution"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for resolution")
		}
		_resolutionErr := writeBuffer.WriteSerializable(ctx, m.GetResolution())
		if popErr := writeBuffer.PopContext("resolution"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for resolution")
		}
		if _resolutionErr != nil {
			return errors.Wrap(_resolutionErr, "Error serializing 'resolution' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataIntegerValueResolution"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataIntegerValueResolution")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataIntegerValueResolution) isBACnetConstructedDataIntegerValueResolution() bool {
	return true
}

func (m *_BACnetConstructedDataIntegerValueResolution) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
