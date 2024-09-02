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

// BACnetPriorityValueCharacterString is the corresponding interface of BACnetPriorityValueCharacterString
type BACnetPriorityValueCharacterString interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetPriorityValue
	// GetCharacterStringValue returns CharacterStringValue (property field)
	GetCharacterStringValue() BACnetApplicationTagCharacterString
}

// BACnetPriorityValueCharacterStringExactly can be used when we want exactly this type and not a type which fulfills BACnetPriorityValueCharacterString.
// This is useful for switch cases.
type BACnetPriorityValueCharacterStringExactly interface {
	BACnetPriorityValueCharacterString
	isBACnetPriorityValueCharacterString() bool
}

// _BACnetPriorityValueCharacterString is the data-structure of this message
type _BACnetPriorityValueCharacterString struct {
	BACnetPriorityValueContract
	CharacterStringValue BACnetApplicationTagCharacterString
}

var _ BACnetPriorityValueCharacterString = (*_BACnetPriorityValueCharacterString)(nil)
var _ BACnetPriorityValueRequirements = (*_BACnetPriorityValueCharacterString)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPriorityValueCharacterString) GetParent() BACnetPriorityValueContract {
	return m.BACnetPriorityValueContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPriorityValueCharacterString) GetCharacterStringValue() BACnetApplicationTagCharacterString {
	return m.CharacterStringValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPriorityValueCharacterString factory function for _BACnetPriorityValueCharacterString
func NewBACnetPriorityValueCharacterString(characterStringValue BACnetApplicationTagCharacterString, peekedTagHeader BACnetTagHeader, objectTypeArgument BACnetObjectType) *_BACnetPriorityValueCharacterString {
	_result := &_BACnetPriorityValueCharacterString{
		BACnetPriorityValueContract: NewBACnetPriorityValue(peekedTagHeader, objectTypeArgument),
		CharacterStringValue:        characterStringValue,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetPriorityValueCharacterString(structType any) BACnetPriorityValueCharacterString {
	if casted, ok := structType.(BACnetPriorityValueCharacterString); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPriorityValueCharacterString); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPriorityValueCharacterString) GetTypeName() string {
	return "BACnetPriorityValueCharacterString"
}

func (m *_BACnetPriorityValueCharacterString) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetPriorityValueContract.(*_BACnetPriorityValue).getLengthInBits(ctx))

	// Simple field (characterStringValue)
	lengthInBits += m.CharacterStringValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPriorityValueCharacterString) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetPriorityValueCharacterString) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetPriorityValue, objectTypeArgument BACnetObjectType) (__bACnetPriorityValueCharacterString BACnetPriorityValueCharacterString, err error) {
	m.BACnetPriorityValueContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPriorityValueCharacterString"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPriorityValueCharacterString")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	characterStringValue, err := ReadSimpleField[BACnetApplicationTagCharacterString](ctx, "characterStringValue", ReadComplex[BACnetApplicationTagCharacterString](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagCharacterString](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'characterStringValue' field"))
	}
	m.CharacterStringValue = characterStringValue

	if closeErr := readBuffer.CloseContext("BACnetPriorityValueCharacterString"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPriorityValueCharacterString")
	}

	return m, nil
}

func (m *_BACnetPriorityValueCharacterString) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPriorityValueCharacterString) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPriorityValueCharacterString"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPriorityValueCharacterString")
		}

		if err := WriteSimpleField[BACnetApplicationTagCharacterString](ctx, "characterStringValue", m.GetCharacterStringValue(), WriteComplex[BACnetApplicationTagCharacterString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'characterStringValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPriorityValueCharacterString"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPriorityValueCharacterString")
		}
		return nil
	}
	return m.BACnetPriorityValueContract.(*_BACnetPriorityValue).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPriorityValueCharacterString) isBACnetPriorityValueCharacterString() bool {
	return true
}

func (m *_BACnetPriorityValueCharacterString) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
