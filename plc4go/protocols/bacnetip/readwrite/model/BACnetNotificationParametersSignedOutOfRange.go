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

// BACnetNotificationParametersSignedOutOfRange is the corresponding interface of BACnetNotificationParametersSignedOutOfRange
type BACnetNotificationParametersSignedOutOfRange interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetNotificationParameters
	// GetInnerOpeningTag returns InnerOpeningTag (property field)
	GetInnerOpeningTag() BACnetOpeningTag
	// GetExceedingValue returns ExceedingValue (property field)
	GetExceedingValue() BACnetContextTagSignedInteger
	// GetStatusFlags returns StatusFlags (property field)
	GetStatusFlags() BACnetStatusFlagsTagged
	// GetDeadband returns Deadband (property field)
	GetDeadband() BACnetContextTagUnsignedInteger
	// GetExceededLimit returns ExceededLimit (property field)
	GetExceededLimit() BACnetContextTagSignedInteger
	// GetInnerClosingTag returns InnerClosingTag (property field)
	GetInnerClosingTag() BACnetClosingTag
}

// BACnetNotificationParametersSignedOutOfRangeExactly can be used when we want exactly this type and not a type which fulfills BACnetNotificationParametersSignedOutOfRange.
// This is useful for switch cases.
type BACnetNotificationParametersSignedOutOfRangeExactly interface {
	BACnetNotificationParametersSignedOutOfRange
	isBACnetNotificationParametersSignedOutOfRange() bool
}

// _BACnetNotificationParametersSignedOutOfRange is the data-structure of this message
type _BACnetNotificationParametersSignedOutOfRange struct {
	*_BACnetNotificationParameters
	InnerOpeningTag BACnetOpeningTag
	ExceedingValue  BACnetContextTagSignedInteger
	StatusFlags     BACnetStatusFlagsTagged
	Deadband        BACnetContextTagUnsignedInteger
	ExceededLimit   BACnetContextTagSignedInteger
	InnerClosingTag BACnetClosingTag
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetNotificationParametersSignedOutOfRange) InitializeParent(parent BACnetNotificationParameters, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetNotificationParametersSignedOutOfRange) GetParent() BACnetNotificationParameters {
	return m._BACnetNotificationParameters
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetNotificationParametersSignedOutOfRange) GetInnerOpeningTag() BACnetOpeningTag {
	return m.InnerOpeningTag
}

func (m *_BACnetNotificationParametersSignedOutOfRange) GetExceedingValue() BACnetContextTagSignedInteger {
	return m.ExceedingValue
}

func (m *_BACnetNotificationParametersSignedOutOfRange) GetStatusFlags() BACnetStatusFlagsTagged {
	return m.StatusFlags
}

func (m *_BACnetNotificationParametersSignedOutOfRange) GetDeadband() BACnetContextTagUnsignedInteger {
	return m.Deadband
}

func (m *_BACnetNotificationParametersSignedOutOfRange) GetExceededLimit() BACnetContextTagSignedInteger {
	return m.ExceededLimit
}

func (m *_BACnetNotificationParametersSignedOutOfRange) GetInnerClosingTag() BACnetClosingTag {
	return m.InnerClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetNotificationParametersSignedOutOfRange factory function for _BACnetNotificationParametersSignedOutOfRange
func NewBACnetNotificationParametersSignedOutOfRange(innerOpeningTag BACnetOpeningTag, exceedingValue BACnetContextTagSignedInteger, statusFlags BACnetStatusFlagsTagged, deadband BACnetContextTagUnsignedInteger, exceededLimit BACnetContextTagSignedInteger, innerClosingTag BACnetClosingTag, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, objectTypeArgument BACnetObjectType) *_BACnetNotificationParametersSignedOutOfRange {
	_result := &_BACnetNotificationParametersSignedOutOfRange{
		InnerOpeningTag:               innerOpeningTag,
		ExceedingValue:                exceedingValue,
		StatusFlags:                   statusFlags,
		Deadband:                      deadband,
		ExceededLimit:                 exceededLimit,
		InnerClosingTag:               innerClosingTag,
		_BACnetNotificationParameters: NewBACnetNotificationParameters(openingTag, peekedTagHeader, closingTag, tagNumber, objectTypeArgument),
	}
	_result._BACnetNotificationParameters._BACnetNotificationParametersChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetNotificationParametersSignedOutOfRange(structType any) BACnetNotificationParametersSignedOutOfRange {
	if casted, ok := structType.(BACnetNotificationParametersSignedOutOfRange); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetNotificationParametersSignedOutOfRange); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetNotificationParametersSignedOutOfRange) GetTypeName() string {
	return "BACnetNotificationParametersSignedOutOfRange"
}

func (m *_BACnetNotificationParametersSignedOutOfRange) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (innerOpeningTag)
	lengthInBits += m.InnerOpeningTag.GetLengthInBits(ctx)

	// Simple field (exceedingValue)
	lengthInBits += m.ExceedingValue.GetLengthInBits(ctx)

	// Simple field (statusFlags)
	lengthInBits += m.StatusFlags.GetLengthInBits(ctx)

	// Simple field (deadband)
	lengthInBits += m.Deadband.GetLengthInBits(ctx)

	// Simple field (exceededLimit)
	lengthInBits += m.ExceededLimit.GetLengthInBits(ctx)

	// Simple field (innerClosingTag)
	lengthInBits += m.InnerClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetNotificationParametersSignedOutOfRange) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetNotificationParametersSignedOutOfRangeParse(ctx context.Context, theBytes []byte, peekedTagNumber uint8, tagNumber uint8, objectTypeArgument BACnetObjectType) (BACnetNotificationParametersSignedOutOfRange, error) {
	return BACnetNotificationParametersSignedOutOfRangeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), peekedTagNumber, tagNumber, objectTypeArgument)
}

func BACnetNotificationParametersSignedOutOfRangeParseWithBufferProducer(peekedTagNumber uint8, tagNumber uint8, objectTypeArgument BACnetObjectType) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetNotificationParametersSignedOutOfRange, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetNotificationParametersSignedOutOfRange, error) {
		return BACnetNotificationParametersSignedOutOfRangeParseWithBuffer(ctx, readBuffer, peekedTagNumber, tagNumber, objectTypeArgument)
	}
}

func BACnetNotificationParametersSignedOutOfRangeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, peekedTagNumber uint8, tagNumber uint8, objectTypeArgument BACnetObjectType) (BACnetNotificationParametersSignedOutOfRange, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetNotificationParametersSignedOutOfRange"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetNotificationParametersSignedOutOfRange")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	innerOpeningTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "innerOpeningTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(peekedTagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'innerOpeningTag' field"))
	}

	exceedingValue, err := ReadSimpleField[BACnetContextTagSignedInteger](ctx, "exceedingValue", ReadComplex[BACnetContextTagSignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagSignedInteger]((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_SIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'exceedingValue' field"))
	}

	statusFlags, err := ReadSimpleField[BACnetStatusFlagsTagged](ctx, "statusFlags", ReadComplex[BACnetStatusFlagsTagged](BACnetStatusFlagsTaggedParseWithBufferProducer((uint8)(uint8(1)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'statusFlags' field"))
	}

	deadband, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "deadband", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(2)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'deadband' field"))
	}

	exceededLimit, err := ReadSimpleField[BACnetContextTagSignedInteger](ctx, "exceededLimit", ReadComplex[BACnetContextTagSignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagSignedInteger]((uint8)(uint8(3)), (BACnetDataType)(BACnetDataType_SIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'exceededLimit' field"))
	}

	innerClosingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "innerClosingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(peekedTagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'innerClosingTag' field"))
	}

	if closeErr := readBuffer.CloseContext("BACnetNotificationParametersSignedOutOfRange"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetNotificationParametersSignedOutOfRange")
	}

	// Create a partially initialized instance
	_child := &_BACnetNotificationParametersSignedOutOfRange{
		_BACnetNotificationParameters: &_BACnetNotificationParameters{
			TagNumber:          tagNumber,
			ObjectTypeArgument: objectTypeArgument,
		},
		InnerOpeningTag: innerOpeningTag,
		ExceedingValue:  exceedingValue,
		StatusFlags:     statusFlags,
		Deadband:        deadband,
		ExceededLimit:   exceededLimit,
		InnerClosingTag: innerClosingTag,
	}
	_child._BACnetNotificationParameters._BACnetNotificationParametersChildRequirements = _child
	return _child, nil
}

func (m *_BACnetNotificationParametersSignedOutOfRange) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetNotificationParametersSignedOutOfRange) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetNotificationParametersSignedOutOfRange"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetNotificationParametersSignedOutOfRange")
		}

		// Simple Field (innerOpeningTag)
		if pushErr := writeBuffer.PushContext("innerOpeningTag"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for innerOpeningTag")
		}
		_innerOpeningTagErr := writeBuffer.WriteSerializable(ctx, m.GetInnerOpeningTag())
		if popErr := writeBuffer.PopContext("innerOpeningTag"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for innerOpeningTag")
		}
		if _innerOpeningTagErr != nil {
			return errors.Wrap(_innerOpeningTagErr, "Error serializing 'innerOpeningTag' field")
		}

		// Simple Field (exceedingValue)
		if pushErr := writeBuffer.PushContext("exceedingValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for exceedingValue")
		}
		_exceedingValueErr := writeBuffer.WriteSerializable(ctx, m.GetExceedingValue())
		if popErr := writeBuffer.PopContext("exceedingValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for exceedingValue")
		}
		if _exceedingValueErr != nil {
			return errors.Wrap(_exceedingValueErr, "Error serializing 'exceedingValue' field")
		}

		// Simple Field (statusFlags)
		if pushErr := writeBuffer.PushContext("statusFlags"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for statusFlags")
		}
		_statusFlagsErr := writeBuffer.WriteSerializable(ctx, m.GetStatusFlags())
		if popErr := writeBuffer.PopContext("statusFlags"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for statusFlags")
		}
		if _statusFlagsErr != nil {
			return errors.Wrap(_statusFlagsErr, "Error serializing 'statusFlags' field")
		}

		// Simple Field (deadband)
		if pushErr := writeBuffer.PushContext("deadband"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for deadband")
		}
		_deadbandErr := writeBuffer.WriteSerializable(ctx, m.GetDeadband())
		if popErr := writeBuffer.PopContext("deadband"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for deadband")
		}
		if _deadbandErr != nil {
			return errors.Wrap(_deadbandErr, "Error serializing 'deadband' field")
		}

		// Simple Field (exceededLimit)
		if pushErr := writeBuffer.PushContext("exceededLimit"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for exceededLimit")
		}
		_exceededLimitErr := writeBuffer.WriteSerializable(ctx, m.GetExceededLimit())
		if popErr := writeBuffer.PopContext("exceededLimit"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for exceededLimit")
		}
		if _exceededLimitErr != nil {
			return errors.Wrap(_exceededLimitErr, "Error serializing 'exceededLimit' field")
		}

		// Simple Field (innerClosingTag)
		if pushErr := writeBuffer.PushContext("innerClosingTag"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for innerClosingTag")
		}
		_innerClosingTagErr := writeBuffer.WriteSerializable(ctx, m.GetInnerClosingTag())
		if popErr := writeBuffer.PopContext("innerClosingTag"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for innerClosingTag")
		}
		if _innerClosingTagErr != nil {
			return errors.Wrap(_innerClosingTagErr, "Error serializing 'innerClosingTag' field")
		}

		if popErr := writeBuffer.PopContext("BACnetNotificationParametersSignedOutOfRange"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetNotificationParametersSignedOutOfRange")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetNotificationParametersSignedOutOfRange) isBACnetNotificationParametersSignedOutOfRange() bool {
	return true
}

func (m *_BACnetNotificationParametersSignedOutOfRange) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
