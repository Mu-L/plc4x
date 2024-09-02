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

// CreateSubscriptionRequest is the corresponding interface of CreateSubscriptionRequest
type CreateSubscriptionRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetRequestHeader returns RequestHeader (property field)
	GetRequestHeader() ExtensionObjectDefinition
	// GetRequestedPublishingInterval returns RequestedPublishingInterval (property field)
	GetRequestedPublishingInterval() float64
	// GetRequestedLifetimeCount returns RequestedLifetimeCount (property field)
	GetRequestedLifetimeCount() uint32
	// GetRequestedMaxKeepAliveCount returns RequestedMaxKeepAliveCount (property field)
	GetRequestedMaxKeepAliveCount() uint32
	// GetMaxNotificationsPerPublish returns MaxNotificationsPerPublish (property field)
	GetMaxNotificationsPerPublish() uint32
	// GetPublishingEnabled returns PublishingEnabled (property field)
	GetPublishingEnabled() bool
	// GetPriority returns Priority (property field)
	GetPriority() uint8
	// IsCreateSubscriptionRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsCreateSubscriptionRequest()
}

// _CreateSubscriptionRequest is the data-structure of this message
type _CreateSubscriptionRequest struct {
	ExtensionObjectDefinitionContract
	RequestHeader               ExtensionObjectDefinition
	RequestedPublishingInterval float64
	RequestedLifetimeCount      uint32
	RequestedMaxKeepAliveCount  uint32
	MaxNotificationsPerPublish  uint32
	PublishingEnabled           bool
	Priority                    uint8
	// Reserved Fields
	reservedField0 *uint8
}

var _ CreateSubscriptionRequest = (*_CreateSubscriptionRequest)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_CreateSubscriptionRequest)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_CreateSubscriptionRequest) GetIdentifier() string {
	return "787"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CreateSubscriptionRequest) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CreateSubscriptionRequest) GetRequestHeader() ExtensionObjectDefinition {
	return m.RequestHeader
}

func (m *_CreateSubscriptionRequest) GetRequestedPublishingInterval() float64 {
	return m.RequestedPublishingInterval
}

func (m *_CreateSubscriptionRequest) GetRequestedLifetimeCount() uint32 {
	return m.RequestedLifetimeCount
}

func (m *_CreateSubscriptionRequest) GetRequestedMaxKeepAliveCount() uint32 {
	return m.RequestedMaxKeepAliveCount
}

func (m *_CreateSubscriptionRequest) GetMaxNotificationsPerPublish() uint32 {
	return m.MaxNotificationsPerPublish
}

func (m *_CreateSubscriptionRequest) GetPublishingEnabled() bool {
	return m.PublishingEnabled
}

func (m *_CreateSubscriptionRequest) GetPriority() uint8 {
	return m.Priority
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCreateSubscriptionRequest factory function for _CreateSubscriptionRequest
func NewCreateSubscriptionRequest(requestHeader ExtensionObjectDefinition, requestedPublishingInterval float64, requestedLifetimeCount uint32, requestedMaxKeepAliveCount uint32, maxNotificationsPerPublish uint32, publishingEnabled bool, priority uint8) *_CreateSubscriptionRequest {
	_result := &_CreateSubscriptionRequest{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		RequestHeader:                     requestHeader,
		RequestedPublishingInterval:       requestedPublishingInterval,
		RequestedLifetimeCount:            requestedLifetimeCount,
		RequestedMaxKeepAliveCount:        requestedMaxKeepAliveCount,
		MaxNotificationsPerPublish:        maxNotificationsPerPublish,
		PublishingEnabled:                 publishingEnabled,
		Priority:                          priority,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastCreateSubscriptionRequest(structType any) CreateSubscriptionRequest {
	if casted, ok := structType.(CreateSubscriptionRequest); ok {
		return casted
	}
	if casted, ok := structType.(*CreateSubscriptionRequest); ok {
		return *casted
	}
	return nil
}

func (m *_CreateSubscriptionRequest) GetTypeName() string {
	return "CreateSubscriptionRequest"
}

func (m *_CreateSubscriptionRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (requestHeader)
	lengthInBits += m.RequestHeader.GetLengthInBits(ctx)

	// Simple field (requestedPublishingInterval)
	lengthInBits += 64

	// Simple field (requestedLifetimeCount)
	lengthInBits += 32

	// Simple field (requestedMaxKeepAliveCount)
	lengthInBits += 32

	// Simple field (maxNotificationsPerPublish)
	lengthInBits += 32

	// Reserved Field (reserved)
	lengthInBits += 7

	// Simple field (publishingEnabled)
	lengthInBits += 1

	// Simple field (priority)
	lengthInBits += 8

	return lengthInBits
}

func (m *_CreateSubscriptionRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_CreateSubscriptionRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__createSubscriptionRequest CreateSubscriptionRequest, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CreateSubscriptionRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CreateSubscriptionRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	requestHeader, err := ReadSimpleField[ExtensionObjectDefinition](ctx, "requestHeader", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("391")), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestHeader' field"))
	}
	m.RequestHeader = requestHeader

	requestedPublishingInterval, err := ReadSimpleField(ctx, "requestedPublishingInterval", ReadDouble(readBuffer, uint8(64)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestedPublishingInterval' field"))
	}
	m.RequestedPublishingInterval = requestedPublishingInterval

	requestedLifetimeCount, err := ReadSimpleField(ctx, "requestedLifetimeCount", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestedLifetimeCount' field"))
	}
	m.RequestedLifetimeCount = requestedLifetimeCount

	requestedMaxKeepAliveCount, err := ReadSimpleField(ctx, "requestedMaxKeepAliveCount", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestedMaxKeepAliveCount' field"))
	}
	m.RequestedMaxKeepAliveCount = requestedMaxKeepAliveCount

	maxNotificationsPerPublish, err := ReadSimpleField(ctx, "maxNotificationsPerPublish", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'maxNotificationsPerPublish' field"))
	}
	m.MaxNotificationsPerPublish = maxNotificationsPerPublish

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(7)), uint8(0x00))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	publishingEnabled, err := ReadSimpleField(ctx, "publishingEnabled", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'publishingEnabled' field"))
	}
	m.PublishingEnabled = publishingEnabled

	priority, err := ReadSimpleField(ctx, "priority", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'priority' field"))
	}
	m.Priority = priority

	if closeErr := readBuffer.CloseContext("CreateSubscriptionRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CreateSubscriptionRequest")
	}

	return m, nil
}

func (m *_CreateSubscriptionRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CreateSubscriptionRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CreateSubscriptionRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CreateSubscriptionRequest")
		}

		if err := WriteSimpleField[ExtensionObjectDefinition](ctx, "requestHeader", m.GetRequestHeader(), WriteComplex[ExtensionObjectDefinition](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'requestHeader' field")
		}

		if err := WriteSimpleField[float64](ctx, "requestedPublishingInterval", m.GetRequestedPublishingInterval(), WriteDouble(writeBuffer, 64)); err != nil {
			return errors.Wrap(err, "Error serializing 'requestedPublishingInterval' field")
		}

		if err := WriteSimpleField[uint32](ctx, "requestedLifetimeCount", m.GetRequestedLifetimeCount(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'requestedLifetimeCount' field")
		}

		if err := WriteSimpleField[uint32](ctx, "requestedMaxKeepAliveCount", m.GetRequestedMaxKeepAliveCount(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'requestedMaxKeepAliveCount' field")
		}

		if err := WriteSimpleField[uint32](ctx, "maxNotificationsPerPublish", m.GetMaxNotificationsPerPublish(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'maxNotificationsPerPublish' field")
		}

		if err := WriteReservedField[uint8](ctx, "reserved", uint8(0x00), WriteUnsignedByte(writeBuffer, 7)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 1")
		}

		if err := WriteSimpleField[bool](ctx, "publishingEnabled", m.GetPublishingEnabled(), WriteBoolean(writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'publishingEnabled' field")
		}

		if err := WriteSimpleField[uint8](ctx, "priority", m.GetPriority(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'priority' field")
		}

		if popErr := writeBuffer.PopContext("CreateSubscriptionRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CreateSubscriptionRequest")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CreateSubscriptionRequest) IsCreateSubscriptionRequest() {}

func (m *_CreateSubscriptionRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
