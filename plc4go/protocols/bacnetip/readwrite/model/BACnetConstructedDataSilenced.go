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

// BACnetConstructedDataSilenced is the corresponding interface of BACnetConstructedDataSilenced
type BACnetConstructedDataSilenced interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetSilenced returns Silenced (property field)
	GetSilenced() BACnetSilencedStateTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetSilencedStateTagged
	// IsBACnetConstructedDataSilenced is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataSilenced()
}

// _BACnetConstructedDataSilenced is the data-structure of this message
type _BACnetConstructedDataSilenced struct {
	BACnetConstructedDataContract
	Silenced BACnetSilencedStateTagged
}

var _ BACnetConstructedDataSilenced = (*_BACnetConstructedDataSilenced)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataSilenced)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataSilenced) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataSilenced) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_SILENCED
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataSilenced) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataSilenced) GetSilenced() BACnetSilencedStateTagged {
	return m.Silenced
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataSilenced) GetActualValue() BACnetSilencedStateTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetSilencedStateTagged(m.GetSilenced())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataSilenced factory function for _BACnetConstructedDataSilenced
func NewBACnetConstructedDataSilenced(silenced BACnetSilencedStateTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataSilenced {
	_result := &_BACnetConstructedDataSilenced{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		Silenced:                      silenced,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataSilenced(structType any) BACnetConstructedDataSilenced {
	if casted, ok := structType.(BACnetConstructedDataSilenced); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataSilenced); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataSilenced) GetTypeName() string {
	return "BACnetConstructedDataSilenced"
}

func (m *_BACnetConstructedDataSilenced) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (silenced)
	lengthInBits += m.Silenced.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataSilenced) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataSilenced) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataSilenced BACnetConstructedDataSilenced, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataSilenced"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataSilenced")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	silenced, err := ReadSimpleField[BACnetSilencedStateTagged](ctx, "silenced", ReadComplex[BACnetSilencedStateTagged](BACnetSilencedStateTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'silenced' field"))
	}
	m.Silenced = silenced

	actualValue, err := ReadVirtualField[BACnetSilencedStateTagged](ctx, "actualValue", (*BACnetSilencedStateTagged)(nil), silenced)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataSilenced"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataSilenced")
	}

	return m, nil
}

func (m *_BACnetConstructedDataSilenced) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataSilenced) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataSilenced"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataSilenced")
		}

		if err := WriteSimpleField[BACnetSilencedStateTagged](ctx, "silenced", m.GetSilenced(), WriteComplex[BACnetSilencedStateTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'silenced' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataSilenced"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataSilenced")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataSilenced) IsBACnetConstructedDataSilenced() {}

func (m *_BACnetConstructedDataSilenced) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
