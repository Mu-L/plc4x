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

// SubscriptionDiagnosticsDataType is the corresponding interface of SubscriptionDiagnosticsDataType
type SubscriptionDiagnosticsDataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetSessionId returns SessionId (property field)
	GetSessionId() NodeId
	// GetSubscriptionId returns SubscriptionId (property field)
	GetSubscriptionId() uint32
	// GetPriority returns Priority (property field)
	GetPriority() uint8
	// GetPublishingInterval returns PublishingInterval (property field)
	GetPublishingInterval() float64
	// GetMaxKeepAliveCount returns MaxKeepAliveCount (property field)
	GetMaxKeepAliveCount() uint32
	// GetMaxLifetimeCount returns MaxLifetimeCount (property field)
	GetMaxLifetimeCount() uint32
	// GetMaxNotificationsPerPublish returns MaxNotificationsPerPublish (property field)
	GetMaxNotificationsPerPublish() uint32
	// GetPublishingEnabled returns PublishingEnabled (property field)
	GetPublishingEnabled() bool
	// GetModifyCount returns ModifyCount (property field)
	GetModifyCount() uint32
	// GetEnableCount returns EnableCount (property field)
	GetEnableCount() uint32
	// GetDisableCount returns DisableCount (property field)
	GetDisableCount() uint32
	// GetRepublishRequestCount returns RepublishRequestCount (property field)
	GetRepublishRequestCount() uint32
	// GetRepublishMessageRequestCount returns RepublishMessageRequestCount (property field)
	GetRepublishMessageRequestCount() uint32
	// GetRepublishMessageCount returns RepublishMessageCount (property field)
	GetRepublishMessageCount() uint32
	// GetTransferRequestCount returns TransferRequestCount (property field)
	GetTransferRequestCount() uint32
	// GetTransferredToAltClientCount returns TransferredToAltClientCount (property field)
	GetTransferredToAltClientCount() uint32
	// GetTransferredToSameClientCount returns TransferredToSameClientCount (property field)
	GetTransferredToSameClientCount() uint32
	// GetPublishRequestCount returns PublishRequestCount (property field)
	GetPublishRequestCount() uint32
	// GetDataChangeNotificationsCount returns DataChangeNotificationsCount (property field)
	GetDataChangeNotificationsCount() uint32
	// GetEventNotificationsCount returns EventNotificationsCount (property field)
	GetEventNotificationsCount() uint32
	// GetNotificationsCount returns NotificationsCount (property field)
	GetNotificationsCount() uint32
	// GetLatePublishRequestCount returns LatePublishRequestCount (property field)
	GetLatePublishRequestCount() uint32
	// GetCurrentKeepAliveCount returns CurrentKeepAliveCount (property field)
	GetCurrentKeepAliveCount() uint32
	// GetCurrentLifetimeCount returns CurrentLifetimeCount (property field)
	GetCurrentLifetimeCount() uint32
	// GetUnacknowledgedMessageCount returns UnacknowledgedMessageCount (property field)
	GetUnacknowledgedMessageCount() uint32
	// GetDiscardedMessageCount returns DiscardedMessageCount (property field)
	GetDiscardedMessageCount() uint32
	// GetMonitoredItemCount returns MonitoredItemCount (property field)
	GetMonitoredItemCount() uint32
	// GetDisabledMonitoredItemCount returns DisabledMonitoredItemCount (property field)
	GetDisabledMonitoredItemCount() uint32
	// GetMonitoringQueueOverflowCount returns MonitoringQueueOverflowCount (property field)
	GetMonitoringQueueOverflowCount() uint32
	// GetNextSequenceNumber returns NextSequenceNumber (property field)
	GetNextSequenceNumber() uint32
	// GetEventQueueOverFlowCount returns EventQueueOverFlowCount (property field)
	GetEventQueueOverFlowCount() uint32
}

// SubscriptionDiagnosticsDataTypeExactly can be used when we want exactly this type and not a type which fulfills SubscriptionDiagnosticsDataType.
// This is useful for switch cases.
type SubscriptionDiagnosticsDataTypeExactly interface {
	SubscriptionDiagnosticsDataType
	isSubscriptionDiagnosticsDataType() bool
}

// _SubscriptionDiagnosticsDataType is the data-structure of this message
type _SubscriptionDiagnosticsDataType struct {
	ExtensionObjectDefinitionContract
	SessionId                    NodeId
	SubscriptionId               uint32
	Priority                     uint8
	PublishingInterval           float64
	MaxKeepAliveCount            uint32
	MaxLifetimeCount             uint32
	MaxNotificationsPerPublish   uint32
	PublishingEnabled            bool
	ModifyCount                  uint32
	EnableCount                  uint32
	DisableCount                 uint32
	RepublishRequestCount        uint32
	RepublishMessageRequestCount uint32
	RepublishMessageCount        uint32
	TransferRequestCount         uint32
	TransferredToAltClientCount  uint32
	TransferredToSameClientCount uint32
	PublishRequestCount          uint32
	DataChangeNotificationsCount uint32
	EventNotificationsCount      uint32
	NotificationsCount           uint32
	LatePublishRequestCount      uint32
	CurrentKeepAliveCount        uint32
	CurrentLifetimeCount         uint32
	UnacknowledgedMessageCount   uint32
	DiscardedMessageCount        uint32
	MonitoredItemCount           uint32
	DisabledMonitoredItemCount   uint32
	MonitoringQueueOverflowCount uint32
	NextSequenceNumber           uint32
	EventQueueOverFlowCount      uint32
	// Reserved Fields
	reservedField0 *uint8
}

var _ SubscriptionDiagnosticsDataType = (*_SubscriptionDiagnosticsDataType)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_SubscriptionDiagnosticsDataType)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SubscriptionDiagnosticsDataType) GetIdentifier() string {
	return "876"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SubscriptionDiagnosticsDataType) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SubscriptionDiagnosticsDataType) GetSessionId() NodeId {
	return m.SessionId
}

func (m *_SubscriptionDiagnosticsDataType) GetSubscriptionId() uint32 {
	return m.SubscriptionId
}

func (m *_SubscriptionDiagnosticsDataType) GetPriority() uint8 {
	return m.Priority
}

func (m *_SubscriptionDiagnosticsDataType) GetPublishingInterval() float64 {
	return m.PublishingInterval
}

func (m *_SubscriptionDiagnosticsDataType) GetMaxKeepAliveCount() uint32 {
	return m.MaxKeepAliveCount
}

func (m *_SubscriptionDiagnosticsDataType) GetMaxLifetimeCount() uint32 {
	return m.MaxLifetimeCount
}

func (m *_SubscriptionDiagnosticsDataType) GetMaxNotificationsPerPublish() uint32 {
	return m.MaxNotificationsPerPublish
}

func (m *_SubscriptionDiagnosticsDataType) GetPublishingEnabled() bool {
	return m.PublishingEnabled
}

func (m *_SubscriptionDiagnosticsDataType) GetModifyCount() uint32 {
	return m.ModifyCount
}

func (m *_SubscriptionDiagnosticsDataType) GetEnableCount() uint32 {
	return m.EnableCount
}

func (m *_SubscriptionDiagnosticsDataType) GetDisableCount() uint32 {
	return m.DisableCount
}

func (m *_SubscriptionDiagnosticsDataType) GetRepublishRequestCount() uint32 {
	return m.RepublishRequestCount
}

func (m *_SubscriptionDiagnosticsDataType) GetRepublishMessageRequestCount() uint32 {
	return m.RepublishMessageRequestCount
}

func (m *_SubscriptionDiagnosticsDataType) GetRepublishMessageCount() uint32 {
	return m.RepublishMessageCount
}

func (m *_SubscriptionDiagnosticsDataType) GetTransferRequestCount() uint32 {
	return m.TransferRequestCount
}

func (m *_SubscriptionDiagnosticsDataType) GetTransferredToAltClientCount() uint32 {
	return m.TransferredToAltClientCount
}

func (m *_SubscriptionDiagnosticsDataType) GetTransferredToSameClientCount() uint32 {
	return m.TransferredToSameClientCount
}

func (m *_SubscriptionDiagnosticsDataType) GetPublishRequestCount() uint32 {
	return m.PublishRequestCount
}

func (m *_SubscriptionDiagnosticsDataType) GetDataChangeNotificationsCount() uint32 {
	return m.DataChangeNotificationsCount
}

func (m *_SubscriptionDiagnosticsDataType) GetEventNotificationsCount() uint32 {
	return m.EventNotificationsCount
}

func (m *_SubscriptionDiagnosticsDataType) GetNotificationsCount() uint32 {
	return m.NotificationsCount
}

func (m *_SubscriptionDiagnosticsDataType) GetLatePublishRequestCount() uint32 {
	return m.LatePublishRequestCount
}

func (m *_SubscriptionDiagnosticsDataType) GetCurrentKeepAliveCount() uint32 {
	return m.CurrentKeepAliveCount
}

func (m *_SubscriptionDiagnosticsDataType) GetCurrentLifetimeCount() uint32 {
	return m.CurrentLifetimeCount
}

func (m *_SubscriptionDiagnosticsDataType) GetUnacknowledgedMessageCount() uint32 {
	return m.UnacknowledgedMessageCount
}

func (m *_SubscriptionDiagnosticsDataType) GetDiscardedMessageCount() uint32 {
	return m.DiscardedMessageCount
}

func (m *_SubscriptionDiagnosticsDataType) GetMonitoredItemCount() uint32 {
	return m.MonitoredItemCount
}

func (m *_SubscriptionDiagnosticsDataType) GetDisabledMonitoredItemCount() uint32 {
	return m.DisabledMonitoredItemCount
}

func (m *_SubscriptionDiagnosticsDataType) GetMonitoringQueueOverflowCount() uint32 {
	return m.MonitoringQueueOverflowCount
}

func (m *_SubscriptionDiagnosticsDataType) GetNextSequenceNumber() uint32 {
	return m.NextSequenceNumber
}

func (m *_SubscriptionDiagnosticsDataType) GetEventQueueOverFlowCount() uint32 {
	return m.EventQueueOverFlowCount
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewSubscriptionDiagnosticsDataType factory function for _SubscriptionDiagnosticsDataType
func NewSubscriptionDiagnosticsDataType(sessionId NodeId, subscriptionId uint32, priority uint8, publishingInterval float64, maxKeepAliveCount uint32, maxLifetimeCount uint32, maxNotificationsPerPublish uint32, publishingEnabled bool, modifyCount uint32, enableCount uint32, disableCount uint32, republishRequestCount uint32, republishMessageRequestCount uint32, republishMessageCount uint32, transferRequestCount uint32, transferredToAltClientCount uint32, transferredToSameClientCount uint32, publishRequestCount uint32, dataChangeNotificationsCount uint32, eventNotificationsCount uint32, notificationsCount uint32, latePublishRequestCount uint32, currentKeepAliveCount uint32, currentLifetimeCount uint32, unacknowledgedMessageCount uint32, discardedMessageCount uint32, monitoredItemCount uint32, disabledMonitoredItemCount uint32, monitoringQueueOverflowCount uint32, nextSequenceNumber uint32, eventQueueOverFlowCount uint32) *_SubscriptionDiagnosticsDataType {
	_result := &_SubscriptionDiagnosticsDataType{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		SessionId:                         sessionId,
		SubscriptionId:                    subscriptionId,
		Priority:                          priority,
		PublishingInterval:                publishingInterval,
		MaxKeepAliveCount:                 maxKeepAliveCount,
		MaxLifetimeCount:                  maxLifetimeCount,
		MaxNotificationsPerPublish:        maxNotificationsPerPublish,
		PublishingEnabled:                 publishingEnabled,
		ModifyCount:                       modifyCount,
		EnableCount:                       enableCount,
		DisableCount:                      disableCount,
		RepublishRequestCount:             republishRequestCount,
		RepublishMessageRequestCount:      republishMessageRequestCount,
		RepublishMessageCount:             republishMessageCount,
		TransferRequestCount:              transferRequestCount,
		TransferredToAltClientCount:       transferredToAltClientCount,
		TransferredToSameClientCount:      transferredToSameClientCount,
		PublishRequestCount:               publishRequestCount,
		DataChangeNotificationsCount:      dataChangeNotificationsCount,
		EventNotificationsCount:           eventNotificationsCount,
		NotificationsCount:                notificationsCount,
		LatePublishRequestCount:           latePublishRequestCount,
		CurrentKeepAliveCount:             currentKeepAliveCount,
		CurrentLifetimeCount:              currentLifetimeCount,
		UnacknowledgedMessageCount:        unacknowledgedMessageCount,
		DiscardedMessageCount:             discardedMessageCount,
		MonitoredItemCount:                monitoredItemCount,
		DisabledMonitoredItemCount:        disabledMonitoredItemCount,
		MonitoringQueueOverflowCount:      monitoringQueueOverflowCount,
		NextSequenceNumber:                nextSequenceNumber,
		EventQueueOverFlowCount:           eventQueueOverFlowCount,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastSubscriptionDiagnosticsDataType(structType any) SubscriptionDiagnosticsDataType {
	if casted, ok := structType.(SubscriptionDiagnosticsDataType); ok {
		return casted
	}
	if casted, ok := structType.(*SubscriptionDiagnosticsDataType); ok {
		return *casted
	}
	return nil
}

func (m *_SubscriptionDiagnosticsDataType) GetTypeName() string {
	return "SubscriptionDiagnosticsDataType"
}

func (m *_SubscriptionDiagnosticsDataType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (sessionId)
	lengthInBits += m.SessionId.GetLengthInBits(ctx)

	// Simple field (subscriptionId)
	lengthInBits += 32

	// Simple field (priority)
	lengthInBits += 8

	// Simple field (publishingInterval)
	lengthInBits += 64

	// Simple field (maxKeepAliveCount)
	lengthInBits += 32

	// Simple field (maxLifetimeCount)
	lengthInBits += 32

	// Simple field (maxNotificationsPerPublish)
	lengthInBits += 32

	// Reserved Field (reserved)
	lengthInBits += 7

	// Simple field (publishingEnabled)
	lengthInBits += 1

	// Simple field (modifyCount)
	lengthInBits += 32

	// Simple field (enableCount)
	lengthInBits += 32

	// Simple field (disableCount)
	lengthInBits += 32

	// Simple field (republishRequestCount)
	lengthInBits += 32

	// Simple field (republishMessageRequestCount)
	lengthInBits += 32

	// Simple field (republishMessageCount)
	lengthInBits += 32

	// Simple field (transferRequestCount)
	lengthInBits += 32

	// Simple field (transferredToAltClientCount)
	lengthInBits += 32

	// Simple field (transferredToSameClientCount)
	lengthInBits += 32

	// Simple field (publishRequestCount)
	lengthInBits += 32

	// Simple field (dataChangeNotificationsCount)
	lengthInBits += 32

	// Simple field (eventNotificationsCount)
	lengthInBits += 32

	// Simple field (notificationsCount)
	lengthInBits += 32

	// Simple field (latePublishRequestCount)
	lengthInBits += 32

	// Simple field (currentKeepAliveCount)
	lengthInBits += 32

	// Simple field (currentLifetimeCount)
	lengthInBits += 32

	// Simple field (unacknowledgedMessageCount)
	lengthInBits += 32

	// Simple field (discardedMessageCount)
	lengthInBits += 32

	// Simple field (monitoredItemCount)
	lengthInBits += 32

	// Simple field (disabledMonitoredItemCount)
	lengthInBits += 32

	// Simple field (monitoringQueueOverflowCount)
	lengthInBits += 32

	// Simple field (nextSequenceNumber)
	lengthInBits += 32

	// Simple field (eventQueueOverFlowCount)
	lengthInBits += 32

	return lengthInBits
}

func (m *_SubscriptionDiagnosticsDataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SubscriptionDiagnosticsDataType) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__subscriptionDiagnosticsDataType SubscriptionDiagnosticsDataType, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SubscriptionDiagnosticsDataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SubscriptionDiagnosticsDataType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	sessionId, err := ReadSimpleField[NodeId](ctx, "sessionId", ReadComplex[NodeId](NodeIdParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'sessionId' field"))
	}
	m.SessionId = sessionId

	subscriptionId, err := ReadSimpleField(ctx, "subscriptionId", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'subscriptionId' field"))
	}
	m.SubscriptionId = subscriptionId

	priority, err := ReadSimpleField(ctx, "priority", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'priority' field"))
	}
	m.Priority = priority

	publishingInterval, err := ReadSimpleField(ctx, "publishingInterval", ReadDouble(readBuffer, uint8(64)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'publishingInterval' field"))
	}
	m.PublishingInterval = publishingInterval

	maxKeepAliveCount, err := ReadSimpleField(ctx, "maxKeepAliveCount", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'maxKeepAliveCount' field"))
	}
	m.MaxKeepAliveCount = maxKeepAliveCount

	maxLifetimeCount, err := ReadSimpleField(ctx, "maxLifetimeCount", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'maxLifetimeCount' field"))
	}
	m.MaxLifetimeCount = maxLifetimeCount

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

	modifyCount, err := ReadSimpleField(ctx, "modifyCount", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'modifyCount' field"))
	}
	m.ModifyCount = modifyCount

	enableCount, err := ReadSimpleField(ctx, "enableCount", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'enableCount' field"))
	}
	m.EnableCount = enableCount

	disableCount, err := ReadSimpleField(ctx, "disableCount", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'disableCount' field"))
	}
	m.DisableCount = disableCount

	republishRequestCount, err := ReadSimpleField(ctx, "republishRequestCount", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'republishRequestCount' field"))
	}
	m.RepublishRequestCount = republishRequestCount

	republishMessageRequestCount, err := ReadSimpleField(ctx, "republishMessageRequestCount", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'republishMessageRequestCount' field"))
	}
	m.RepublishMessageRequestCount = republishMessageRequestCount

	republishMessageCount, err := ReadSimpleField(ctx, "republishMessageCount", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'republishMessageCount' field"))
	}
	m.RepublishMessageCount = republishMessageCount

	transferRequestCount, err := ReadSimpleField(ctx, "transferRequestCount", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'transferRequestCount' field"))
	}
	m.TransferRequestCount = transferRequestCount

	transferredToAltClientCount, err := ReadSimpleField(ctx, "transferredToAltClientCount", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'transferredToAltClientCount' field"))
	}
	m.TransferredToAltClientCount = transferredToAltClientCount

	transferredToSameClientCount, err := ReadSimpleField(ctx, "transferredToSameClientCount", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'transferredToSameClientCount' field"))
	}
	m.TransferredToSameClientCount = transferredToSameClientCount

	publishRequestCount, err := ReadSimpleField(ctx, "publishRequestCount", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'publishRequestCount' field"))
	}
	m.PublishRequestCount = publishRequestCount

	dataChangeNotificationsCount, err := ReadSimpleField(ctx, "dataChangeNotificationsCount", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dataChangeNotificationsCount' field"))
	}
	m.DataChangeNotificationsCount = dataChangeNotificationsCount

	eventNotificationsCount, err := ReadSimpleField(ctx, "eventNotificationsCount", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'eventNotificationsCount' field"))
	}
	m.EventNotificationsCount = eventNotificationsCount

	notificationsCount, err := ReadSimpleField(ctx, "notificationsCount", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'notificationsCount' field"))
	}
	m.NotificationsCount = notificationsCount

	latePublishRequestCount, err := ReadSimpleField(ctx, "latePublishRequestCount", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'latePublishRequestCount' field"))
	}
	m.LatePublishRequestCount = latePublishRequestCount

	currentKeepAliveCount, err := ReadSimpleField(ctx, "currentKeepAliveCount", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'currentKeepAliveCount' field"))
	}
	m.CurrentKeepAliveCount = currentKeepAliveCount

	currentLifetimeCount, err := ReadSimpleField(ctx, "currentLifetimeCount", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'currentLifetimeCount' field"))
	}
	m.CurrentLifetimeCount = currentLifetimeCount

	unacknowledgedMessageCount, err := ReadSimpleField(ctx, "unacknowledgedMessageCount", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'unacknowledgedMessageCount' field"))
	}
	m.UnacknowledgedMessageCount = unacknowledgedMessageCount

	discardedMessageCount, err := ReadSimpleField(ctx, "discardedMessageCount", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'discardedMessageCount' field"))
	}
	m.DiscardedMessageCount = discardedMessageCount

	monitoredItemCount, err := ReadSimpleField(ctx, "monitoredItemCount", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'monitoredItemCount' field"))
	}
	m.MonitoredItemCount = monitoredItemCount

	disabledMonitoredItemCount, err := ReadSimpleField(ctx, "disabledMonitoredItemCount", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'disabledMonitoredItemCount' field"))
	}
	m.DisabledMonitoredItemCount = disabledMonitoredItemCount

	monitoringQueueOverflowCount, err := ReadSimpleField(ctx, "monitoringQueueOverflowCount", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'monitoringQueueOverflowCount' field"))
	}
	m.MonitoringQueueOverflowCount = monitoringQueueOverflowCount

	nextSequenceNumber, err := ReadSimpleField(ctx, "nextSequenceNumber", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'nextSequenceNumber' field"))
	}
	m.NextSequenceNumber = nextSequenceNumber

	eventQueueOverFlowCount, err := ReadSimpleField(ctx, "eventQueueOverFlowCount", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'eventQueueOverFlowCount' field"))
	}
	m.EventQueueOverFlowCount = eventQueueOverFlowCount

	if closeErr := readBuffer.CloseContext("SubscriptionDiagnosticsDataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SubscriptionDiagnosticsDataType")
	}

	return m, nil
}

func (m *_SubscriptionDiagnosticsDataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SubscriptionDiagnosticsDataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SubscriptionDiagnosticsDataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SubscriptionDiagnosticsDataType")
		}

		if err := WriteSimpleField[NodeId](ctx, "sessionId", m.GetSessionId(), WriteComplex[NodeId](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'sessionId' field")
		}

		if err := WriteSimpleField[uint32](ctx, "subscriptionId", m.GetSubscriptionId(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'subscriptionId' field")
		}

		if err := WriteSimpleField[uint8](ctx, "priority", m.GetPriority(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'priority' field")
		}

		if err := WriteSimpleField[float64](ctx, "publishingInterval", m.GetPublishingInterval(), WriteDouble(writeBuffer, 64)); err != nil {
			return errors.Wrap(err, "Error serializing 'publishingInterval' field")
		}

		if err := WriteSimpleField[uint32](ctx, "maxKeepAliveCount", m.GetMaxKeepAliveCount(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'maxKeepAliveCount' field")
		}

		if err := WriteSimpleField[uint32](ctx, "maxLifetimeCount", m.GetMaxLifetimeCount(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'maxLifetimeCount' field")
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

		if err := WriteSimpleField[uint32](ctx, "modifyCount", m.GetModifyCount(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'modifyCount' field")
		}

		if err := WriteSimpleField[uint32](ctx, "enableCount", m.GetEnableCount(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'enableCount' field")
		}

		if err := WriteSimpleField[uint32](ctx, "disableCount", m.GetDisableCount(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'disableCount' field")
		}

		if err := WriteSimpleField[uint32](ctx, "republishRequestCount", m.GetRepublishRequestCount(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'republishRequestCount' field")
		}

		if err := WriteSimpleField[uint32](ctx, "republishMessageRequestCount", m.GetRepublishMessageRequestCount(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'republishMessageRequestCount' field")
		}

		if err := WriteSimpleField[uint32](ctx, "republishMessageCount", m.GetRepublishMessageCount(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'republishMessageCount' field")
		}

		if err := WriteSimpleField[uint32](ctx, "transferRequestCount", m.GetTransferRequestCount(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'transferRequestCount' field")
		}

		if err := WriteSimpleField[uint32](ctx, "transferredToAltClientCount", m.GetTransferredToAltClientCount(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'transferredToAltClientCount' field")
		}

		if err := WriteSimpleField[uint32](ctx, "transferredToSameClientCount", m.GetTransferredToSameClientCount(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'transferredToSameClientCount' field")
		}

		if err := WriteSimpleField[uint32](ctx, "publishRequestCount", m.GetPublishRequestCount(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'publishRequestCount' field")
		}

		if err := WriteSimpleField[uint32](ctx, "dataChangeNotificationsCount", m.GetDataChangeNotificationsCount(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'dataChangeNotificationsCount' field")
		}

		if err := WriteSimpleField[uint32](ctx, "eventNotificationsCount", m.GetEventNotificationsCount(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'eventNotificationsCount' field")
		}

		if err := WriteSimpleField[uint32](ctx, "notificationsCount", m.GetNotificationsCount(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'notificationsCount' field")
		}

		if err := WriteSimpleField[uint32](ctx, "latePublishRequestCount", m.GetLatePublishRequestCount(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'latePublishRequestCount' field")
		}

		if err := WriteSimpleField[uint32](ctx, "currentKeepAliveCount", m.GetCurrentKeepAliveCount(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'currentKeepAliveCount' field")
		}

		if err := WriteSimpleField[uint32](ctx, "currentLifetimeCount", m.GetCurrentLifetimeCount(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'currentLifetimeCount' field")
		}

		if err := WriteSimpleField[uint32](ctx, "unacknowledgedMessageCount", m.GetUnacknowledgedMessageCount(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'unacknowledgedMessageCount' field")
		}

		if err := WriteSimpleField[uint32](ctx, "discardedMessageCount", m.GetDiscardedMessageCount(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'discardedMessageCount' field")
		}

		if err := WriteSimpleField[uint32](ctx, "monitoredItemCount", m.GetMonitoredItemCount(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'monitoredItemCount' field")
		}

		if err := WriteSimpleField[uint32](ctx, "disabledMonitoredItemCount", m.GetDisabledMonitoredItemCount(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'disabledMonitoredItemCount' field")
		}

		if err := WriteSimpleField[uint32](ctx, "monitoringQueueOverflowCount", m.GetMonitoringQueueOverflowCount(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'monitoringQueueOverflowCount' field")
		}

		if err := WriteSimpleField[uint32](ctx, "nextSequenceNumber", m.GetNextSequenceNumber(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'nextSequenceNumber' field")
		}

		if err := WriteSimpleField[uint32](ctx, "eventQueueOverFlowCount", m.GetEventQueueOverFlowCount(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'eventQueueOverFlowCount' field")
		}

		if popErr := writeBuffer.PopContext("SubscriptionDiagnosticsDataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SubscriptionDiagnosticsDataType")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SubscriptionDiagnosticsDataType) isSubscriptionDiagnosticsDataType() bool {
	return true
}

func (m *_SubscriptionDiagnosticsDataType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
