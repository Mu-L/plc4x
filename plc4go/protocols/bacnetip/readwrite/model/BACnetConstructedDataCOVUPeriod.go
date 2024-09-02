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

// BACnetConstructedDataCOVUPeriod is the corresponding interface of BACnetConstructedDataCOVUPeriod
type BACnetConstructedDataCOVUPeriod interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetCovuPeriod returns CovuPeriod (property field)
	GetCovuPeriod() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataCOVUPeriodExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataCOVUPeriod.
// This is useful for switch cases.
type BACnetConstructedDataCOVUPeriodExactly interface {
	BACnetConstructedDataCOVUPeriod
	isBACnetConstructedDataCOVUPeriod() bool
}

// _BACnetConstructedDataCOVUPeriod is the data-structure of this message
type _BACnetConstructedDataCOVUPeriod struct {
	BACnetConstructedDataContract
	CovuPeriod BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataCOVUPeriod = (*_BACnetConstructedDataCOVUPeriod)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataCOVUPeriod)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataCOVUPeriod) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataCOVUPeriod) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_COVU_PERIOD
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataCOVUPeriod) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataCOVUPeriod) GetCovuPeriod() BACnetApplicationTagUnsignedInteger {
	return m.CovuPeriod
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataCOVUPeriod) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetCovuPeriod())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataCOVUPeriod factory function for _BACnetConstructedDataCOVUPeriod
func NewBACnetConstructedDataCOVUPeriod(covuPeriod BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataCOVUPeriod {
	_result := &_BACnetConstructedDataCOVUPeriod{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		CovuPeriod:                    covuPeriod,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataCOVUPeriod(structType any) BACnetConstructedDataCOVUPeriod {
	if casted, ok := structType.(BACnetConstructedDataCOVUPeriod); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataCOVUPeriod); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataCOVUPeriod) GetTypeName() string {
	return "BACnetConstructedDataCOVUPeriod"
}

func (m *_BACnetConstructedDataCOVUPeriod) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (covuPeriod)
	lengthInBits += m.CovuPeriod.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataCOVUPeriod) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataCOVUPeriod) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataCOVUPeriod BACnetConstructedDataCOVUPeriod, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataCOVUPeriod"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataCOVUPeriod")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	covuPeriod, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "covuPeriod", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'covuPeriod' field"))
	}
	m.CovuPeriod = covuPeriod

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), covuPeriod)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataCOVUPeriod"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataCOVUPeriod")
	}

	return m, nil
}

func (m *_BACnetConstructedDataCOVUPeriod) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataCOVUPeriod) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataCOVUPeriod"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataCOVUPeriod")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "covuPeriod", m.GetCovuPeriod(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'covuPeriod' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataCOVUPeriod"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataCOVUPeriod")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataCOVUPeriod) isBACnetConstructedDataCOVUPeriod() bool {
	return true
}

func (m *_BACnetConstructedDataCOVUPeriod) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
