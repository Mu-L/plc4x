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

// BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter is the corresponding interface of BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter
type BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetMinPriority returns MinPriority (property field)
	GetMinPriority() BACnetContextTagUnsignedInteger
	// GetMaxPriority returns MaxPriority (property field)
	GetMaxPriority() BACnetContextTagUnsignedInteger
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
}

// BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilterExactly can be used when we want exactly this type and not a type which fulfills BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter.
// This is useful for switch cases.
type BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilterExactly interface {
	BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter
	isBACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter() bool
}

// _BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter is the data-structure of this message
type _BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter struct {
	OpeningTag  BACnetOpeningTag
	MinPriority BACnetContextTagUnsignedInteger
	MaxPriority BACnetContextTagUnsignedInteger
	ClosingTag  BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter) GetMinPriority() BACnetContextTagUnsignedInteger {
	return m.MinPriority
}

func (m *_BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter) GetMaxPriority() BACnetContextTagUnsignedInteger {
	return m.MaxPriority
}

func (m *_BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter factory function for _BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter
func NewBACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter(openingTag BACnetOpeningTag, minPriority BACnetContextTagUnsignedInteger, maxPriority BACnetContextTagUnsignedInteger, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter {
	return &_BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter{OpeningTag: openingTag, MinPriority: minPriority, MaxPriority: maxPriority, ClosingTag: closingTag, TagNumber: tagNumber}
}

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter(structType any) BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter {
	if casted, ok := structType.(BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter) GetTypeName() string {
	return "BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter"
}

func (m *_BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Simple field (minPriority)
	lengthInBits += m.MinPriority.GetLengthInBits(ctx)

	// Simple field (maxPriority)
	lengthInBits += m.MaxPriority.GetLengthInBits(ctx)

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilterParse(ctx context.Context, theBytes []byte, tagNumber uint8) (BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter, error) {
	return BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilterParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilterParseWithBufferProducer(tagNumber uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter, error) {
		return BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilterParseWithBuffer(ctx, readBuffer, tagNumber)
	}
}

func BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilterParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}

	minPriority, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "minPriority", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'minPriority' field"))
	}

	maxPriority, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "maxPriority", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(1)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'maxPriority' field"))
	}

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter")
	}

	// Create the instance
	return &_BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter{
		TagNumber:   tagNumber,
		OpeningTag:  openingTag,
		MinPriority: minPriority,
		MaxPriority: maxPriority,
		ClosingTag:  closingTag,
	}, nil
}

func (m *_BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter")
	}

	// Simple Field (openingTag)
	if pushErr := writeBuffer.PushContext("openingTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for openingTag")
	}
	_openingTagErr := writeBuffer.WriteSerializable(ctx, m.GetOpeningTag())
	if popErr := writeBuffer.PopContext("openingTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for openingTag")
	}
	if _openingTagErr != nil {
		return errors.Wrap(_openingTagErr, "Error serializing 'openingTag' field")
	}

	// Simple Field (minPriority)
	if pushErr := writeBuffer.PushContext("minPriority"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for minPriority")
	}
	_minPriorityErr := writeBuffer.WriteSerializable(ctx, m.GetMinPriority())
	if popErr := writeBuffer.PopContext("minPriority"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for minPriority")
	}
	if _minPriorityErr != nil {
		return errors.Wrap(_minPriorityErr, "Error serializing 'minPriority' field")
	}

	// Simple Field (maxPriority)
	if pushErr := writeBuffer.PushContext("maxPriority"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for maxPriority")
	}
	_maxPriorityErr := writeBuffer.WriteSerializable(ctx, m.GetMaxPriority())
	if popErr := writeBuffer.PopContext("maxPriority"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for maxPriority")
	}
	if _maxPriorityErr != nil {
		return errors.Wrap(_maxPriorityErr, "Error serializing 'maxPriority' field")
	}

	// Simple Field (closingTag)
	if pushErr := writeBuffer.PushContext("closingTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for closingTag")
	}
	_closingTagErr := writeBuffer.WriteSerializable(ctx, m.GetClosingTag())
	if popErr := writeBuffer.PopContext("closingTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for closingTag")
	}
	if _closingTagErr != nil {
		return errors.Wrap(_closingTagErr, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter) isBACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter() bool {
	return true
}

func (m *_BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
