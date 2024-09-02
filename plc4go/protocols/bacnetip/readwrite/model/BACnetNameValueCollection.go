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

// BACnetNameValueCollection is the corresponding interface of BACnetNameValueCollection
type BACnetNameValueCollection interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetMembers returns Members (property field)
	GetMembers() []BACnetNameValue
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
	// IsBACnetNameValueCollection is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetNameValueCollection()
}

// _BACnetNameValueCollection is the data-structure of this message
type _BACnetNameValueCollection struct {
	OpeningTag BACnetOpeningTag
	Members    []BACnetNameValue
	ClosingTag BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

var _ BACnetNameValueCollection = (*_BACnetNameValueCollection)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetNameValueCollection) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetNameValueCollection) GetMembers() []BACnetNameValue {
	return m.Members
}

func (m *_BACnetNameValueCollection) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetNameValueCollection factory function for _BACnetNameValueCollection
func NewBACnetNameValueCollection(openingTag BACnetOpeningTag, members []BACnetNameValue, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetNameValueCollection {
	return &_BACnetNameValueCollection{OpeningTag: openingTag, Members: members, ClosingTag: closingTag, TagNumber: tagNumber}
}

// Deprecated: use the interface for direct cast
func CastBACnetNameValueCollection(structType any) BACnetNameValueCollection {
	if casted, ok := structType.(BACnetNameValueCollection); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetNameValueCollection); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetNameValueCollection) GetTypeName() string {
	return "BACnetNameValueCollection"
}

func (m *_BACnetNameValueCollection) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Array field
	if len(m.Members) > 0 {
		for _, element := range m.Members {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetNameValueCollection) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetNameValueCollectionParse(ctx context.Context, theBytes []byte, tagNumber uint8) (BACnetNameValueCollection, error) {
	return BACnetNameValueCollectionParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetNameValueCollectionParseWithBufferProducer(tagNumber uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetNameValueCollection, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetNameValueCollection, error) {
		return BACnetNameValueCollectionParseWithBuffer(ctx, readBuffer, tagNumber)
	}
}

func BACnetNameValueCollectionParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetNameValueCollection, error) {
	v, err := (&_BACnetNameValueCollection{TagNumber: tagNumber}).parse(ctx, readBuffer, tagNumber)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *_BACnetNameValueCollection) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (__bACnetNameValueCollection BACnetNameValueCollection, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetNameValueCollection"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetNameValueCollection")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}
	m.OpeningTag = openingTag

	members, err := ReadTerminatedArrayField[BACnetNameValue](ctx, "members", ReadComplex[BACnetNameValue](BACnetNameValueParseWithBuffer, readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'members' field"))
	}
	m.Members = members

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	m.ClosingTag = closingTag

	if closeErr := readBuffer.CloseContext("BACnetNameValueCollection"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetNameValueCollection")
	}

	return m, nil
}

func (m *_BACnetNameValueCollection) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetNameValueCollection) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetNameValueCollection"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetNameValueCollection")
	}

	if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'openingTag' field")
	}

	if err := WriteComplexTypeArrayField(ctx, "members", m.GetMembers(), writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'members' field")
	}

	if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetNameValueCollection"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetNameValueCollection")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetNameValueCollection) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_BACnetNameValueCollection) IsBACnetNameValueCollection() {}

func (m *_BACnetNameValueCollection) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
