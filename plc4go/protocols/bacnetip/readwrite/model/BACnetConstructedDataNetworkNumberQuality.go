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

// BACnetConstructedDataNetworkNumberQuality is the corresponding interface of BACnetConstructedDataNetworkNumberQuality
type BACnetConstructedDataNetworkNumberQuality interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetNetworkNumberQuality returns NetworkNumberQuality (property field)
	GetNetworkNumberQuality() BACnetNetworkNumberQualityTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetNetworkNumberQualityTagged
}

// BACnetConstructedDataNetworkNumberQualityExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataNetworkNumberQuality.
// This is useful for switch cases.
type BACnetConstructedDataNetworkNumberQualityExactly interface {
	BACnetConstructedDataNetworkNumberQuality
	isBACnetConstructedDataNetworkNumberQuality() bool
}

// _BACnetConstructedDataNetworkNumberQuality is the data-structure of this message
type _BACnetConstructedDataNetworkNumberQuality struct {
	BACnetConstructedDataContract
	NetworkNumberQuality BACnetNetworkNumberQualityTagged
}

var _ BACnetConstructedDataNetworkNumberQuality = (*_BACnetConstructedDataNetworkNumberQuality)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataNetworkNumberQuality)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataNetworkNumberQuality) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataNetworkNumberQuality) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_NETWORK_NUMBER_QUALITY
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataNetworkNumberQuality) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataNetworkNumberQuality) GetNetworkNumberQuality() BACnetNetworkNumberQualityTagged {
	return m.NetworkNumberQuality
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataNetworkNumberQuality) GetActualValue() BACnetNetworkNumberQualityTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetNetworkNumberQualityTagged(m.GetNetworkNumberQuality())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataNetworkNumberQuality factory function for _BACnetConstructedDataNetworkNumberQuality
func NewBACnetConstructedDataNetworkNumberQuality(networkNumberQuality BACnetNetworkNumberQualityTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataNetworkNumberQuality {
	_result := &_BACnetConstructedDataNetworkNumberQuality{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		NetworkNumberQuality:          networkNumberQuality,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataNetworkNumberQuality(structType any) BACnetConstructedDataNetworkNumberQuality {
	if casted, ok := structType.(BACnetConstructedDataNetworkNumberQuality); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataNetworkNumberQuality); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataNetworkNumberQuality) GetTypeName() string {
	return "BACnetConstructedDataNetworkNumberQuality"
}

func (m *_BACnetConstructedDataNetworkNumberQuality) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (networkNumberQuality)
	lengthInBits += m.NetworkNumberQuality.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataNetworkNumberQuality) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataNetworkNumberQuality) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataNetworkNumberQuality BACnetConstructedDataNetworkNumberQuality, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataNetworkNumberQuality"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataNetworkNumberQuality")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	networkNumberQuality, err := ReadSimpleField[BACnetNetworkNumberQualityTagged](ctx, "networkNumberQuality", ReadComplex[BACnetNetworkNumberQualityTagged](BACnetNetworkNumberQualityTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'networkNumberQuality' field"))
	}
	m.NetworkNumberQuality = networkNumberQuality

	actualValue, err := ReadVirtualField[BACnetNetworkNumberQualityTagged](ctx, "actualValue", (*BACnetNetworkNumberQualityTagged)(nil), networkNumberQuality)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataNetworkNumberQuality"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataNetworkNumberQuality")
	}

	return m, nil
}

func (m *_BACnetConstructedDataNetworkNumberQuality) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataNetworkNumberQuality) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataNetworkNumberQuality"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataNetworkNumberQuality")
		}

		if err := WriteSimpleField[BACnetNetworkNumberQualityTagged](ctx, "networkNumberQuality", m.GetNetworkNumberQuality(), WriteComplex[BACnetNetworkNumberQualityTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'networkNumberQuality' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataNetworkNumberQuality"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataNetworkNumberQuality")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataNetworkNumberQuality) isBACnetConstructedDataNetworkNumberQuality() bool {
	return true
}

func (m *_BACnetConstructedDataNetworkNumberQuality) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
