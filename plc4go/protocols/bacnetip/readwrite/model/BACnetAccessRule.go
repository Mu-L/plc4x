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

// BACnetAccessRule is the corresponding interface of BACnetAccessRule
type BACnetAccessRule interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetTimeRangeSpecifier returns TimeRangeSpecifier (property field)
	GetTimeRangeSpecifier() BACnetAccessRuleTimeRangeSpecifierTagged
	// GetTimeRange returns TimeRange (property field)
	GetTimeRange() BACnetDeviceObjectPropertyReferenceEnclosed
	// GetLocationSpecifier returns LocationSpecifier (property field)
	GetLocationSpecifier() BACnetAccessRuleLocationSpecifierTagged
	// GetLocation returns Location (property field)
	GetLocation() BACnetDeviceObjectReferenceEnclosed
	// GetEnable returns Enable (property field)
	GetEnable() BACnetContextTagBoolean
	// IsBACnetAccessRule is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetAccessRule()
}

// _BACnetAccessRule is the data-structure of this message
type _BACnetAccessRule struct {
	TimeRangeSpecifier BACnetAccessRuleTimeRangeSpecifierTagged
	TimeRange          BACnetDeviceObjectPropertyReferenceEnclosed
	LocationSpecifier  BACnetAccessRuleLocationSpecifierTagged
	Location           BACnetDeviceObjectReferenceEnclosed
	Enable             BACnetContextTagBoolean
}

var _ BACnetAccessRule = (*_BACnetAccessRule)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetAccessRule) GetTimeRangeSpecifier() BACnetAccessRuleTimeRangeSpecifierTagged {
	return m.TimeRangeSpecifier
}

func (m *_BACnetAccessRule) GetTimeRange() BACnetDeviceObjectPropertyReferenceEnclosed {
	return m.TimeRange
}

func (m *_BACnetAccessRule) GetLocationSpecifier() BACnetAccessRuleLocationSpecifierTagged {
	return m.LocationSpecifier
}

func (m *_BACnetAccessRule) GetLocation() BACnetDeviceObjectReferenceEnclosed {
	return m.Location
}

func (m *_BACnetAccessRule) GetEnable() BACnetContextTagBoolean {
	return m.Enable
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetAccessRule factory function for _BACnetAccessRule
func NewBACnetAccessRule(timeRangeSpecifier BACnetAccessRuleTimeRangeSpecifierTagged, timeRange BACnetDeviceObjectPropertyReferenceEnclosed, locationSpecifier BACnetAccessRuleLocationSpecifierTagged, location BACnetDeviceObjectReferenceEnclosed, enable BACnetContextTagBoolean) *_BACnetAccessRule {
	return &_BACnetAccessRule{TimeRangeSpecifier: timeRangeSpecifier, TimeRange: timeRange, LocationSpecifier: locationSpecifier, Location: location, Enable: enable}
}

// Deprecated: use the interface for direct cast
func CastBACnetAccessRule(structType any) BACnetAccessRule {
	if casted, ok := structType.(BACnetAccessRule); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetAccessRule); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetAccessRule) GetTypeName() string {
	return "BACnetAccessRule"
}

func (m *_BACnetAccessRule) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (timeRangeSpecifier)
	lengthInBits += m.TimeRangeSpecifier.GetLengthInBits(ctx)

	// Optional Field (timeRange)
	if m.TimeRange != nil {
		lengthInBits += m.TimeRange.GetLengthInBits(ctx)
	}

	// Simple field (locationSpecifier)
	lengthInBits += m.LocationSpecifier.GetLengthInBits(ctx)

	// Optional Field (location)
	if m.Location != nil {
		lengthInBits += m.Location.GetLengthInBits(ctx)
	}

	// Simple field (enable)
	lengthInBits += m.Enable.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetAccessRule) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetAccessRuleParse(ctx context.Context, theBytes []byte) (BACnetAccessRule, error) {
	return BACnetAccessRuleParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetAccessRuleParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAccessRule, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAccessRule, error) {
		return BACnetAccessRuleParseWithBuffer(ctx, readBuffer)
	}
}

func BACnetAccessRuleParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAccessRule, error) {
	v, err := (&_BACnetAccessRule{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *_BACnetAccessRule) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__bACnetAccessRule BACnetAccessRule, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetAccessRule"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetAccessRule")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	timeRangeSpecifier, err := ReadSimpleField[BACnetAccessRuleTimeRangeSpecifierTagged](ctx, "timeRangeSpecifier", ReadComplex[BACnetAccessRuleTimeRangeSpecifierTagged](BACnetAccessRuleTimeRangeSpecifierTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timeRangeSpecifier' field"))
	}
	m.TimeRangeSpecifier = timeRangeSpecifier

	var timeRange BACnetDeviceObjectPropertyReferenceEnclosed
	_timeRange, err := ReadOptionalField[BACnetDeviceObjectPropertyReferenceEnclosed](ctx, "timeRange", ReadComplex[BACnetDeviceObjectPropertyReferenceEnclosed](BACnetDeviceObjectPropertyReferenceEnclosedParseWithBufferProducer((uint8)(uint8(1))), readBuffer), bool((timeRangeSpecifier) != (nil)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timeRange' field"))
	}
	if _timeRange != nil {
		timeRange = *_timeRange
		m.TimeRange = timeRange
	}

	locationSpecifier, err := ReadSimpleField[BACnetAccessRuleLocationSpecifierTagged](ctx, "locationSpecifier", ReadComplex[BACnetAccessRuleLocationSpecifierTagged](BACnetAccessRuleLocationSpecifierTaggedParseWithBufferProducer((uint8)(uint8(2)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'locationSpecifier' field"))
	}
	m.LocationSpecifier = locationSpecifier

	var location BACnetDeviceObjectReferenceEnclosed
	_location, err := ReadOptionalField[BACnetDeviceObjectReferenceEnclosed](ctx, "location", ReadComplex[BACnetDeviceObjectReferenceEnclosed](BACnetDeviceObjectReferenceEnclosedParseWithBufferProducer((uint8)(uint8(3))), readBuffer), bool((locationSpecifier) != (nil)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'location' field"))
	}
	if _location != nil {
		location = *_location
		m.Location = location
	}

	enable, err := ReadSimpleField[BACnetContextTagBoolean](ctx, "enable", ReadComplex[BACnetContextTagBoolean](BACnetContextTagParseWithBufferProducer[BACnetContextTagBoolean]((uint8)(uint8(4)), (BACnetDataType)(BACnetDataType_BOOLEAN)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'enable' field"))
	}
	m.Enable = enable

	if closeErr := readBuffer.CloseContext("BACnetAccessRule"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetAccessRule")
	}

	return m, nil
}

func (m *_BACnetAccessRule) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetAccessRule) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetAccessRule"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetAccessRule")
	}

	if err := WriteSimpleField[BACnetAccessRuleTimeRangeSpecifierTagged](ctx, "timeRangeSpecifier", m.GetTimeRangeSpecifier(), WriteComplex[BACnetAccessRuleTimeRangeSpecifierTagged](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'timeRangeSpecifier' field")
	}

	if err := WriteOptionalField[BACnetDeviceObjectPropertyReferenceEnclosed](ctx, "timeRange", GetRef(m.GetTimeRange()), WriteComplex[BACnetDeviceObjectPropertyReferenceEnclosed](writeBuffer), true); err != nil {
		return errors.Wrap(err, "Error serializing 'timeRange' field")
	}

	if err := WriteSimpleField[BACnetAccessRuleLocationSpecifierTagged](ctx, "locationSpecifier", m.GetLocationSpecifier(), WriteComplex[BACnetAccessRuleLocationSpecifierTagged](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'locationSpecifier' field")
	}

	if err := WriteOptionalField[BACnetDeviceObjectReferenceEnclosed](ctx, "location", GetRef(m.GetLocation()), WriteComplex[BACnetDeviceObjectReferenceEnclosed](writeBuffer), true); err != nil {
		return errors.Wrap(err, "Error serializing 'location' field")
	}

	if err := WriteSimpleField[BACnetContextTagBoolean](ctx, "enable", m.GetEnable(), WriteComplex[BACnetContextTagBoolean](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'enable' field")
	}

	if popErr := writeBuffer.PopContext("BACnetAccessRule"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetAccessRule")
	}
	return nil
}

func (m *_BACnetAccessRule) IsBACnetAccessRule() {}

func (m *_BACnetAccessRule) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
