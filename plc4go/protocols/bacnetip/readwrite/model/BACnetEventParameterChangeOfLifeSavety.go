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

// BACnetEventParameterChangeOfLifeSavety is the corresponding interface of BACnetEventParameterChangeOfLifeSavety
type BACnetEventParameterChangeOfLifeSavety interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetEventParameter
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetTimeDelay returns TimeDelay (property field)
	GetTimeDelay() BACnetContextTagUnsignedInteger
	// GetListOfLifeSavetyAlarmValues returns ListOfLifeSavetyAlarmValues (property field)
	GetListOfLifeSavetyAlarmValues() BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues
	// GetListOfAlarmValues returns ListOfAlarmValues (property field)
	GetListOfAlarmValues() BACnetEventParameterChangeOfLifeSavetyListOfAlarmValues
	// GetModePropertyReference returns ModePropertyReference (property field)
	GetModePropertyReference() BACnetDeviceObjectPropertyReferenceEnclosed
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
	// IsBACnetEventParameterChangeOfLifeSavety is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetEventParameterChangeOfLifeSavety()
}

// _BACnetEventParameterChangeOfLifeSavety is the data-structure of this message
type _BACnetEventParameterChangeOfLifeSavety struct {
	BACnetEventParameterContract
	OpeningTag                  BACnetOpeningTag
	TimeDelay                   BACnetContextTagUnsignedInteger
	ListOfLifeSavetyAlarmValues BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues
	ListOfAlarmValues           BACnetEventParameterChangeOfLifeSavetyListOfAlarmValues
	ModePropertyReference       BACnetDeviceObjectPropertyReferenceEnclosed
	ClosingTag                  BACnetClosingTag
}

var _ BACnetEventParameterChangeOfLifeSavety = (*_BACnetEventParameterChangeOfLifeSavety)(nil)
var _ BACnetEventParameterRequirements = (*_BACnetEventParameterChangeOfLifeSavety)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetEventParameterChangeOfLifeSavety) GetParent() BACnetEventParameterContract {
	return m.BACnetEventParameterContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetEventParameterChangeOfLifeSavety) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetEventParameterChangeOfLifeSavety) GetTimeDelay() BACnetContextTagUnsignedInteger {
	return m.TimeDelay
}

func (m *_BACnetEventParameterChangeOfLifeSavety) GetListOfLifeSavetyAlarmValues() BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues {
	return m.ListOfLifeSavetyAlarmValues
}

func (m *_BACnetEventParameterChangeOfLifeSavety) GetListOfAlarmValues() BACnetEventParameterChangeOfLifeSavetyListOfAlarmValues {
	return m.ListOfAlarmValues
}

func (m *_BACnetEventParameterChangeOfLifeSavety) GetModePropertyReference() BACnetDeviceObjectPropertyReferenceEnclosed {
	return m.ModePropertyReference
}

func (m *_BACnetEventParameterChangeOfLifeSavety) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetEventParameterChangeOfLifeSavety factory function for _BACnetEventParameterChangeOfLifeSavety
func NewBACnetEventParameterChangeOfLifeSavety(openingTag BACnetOpeningTag, timeDelay BACnetContextTagUnsignedInteger, listOfLifeSavetyAlarmValues BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues, listOfAlarmValues BACnetEventParameterChangeOfLifeSavetyListOfAlarmValues, modePropertyReference BACnetDeviceObjectPropertyReferenceEnclosed, closingTag BACnetClosingTag, peekedTagHeader BACnetTagHeader) *_BACnetEventParameterChangeOfLifeSavety {
	_result := &_BACnetEventParameterChangeOfLifeSavety{
		BACnetEventParameterContract: NewBACnetEventParameter(peekedTagHeader),
		OpeningTag:                   openingTag,
		TimeDelay:                    timeDelay,
		ListOfLifeSavetyAlarmValues:  listOfLifeSavetyAlarmValues,
		ListOfAlarmValues:            listOfAlarmValues,
		ModePropertyReference:        modePropertyReference,
		ClosingTag:                   closingTag,
	}
	_result.BACnetEventParameterContract.(*_BACnetEventParameter)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetEventParameterChangeOfLifeSavety(structType any) BACnetEventParameterChangeOfLifeSavety {
	if casted, ok := structType.(BACnetEventParameterChangeOfLifeSavety); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetEventParameterChangeOfLifeSavety); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetEventParameterChangeOfLifeSavety) GetTypeName() string {
	return "BACnetEventParameterChangeOfLifeSavety"
}

func (m *_BACnetEventParameterChangeOfLifeSavety) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetEventParameterContract.(*_BACnetEventParameter).getLengthInBits(ctx))

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Simple field (timeDelay)
	lengthInBits += m.TimeDelay.GetLengthInBits(ctx)

	// Simple field (listOfLifeSavetyAlarmValues)
	lengthInBits += m.ListOfLifeSavetyAlarmValues.GetLengthInBits(ctx)

	// Simple field (listOfAlarmValues)
	lengthInBits += m.ListOfAlarmValues.GetLengthInBits(ctx)

	// Simple field (modePropertyReference)
	lengthInBits += m.ModePropertyReference.GetLengthInBits(ctx)

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetEventParameterChangeOfLifeSavety) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetEventParameterChangeOfLifeSavety) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetEventParameter) (__bACnetEventParameterChangeOfLifeSavety BACnetEventParameterChangeOfLifeSavety, err error) {
	m.BACnetEventParameterContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetEventParameterChangeOfLifeSavety"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetEventParameterChangeOfLifeSavety")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(uint8(8))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}
	m.OpeningTag = openingTag

	timeDelay, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "timeDelay", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timeDelay' field"))
	}
	m.TimeDelay = timeDelay

	listOfLifeSavetyAlarmValues, err := ReadSimpleField[BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues](ctx, "listOfLifeSavetyAlarmValues", ReadComplex[BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues](BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValuesParseWithBufferProducer((uint8)(uint8(1))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'listOfLifeSavetyAlarmValues' field"))
	}
	m.ListOfLifeSavetyAlarmValues = listOfLifeSavetyAlarmValues

	listOfAlarmValues, err := ReadSimpleField[BACnetEventParameterChangeOfLifeSavetyListOfAlarmValues](ctx, "listOfAlarmValues", ReadComplex[BACnetEventParameterChangeOfLifeSavetyListOfAlarmValues](BACnetEventParameterChangeOfLifeSavetyListOfAlarmValuesParseWithBufferProducer((uint8)(uint8(2))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'listOfAlarmValues' field"))
	}
	m.ListOfAlarmValues = listOfAlarmValues

	modePropertyReference, err := ReadSimpleField[BACnetDeviceObjectPropertyReferenceEnclosed](ctx, "modePropertyReference", ReadComplex[BACnetDeviceObjectPropertyReferenceEnclosed](BACnetDeviceObjectPropertyReferenceEnclosedParseWithBufferProducer((uint8)(uint8(4))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'modePropertyReference' field"))
	}
	m.ModePropertyReference = modePropertyReference

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(uint8(8))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	m.ClosingTag = closingTag

	if closeErr := readBuffer.CloseContext("BACnetEventParameterChangeOfLifeSavety"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetEventParameterChangeOfLifeSavety")
	}

	return m, nil
}

func (m *_BACnetEventParameterChangeOfLifeSavety) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetEventParameterChangeOfLifeSavety) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetEventParameterChangeOfLifeSavety"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetEventParameterChangeOfLifeSavety")
		}

		if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'openingTag' field")
		}

		if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "timeDelay", m.GetTimeDelay(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'timeDelay' field")
		}

		if err := WriteSimpleField[BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues](ctx, "listOfLifeSavetyAlarmValues", m.GetListOfLifeSavetyAlarmValues(), WriteComplex[BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'listOfLifeSavetyAlarmValues' field")
		}

		if err := WriteSimpleField[BACnetEventParameterChangeOfLifeSavetyListOfAlarmValues](ctx, "listOfAlarmValues", m.GetListOfAlarmValues(), WriteComplex[BACnetEventParameterChangeOfLifeSavetyListOfAlarmValues](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'listOfAlarmValues' field")
		}

		if err := WriteSimpleField[BACnetDeviceObjectPropertyReferenceEnclosed](ctx, "modePropertyReference", m.GetModePropertyReference(), WriteComplex[BACnetDeviceObjectPropertyReferenceEnclosed](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'modePropertyReference' field")
		}

		if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'closingTag' field")
		}

		if popErr := writeBuffer.PopContext("BACnetEventParameterChangeOfLifeSavety"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetEventParameterChangeOfLifeSavety")
		}
		return nil
	}
	return m.BACnetEventParameterContract.(*_BACnetEventParameter).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetEventParameterChangeOfLifeSavety) IsBACnetEventParameterChangeOfLifeSavety() {}

func (m *_BACnetEventParameterChangeOfLifeSavety) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
