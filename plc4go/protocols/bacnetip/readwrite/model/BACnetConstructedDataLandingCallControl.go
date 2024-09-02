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

// BACnetConstructedDataLandingCallControl is the corresponding interface of BACnetConstructedDataLandingCallControl
type BACnetConstructedDataLandingCallControl interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetLandingCallControl returns LandingCallControl (property field)
	GetLandingCallControl() BACnetLandingCallStatus
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetLandingCallStatus
	// IsBACnetConstructedDataLandingCallControl is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataLandingCallControl()
}

// _BACnetConstructedDataLandingCallControl is the data-structure of this message
type _BACnetConstructedDataLandingCallControl struct {
	BACnetConstructedDataContract
	LandingCallControl BACnetLandingCallStatus
}

var _ BACnetConstructedDataLandingCallControl = (*_BACnetConstructedDataLandingCallControl)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataLandingCallControl)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLandingCallControl) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataLandingCallControl) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LANDING_CALL_CONTROL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLandingCallControl) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLandingCallControl) GetLandingCallControl() BACnetLandingCallStatus {
	return m.LandingCallControl
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataLandingCallControl) GetActualValue() BACnetLandingCallStatus {
	ctx := context.Background()
	_ = ctx
	return CastBACnetLandingCallStatus(m.GetLandingCallControl())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataLandingCallControl factory function for _BACnetConstructedDataLandingCallControl
func NewBACnetConstructedDataLandingCallControl(landingCallControl BACnetLandingCallStatus, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLandingCallControl {
	_result := &_BACnetConstructedDataLandingCallControl{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		LandingCallControl:            landingCallControl,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLandingCallControl(structType any) BACnetConstructedDataLandingCallControl {
	if casted, ok := structType.(BACnetConstructedDataLandingCallControl); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLandingCallControl); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLandingCallControl) GetTypeName() string {
	return "BACnetConstructedDataLandingCallControl"
}

func (m *_BACnetConstructedDataLandingCallControl) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (landingCallControl)
	lengthInBits += m.LandingCallControl.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataLandingCallControl) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataLandingCallControl) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataLandingCallControl BACnetConstructedDataLandingCallControl, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLandingCallControl"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLandingCallControl")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	landingCallControl, err := ReadSimpleField[BACnetLandingCallStatus](ctx, "landingCallControl", ReadComplex[BACnetLandingCallStatus](BACnetLandingCallStatusParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'landingCallControl' field"))
	}
	m.LandingCallControl = landingCallControl

	actualValue, err := ReadVirtualField[BACnetLandingCallStatus](ctx, "actualValue", (*BACnetLandingCallStatus)(nil), landingCallControl)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLandingCallControl"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLandingCallControl")
	}

	return m, nil
}

func (m *_BACnetConstructedDataLandingCallControl) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataLandingCallControl) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLandingCallControl"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLandingCallControl")
		}

		if err := WriteSimpleField[BACnetLandingCallStatus](ctx, "landingCallControl", m.GetLandingCallControl(), WriteComplex[BACnetLandingCallStatus](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'landingCallControl' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLandingCallControl"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLandingCallControl")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLandingCallControl) IsBACnetConstructedDataLandingCallControl() {}

func (m *_BACnetConstructedDataLandingCallControl) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
