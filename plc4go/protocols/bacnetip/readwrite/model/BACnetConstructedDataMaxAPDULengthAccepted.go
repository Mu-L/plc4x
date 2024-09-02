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

// BACnetConstructedDataMaxAPDULengthAccepted is the corresponding interface of BACnetConstructedDataMaxAPDULengthAccepted
type BACnetConstructedDataMaxAPDULengthAccepted interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetMaxApduLengthAccepted returns MaxApduLengthAccepted (property field)
	GetMaxApduLengthAccepted() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetConstructedDataMaxAPDULengthAccepted is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataMaxAPDULengthAccepted()
}

// _BACnetConstructedDataMaxAPDULengthAccepted is the data-structure of this message
type _BACnetConstructedDataMaxAPDULengthAccepted struct {
	BACnetConstructedDataContract
	MaxApduLengthAccepted BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataMaxAPDULengthAccepted = (*_BACnetConstructedDataMaxAPDULengthAccepted)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataMaxAPDULengthAccepted)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataMaxAPDULengthAccepted) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataMaxAPDULengthAccepted) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_MAX_APDU_LENGTH_ACCEPTED
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataMaxAPDULengthAccepted) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataMaxAPDULengthAccepted) GetMaxApduLengthAccepted() BACnetApplicationTagUnsignedInteger {
	return m.MaxApduLengthAccepted
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataMaxAPDULengthAccepted) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetMaxApduLengthAccepted())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataMaxAPDULengthAccepted factory function for _BACnetConstructedDataMaxAPDULengthAccepted
func NewBACnetConstructedDataMaxAPDULengthAccepted(maxApduLengthAccepted BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataMaxAPDULengthAccepted {
	_result := &_BACnetConstructedDataMaxAPDULengthAccepted{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		MaxApduLengthAccepted:         maxApduLengthAccepted,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataMaxAPDULengthAccepted(structType any) BACnetConstructedDataMaxAPDULengthAccepted {
	if casted, ok := structType.(BACnetConstructedDataMaxAPDULengthAccepted); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataMaxAPDULengthAccepted); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataMaxAPDULengthAccepted) GetTypeName() string {
	return "BACnetConstructedDataMaxAPDULengthAccepted"
}

func (m *_BACnetConstructedDataMaxAPDULengthAccepted) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (maxApduLengthAccepted)
	lengthInBits += m.MaxApduLengthAccepted.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataMaxAPDULengthAccepted) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataMaxAPDULengthAccepted) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataMaxAPDULengthAccepted BACnetConstructedDataMaxAPDULengthAccepted, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataMaxAPDULengthAccepted"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataMaxAPDULengthAccepted")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	maxApduLengthAccepted, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "maxApduLengthAccepted", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'maxApduLengthAccepted' field"))
	}
	m.MaxApduLengthAccepted = maxApduLengthAccepted

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), maxApduLengthAccepted)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataMaxAPDULengthAccepted"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataMaxAPDULengthAccepted")
	}

	return m, nil
}

func (m *_BACnetConstructedDataMaxAPDULengthAccepted) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataMaxAPDULengthAccepted) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataMaxAPDULengthAccepted"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataMaxAPDULengthAccepted")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "maxApduLengthAccepted", m.GetMaxApduLengthAccepted(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'maxApduLengthAccepted' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataMaxAPDULengthAccepted"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataMaxAPDULengthAccepted")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataMaxAPDULengthAccepted) IsBACnetConstructedDataMaxAPDULengthAccepted() {
}

func (m *_BACnetConstructedDataMaxAPDULengthAccepted) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
