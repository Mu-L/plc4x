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

// BACnetRecipientEnclosed is the corresponding interface of BACnetRecipientEnclosed
type BACnetRecipientEnclosed interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetRecipient returns Recipient (property field)
	GetRecipient() BACnetRecipient
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
	// IsBACnetRecipientEnclosed is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetRecipientEnclosed()
}

// _BACnetRecipientEnclosed is the data-structure of this message
type _BACnetRecipientEnclosed struct {
	OpeningTag BACnetOpeningTag
	Recipient  BACnetRecipient
	ClosingTag BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

var _ BACnetRecipientEnclosed = (*_BACnetRecipientEnclosed)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetRecipientEnclosed) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetRecipientEnclosed) GetRecipient() BACnetRecipient {
	return m.Recipient
}

func (m *_BACnetRecipientEnclosed) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetRecipientEnclosed factory function for _BACnetRecipientEnclosed
func NewBACnetRecipientEnclosed(openingTag BACnetOpeningTag, recipient BACnetRecipient, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetRecipientEnclosed {
	return &_BACnetRecipientEnclosed{OpeningTag: openingTag, Recipient: recipient, ClosingTag: closingTag, TagNumber: tagNumber}
}

// Deprecated: use the interface for direct cast
func CastBACnetRecipientEnclosed(structType any) BACnetRecipientEnclosed {
	if casted, ok := structType.(BACnetRecipientEnclosed); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetRecipientEnclosed); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetRecipientEnclosed) GetTypeName() string {
	return "BACnetRecipientEnclosed"
}

func (m *_BACnetRecipientEnclosed) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Simple field (recipient)
	lengthInBits += m.Recipient.GetLengthInBits(ctx)

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetRecipientEnclosed) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetRecipientEnclosedParse(ctx context.Context, theBytes []byte, tagNumber uint8) (BACnetRecipientEnclosed, error) {
	return BACnetRecipientEnclosedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetRecipientEnclosedParseWithBufferProducer(tagNumber uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetRecipientEnclosed, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetRecipientEnclosed, error) {
		return BACnetRecipientEnclosedParseWithBuffer(ctx, readBuffer, tagNumber)
	}
}

func BACnetRecipientEnclosedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetRecipientEnclosed, error) {
	v, err := (&_BACnetRecipientEnclosed{TagNumber: tagNumber}).parse(ctx, readBuffer, tagNumber)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *_BACnetRecipientEnclosed) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (__bACnetRecipientEnclosed BACnetRecipientEnclosed, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetRecipientEnclosed"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetRecipientEnclosed")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}
	m.OpeningTag = openingTag

	recipient, err := ReadSimpleField[BACnetRecipient](ctx, "recipient", ReadComplex[BACnetRecipient](BACnetRecipientParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'recipient' field"))
	}
	m.Recipient = recipient

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	m.ClosingTag = closingTag

	if closeErr := readBuffer.CloseContext("BACnetRecipientEnclosed"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetRecipientEnclosed")
	}

	return m, nil
}

func (m *_BACnetRecipientEnclosed) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetRecipientEnclosed) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetRecipientEnclosed"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetRecipientEnclosed")
	}

	if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'openingTag' field")
	}

	if err := WriteSimpleField[BACnetRecipient](ctx, "recipient", m.GetRecipient(), WriteComplex[BACnetRecipient](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'recipient' field")
	}

	if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetRecipientEnclosed"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetRecipientEnclosed")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetRecipientEnclosed) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_BACnetRecipientEnclosed) IsBACnetRecipientEnclosed() {}

func (m *_BACnetRecipientEnclosed) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
