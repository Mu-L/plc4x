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

// BACnetConstructedDataIPSubnetMask is the corresponding interface of BACnetConstructedDataIPSubnetMask
type BACnetConstructedDataIPSubnetMask interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetIpSubnetMask returns IpSubnetMask (property field)
	GetIpSubnetMask() BACnetApplicationTagOctetString
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagOctetString
	// IsBACnetConstructedDataIPSubnetMask is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataIPSubnetMask()
}

// _BACnetConstructedDataIPSubnetMask is the data-structure of this message
type _BACnetConstructedDataIPSubnetMask struct {
	BACnetConstructedDataContract
	IpSubnetMask BACnetApplicationTagOctetString
}

var _ BACnetConstructedDataIPSubnetMask = (*_BACnetConstructedDataIPSubnetMask)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataIPSubnetMask)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataIPSubnetMask) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataIPSubnetMask) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_IP_SUBNET_MASK
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataIPSubnetMask) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataIPSubnetMask) GetIpSubnetMask() BACnetApplicationTagOctetString {
	return m.IpSubnetMask
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataIPSubnetMask) GetActualValue() BACnetApplicationTagOctetString {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagOctetString(m.GetIpSubnetMask())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataIPSubnetMask factory function for _BACnetConstructedDataIPSubnetMask
func NewBACnetConstructedDataIPSubnetMask(ipSubnetMask BACnetApplicationTagOctetString, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataIPSubnetMask {
	_result := &_BACnetConstructedDataIPSubnetMask{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		IpSubnetMask:                  ipSubnetMask,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataIPSubnetMask(structType any) BACnetConstructedDataIPSubnetMask {
	if casted, ok := structType.(BACnetConstructedDataIPSubnetMask); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataIPSubnetMask); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataIPSubnetMask) GetTypeName() string {
	return "BACnetConstructedDataIPSubnetMask"
}

func (m *_BACnetConstructedDataIPSubnetMask) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (ipSubnetMask)
	lengthInBits += m.IpSubnetMask.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataIPSubnetMask) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataIPSubnetMask) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataIPSubnetMask BACnetConstructedDataIPSubnetMask, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataIPSubnetMask"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataIPSubnetMask")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	ipSubnetMask, err := ReadSimpleField[BACnetApplicationTagOctetString](ctx, "ipSubnetMask", ReadComplex[BACnetApplicationTagOctetString](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagOctetString](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'ipSubnetMask' field"))
	}
	m.IpSubnetMask = ipSubnetMask

	actualValue, err := ReadVirtualField[BACnetApplicationTagOctetString](ctx, "actualValue", (*BACnetApplicationTagOctetString)(nil), ipSubnetMask)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataIPSubnetMask"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataIPSubnetMask")
	}

	return m, nil
}

func (m *_BACnetConstructedDataIPSubnetMask) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataIPSubnetMask) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataIPSubnetMask"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataIPSubnetMask")
		}

		if err := WriteSimpleField[BACnetApplicationTagOctetString](ctx, "ipSubnetMask", m.GetIpSubnetMask(), WriteComplex[BACnetApplicationTagOctetString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'ipSubnetMask' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataIPSubnetMask"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataIPSubnetMask")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataIPSubnetMask) IsBACnetConstructedDataIPSubnetMask() {}

func (m *_BACnetConstructedDataIPSubnetMask) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
