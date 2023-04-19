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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataAccessDoorFaultValues is the corresponding interface of BACnetConstructedDataAccessDoorFaultValues
type BACnetConstructedDataAccessDoorFaultValues interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetFaultValues returns FaultValues (property field)
	GetFaultValues() []BACnetDoorAlarmStateTagged
}

// BACnetConstructedDataAccessDoorFaultValuesExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataAccessDoorFaultValues.
// This is useful for switch cases.
type BACnetConstructedDataAccessDoorFaultValuesExactly interface {
	BACnetConstructedDataAccessDoorFaultValues
	isBACnetConstructedDataAccessDoorFaultValues() bool
}

// _BACnetConstructedDataAccessDoorFaultValues is the data-structure of this message
type _BACnetConstructedDataAccessDoorFaultValues struct {
	*_BACnetConstructedData
	FaultValues []BACnetDoorAlarmStateTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAccessDoorFaultValues) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_ACCESS_DOOR
}

func (m *_BACnetConstructedDataAccessDoorFaultValues) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_FAULT_VALUES
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAccessDoorFaultValues) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataAccessDoorFaultValues) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataAccessDoorFaultValues) GetFaultValues() []BACnetDoorAlarmStateTagged {
	return m.FaultValues
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataAccessDoorFaultValues factory function for _BACnetConstructedDataAccessDoorFaultValues
func NewBACnetConstructedDataAccessDoorFaultValues(faultValues []BACnetDoorAlarmStateTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAccessDoorFaultValues {
	_result := &_BACnetConstructedDataAccessDoorFaultValues{
		FaultValues:            faultValues,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAccessDoorFaultValues(structType interface{}) BACnetConstructedDataAccessDoorFaultValues {
	if casted, ok := structType.(BACnetConstructedDataAccessDoorFaultValues); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAccessDoorFaultValues); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAccessDoorFaultValues) GetTypeName() string {
	return "BACnetConstructedDataAccessDoorFaultValues"
}

func (m *_BACnetConstructedDataAccessDoorFaultValues) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Array field
	if len(m.FaultValues) > 0 {
		for _, element := range m.FaultValues {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataAccessDoorFaultValues) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataAccessDoorFaultValuesParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataAccessDoorFaultValues, error) {
	return BACnetConstructedDataAccessDoorFaultValuesParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataAccessDoorFaultValuesParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataAccessDoorFaultValues, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAccessDoorFaultValues"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAccessDoorFaultValues")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Array field (faultValues)
	if pullErr := readBuffer.PullContext("faultValues", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for faultValues")
	}
	// Terminated array
	var faultValues []BACnetDoorAlarmStateTagged
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetDoorAlarmStateTaggedParseWithBuffer(ctx, readBuffer, uint8(0), TagClass_APPLICATION_TAGS)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'faultValues' field of BACnetConstructedDataAccessDoorFaultValues")
			}
			faultValues = append(faultValues, _item.(BACnetDoorAlarmStateTagged))
		}
	}
	if closeErr := readBuffer.CloseContext("faultValues", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for faultValues")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAccessDoorFaultValues"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAccessDoorFaultValues")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataAccessDoorFaultValues{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		FaultValues: faultValues,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataAccessDoorFaultValues) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataAccessDoorFaultValues) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAccessDoorFaultValues"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAccessDoorFaultValues")
		}

		// Array Field (faultValues)
		if pushErr := writeBuffer.PushContext("faultValues", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for faultValues")
		}
		for _curItem, _element := range m.GetFaultValues() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetFaultValues()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'faultValues' field")
			}
		}
		if popErr := writeBuffer.PopContext("faultValues", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for faultValues")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAccessDoorFaultValues"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAccessDoorFaultValues")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAccessDoorFaultValues) isBACnetConstructedDataAccessDoorFaultValues() bool {
	return true
}

func (m *_BACnetConstructedDataAccessDoorFaultValues) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
