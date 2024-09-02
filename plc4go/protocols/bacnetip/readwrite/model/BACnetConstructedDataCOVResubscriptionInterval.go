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

// BACnetConstructedDataCOVResubscriptionInterval is the corresponding interface of BACnetConstructedDataCOVResubscriptionInterval
type BACnetConstructedDataCOVResubscriptionInterval interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetCovResubscriptionInterval returns CovResubscriptionInterval (property field)
	GetCovResubscriptionInterval() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetConstructedDataCOVResubscriptionInterval is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataCOVResubscriptionInterval()
}

// _BACnetConstructedDataCOVResubscriptionInterval is the data-structure of this message
type _BACnetConstructedDataCOVResubscriptionInterval struct {
	BACnetConstructedDataContract
	CovResubscriptionInterval BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataCOVResubscriptionInterval = (*_BACnetConstructedDataCOVResubscriptionInterval)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataCOVResubscriptionInterval)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataCOVResubscriptionInterval) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataCOVResubscriptionInterval) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_COV_RESUBSCRIPTION_INTERVAL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataCOVResubscriptionInterval) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataCOVResubscriptionInterval) GetCovResubscriptionInterval() BACnetApplicationTagUnsignedInteger {
	return m.CovResubscriptionInterval
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataCOVResubscriptionInterval) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetCovResubscriptionInterval())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataCOVResubscriptionInterval factory function for _BACnetConstructedDataCOVResubscriptionInterval
func NewBACnetConstructedDataCOVResubscriptionInterval(covResubscriptionInterval BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataCOVResubscriptionInterval {
	_result := &_BACnetConstructedDataCOVResubscriptionInterval{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		CovResubscriptionInterval:     covResubscriptionInterval,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataCOVResubscriptionInterval(structType any) BACnetConstructedDataCOVResubscriptionInterval {
	if casted, ok := structType.(BACnetConstructedDataCOVResubscriptionInterval); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataCOVResubscriptionInterval); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataCOVResubscriptionInterval) GetTypeName() string {
	return "BACnetConstructedDataCOVResubscriptionInterval"
}

func (m *_BACnetConstructedDataCOVResubscriptionInterval) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (covResubscriptionInterval)
	lengthInBits += m.CovResubscriptionInterval.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataCOVResubscriptionInterval) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataCOVResubscriptionInterval) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataCOVResubscriptionInterval BACnetConstructedDataCOVResubscriptionInterval, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataCOVResubscriptionInterval"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataCOVResubscriptionInterval")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	covResubscriptionInterval, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "covResubscriptionInterval", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'covResubscriptionInterval' field"))
	}
	m.CovResubscriptionInterval = covResubscriptionInterval

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), covResubscriptionInterval)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataCOVResubscriptionInterval"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataCOVResubscriptionInterval")
	}

	return m, nil
}

func (m *_BACnetConstructedDataCOVResubscriptionInterval) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataCOVResubscriptionInterval) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataCOVResubscriptionInterval"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataCOVResubscriptionInterval")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "covResubscriptionInterval", m.GetCovResubscriptionInterval(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'covResubscriptionInterval' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataCOVResubscriptionInterval"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataCOVResubscriptionInterval")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataCOVResubscriptionInterval) IsBACnetConstructedDataCOVResubscriptionInterval() {
}

func (m *_BACnetConstructedDataCOVResubscriptionInterval) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
