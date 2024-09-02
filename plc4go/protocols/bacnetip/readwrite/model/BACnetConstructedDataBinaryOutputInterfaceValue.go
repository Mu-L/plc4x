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

// BACnetConstructedDataBinaryOutputInterfaceValue is the corresponding interface of BACnetConstructedDataBinaryOutputInterfaceValue
type BACnetConstructedDataBinaryOutputInterfaceValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetInterfaceValue returns InterfaceValue (property field)
	GetInterfaceValue() BACnetOptionalBinaryPV
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetOptionalBinaryPV
}

// BACnetConstructedDataBinaryOutputInterfaceValueExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataBinaryOutputInterfaceValue.
// This is useful for switch cases.
type BACnetConstructedDataBinaryOutputInterfaceValueExactly interface {
	BACnetConstructedDataBinaryOutputInterfaceValue
	isBACnetConstructedDataBinaryOutputInterfaceValue() bool
}

// _BACnetConstructedDataBinaryOutputInterfaceValue is the data-structure of this message
type _BACnetConstructedDataBinaryOutputInterfaceValue struct {
	BACnetConstructedDataContract
	InterfaceValue BACnetOptionalBinaryPV
}

var _ BACnetConstructedDataBinaryOutputInterfaceValue = (*_BACnetConstructedDataBinaryOutputInterfaceValue)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataBinaryOutputInterfaceValue)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataBinaryOutputInterfaceValue) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_BINARY_OUTPUT
}

func (m *_BACnetConstructedDataBinaryOutputInterfaceValue) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_INTERFACE_VALUE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataBinaryOutputInterfaceValue) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataBinaryOutputInterfaceValue) GetInterfaceValue() BACnetOptionalBinaryPV {
	return m.InterfaceValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataBinaryOutputInterfaceValue) GetActualValue() BACnetOptionalBinaryPV {
	ctx := context.Background()
	_ = ctx
	return CastBACnetOptionalBinaryPV(m.GetInterfaceValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataBinaryOutputInterfaceValue factory function for _BACnetConstructedDataBinaryOutputInterfaceValue
func NewBACnetConstructedDataBinaryOutputInterfaceValue(interfaceValue BACnetOptionalBinaryPV, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataBinaryOutputInterfaceValue {
	_result := &_BACnetConstructedDataBinaryOutputInterfaceValue{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		InterfaceValue:                interfaceValue,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataBinaryOutputInterfaceValue(structType any) BACnetConstructedDataBinaryOutputInterfaceValue {
	if casted, ok := structType.(BACnetConstructedDataBinaryOutputInterfaceValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataBinaryOutputInterfaceValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataBinaryOutputInterfaceValue) GetTypeName() string {
	return "BACnetConstructedDataBinaryOutputInterfaceValue"
}

func (m *_BACnetConstructedDataBinaryOutputInterfaceValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (interfaceValue)
	lengthInBits += m.InterfaceValue.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataBinaryOutputInterfaceValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataBinaryOutputInterfaceValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataBinaryOutputInterfaceValue BACnetConstructedDataBinaryOutputInterfaceValue, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataBinaryOutputInterfaceValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataBinaryOutputInterfaceValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	interfaceValue, err := ReadSimpleField[BACnetOptionalBinaryPV](ctx, "interfaceValue", ReadComplex[BACnetOptionalBinaryPV](BACnetOptionalBinaryPVParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'interfaceValue' field"))
	}
	m.InterfaceValue = interfaceValue

	actualValue, err := ReadVirtualField[BACnetOptionalBinaryPV](ctx, "actualValue", (*BACnetOptionalBinaryPV)(nil), interfaceValue)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataBinaryOutputInterfaceValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataBinaryOutputInterfaceValue")
	}

	return m, nil
}

func (m *_BACnetConstructedDataBinaryOutputInterfaceValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataBinaryOutputInterfaceValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataBinaryOutputInterfaceValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataBinaryOutputInterfaceValue")
		}

		if err := WriteSimpleField[BACnetOptionalBinaryPV](ctx, "interfaceValue", m.GetInterfaceValue(), WriteComplex[BACnetOptionalBinaryPV](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'interfaceValue' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataBinaryOutputInterfaceValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataBinaryOutputInterfaceValue")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataBinaryOutputInterfaceValue) isBACnetConstructedDataBinaryOutputInterfaceValue() bool {
	return true
}

func (m *_BACnetConstructedDataBinaryOutputInterfaceValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
