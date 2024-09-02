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

// BACnetFaultParameterFaultStateListOfFaultValues is the corresponding interface of BACnetFaultParameterFaultStateListOfFaultValues
type BACnetFaultParameterFaultStateListOfFaultValues interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetListIfFaultValues returns ListIfFaultValues (property field)
	GetListIfFaultValues() []BACnetPropertyStates
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
	// IsBACnetFaultParameterFaultStateListOfFaultValues is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetFaultParameterFaultStateListOfFaultValues()
}

// _BACnetFaultParameterFaultStateListOfFaultValues is the data-structure of this message
type _BACnetFaultParameterFaultStateListOfFaultValues struct {
	OpeningTag        BACnetOpeningTag
	ListIfFaultValues []BACnetPropertyStates
	ClosingTag        BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

var _ BACnetFaultParameterFaultStateListOfFaultValues = (*_BACnetFaultParameterFaultStateListOfFaultValues)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetFaultParameterFaultStateListOfFaultValues) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetFaultParameterFaultStateListOfFaultValues) GetListIfFaultValues() []BACnetPropertyStates {
	return m.ListIfFaultValues
}

func (m *_BACnetFaultParameterFaultStateListOfFaultValues) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetFaultParameterFaultStateListOfFaultValues factory function for _BACnetFaultParameterFaultStateListOfFaultValues
func NewBACnetFaultParameterFaultStateListOfFaultValues(openingTag BACnetOpeningTag, listIfFaultValues []BACnetPropertyStates, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetFaultParameterFaultStateListOfFaultValues {
	return &_BACnetFaultParameterFaultStateListOfFaultValues{OpeningTag: openingTag, ListIfFaultValues: listIfFaultValues, ClosingTag: closingTag, TagNumber: tagNumber}
}

// Deprecated: use the interface for direct cast
func CastBACnetFaultParameterFaultStateListOfFaultValues(structType any) BACnetFaultParameterFaultStateListOfFaultValues {
	if casted, ok := structType.(BACnetFaultParameterFaultStateListOfFaultValues); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetFaultParameterFaultStateListOfFaultValues); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetFaultParameterFaultStateListOfFaultValues) GetTypeName() string {
	return "BACnetFaultParameterFaultStateListOfFaultValues"
}

func (m *_BACnetFaultParameterFaultStateListOfFaultValues) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Array field
	if len(m.ListIfFaultValues) > 0 {
		for _, element := range m.ListIfFaultValues {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetFaultParameterFaultStateListOfFaultValues) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetFaultParameterFaultStateListOfFaultValuesParse(ctx context.Context, theBytes []byte, tagNumber uint8) (BACnetFaultParameterFaultStateListOfFaultValues, error) {
	return BACnetFaultParameterFaultStateListOfFaultValuesParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetFaultParameterFaultStateListOfFaultValuesParseWithBufferProducer(tagNumber uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetFaultParameterFaultStateListOfFaultValues, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetFaultParameterFaultStateListOfFaultValues, error) {
		return BACnetFaultParameterFaultStateListOfFaultValuesParseWithBuffer(ctx, readBuffer, tagNumber)
	}
}

func BACnetFaultParameterFaultStateListOfFaultValuesParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetFaultParameterFaultStateListOfFaultValues, error) {
	v, err := (&_BACnetFaultParameterFaultStateListOfFaultValues{TagNumber: tagNumber}).parse(ctx, readBuffer, tagNumber)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *_BACnetFaultParameterFaultStateListOfFaultValues) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (__bACnetFaultParameterFaultStateListOfFaultValues BACnetFaultParameterFaultStateListOfFaultValues, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetFaultParameterFaultStateListOfFaultValues"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetFaultParameterFaultStateListOfFaultValues")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}
	m.OpeningTag = openingTag

	listIfFaultValues, err := ReadTerminatedArrayField[BACnetPropertyStates](ctx, "listIfFaultValues", ReadComplex[BACnetPropertyStates](BACnetPropertyStatesParseWithBuffer, readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'listIfFaultValues' field"))
	}
	m.ListIfFaultValues = listIfFaultValues

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	m.ClosingTag = closingTag

	if closeErr := readBuffer.CloseContext("BACnetFaultParameterFaultStateListOfFaultValues"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetFaultParameterFaultStateListOfFaultValues")
	}

	return m, nil
}

func (m *_BACnetFaultParameterFaultStateListOfFaultValues) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetFaultParameterFaultStateListOfFaultValues) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetFaultParameterFaultStateListOfFaultValues"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetFaultParameterFaultStateListOfFaultValues")
	}

	if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'openingTag' field")
	}

	if err := WriteComplexTypeArrayField(ctx, "listIfFaultValues", m.GetListIfFaultValues(), writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'listIfFaultValues' field")
	}

	if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetFaultParameterFaultStateListOfFaultValues"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetFaultParameterFaultStateListOfFaultValues")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetFaultParameterFaultStateListOfFaultValues) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_BACnetFaultParameterFaultStateListOfFaultValues) IsBACnetFaultParameterFaultStateListOfFaultValues() {
}

func (m *_BACnetFaultParameterFaultStateListOfFaultValues) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
