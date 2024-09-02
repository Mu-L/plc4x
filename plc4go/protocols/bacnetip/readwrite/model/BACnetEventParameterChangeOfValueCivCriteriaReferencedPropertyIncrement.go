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

// BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement is the corresponding interface of BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement
type BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetEventParameterChangeOfValueCivCriteria
	// GetReferencedPropertyIncrement returns ReferencedPropertyIncrement (property field)
	GetReferencedPropertyIncrement() BACnetContextTagReal
	// IsBACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement()
}

// _BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement is the data-structure of this message
type _BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement struct {
	BACnetEventParameterChangeOfValueCivCriteriaContract
	ReferencedPropertyIncrement BACnetContextTagReal
}

var _ BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement = (*_BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement)(nil)
var _ BACnetEventParameterChangeOfValueCivCriteriaRequirements = (*_BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement) GetParent() BACnetEventParameterChangeOfValueCivCriteriaContract {
	return m.BACnetEventParameterChangeOfValueCivCriteriaContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement) GetReferencedPropertyIncrement() BACnetContextTagReal {
	return m.ReferencedPropertyIncrement
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement factory function for _BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement
func NewBACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement(referencedPropertyIncrement BACnetContextTagReal, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement {
	_result := &_BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement{
		BACnetEventParameterChangeOfValueCivCriteriaContract: NewBACnetEventParameterChangeOfValueCivCriteria(openingTag, peekedTagHeader, closingTag, tagNumber),
		ReferencedPropertyIncrement:                          referencedPropertyIncrement,
	}
	_result.BACnetEventParameterChangeOfValueCivCriteriaContract.(*_BACnetEventParameterChangeOfValueCivCriteria)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement(structType any) BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement {
	if casted, ok := structType.(BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement) GetTypeName() string {
	return "BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement"
}

func (m *_BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetEventParameterChangeOfValueCivCriteriaContract.(*_BACnetEventParameterChangeOfValueCivCriteria).getLengthInBits(ctx))

	// Simple field (referencedPropertyIncrement)
	lengthInBits += m.ReferencedPropertyIncrement.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetEventParameterChangeOfValueCivCriteria, tagNumber uint8) (__bACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement, err error) {
	m.BACnetEventParameterChangeOfValueCivCriteriaContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	referencedPropertyIncrement, err := ReadSimpleField[BACnetContextTagReal](ctx, "referencedPropertyIncrement", ReadComplex[BACnetContextTagReal](BACnetContextTagParseWithBufferProducer[BACnetContextTagReal]((uint8)(uint8(1)), (BACnetDataType)(BACnetDataType_REAL)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'referencedPropertyIncrement' field"))
	}
	m.ReferencedPropertyIncrement = referencedPropertyIncrement

	if closeErr := readBuffer.CloseContext("BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement")
	}

	return m, nil
}

func (m *_BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement")
		}

		if err := WriteSimpleField[BACnetContextTagReal](ctx, "referencedPropertyIncrement", m.GetReferencedPropertyIncrement(), WriteComplex[BACnetContextTagReal](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'referencedPropertyIncrement' field")
		}

		if popErr := writeBuffer.PopContext("BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement")
		}
		return nil
	}
	return m.BACnetEventParameterChangeOfValueCivCriteriaContract.(*_BACnetEventParameterChangeOfValueCivCriteria).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement) IsBACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement() {
}

func (m *_BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
