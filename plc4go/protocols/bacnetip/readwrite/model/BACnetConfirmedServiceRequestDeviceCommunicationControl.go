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

// BACnetConfirmedServiceRequestDeviceCommunicationControl is the corresponding interface of BACnetConfirmedServiceRequestDeviceCommunicationControl
type BACnetConfirmedServiceRequestDeviceCommunicationControl interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConfirmedServiceRequest
	// GetTimeDuration returns TimeDuration (property field)
	GetTimeDuration() BACnetContextTagUnsignedInteger
	// GetEnableDisable returns EnableDisable (property field)
	GetEnableDisable() BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTagged
	// GetPassword returns Password (property field)
	GetPassword() BACnetContextTagCharacterString
}

// BACnetConfirmedServiceRequestDeviceCommunicationControlExactly can be used when we want exactly this type and not a type which fulfills BACnetConfirmedServiceRequestDeviceCommunicationControl.
// This is useful for switch cases.
type BACnetConfirmedServiceRequestDeviceCommunicationControlExactly interface {
	BACnetConfirmedServiceRequestDeviceCommunicationControl
	isBACnetConfirmedServiceRequestDeviceCommunicationControl() bool
}

// _BACnetConfirmedServiceRequestDeviceCommunicationControl is the data-structure of this message
type _BACnetConfirmedServiceRequestDeviceCommunicationControl struct {
	BACnetConfirmedServiceRequestContract
	TimeDuration  BACnetContextTagUnsignedInteger
	EnableDisable BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTagged
	Password      BACnetContextTagCharacterString
}

var _ BACnetConfirmedServiceRequestDeviceCommunicationControl = (*_BACnetConfirmedServiceRequestDeviceCommunicationControl)(nil)
var _ BACnetConfirmedServiceRequestRequirements = (*_BACnetConfirmedServiceRequestDeviceCommunicationControl)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConfirmedServiceRequestDeviceCommunicationControl) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_DEVICE_COMMUNICATION_CONTROL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConfirmedServiceRequestDeviceCommunicationControl) GetParent() BACnetConfirmedServiceRequestContract {
	return m.BACnetConfirmedServiceRequestContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestDeviceCommunicationControl) GetTimeDuration() BACnetContextTagUnsignedInteger {
	return m.TimeDuration
}

func (m *_BACnetConfirmedServiceRequestDeviceCommunicationControl) GetEnableDisable() BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTagged {
	return m.EnableDisable
}

func (m *_BACnetConfirmedServiceRequestDeviceCommunicationControl) GetPassword() BACnetContextTagCharacterString {
	return m.Password
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConfirmedServiceRequestDeviceCommunicationControl factory function for _BACnetConfirmedServiceRequestDeviceCommunicationControl
func NewBACnetConfirmedServiceRequestDeviceCommunicationControl(timeDuration BACnetContextTagUnsignedInteger, enableDisable BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTagged, password BACnetContextTagCharacterString, serviceRequestLength uint32) *_BACnetConfirmedServiceRequestDeviceCommunicationControl {
	_result := &_BACnetConfirmedServiceRequestDeviceCommunicationControl{
		BACnetConfirmedServiceRequestContract: NewBACnetConfirmedServiceRequest(serviceRequestLength),
		TimeDuration:                          timeDuration,
		EnableDisable:                         enableDisable,
		Password:                              password,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestDeviceCommunicationControl(structType any) BACnetConfirmedServiceRequestDeviceCommunicationControl {
	if casted, ok := structType.(BACnetConfirmedServiceRequestDeviceCommunicationControl); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestDeviceCommunicationControl); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestDeviceCommunicationControl) GetTypeName() string {
	return "BACnetConfirmedServiceRequestDeviceCommunicationControl"
}

func (m *_BACnetConfirmedServiceRequestDeviceCommunicationControl) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).getLengthInBits(ctx))

	// Optional Field (timeDuration)
	if m.TimeDuration != nil {
		lengthInBits += m.TimeDuration.GetLengthInBits(ctx)
	}

	// Simple field (enableDisable)
	lengthInBits += m.EnableDisable.GetLengthInBits(ctx)

	// Optional Field (password)
	if m.Password != nil {
		lengthInBits += m.Password.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestDeviceCommunicationControl) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConfirmedServiceRequestDeviceCommunicationControl) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConfirmedServiceRequest, serviceRequestLength uint32) (__bACnetConfirmedServiceRequestDeviceCommunicationControl BACnetConfirmedServiceRequestDeviceCommunicationControl, err error) {
	m.BACnetConfirmedServiceRequestContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestDeviceCommunicationControl"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestDeviceCommunicationControl")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	var timeDuration BACnetContextTagUnsignedInteger
	_timeDuration, err := ReadOptionalField[BACnetContextTagUnsignedInteger](ctx, "timeDuration", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timeDuration' field"))
	}
	if _timeDuration != nil {
		timeDuration = *_timeDuration
		m.TimeDuration = timeDuration
	}

	enableDisable, err := ReadSimpleField[BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTagged](ctx, "enableDisable", ReadComplex[BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTagged](BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTaggedParseWithBufferProducer((uint8)(uint8(1)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'enableDisable' field"))
	}
	m.EnableDisable = enableDisable

	var password BACnetContextTagCharacterString
	_password, err := ReadOptionalField[BACnetContextTagCharacterString](ctx, "password", ReadComplex[BACnetContextTagCharacterString](BACnetContextTagParseWithBufferProducer[BACnetContextTagCharacterString]((uint8)(uint8(2)), (BACnetDataType)(BACnetDataType_CHARACTER_STRING)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'password' field"))
	}
	if _password != nil {
		password = *_password
		m.Password = password
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestDeviceCommunicationControl"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestDeviceCommunicationControl")
	}

	return m, nil
}

func (m *_BACnetConfirmedServiceRequestDeviceCommunicationControl) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConfirmedServiceRequestDeviceCommunicationControl) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestDeviceCommunicationControl"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestDeviceCommunicationControl")
		}

		if err := WriteOptionalField[BACnetContextTagUnsignedInteger](ctx, "timeDuration", GetRef(m.GetTimeDuration()), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'timeDuration' field")
		}

		if err := WriteSimpleField[BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTagged](ctx, "enableDisable", m.GetEnableDisable(), WriteComplex[BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'enableDisable' field")
		}

		if err := WriteOptionalField[BACnetContextTagCharacterString](ctx, "password", GetRef(m.GetPassword()), WriteComplex[BACnetContextTagCharacterString](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'password' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestDeviceCommunicationControl"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestDeviceCommunicationControl")
		}
		return nil
	}
	return m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConfirmedServiceRequestDeviceCommunicationControl) isBACnetConfirmedServiceRequestDeviceCommunicationControl() bool {
	return true
}

func (m *_BACnetConfirmedServiceRequestDeviceCommunicationControl) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
