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

// SetTriggeringRequest is the corresponding interface of SetTriggeringRequest
type SetTriggeringRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetRequestHeader returns RequestHeader (property field)
	GetRequestHeader() ExtensionObjectDefinition
	// GetSubscriptionId returns SubscriptionId (property field)
	GetSubscriptionId() uint32
	// GetTriggeringItemId returns TriggeringItemId (property field)
	GetTriggeringItemId() uint32
	// GetNoOfLinksToAdd returns NoOfLinksToAdd (property field)
	GetNoOfLinksToAdd() int32
	// GetLinksToAdd returns LinksToAdd (property field)
	GetLinksToAdd() []uint32
	// GetNoOfLinksToRemove returns NoOfLinksToRemove (property field)
	GetNoOfLinksToRemove() int32
	// GetLinksToRemove returns LinksToRemove (property field)
	GetLinksToRemove() []uint32
}

// SetTriggeringRequestExactly can be used when we want exactly this type and not a type which fulfills SetTriggeringRequest.
// This is useful for switch cases.
type SetTriggeringRequestExactly interface {
	SetTriggeringRequest
	isSetTriggeringRequest() bool
}

// _SetTriggeringRequest is the data-structure of this message
type _SetTriggeringRequest struct {
	ExtensionObjectDefinitionContract
	RequestHeader     ExtensionObjectDefinition
	SubscriptionId    uint32
	TriggeringItemId  uint32
	NoOfLinksToAdd    int32
	LinksToAdd        []uint32
	NoOfLinksToRemove int32
	LinksToRemove     []uint32
}

var _ SetTriggeringRequest = (*_SetTriggeringRequest)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_SetTriggeringRequest)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SetTriggeringRequest) GetIdentifier() string {
	return "775"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SetTriggeringRequest) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SetTriggeringRequest) GetRequestHeader() ExtensionObjectDefinition {
	return m.RequestHeader
}

func (m *_SetTriggeringRequest) GetSubscriptionId() uint32 {
	return m.SubscriptionId
}

func (m *_SetTriggeringRequest) GetTriggeringItemId() uint32 {
	return m.TriggeringItemId
}

func (m *_SetTriggeringRequest) GetNoOfLinksToAdd() int32 {
	return m.NoOfLinksToAdd
}

func (m *_SetTriggeringRequest) GetLinksToAdd() []uint32 {
	return m.LinksToAdd
}

func (m *_SetTriggeringRequest) GetNoOfLinksToRemove() int32 {
	return m.NoOfLinksToRemove
}

func (m *_SetTriggeringRequest) GetLinksToRemove() []uint32 {
	return m.LinksToRemove
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewSetTriggeringRequest factory function for _SetTriggeringRequest
func NewSetTriggeringRequest(requestHeader ExtensionObjectDefinition, subscriptionId uint32, triggeringItemId uint32, noOfLinksToAdd int32, linksToAdd []uint32, noOfLinksToRemove int32, linksToRemove []uint32) *_SetTriggeringRequest {
	_result := &_SetTriggeringRequest{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		RequestHeader:                     requestHeader,
		SubscriptionId:                    subscriptionId,
		TriggeringItemId:                  triggeringItemId,
		NoOfLinksToAdd:                    noOfLinksToAdd,
		LinksToAdd:                        linksToAdd,
		NoOfLinksToRemove:                 noOfLinksToRemove,
		LinksToRemove:                     linksToRemove,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastSetTriggeringRequest(structType any) SetTriggeringRequest {
	if casted, ok := structType.(SetTriggeringRequest); ok {
		return casted
	}
	if casted, ok := structType.(*SetTriggeringRequest); ok {
		return *casted
	}
	return nil
}

func (m *_SetTriggeringRequest) GetTypeName() string {
	return "SetTriggeringRequest"
}

func (m *_SetTriggeringRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (requestHeader)
	lengthInBits += m.RequestHeader.GetLengthInBits(ctx)

	// Simple field (subscriptionId)
	lengthInBits += 32

	// Simple field (triggeringItemId)
	lengthInBits += 32

	// Simple field (noOfLinksToAdd)
	lengthInBits += 32

	// Array field
	if len(m.LinksToAdd) > 0 {
		lengthInBits += 32 * uint16(len(m.LinksToAdd))
	}

	// Simple field (noOfLinksToRemove)
	lengthInBits += 32

	// Array field
	if len(m.LinksToRemove) > 0 {
		lengthInBits += 32 * uint16(len(m.LinksToRemove))
	}

	return lengthInBits
}

func (m *_SetTriggeringRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SetTriggeringRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__setTriggeringRequest SetTriggeringRequest, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SetTriggeringRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SetTriggeringRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	requestHeader, err := ReadSimpleField[ExtensionObjectDefinition](ctx, "requestHeader", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("391")), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestHeader' field"))
	}
	m.RequestHeader = requestHeader

	subscriptionId, err := ReadSimpleField(ctx, "subscriptionId", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'subscriptionId' field"))
	}
	m.SubscriptionId = subscriptionId

	triggeringItemId, err := ReadSimpleField(ctx, "triggeringItemId", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'triggeringItemId' field"))
	}
	m.TriggeringItemId = triggeringItemId

	noOfLinksToAdd, err := ReadSimpleField(ctx, "noOfLinksToAdd", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfLinksToAdd' field"))
	}
	m.NoOfLinksToAdd = noOfLinksToAdd

	linksToAdd, err := ReadCountArrayField[uint32](ctx, "linksToAdd", ReadUnsignedInt(readBuffer, uint8(32)), uint64(noOfLinksToAdd))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'linksToAdd' field"))
	}
	m.LinksToAdd = linksToAdd

	noOfLinksToRemove, err := ReadSimpleField(ctx, "noOfLinksToRemove", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfLinksToRemove' field"))
	}
	m.NoOfLinksToRemove = noOfLinksToRemove

	linksToRemove, err := ReadCountArrayField[uint32](ctx, "linksToRemove", ReadUnsignedInt(readBuffer, uint8(32)), uint64(noOfLinksToRemove))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'linksToRemove' field"))
	}
	m.LinksToRemove = linksToRemove

	if closeErr := readBuffer.CloseContext("SetTriggeringRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SetTriggeringRequest")
	}

	return m, nil
}

func (m *_SetTriggeringRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SetTriggeringRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SetTriggeringRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SetTriggeringRequest")
		}

		if err := WriteSimpleField[ExtensionObjectDefinition](ctx, "requestHeader", m.GetRequestHeader(), WriteComplex[ExtensionObjectDefinition](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'requestHeader' field")
		}

		if err := WriteSimpleField[uint32](ctx, "subscriptionId", m.GetSubscriptionId(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'subscriptionId' field")
		}

		if err := WriteSimpleField[uint32](ctx, "triggeringItemId", m.GetTriggeringItemId(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'triggeringItemId' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfLinksToAdd", m.GetNoOfLinksToAdd(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfLinksToAdd' field")
		}

		if err := WriteSimpleTypeArrayField(ctx, "linksToAdd", m.GetLinksToAdd(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'linksToAdd' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfLinksToRemove", m.GetNoOfLinksToRemove(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfLinksToRemove' field")
		}

		if err := WriteSimpleTypeArrayField(ctx, "linksToRemove", m.GetLinksToRemove(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'linksToRemove' field")
		}

		if popErr := writeBuffer.PopContext("SetTriggeringRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SetTriggeringRequest")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SetTriggeringRequest) isSetTriggeringRequest() bool {
	return true
}

func (m *_SetTriggeringRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
