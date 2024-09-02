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

// BACnetLightingCommandEnclosed is the corresponding interface of BACnetLightingCommandEnclosed
type BACnetLightingCommandEnclosed interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetLightingCommand returns LightingCommand (property field)
	GetLightingCommand() BACnetLightingCommand
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
	// IsBACnetLightingCommandEnclosed is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetLightingCommandEnclosed()
}

// _BACnetLightingCommandEnclosed is the data-structure of this message
type _BACnetLightingCommandEnclosed struct {
	OpeningTag      BACnetOpeningTag
	LightingCommand BACnetLightingCommand
	ClosingTag      BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

var _ BACnetLightingCommandEnclosed = (*_BACnetLightingCommandEnclosed)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetLightingCommandEnclosed) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetLightingCommandEnclosed) GetLightingCommand() BACnetLightingCommand {
	return m.LightingCommand
}

func (m *_BACnetLightingCommandEnclosed) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetLightingCommandEnclosed factory function for _BACnetLightingCommandEnclosed
func NewBACnetLightingCommandEnclosed(openingTag BACnetOpeningTag, lightingCommand BACnetLightingCommand, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetLightingCommandEnclosed {
	return &_BACnetLightingCommandEnclosed{OpeningTag: openingTag, LightingCommand: lightingCommand, ClosingTag: closingTag, TagNumber: tagNumber}
}

// Deprecated: use the interface for direct cast
func CastBACnetLightingCommandEnclosed(structType any) BACnetLightingCommandEnclosed {
	if casted, ok := structType.(BACnetLightingCommandEnclosed); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetLightingCommandEnclosed); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetLightingCommandEnclosed) GetTypeName() string {
	return "BACnetLightingCommandEnclosed"
}

func (m *_BACnetLightingCommandEnclosed) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Simple field (lightingCommand)
	lengthInBits += m.LightingCommand.GetLengthInBits(ctx)

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetLightingCommandEnclosed) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetLightingCommandEnclosedParse(ctx context.Context, theBytes []byte, tagNumber uint8) (BACnetLightingCommandEnclosed, error) {
	return BACnetLightingCommandEnclosedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetLightingCommandEnclosedParseWithBufferProducer(tagNumber uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetLightingCommandEnclosed, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetLightingCommandEnclosed, error) {
		return BACnetLightingCommandEnclosedParseWithBuffer(ctx, readBuffer, tagNumber)
	}
}

func BACnetLightingCommandEnclosedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetLightingCommandEnclosed, error) {
	v, err := (&_BACnetLightingCommandEnclosed{TagNumber: tagNumber}).parse(ctx, readBuffer, tagNumber)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *_BACnetLightingCommandEnclosed) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (__bACnetLightingCommandEnclosed BACnetLightingCommandEnclosed, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetLightingCommandEnclosed"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetLightingCommandEnclosed")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}
	m.OpeningTag = openingTag

	lightingCommand, err := ReadSimpleField[BACnetLightingCommand](ctx, "lightingCommand", ReadComplex[BACnetLightingCommand](BACnetLightingCommandParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lightingCommand' field"))
	}
	m.LightingCommand = lightingCommand

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	m.ClosingTag = closingTag

	if closeErr := readBuffer.CloseContext("BACnetLightingCommandEnclosed"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetLightingCommandEnclosed")
	}

	return m, nil
}

func (m *_BACnetLightingCommandEnclosed) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetLightingCommandEnclosed) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetLightingCommandEnclosed"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetLightingCommandEnclosed")
	}

	if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'openingTag' field")
	}

	if err := WriteSimpleField[BACnetLightingCommand](ctx, "lightingCommand", m.GetLightingCommand(), WriteComplex[BACnetLightingCommand](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'lightingCommand' field")
	}

	if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetLightingCommandEnclosed"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetLightingCommandEnclosed")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetLightingCommandEnclosed) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_BACnetLightingCommandEnclosed) IsBACnetLightingCommandEnclosed() {}

func (m *_BACnetLightingCommandEnclosed) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
