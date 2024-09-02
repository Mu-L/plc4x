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

// BACnetConfirmedServiceRequest is the corresponding interface of BACnetConfirmedServiceRequest
type BACnetConfirmedServiceRequest interface {
	BACnetConfirmedServiceRequestContract
	BACnetConfirmedServiceRequestRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

// BACnetConfirmedServiceRequestContract provides a set of functions which can be overwritten by a sub struct
type BACnetConfirmedServiceRequestContract interface {
	// GetServiceRequestPayloadLength returns ServiceRequestPayloadLength (virtual field)
	GetServiceRequestPayloadLength() uint32
	// GetServiceRequestLength() returns a parser argument
	GetServiceRequestLength() uint32
}

// BACnetConfirmedServiceRequestRequirements provides a set of functions which need to be implemented by a sub struct
type BACnetConfirmedServiceRequestRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetServiceChoice returns ServiceChoice (discriminator field)
	GetServiceChoice() BACnetConfirmedServiceChoice
}

// BACnetConfirmedServiceRequestExactly can be used when we want exactly this type and not a type which fulfills BACnetConfirmedServiceRequest.
// This is useful for switch cases.
type BACnetConfirmedServiceRequestExactly interface {
	BACnetConfirmedServiceRequest
	isBACnetConfirmedServiceRequest() bool
}

// _BACnetConfirmedServiceRequest is the data-structure of this message
type _BACnetConfirmedServiceRequest struct {
	_SubType BACnetConfirmedServiceRequest

	// Arguments.
	ServiceRequestLength uint32
}

var _ BACnetConfirmedServiceRequestContract = (*_BACnetConfirmedServiceRequest)(nil)

type BACnetConfirmedServiceRequestChild interface {
	utils.Serializable

	GetParent() *BACnetConfirmedServiceRequest

	GetTypeName() string
	BACnetConfirmedServiceRequest
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (pm *_BACnetConfirmedServiceRequest) GetServiceRequestPayloadLength() uint32 {
	m := pm._SubType
	ctx := context.Background()
	_ = ctx
	return uint32(utils.InlineIf((bool((m.GetServiceRequestLength()) > (0))), func() any { return uint32((uint32(m.GetServiceRequestLength()) - uint32(uint32(1)))) }, func() any { return uint32(uint32(0)) }).(uint32))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConfirmedServiceRequest factory function for _BACnetConfirmedServiceRequest
func NewBACnetConfirmedServiceRequest(serviceRequestLength uint32) *_BACnetConfirmedServiceRequest {
	return &_BACnetConfirmedServiceRequest{ServiceRequestLength: serviceRequestLength}
}

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequest(structType any) BACnetConfirmedServiceRequest {
	if casted, ok := structType.(BACnetConfirmedServiceRequest); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequest); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequest) GetTypeName() string {
	return "BACnetConfirmedServiceRequest"
}

func (m *_BACnetConfirmedServiceRequest) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (serviceChoice)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func BACnetConfirmedServiceRequestParse[T BACnetConfirmedServiceRequest](ctx context.Context, theBytes []byte, serviceRequestLength uint32) (T, error) {
	return BACnetConfirmedServiceRequestParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes), serviceRequestLength)
}

func BACnetConfirmedServiceRequestParseWithBufferProducer[T BACnetConfirmedServiceRequest](serviceRequestLength uint32) func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := BACnetConfirmedServiceRequestParseWithBuffer[T](ctx, readBuffer, serviceRequestLength)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, err
	}
}

func BACnetConfirmedServiceRequestParseWithBuffer[T BACnetConfirmedServiceRequest](ctx context.Context, readBuffer utils.ReadBuffer, serviceRequestLength uint32) (T, error) {
	v, err := (&_BACnetConfirmedServiceRequest{}).parse(ctx, readBuffer, serviceRequestLength)
	if err != nil {
		var zero T
		return zero, err
	}
	return v.(T), err
}

func (m *_BACnetConfirmedServiceRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, serviceRequestLength uint32) (__bACnetConfirmedServiceRequest BACnetConfirmedServiceRequest, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	serviceChoice, err := ReadDiscriminatorEnumField[BACnetConfirmedServiceChoice](ctx, "serviceChoice", "BACnetConfirmedServiceChoice", ReadEnum(BACnetConfirmedServiceChoiceByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'serviceChoice' field"))
	}

	serviceRequestPayloadLength, err := ReadVirtualField[uint32](ctx, "serviceRequestPayloadLength", (*uint32)(nil), utils.InlineIf((bool((serviceRequestLength) > (0))), func() any { return uint32((uint32(serviceRequestLength) - uint32(uint32(1)))) }, func() any { return uint32(uint32(0)) }).(uint32))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'serviceRequestPayloadLength' field"))
	}
	_ = serviceRequestPayloadLength

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child BACnetConfirmedServiceRequest
	switch {
	case serviceChoice == BACnetConfirmedServiceChoice_ACKNOWLEDGE_ALARM: // BACnetConfirmedServiceRequestAcknowledgeAlarm
		if _child, err = (&_BACnetConfirmedServiceRequestAcknowledgeAlarm{}).parse(ctx, readBuffer, m, serviceRequestLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetConfirmedServiceRequestAcknowledgeAlarm for type-switch of BACnetConfirmedServiceRequest")
		}
	case serviceChoice == BACnetConfirmedServiceChoice_CONFIRMED_COV_NOTIFICATION: // BACnetConfirmedServiceRequestConfirmedCOVNotification
		if _child, err = (&_BACnetConfirmedServiceRequestConfirmedCOVNotification{}).parse(ctx, readBuffer, m, serviceRequestLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetConfirmedServiceRequestConfirmedCOVNotification for type-switch of BACnetConfirmedServiceRequest")
		}
	case serviceChoice == BACnetConfirmedServiceChoice_CONFIRMED_COV_NOTIFICATION_MULTIPLE: // BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple
		if _child, err = (&_BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple{}).parse(ctx, readBuffer, m, serviceRequestLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple for type-switch of BACnetConfirmedServiceRequest")
		}
	case serviceChoice == BACnetConfirmedServiceChoice_CONFIRMED_EVENT_NOTIFICATION: // BACnetConfirmedServiceRequestConfirmedEventNotification
		if _child, err = (&_BACnetConfirmedServiceRequestConfirmedEventNotification{}).parse(ctx, readBuffer, m, serviceRequestLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetConfirmedServiceRequestConfirmedEventNotification for type-switch of BACnetConfirmedServiceRequest")
		}
	case serviceChoice == BACnetConfirmedServiceChoice_GET_ENROLLMENT_SUMMARY: // BACnetConfirmedServiceRequestGetEnrollmentSummary
		if _child, err = (&_BACnetConfirmedServiceRequestGetEnrollmentSummary{}).parse(ctx, readBuffer, m, serviceRequestLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetConfirmedServiceRequestGetEnrollmentSummary for type-switch of BACnetConfirmedServiceRequest")
		}
	case serviceChoice == BACnetConfirmedServiceChoice_GET_EVENT_INFORMATION: // BACnetConfirmedServiceRequestGetEventInformation
		if _child, err = (&_BACnetConfirmedServiceRequestGetEventInformation{}).parse(ctx, readBuffer, m, serviceRequestLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetConfirmedServiceRequestGetEventInformation for type-switch of BACnetConfirmedServiceRequest")
		}
	case serviceChoice == BACnetConfirmedServiceChoice_LIFE_SAFETY_OPERATION: // BACnetConfirmedServiceRequestLifeSafetyOperation
		if _child, err = (&_BACnetConfirmedServiceRequestLifeSafetyOperation{}).parse(ctx, readBuffer, m, serviceRequestLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetConfirmedServiceRequestLifeSafetyOperation for type-switch of BACnetConfirmedServiceRequest")
		}
	case serviceChoice == BACnetConfirmedServiceChoice_SUBSCRIBE_COV: // BACnetConfirmedServiceRequestSubscribeCOV
		if _child, err = (&_BACnetConfirmedServiceRequestSubscribeCOV{}).parse(ctx, readBuffer, m, serviceRequestLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetConfirmedServiceRequestSubscribeCOV for type-switch of BACnetConfirmedServiceRequest")
		}
	case serviceChoice == BACnetConfirmedServiceChoice_SUBSCRIBE_COV_PROPERTY: // BACnetConfirmedServiceRequestSubscribeCOVProperty
		if _child, err = (&_BACnetConfirmedServiceRequestSubscribeCOVProperty{}).parse(ctx, readBuffer, m, serviceRequestLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetConfirmedServiceRequestSubscribeCOVProperty for type-switch of BACnetConfirmedServiceRequest")
		}
	case serviceChoice == BACnetConfirmedServiceChoice_SUBSCRIBE_COV_PROPERTY_MULTIPLE: // BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple
		if _child, err = (&_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple{}).parse(ctx, readBuffer, m, serviceRequestLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple for type-switch of BACnetConfirmedServiceRequest")
		}
	case serviceChoice == BACnetConfirmedServiceChoice_ATOMIC_READ_FILE: // BACnetConfirmedServiceRequestAtomicReadFile
		if _child, err = (&_BACnetConfirmedServiceRequestAtomicReadFile{}).parse(ctx, readBuffer, m, serviceRequestLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetConfirmedServiceRequestAtomicReadFile for type-switch of BACnetConfirmedServiceRequest")
		}
	case serviceChoice == BACnetConfirmedServiceChoice_ATOMIC_WRITE_FILE: // BACnetConfirmedServiceRequestAtomicWriteFile
		if _child, err = (&_BACnetConfirmedServiceRequestAtomicWriteFile{}).parse(ctx, readBuffer, m, serviceRequestLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetConfirmedServiceRequestAtomicWriteFile for type-switch of BACnetConfirmedServiceRequest")
		}
	case serviceChoice == BACnetConfirmedServiceChoice_ADD_LIST_ELEMENT: // BACnetConfirmedServiceRequestAddListElement
		if _child, err = (&_BACnetConfirmedServiceRequestAddListElement{}).parse(ctx, readBuffer, m, serviceRequestLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetConfirmedServiceRequestAddListElement for type-switch of BACnetConfirmedServiceRequest")
		}
	case serviceChoice == BACnetConfirmedServiceChoice_REMOVE_LIST_ELEMENT: // BACnetConfirmedServiceRequestRemoveListElement
		if _child, err = (&_BACnetConfirmedServiceRequestRemoveListElement{}).parse(ctx, readBuffer, m, serviceRequestLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetConfirmedServiceRequestRemoveListElement for type-switch of BACnetConfirmedServiceRequest")
		}
	case serviceChoice == BACnetConfirmedServiceChoice_CREATE_OBJECT: // BACnetConfirmedServiceRequestCreateObject
		if _child, err = (&_BACnetConfirmedServiceRequestCreateObject{}).parse(ctx, readBuffer, m, serviceRequestLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetConfirmedServiceRequestCreateObject for type-switch of BACnetConfirmedServiceRequest")
		}
	case serviceChoice == BACnetConfirmedServiceChoice_DELETE_OBJECT: // BACnetConfirmedServiceRequestDeleteObject
		if _child, err = (&_BACnetConfirmedServiceRequestDeleteObject{}).parse(ctx, readBuffer, m, serviceRequestLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetConfirmedServiceRequestDeleteObject for type-switch of BACnetConfirmedServiceRequest")
		}
	case serviceChoice == BACnetConfirmedServiceChoice_READ_PROPERTY: // BACnetConfirmedServiceRequestReadProperty
		if _child, err = (&_BACnetConfirmedServiceRequestReadProperty{}).parse(ctx, readBuffer, m, serviceRequestLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetConfirmedServiceRequestReadProperty for type-switch of BACnetConfirmedServiceRequest")
		}
	case serviceChoice == BACnetConfirmedServiceChoice_READ_PROPERTY_MULTIPLE: // BACnetConfirmedServiceRequestReadPropertyMultiple
		if _child, err = (&_BACnetConfirmedServiceRequestReadPropertyMultiple{}).parse(ctx, readBuffer, m, serviceRequestPayloadLength, serviceRequestLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetConfirmedServiceRequestReadPropertyMultiple for type-switch of BACnetConfirmedServiceRequest")
		}
	case serviceChoice == BACnetConfirmedServiceChoice_READ_RANGE: // BACnetConfirmedServiceRequestReadRange
		if _child, err = (&_BACnetConfirmedServiceRequestReadRange{}).parse(ctx, readBuffer, m, serviceRequestLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetConfirmedServiceRequestReadRange for type-switch of BACnetConfirmedServiceRequest")
		}
	case serviceChoice == BACnetConfirmedServiceChoice_WRITE_PROPERTY: // BACnetConfirmedServiceRequestWriteProperty
		if _child, err = (&_BACnetConfirmedServiceRequestWriteProperty{}).parse(ctx, readBuffer, m, serviceRequestLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetConfirmedServiceRequestWriteProperty for type-switch of BACnetConfirmedServiceRequest")
		}
	case serviceChoice == BACnetConfirmedServiceChoice_WRITE_PROPERTY_MULTIPLE: // BACnetConfirmedServiceRequestWritePropertyMultiple
		if _child, err = (&_BACnetConfirmedServiceRequestWritePropertyMultiple{}).parse(ctx, readBuffer, m, serviceRequestPayloadLength, serviceRequestLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetConfirmedServiceRequestWritePropertyMultiple for type-switch of BACnetConfirmedServiceRequest")
		}
	case serviceChoice == BACnetConfirmedServiceChoice_DEVICE_COMMUNICATION_CONTROL: // BACnetConfirmedServiceRequestDeviceCommunicationControl
		if _child, err = (&_BACnetConfirmedServiceRequestDeviceCommunicationControl{}).parse(ctx, readBuffer, m, serviceRequestLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetConfirmedServiceRequestDeviceCommunicationControl for type-switch of BACnetConfirmedServiceRequest")
		}
	case serviceChoice == BACnetConfirmedServiceChoice_CONFIRMED_PRIVATE_TRANSFER: // BACnetConfirmedServiceRequestConfirmedPrivateTransfer
		if _child, err = (&_BACnetConfirmedServiceRequestConfirmedPrivateTransfer{}).parse(ctx, readBuffer, m, serviceRequestLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetConfirmedServiceRequestConfirmedPrivateTransfer for type-switch of BACnetConfirmedServiceRequest")
		}
	case serviceChoice == BACnetConfirmedServiceChoice_CONFIRMED_TEXT_MESSAGE: // BACnetConfirmedServiceRequestConfirmedTextMessage
		if _child, err = (&_BACnetConfirmedServiceRequestConfirmedTextMessage{}).parse(ctx, readBuffer, m, serviceRequestLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetConfirmedServiceRequestConfirmedTextMessage for type-switch of BACnetConfirmedServiceRequest")
		}
	case serviceChoice == BACnetConfirmedServiceChoice_REINITIALIZE_DEVICE: // BACnetConfirmedServiceRequestReinitializeDevice
		if _child, err = (&_BACnetConfirmedServiceRequestReinitializeDevice{}).parse(ctx, readBuffer, m, serviceRequestLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetConfirmedServiceRequestReinitializeDevice for type-switch of BACnetConfirmedServiceRequest")
		}
	case serviceChoice == BACnetConfirmedServiceChoice_VT_OPEN: // BACnetConfirmedServiceRequestVTOpen
		if _child, err = (&_BACnetConfirmedServiceRequestVTOpen{}).parse(ctx, readBuffer, m, serviceRequestLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetConfirmedServiceRequestVTOpen for type-switch of BACnetConfirmedServiceRequest")
		}
	case serviceChoice == BACnetConfirmedServiceChoice_VT_CLOSE: // BACnetConfirmedServiceRequestVTClose
		if _child, err = (&_BACnetConfirmedServiceRequestVTClose{}).parse(ctx, readBuffer, m, serviceRequestPayloadLength, serviceRequestLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetConfirmedServiceRequestVTClose for type-switch of BACnetConfirmedServiceRequest")
		}
	case serviceChoice == BACnetConfirmedServiceChoice_VT_DATA: // BACnetConfirmedServiceRequestVTData
		if _child, err = (&_BACnetConfirmedServiceRequestVTData{}).parse(ctx, readBuffer, m, serviceRequestLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetConfirmedServiceRequestVTData for type-switch of BACnetConfirmedServiceRequest")
		}
	case serviceChoice == BACnetConfirmedServiceChoice_AUTHENTICATE: // BACnetConfirmedServiceRequestAuthenticate
		if _child, err = (&_BACnetConfirmedServiceRequestAuthenticate{}).parse(ctx, readBuffer, m, serviceRequestPayloadLength, serviceRequestLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetConfirmedServiceRequestAuthenticate for type-switch of BACnetConfirmedServiceRequest")
		}
	case serviceChoice == BACnetConfirmedServiceChoice_REQUEST_KEY: // BACnetConfirmedServiceRequestRequestKey
		if _child, err = (&_BACnetConfirmedServiceRequestRequestKey{}).parse(ctx, readBuffer, m, serviceRequestPayloadLength, serviceRequestLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetConfirmedServiceRequestRequestKey for type-switch of BACnetConfirmedServiceRequest")
		}
	case serviceChoice == BACnetConfirmedServiceChoice_READ_PROPERTY_CONDITIONAL: // BACnetConfirmedServiceRequestReadPropertyConditional
		if _child, err = (&_BACnetConfirmedServiceRequestReadPropertyConditional{}).parse(ctx, readBuffer, m, serviceRequestPayloadLength, serviceRequestLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetConfirmedServiceRequestReadPropertyConditional for type-switch of BACnetConfirmedServiceRequest")
		}
	case 0 == 0: // BACnetConfirmedServiceRequestUnknown
		if _child, err = (&_BACnetConfirmedServiceRequestUnknown{}).parse(ctx, readBuffer, m, serviceRequestPayloadLength, serviceRequestLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetConfirmedServiceRequestUnknown for type-switch of BACnetConfirmedServiceRequest")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [serviceChoice=%v]", serviceChoice)
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequest")
	}

	return _child, nil
}

func (pm *_BACnetConfirmedServiceRequest) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetConfirmedServiceRequest, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequest"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequest")
	}

	if err := WriteDiscriminatorEnumField(ctx, "serviceChoice", "BACnetConfirmedServiceChoice", m.GetServiceChoice(), WriteEnum[BACnetConfirmedServiceChoice, uint8](BACnetConfirmedServiceChoice.GetValue, BACnetConfirmedServiceChoice.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
		return errors.Wrap(err, "Error serializing 'serviceChoice' field")
	}
	// Virtual field
	serviceRequestPayloadLength := m.GetServiceRequestPayloadLength()
	_ = serviceRequestPayloadLength
	if _serviceRequestPayloadLengthErr := writeBuffer.WriteVirtual(ctx, "serviceRequestPayloadLength", m.GetServiceRequestPayloadLength()); _serviceRequestPayloadLengthErr != nil {
		return errors.Wrap(_serviceRequestPayloadLengthErr, "Error serializing 'serviceRequestPayloadLength' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequest"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequest")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetConfirmedServiceRequest) GetServiceRequestLength() uint32 {
	return m.ServiceRequestLength
}

//
////

func (m *_BACnetConfirmedServiceRequest) isBACnetConfirmedServiceRequest() bool {
	return true
}
