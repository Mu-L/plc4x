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

// BACnetConstructedDataFaultParameters is the corresponding interface of BACnetConstructedDataFaultParameters
type BACnetConstructedDataFaultParameters interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetFaultParameters returns FaultParameters (property field)
	GetFaultParameters() BACnetFaultParameter
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetFaultParameter
	// IsBACnetConstructedDataFaultParameters is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataFaultParameters()
}

// _BACnetConstructedDataFaultParameters is the data-structure of this message
type _BACnetConstructedDataFaultParameters struct {
	BACnetConstructedDataContract
	FaultParameters BACnetFaultParameter
}

var _ BACnetConstructedDataFaultParameters = (*_BACnetConstructedDataFaultParameters)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataFaultParameters)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataFaultParameters) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataFaultParameters) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_FAULT_PARAMETERS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataFaultParameters) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataFaultParameters) GetFaultParameters() BACnetFaultParameter {
	return m.FaultParameters
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataFaultParameters) GetActualValue() BACnetFaultParameter {
	ctx := context.Background()
	_ = ctx
	return CastBACnetFaultParameter(m.GetFaultParameters())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataFaultParameters factory function for _BACnetConstructedDataFaultParameters
func NewBACnetConstructedDataFaultParameters(faultParameters BACnetFaultParameter, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataFaultParameters {
	_result := &_BACnetConstructedDataFaultParameters{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		FaultParameters:               faultParameters,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataFaultParameters(structType any) BACnetConstructedDataFaultParameters {
	if casted, ok := structType.(BACnetConstructedDataFaultParameters); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataFaultParameters); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataFaultParameters) GetTypeName() string {
	return "BACnetConstructedDataFaultParameters"
}

func (m *_BACnetConstructedDataFaultParameters) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (faultParameters)
	lengthInBits += m.FaultParameters.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataFaultParameters) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataFaultParameters) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataFaultParameters BACnetConstructedDataFaultParameters, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataFaultParameters"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataFaultParameters")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	faultParameters, err := ReadSimpleField[BACnetFaultParameter](ctx, "faultParameters", ReadComplex[BACnetFaultParameter](BACnetFaultParameterParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'faultParameters' field"))
	}
	m.FaultParameters = faultParameters

	actualValue, err := ReadVirtualField[BACnetFaultParameter](ctx, "actualValue", (*BACnetFaultParameter)(nil), faultParameters)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataFaultParameters"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataFaultParameters")
	}

	return m, nil
}

func (m *_BACnetConstructedDataFaultParameters) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataFaultParameters) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataFaultParameters"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataFaultParameters")
		}

		if err := WriteSimpleField[BACnetFaultParameter](ctx, "faultParameters", m.GetFaultParameters(), WriteComplex[BACnetFaultParameter](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'faultParameters' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataFaultParameters"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataFaultParameters")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataFaultParameters) IsBACnetConstructedDataFaultParameters() {}

func (m *_BACnetConstructedDataFaultParameters) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
