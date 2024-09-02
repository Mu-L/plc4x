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

// BACnetNotificationParametersChangeOfDiscreteValueNewValueObjectidentifier is the corresponding interface of BACnetNotificationParametersChangeOfDiscreteValueNewValueObjectidentifier
type BACnetNotificationParametersChangeOfDiscreteValueNewValueObjectidentifier interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetNotificationParametersChangeOfDiscreteValueNewValue
	// GetObjectidentifierValue returns ObjectidentifierValue (property field)
	GetObjectidentifierValue() BACnetApplicationTagObjectIdentifier
	// IsBACnetNotificationParametersChangeOfDiscreteValueNewValueObjectidentifier is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetNotificationParametersChangeOfDiscreteValueNewValueObjectidentifier()
}

// _BACnetNotificationParametersChangeOfDiscreteValueNewValueObjectidentifier is the data-structure of this message
type _BACnetNotificationParametersChangeOfDiscreteValueNewValueObjectidentifier struct {
	BACnetNotificationParametersChangeOfDiscreteValueNewValueContract
	ObjectidentifierValue BACnetApplicationTagObjectIdentifier
}

var _ BACnetNotificationParametersChangeOfDiscreteValueNewValueObjectidentifier = (*_BACnetNotificationParametersChangeOfDiscreteValueNewValueObjectidentifier)(nil)
var _ BACnetNotificationParametersChangeOfDiscreteValueNewValueRequirements = (*_BACnetNotificationParametersChangeOfDiscreteValueNewValueObjectidentifier)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueObjectidentifier) GetParent() BACnetNotificationParametersChangeOfDiscreteValueNewValueContract {
	return m.BACnetNotificationParametersChangeOfDiscreteValueNewValueContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueObjectidentifier) GetObjectidentifierValue() BACnetApplicationTagObjectIdentifier {
	return m.ObjectidentifierValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetNotificationParametersChangeOfDiscreteValueNewValueObjectidentifier factory function for _BACnetNotificationParametersChangeOfDiscreteValueNewValueObjectidentifier
func NewBACnetNotificationParametersChangeOfDiscreteValueNewValueObjectidentifier(objectidentifierValue BACnetApplicationTagObjectIdentifier, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetNotificationParametersChangeOfDiscreteValueNewValueObjectidentifier {
	_result := &_BACnetNotificationParametersChangeOfDiscreteValueNewValueObjectidentifier{
		BACnetNotificationParametersChangeOfDiscreteValueNewValueContract: NewBACnetNotificationParametersChangeOfDiscreteValueNewValue(openingTag, peekedTagHeader, closingTag, tagNumber),
		ObjectidentifierValue: objectidentifierValue,
	}
	_result.BACnetNotificationParametersChangeOfDiscreteValueNewValueContract.(*_BACnetNotificationParametersChangeOfDiscreteValueNewValue)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetNotificationParametersChangeOfDiscreteValueNewValueObjectidentifier(structType any) BACnetNotificationParametersChangeOfDiscreteValueNewValueObjectidentifier {
	if casted, ok := structType.(BACnetNotificationParametersChangeOfDiscreteValueNewValueObjectidentifier); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetNotificationParametersChangeOfDiscreteValueNewValueObjectidentifier); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueObjectidentifier) GetTypeName() string {
	return "BACnetNotificationParametersChangeOfDiscreteValueNewValueObjectidentifier"
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueObjectidentifier) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetNotificationParametersChangeOfDiscreteValueNewValueContract.(*_BACnetNotificationParametersChangeOfDiscreteValueNewValue).getLengthInBits(ctx))

	// Simple field (objectidentifierValue)
	lengthInBits += m.ObjectidentifierValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueObjectidentifier) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueObjectidentifier) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetNotificationParametersChangeOfDiscreteValueNewValue, tagNumber uint8) (__bACnetNotificationParametersChangeOfDiscreteValueNewValueObjectidentifier BACnetNotificationParametersChangeOfDiscreteValueNewValueObjectidentifier, err error) {
	m.BACnetNotificationParametersChangeOfDiscreteValueNewValueContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetNotificationParametersChangeOfDiscreteValueNewValueObjectidentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetNotificationParametersChangeOfDiscreteValueNewValueObjectidentifier")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	objectidentifierValue, err := ReadSimpleField[BACnetApplicationTagObjectIdentifier](ctx, "objectidentifierValue", ReadComplex[BACnetApplicationTagObjectIdentifier](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagObjectIdentifier](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'objectidentifierValue' field"))
	}
	m.ObjectidentifierValue = objectidentifierValue

	if closeErr := readBuffer.CloseContext("BACnetNotificationParametersChangeOfDiscreteValueNewValueObjectidentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetNotificationParametersChangeOfDiscreteValueNewValueObjectidentifier")
	}

	return m, nil
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueObjectidentifier) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueObjectidentifier) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetNotificationParametersChangeOfDiscreteValueNewValueObjectidentifier"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetNotificationParametersChangeOfDiscreteValueNewValueObjectidentifier")
		}

		if err := WriteSimpleField[BACnetApplicationTagObjectIdentifier](ctx, "objectidentifierValue", m.GetObjectidentifierValue(), WriteComplex[BACnetApplicationTagObjectIdentifier](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'objectidentifierValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetNotificationParametersChangeOfDiscreteValueNewValueObjectidentifier"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetNotificationParametersChangeOfDiscreteValueNewValueObjectidentifier")
		}
		return nil
	}
	return m.BACnetNotificationParametersChangeOfDiscreteValueNewValueContract.(*_BACnetNotificationParametersChangeOfDiscreteValueNewValue).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueObjectidentifier) IsBACnetNotificationParametersChangeOfDiscreteValueNewValueObjectidentifier() {
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueObjectidentifier) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
