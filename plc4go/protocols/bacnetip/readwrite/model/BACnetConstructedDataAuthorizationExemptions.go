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

// BACnetConstructedDataAuthorizationExemptions is the corresponding interface of BACnetConstructedDataAuthorizationExemptions
type BACnetConstructedDataAuthorizationExemptions interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetAuthorizationExemption returns AuthorizationExemption (property field)
	GetAuthorizationExemption() []BACnetAuthorizationExemptionTagged
}

// BACnetConstructedDataAuthorizationExemptionsExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataAuthorizationExemptions.
// This is useful for switch cases.
type BACnetConstructedDataAuthorizationExemptionsExactly interface {
	BACnetConstructedDataAuthorizationExemptions
	isBACnetConstructedDataAuthorizationExemptions() bool
}

// _BACnetConstructedDataAuthorizationExemptions is the data-structure of this message
type _BACnetConstructedDataAuthorizationExemptions struct {
	*_BACnetConstructedData
	AuthorizationExemption []BACnetAuthorizationExemptionTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAuthorizationExemptions) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataAuthorizationExemptions) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_AUTHORIZATION_EXEMPTIONS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAuthorizationExemptions) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataAuthorizationExemptions) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataAuthorizationExemptions) GetAuthorizationExemption() []BACnetAuthorizationExemptionTagged {
	return m.AuthorizationExemption
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataAuthorizationExemptions factory function for _BACnetConstructedDataAuthorizationExemptions
func NewBACnetConstructedDataAuthorizationExemptions(authorizationExemption []BACnetAuthorizationExemptionTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAuthorizationExemptions {
	_result := &_BACnetConstructedDataAuthorizationExemptions{
		AuthorizationExemption: authorizationExemption,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAuthorizationExemptions(structType interface{}) BACnetConstructedDataAuthorizationExemptions {
	if casted, ok := structType.(BACnetConstructedDataAuthorizationExemptions); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAuthorizationExemptions); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAuthorizationExemptions) GetTypeName() string {
	return "BACnetConstructedDataAuthorizationExemptions"
}

func (m *_BACnetConstructedDataAuthorizationExemptions) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Array field
	if len(m.AuthorizationExemption) > 0 {
		for _, element := range m.AuthorizationExemption {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataAuthorizationExemptions) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataAuthorizationExemptionsParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataAuthorizationExemptions, error) {
	return BACnetConstructedDataAuthorizationExemptionsParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataAuthorizationExemptionsParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataAuthorizationExemptions, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAuthorizationExemptions"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAuthorizationExemptions")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Array field (authorizationExemption)
	if pullErr := readBuffer.PullContext("authorizationExemption", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for authorizationExemption")
	}
	// Terminated array
	var authorizationExemption []BACnetAuthorizationExemptionTagged
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetAuthorizationExemptionTaggedParseWithBuffer(ctx, readBuffer, uint8(0), TagClass_APPLICATION_TAGS)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'authorizationExemption' field of BACnetConstructedDataAuthorizationExemptions")
			}
			authorizationExemption = append(authorizationExemption, _item.(BACnetAuthorizationExemptionTagged))
		}
	}
	if closeErr := readBuffer.CloseContext("authorizationExemption", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for authorizationExemption")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAuthorizationExemptions"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAuthorizationExemptions")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataAuthorizationExemptions{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		AuthorizationExemption: authorizationExemption,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataAuthorizationExemptions) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataAuthorizationExemptions) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAuthorizationExemptions"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAuthorizationExemptions")
		}

		// Array Field (authorizationExemption)
		if pushErr := writeBuffer.PushContext("authorizationExemption", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for authorizationExemption")
		}
		for _curItem, _element := range m.GetAuthorizationExemption() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetAuthorizationExemption()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'authorizationExemption' field")
			}
		}
		if popErr := writeBuffer.PopContext("authorizationExemption", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for authorizationExemption")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAuthorizationExemptions"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAuthorizationExemptions")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAuthorizationExemptions) isBACnetConstructedDataAuthorizationExemptions() bool {
	return true
}

func (m *_BACnetConstructedDataAuthorizationExemptions) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
