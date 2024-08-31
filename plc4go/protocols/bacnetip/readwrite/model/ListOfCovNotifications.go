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

// ListOfCovNotifications is the corresponding interface of ListOfCovNotifications
type ListOfCovNotifications interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetMonitoredObjectIdentifier returns MonitoredObjectIdentifier (property field)
	GetMonitoredObjectIdentifier() BACnetContextTagObjectIdentifier
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetListOfValues returns ListOfValues (property field)
	GetListOfValues() []ListOfCovNotificationsValue
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
}

// ListOfCovNotificationsExactly can be used when we want exactly this type and not a type which fulfills ListOfCovNotifications.
// This is useful for switch cases.
type ListOfCovNotificationsExactly interface {
	ListOfCovNotifications
	isListOfCovNotifications() bool
}

// _ListOfCovNotifications is the data-structure of this message
type _ListOfCovNotifications struct {
	MonitoredObjectIdentifier BACnetContextTagObjectIdentifier
	OpeningTag                BACnetOpeningTag
	ListOfValues              []ListOfCovNotificationsValue
	ClosingTag                BACnetClosingTag
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ListOfCovNotifications) GetMonitoredObjectIdentifier() BACnetContextTagObjectIdentifier {
	return m.MonitoredObjectIdentifier
}

func (m *_ListOfCovNotifications) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_ListOfCovNotifications) GetListOfValues() []ListOfCovNotificationsValue {
	return m.ListOfValues
}

func (m *_ListOfCovNotifications) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewListOfCovNotifications factory function for _ListOfCovNotifications
func NewListOfCovNotifications(monitoredObjectIdentifier BACnetContextTagObjectIdentifier, openingTag BACnetOpeningTag, listOfValues []ListOfCovNotificationsValue, closingTag BACnetClosingTag) *_ListOfCovNotifications {
	return &_ListOfCovNotifications{MonitoredObjectIdentifier: monitoredObjectIdentifier, OpeningTag: openingTag, ListOfValues: listOfValues, ClosingTag: closingTag}
}

// Deprecated: use the interface for direct cast
func CastListOfCovNotifications(structType any) ListOfCovNotifications {
	if casted, ok := structType.(ListOfCovNotifications); ok {
		return casted
	}
	if casted, ok := structType.(*ListOfCovNotifications); ok {
		return *casted
	}
	return nil
}

func (m *_ListOfCovNotifications) GetTypeName() string {
	return "ListOfCovNotifications"
}

func (m *_ListOfCovNotifications) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (monitoredObjectIdentifier)
	lengthInBits += m.MonitoredObjectIdentifier.GetLengthInBits(ctx)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Array field
	if len(m.ListOfValues) > 0 {
		for _, element := range m.ListOfValues {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_ListOfCovNotifications) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ListOfCovNotificationsParse(ctx context.Context, theBytes []byte) (ListOfCovNotifications, error) {
	return ListOfCovNotificationsParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func ListOfCovNotificationsParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (ListOfCovNotifications, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (ListOfCovNotifications, error) {
		return ListOfCovNotificationsParseWithBuffer(ctx, readBuffer)
	}
}

func ListOfCovNotificationsParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (ListOfCovNotifications, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("ListOfCovNotifications"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ListOfCovNotifications")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	monitoredObjectIdentifier, err := ReadSimpleField[BACnetContextTagObjectIdentifier](ctx, "monitoredObjectIdentifier", ReadComplex[BACnetContextTagObjectIdentifier](BACnetContextTagParseWithBufferProducer[BACnetContextTagObjectIdentifier]((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_BACNET_OBJECT_IDENTIFIER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'monitoredObjectIdentifier' field"))
	}

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(uint8(1))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}

	listOfValues, err := ReadTerminatedArrayField[ListOfCovNotificationsValue](ctx, "listOfValues", ReadComplex[ListOfCovNotificationsValue](ListOfCovNotificationsValueParseWithBufferProducer((BACnetObjectType)(monitoredObjectIdentifier.GetObjectType())), readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, 1))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'listOfValues' field"))
	}

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(uint8(1))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}

	if closeErr := readBuffer.CloseContext("ListOfCovNotifications"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ListOfCovNotifications")
	}

	// Create the instance
	return &_ListOfCovNotifications{
		MonitoredObjectIdentifier: monitoredObjectIdentifier,
		OpeningTag:                openingTag,
		ListOfValues:              listOfValues,
		ClosingTag:                closingTag,
	}, nil
}

func (m *_ListOfCovNotifications) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ListOfCovNotifications) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("ListOfCovNotifications"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ListOfCovNotifications")
	}

	// Simple Field (monitoredObjectIdentifier)
	if pushErr := writeBuffer.PushContext("monitoredObjectIdentifier"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for monitoredObjectIdentifier")
	}
	_monitoredObjectIdentifierErr := writeBuffer.WriteSerializable(ctx, m.GetMonitoredObjectIdentifier())
	if popErr := writeBuffer.PopContext("monitoredObjectIdentifier"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for monitoredObjectIdentifier")
	}
	if _monitoredObjectIdentifierErr != nil {
		return errors.Wrap(_monitoredObjectIdentifierErr, "Error serializing 'monitoredObjectIdentifier' field")
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

	// Array Field (listOfValues)
	if pushErr := writeBuffer.PushContext("listOfValues", utils.WithRenderAsList(true)); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for listOfValues")
	}
	for _curItem, _element := range m.GetListOfValues() {
		_ = _curItem
		arrayCtx := utils.CreateArrayContext(ctx, len(m.GetListOfValues()), _curItem)
		_ = arrayCtx
		_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
		if _elementErr != nil {
			return errors.Wrap(_elementErr, "Error serializing 'listOfValues' field")
		}
	}
	if popErr := writeBuffer.PopContext("listOfValues", utils.WithRenderAsList(true)); popErr != nil {
		return errors.Wrap(popErr, "Error popping for listOfValues")
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

	if popErr := writeBuffer.PopContext("ListOfCovNotifications"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ListOfCovNotifications")
	}
	return nil
}

func (m *_ListOfCovNotifications) isListOfCovNotifications() bool {
	return true
}

func (m *_ListOfCovNotifications) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
