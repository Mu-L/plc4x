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

// BACnetNotificationParametersChangeOfLifeSafety is the corresponding interface of BACnetNotificationParametersChangeOfLifeSafety
type BACnetNotificationParametersChangeOfLifeSafety interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetNotificationParameters
	// GetInnerOpeningTag returns InnerOpeningTag (property field)
	GetInnerOpeningTag() BACnetOpeningTag
	// GetNewState returns NewState (property field)
	GetNewState() BACnetLifeSafetyStateTagged
	// GetNewMode returns NewMode (property field)
	GetNewMode() BACnetLifeSafetyModeTagged
	// GetStatusFlags returns StatusFlags (property field)
	GetStatusFlags() BACnetStatusFlagsTagged
	// GetOperationExpected returns OperationExpected (property field)
	GetOperationExpected() BACnetLifeSafetyOperationTagged
	// GetInnerClosingTag returns InnerClosingTag (property field)
	GetInnerClosingTag() BACnetClosingTag
}

// BACnetNotificationParametersChangeOfLifeSafetyExactly can be used when we want exactly this type and not a type which fulfills BACnetNotificationParametersChangeOfLifeSafety.
// This is useful for switch cases.
type BACnetNotificationParametersChangeOfLifeSafetyExactly interface {
	BACnetNotificationParametersChangeOfLifeSafety
	isBACnetNotificationParametersChangeOfLifeSafety() bool
}

// _BACnetNotificationParametersChangeOfLifeSafety is the data-structure of this message
type _BACnetNotificationParametersChangeOfLifeSafety struct {
	BACnetNotificationParametersContract
	InnerOpeningTag   BACnetOpeningTag
	NewState          BACnetLifeSafetyStateTagged
	NewMode           BACnetLifeSafetyModeTagged
	StatusFlags       BACnetStatusFlagsTagged
	OperationExpected BACnetLifeSafetyOperationTagged
	InnerClosingTag   BACnetClosingTag
}

var _ BACnetNotificationParametersChangeOfLifeSafety = (*_BACnetNotificationParametersChangeOfLifeSafety)(nil)
var _ BACnetNotificationParametersRequirements = (*_BACnetNotificationParametersChangeOfLifeSafety)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetNotificationParametersChangeOfLifeSafety) GetParent() BACnetNotificationParametersContract {
	return m.BACnetNotificationParametersContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetNotificationParametersChangeOfLifeSafety) GetInnerOpeningTag() BACnetOpeningTag {
	return m.InnerOpeningTag
}

func (m *_BACnetNotificationParametersChangeOfLifeSafety) GetNewState() BACnetLifeSafetyStateTagged {
	return m.NewState
}

func (m *_BACnetNotificationParametersChangeOfLifeSafety) GetNewMode() BACnetLifeSafetyModeTagged {
	return m.NewMode
}

func (m *_BACnetNotificationParametersChangeOfLifeSafety) GetStatusFlags() BACnetStatusFlagsTagged {
	return m.StatusFlags
}

func (m *_BACnetNotificationParametersChangeOfLifeSafety) GetOperationExpected() BACnetLifeSafetyOperationTagged {
	return m.OperationExpected
}

func (m *_BACnetNotificationParametersChangeOfLifeSafety) GetInnerClosingTag() BACnetClosingTag {
	return m.InnerClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetNotificationParametersChangeOfLifeSafety factory function for _BACnetNotificationParametersChangeOfLifeSafety
func NewBACnetNotificationParametersChangeOfLifeSafety(innerOpeningTag BACnetOpeningTag, newState BACnetLifeSafetyStateTagged, newMode BACnetLifeSafetyModeTagged, statusFlags BACnetStatusFlagsTagged, operationExpected BACnetLifeSafetyOperationTagged, innerClosingTag BACnetClosingTag, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, objectTypeArgument BACnetObjectType) *_BACnetNotificationParametersChangeOfLifeSafety {
	_result := &_BACnetNotificationParametersChangeOfLifeSafety{
		BACnetNotificationParametersContract: NewBACnetNotificationParameters(openingTag, peekedTagHeader, closingTag, tagNumber, objectTypeArgument),
		InnerOpeningTag:                      innerOpeningTag,
		NewState:                             newState,
		NewMode:                              newMode,
		StatusFlags:                          statusFlags,
		OperationExpected:                    operationExpected,
		InnerClosingTag:                      innerClosingTag,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetNotificationParametersChangeOfLifeSafety(structType any) BACnetNotificationParametersChangeOfLifeSafety {
	if casted, ok := structType.(BACnetNotificationParametersChangeOfLifeSafety); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetNotificationParametersChangeOfLifeSafety); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetNotificationParametersChangeOfLifeSafety) GetTypeName() string {
	return "BACnetNotificationParametersChangeOfLifeSafety"
}

func (m *_BACnetNotificationParametersChangeOfLifeSafety) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetNotificationParametersContract.(*_BACnetNotificationParameters).getLengthInBits(ctx))

	// Simple field (innerOpeningTag)
	lengthInBits += m.InnerOpeningTag.GetLengthInBits(ctx)

	// Simple field (newState)
	lengthInBits += m.NewState.GetLengthInBits(ctx)

	// Simple field (newMode)
	lengthInBits += m.NewMode.GetLengthInBits(ctx)

	// Simple field (statusFlags)
	lengthInBits += m.StatusFlags.GetLengthInBits(ctx)

	// Simple field (operationExpected)
	lengthInBits += m.OperationExpected.GetLengthInBits(ctx)

	// Simple field (innerClosingTag)
	lengthInBits += m.InnerClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetNotificationParametersChangeOfLifeSafety) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetNotificationParametersChangeOfLifeSafety) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetNotificationParameters, peekedTagNumber uint8, tagNumber uint8, objectTypeArgument BACnetObjectType) (__bACnetNotificationParametersChangeOfLifeSafety BACnetNotificationParametersChangeOfLifeSafety, err error) {
	m.BACnetNotificationParametersContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetNotificationParametersChangeOfLifeSafety"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetNotificationParametersChangeOfLifeSafety")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	innerOpeningTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "innerOpeningTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(peekedTagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'innerOpeningTag' field"))
	}
	m.InnerOpeningTag = innerOpeningTag

	newState, err := ReadSimpleField[BACnetLifeSafetyStateTagged](ctx, "newState", ReadComplex[BACnetLifeSafetyStateTagged](BACnetLifeSafetyStateTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'newState' field"))
	}
	m.NewState = newState

	newMode, err := ReadSimpleField[BACnetLifeSafetyModeTagged](ctx, "newMode", ReadComplex[BACnetLifeSafetyModeTagged](BACnetLifeSafetyModeTaggedParseWithBufferProducer((uint8)(uint8(1)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'newMode' field"))
	}
	m.NewMode = newMode

	statusFlags, err := ReadSimpleField[BACnetStatusFlagsTagged](ctx, "statusFlags", ReadComplex[BACnetStatusFlagsTagged](BACnetStatusFlagsTaggedParseWithBufferProducer((uint8)(uint8(2)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'statusFlags' field"))
	}
	m.StatusFlags = statusFlags

	operationExpected, err := ReadSimpleField[BACnetLifeSafetyOperationTagged](ctx, "operationExpected", ReadComplex[BACnetLifeSafetyOperationTagged](BACnetLifeSafetyOperationTaggedParseWithBufferProducer((uint8)(uint8(3)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'operationExpected' field"))
	}
	m.OperationExpected = operationExpected

	innerClosingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "innerClosingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(peekedTagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'innerClosingTag' field"))
	}
	m.InnerClosingTag = innerClosingTag

	if closeErr := readBuffer.CloseContext("BACnetNotificationParametersChangeOfLifeSafety"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetNotificationParametersChangeOfLifeSafety")
	}

	return m, nil
}

func (m *_BACnetNotificationParametersChangeOfLifeSafety) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetNotificationParametersChangeOfLifeSafety) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetNotificationParametersChangeOfLifeSafety"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetNotificationParametersChangeOfLifeSafety")
		}

		if err := WriteSimpleField[BACnetOpeningTag](ctx, "innerOpeningTag", m.GetInnerOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'innerOpeningTag' field")
		}

		if err := WriteSimpleField[BACnetLifeSafetyStateTagged](ctx, "newState", m.GetNewState(), WriteComplex[BACnetLifeSafetyStateTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'newState' field")
		}

		if err := WriteSimpleField[BACnetLifeSafetyModeTagged](ctx, "newMode", m.GetNewMode(), WriteComplex[BACnetLifeSafetyModeTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'newMode' field")
		}

		if err := WriteSimpleField[BACnetStatusFlagsTagged](ctx, "statusFlags", m.GetStatusFlags(), WriteComplex[BACnetStatusFlagsTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'statusFlags' field")
		}

		if err := WriteSimpleField[BACnetLifeSafetyOperationTagged](ctx, "operationExpected", m.GetOperationExpected(), WriteComplex[BACnetLifeSafetyOperationTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'operationExpected' field")
		}

		if err := WriteSimpleField[BACnetClosingTag](ctx, "innerClosingTag", m.GetInnerClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'innerClosingTag' field")
		}

		if popErr := writeBuffer.PopContext("BACnetNotificationParametersChangeOfLifeSafety"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetNotificationParametersChangeOfLifeSafety")
		}
		return nil
	}
	return m.BACnetNotificationParametersContract.(*_BACnetNotificationParameters).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetNotificationParametersChangeOfLifeSafety) isBACnetNotificationParametersChangeOfLifeSafety() bool {
	return true
}

func (m *_BACnetNotificationParametersChangeOfLifeSafety) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
