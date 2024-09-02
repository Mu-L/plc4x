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

// BACnetConstructedDataPrescale is the corresponding interface of BACnetConstructedDataPrescale
type BACnetConstructedDataPrescale interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetPrescale returns Prescale (property field)
	GetPrescale() BACnetPrescale
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetPrescale
	// IsBACnetConstructedDataPrescale is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataPrescale()
}

// _BACnetConstructedDataPrescale is the data-structure of this message
type _BACnetConstructedDataPrescale struct {
	BACnetConstructedDataContract
	Prescale BACnetPrescale
}

var _ BACnetConstructedDataPrescale = (*_BACnetConstructedDataPrescale)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataPrescale)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataPrescale) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataPrescale) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_PRESCALE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataPrescale) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataPrescale) GetPrescale() BACnetPrescale {
	return m.Prescale
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataPrescale) GetActualValue() BACnetPrescale {
	ctx := context.Background()
	_ = ctx
	return CastBACnetPrescale(m.GetPrescale())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataPrescale factory function for _BACnetConstructedDataPrescale
func NewBACnetConstructedDataPrescale(prescale BACnetPrescale, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataPrescale {
	_result := &_BACnetConstructedDataPrescale{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		Prescale:                      prescale,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataPrescale(structType any) BACnetConstructedDataPrescale {
	if casted, ok := structType.(BACnetConstructedDataPrescale); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataPrescale); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataPrescale) GetTypeName() string {
	return "BACnetConstructedDataPrescale"
}

func (m *_BACnetConstructedDataPrescale) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (prescale)
	lengthInBits += m.Prescale.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataPrescale) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataPrescale) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataPrescale BACnetConstructedDataPrescale, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataPrescale"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataPrescale")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	prescale, err := ReadSimpleField[BACnetPrescale](ctx, "prescale", ReadComplex[BACnetPrescale](BACnetPrescaleParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'prescale' field"))
	}
	m.Prescale = prescale

	actualValue, err := ReadVirtualField[BACnetPrescale](ctx, "actualValue", (*BACnetPrescale)(nil), prescale)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataPrescale"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataPrescale")
	}

	return m, nil
}

func (m *_BACnetConstructedDataPrescale) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataPrescale) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataPrescale"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataPrescale")
		}

		if err := WriteSimpleField[BACnetPrescale](ctx, "prescale", m.GetPrescale(), WriteComplex[BACnetPrescale](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'prescale' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataPrescale"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataPrescale")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataPrescale) IsBACnetConstructedDataPrescale() {}

func (m *_BACnetConstructedDataPrescale) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
