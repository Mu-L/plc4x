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

// BACnetConstructedDataAckRequired is the corresponding interface of BACnetConstructedDataAckRequired
type BACnetConstructedDataAckRequired interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetAckRequired returns AckRequired (property field)
	GetAckRequired() BACnetEventTransitionBitsTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetEventTransitionBitsTagged
}

// BACnetConstructedDataAckRequiredExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataAckRequired.
// This is useful for switch cases.
type BACnetConstructedDataAckRequiredExactly interface {
	BACnetConstructedDataAckRequired
	isBACnetConstructedDataAckRequired() bool
}

// _BACnetConstructedDataAckRequired is the data-structure of this message
type _BACnetConstructedDataAckRequired struct {
	BACnetConstructedDataContract
	AckRequired BACnetEventTransitionBitsTagged
}

var _ BACnetConstructedDataAckRequired = (*_BACnetConstructedDataAckRequired)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataAckRequired)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAckRequired) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataAckRequired) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ACK_REQUIRED
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAckRequired) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataAckRequired) GetAckRequired() BACnetEventTransitionBitsTagged {
	return m.AckRequired
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataAckRequired) GetActualValue() BACnetEventTransitionBitsTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetEventTransitionBitsTagged(m.GetAckRequired())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataAckRequired factory function for _BACnetConstructedDataAckRequired
func NewBACnetConstructedDataAckRequired(ackRequired BACnetEventTransitionBitsTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAckRequired {
	_result := &_BACnetConstructedDataAckRequired{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		AckRequired:                   ackRequired,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAckRequired(structType any) BACnetConstructedDataAckRequired {
	if casted, ok := structType.(BACnetConstructedDataAckRequired); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAckRequired); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAckRequired) GetTypeName() string {
	return "BACnetConstructedDataAckRequired"
}

func (m *_BACnetConstructedDataAckRequired) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (ackRequired)
	lengthInBits += m.AckRequired.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataAckRequired) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataAckRequired) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataAckRequired BACnetConstructedDataAckRequired, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAckRequired"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAckRequired")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	ackRequired, err := ReadSimpleField[BACnetEventTransitionBitsTagged](ctx, "ackRequired", ReadComplex[BACnetEventTransitionBitsTagged](BACnetEventTransitionBitsTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'ackRequired' field"))
	}
	m.AckRequired = ackRequired

	actualValue, err := ReadVirtualField[BACnetEventTransitionBitsTagged](ctx, "actualValue", (*BACnetEventTransitionBitsTagged)(nil), ackRequired)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAckRequired"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAckRequired")
	}

	return m, nil
}

func (m *_BACnetConstructedDataAckRequired) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataAckRequired) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAckRequired"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAckRequired")
		}

		if err := WriteSimpleField[BACnetEventTransitionBitsTagged](ctx, "ackRequired", m.GetAckRequired(), WriteComplex[BACnetEventTransitionBitsTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'ackRequired' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAckRequired"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAckRequired")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAckRequired) isBACnetConstructedDataAckRequired() bool {
	return true
}

func (m *_BACnetConstructedDataAckRequired) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
