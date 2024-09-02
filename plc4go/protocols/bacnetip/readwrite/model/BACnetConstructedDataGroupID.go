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

// BACnetConstructedDataGroupID is the corresponding interface of BACnetConstructedDataGroupID
type BACnetConstructedDataGroupID interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetGroupId returns GroupId (property field)
	GetGroupId() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataGroupIDExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataGroupID.
// This is useful for switch cases.
type BACnetConstructedDataGroupIDExactly interface {
	BACnetConstructedDataGroupID
	isBACnetConstructedDataGroupID() bool
}

// _BACnetConstructedDataGroupID is the data-structure of this message
type _BACnetConstructedDataGroupID struct {
	BACnetConstructedDataContract
	GroupId BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataGroupID = (*_BACnetConstructedDataGroupID)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataGroupID)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataGroupID) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataGroupID) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_GROUP_ID
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataGroupID) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataGroupID) GetGroupId() BACnetApplicationTagUnsignedInteger {
	return m.GroupId
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataGroupID) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetGroupId())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataGroupID factory function for _BACnetConstructedDataGroupID
func NewBACnetConstructedDataGroupID(groupId BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataGroupID {
	_result := &_BACnetConstructedDataGroupID{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		GroupId:                       groupId,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataGroupID(structType any) BACnetConstructedDataGroupID {
	if casted, ok := structType.(BACnetConstructedDataGroupID); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataGroupID); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataGroupID) GetTypeName() string {
	return "BACnetConstructedDataGroupID"
}

func (m *_BACnetConstructedDataGroupID) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (groupId)
	lengthInBits += m.GroupId.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataGroupID) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataGroupID) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataGroupID BACnetConstructedDataGroupID, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataGroupID"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataGroupID")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	groupId, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "groupId", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'groupId' field"))
	}
	m.GroupId = groupId

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), groupId)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataGroupID"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataGroupID")
	}

	return m, nil
}

func (m *_BACnetConstructedDataGroupID) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataGroupID) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataGroupID"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataGroupID")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "groupId", m.GetGroupId(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'groupId' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataGroupID"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataGroupID")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataGroupID) isBACnetConstructedDataGroupID() bool {
	return true
}

func (m *_BACnetConstructedDataGroupID) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
