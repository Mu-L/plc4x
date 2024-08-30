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

// BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues is the corresponding interface of BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues
type BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetListOfLifeSavetyAlarmValues returns ListOfLifeSavetyAlarmValues (property field)
	GetListOfLifeSavetyAlarmValues() []BACnetLifeSafetyStateTagged
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
}

// BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValuesExactly can be used when we want exactly this type and not a type which fulfills BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues.
// This is useful for switch cases.
type BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValuesExactly interface {
	BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues
	isBACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues() bool
}

// _BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues is the data-structure of this message
type _BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues struct {
	OpeningTag                  BACnetOpeningTag
	ListOfLifeSavetyAlarmValues []BACnetLifeSafetyStateTagged
	ClosingTag                  BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues) GetListOfLifeSavetyAlarmValues() []BACnetLifeSafetyStateTagged {
	return m.ListOfLifeSavetyAlarmValues
}

func (m *_BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues factory function for _BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues
func NewBACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues(openingTag BACnetOpeningTag, listOfLifeSavetyAlarmValues []BACnetLifeSafetyStateTagged, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues {
	return &_BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues{OpeningTag: openingTag, ListOfLifeSavetyAlarmValues: listOfLifeSavetyAlarmValues, ClosingTag: closingTag, TagNumber: tagNumber}
}

// Deprecated: use the interface for direct cast
func CastBACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues(structType any) BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues {
	if casted, ok := structType.(BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues) GetTypeName() string {
	return "BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues"
}

func (m *_BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Array field
	if len(m.ListOfLifeSavetyAlarmValues) > 0 {
		for _, element := range m.ListOfLifeSavetyAlarmValues {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValuesParse(ctx context.Context, theBytes []byte, tagNumber uint8) (BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues, error) {
	return BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValuesParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValuesParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for openingTag")
	}
	_openingTag, _openingTagErr := BACnetOpeningTagParseWithBuffer(ctx, readBuffer, uint8(tagNumber))
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field of BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues")
	}
	openingTag := _openingTag.(BACnetOpeningTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for openingTag")
	}

	listOfLifeSavetyAlarmValues, err := ReadTerminatedArrayField[BACnetLifeSafetyStateTagged](ctx, "listOfLifeSavetyAlarmValues", ReadComplex[BACnetLifeSafetyStateTagged](func(ctx context.Context, buffer utils.ReadBuffer) (BACnetLifeSafetyStateTagged, error) {
		v, err := BACnetLifeSafetyStateTaggedParseWithBuffer(ctx, readBuffer, (uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS))
		if err != nil {
			return nil, err
		}
		return v.(BACnetLifeSafetyStateTagged), nil
	}, readBuffer), func() bool { return IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber) })
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'listOfLifeSavetyAlarmValues' field"))
	}

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for closingTag")
	}
	_closingTag, _closingTagErr := BACnetClosingTagParseWithBuffer(ctx, readBuffer, uint8(tagNumber))
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field of BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues")
	}
	closingTag := _closingTag.(BACnetClosingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for closingTag")
	}

	if closeErr := readBuffer.CloseContext("BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues")
	}

	// Create the instance
	return &_BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues{
		TagNumber:                   tagNumber,
		OpeningTag:                  openingTag,
		ListOfLifeSavetyAlarmValues: listOfLifeSavetyAlarmValues,
		ClosingTag:                  closingTag,
	}, nil
}

func (m *_BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues")
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

	// Array Field (listOfLifeSavetyAlarmValues)
	if pushErr := writeBuffer.PushContext("listOfLifeSavetyAlarmValues", utils.WithRenderAsList(true)); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for listOfLifeSavetyAlarmValues")
	}
	for _curItem, _element := range m.GetListOfLifeSavetyAlarmValues() {
		_ = _curItem
		arrayCtx := utils.CreateArrayContext(ctx, len(m.GetListOfLifeSavetyAlarmValues()), _curItem)
		_ = arrayCtx
		_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
		if _elementErr != nil {
			return errors.Wrap(_elementErr, "Error serializing 'listOfLifeSavetyAlarmValues' field")
		}
	}
	if popErr := writeBuffer.PopContext("listOfLifeSavetyAlarmValues", utils.WithRenderAsList(true)); popErr != nil {
		return errors.Wrap(popErr, "Error popping for listOfLifeSavetyAlarmValues")
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

	if popErr := writeBuffer.PopContext("BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues) isBACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues() bool {
	return true
}

func (m *_BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
