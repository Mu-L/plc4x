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

// BACnetConfirmedServiceRequestReinitializeDevice is the corresponding interface of BACnetConfirmedServiceRequestReinitializeDevice
type BACnetConfirmedServiceRequestReinitializeDevice interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConfirmedServiceRequest
	// GetReinitializedStateOfDevice returns ReinitializedStateOfDevice (property field)
	GetReinitializedStateOfDevice() BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged
	// GetPassword returns Password (property field)
	GetPassword() BACnetContextTagCharacterString
	// IsBACnetConfirmedServiceRequestReinitializeDevice is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConfirmedServiceRequestReinitializeDevice()
}

// _BACnetConfirmedServiceRequestReinitializeDevice is the data-structure of this message
type _BACnetConfirmedServiceRequestReinitializeDevice struct {
	BACnetConfirmedServiceRequestContract
	ReinitializedStateOfDevice BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged
	Password                   BACnetContextTagCharacterString
}

var _ BACnetConfirmedServiceRequestReinitializeDevice = (*_BACnetConfirmedServiceRequestReinitializeDevice)(nil)
var _ BACnetConfirmedServiceRequestRequirements = (*_BACnetConfirmedServiceRequestReinitializeDevice)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConfirmedServiceRequestReinitializeDevice) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_REINITIALIZE_DEVICE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConfirmedServiceRequestReinitializeDevice) GetParent() BACnetConfirmedServiceRequestContract {
	return m.BACnetConfirmedServiceRequestContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestReinitializeDevice) GetReinitializedStateOfDevice() BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged {
	return m.ReinitializedStateOfDevice
}

func (m *_BACnetConfirmedServiceRequestReinitializeDevice) GetPassword() BACnetContextTagCharacterString {
	return m.Password
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConfirmedServiceRequestReinitializeDevice factory function for _BACnetConfirmedServiceRequestReinitializeDevice
func NewBACnetConfirmedServiceRequestReinitializeDevice(reinitializedStateOfDevice BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged, password BACnetContextTagCharacterString, serviceRequestLength uint32) *_BACnetConfirmedServiceRequestReinitializeDevice {
	_result := &_BACnetConfirmedServiceRequestReinitializeDevice{
		BACnetConfirmedServiceRequestContract: NewBACnetConfirmedServiceRequest(serviceRequestLength),
		ReinitializedStateOfDevice:            reinitializedStateOfDevice,
		Password:                              password,
	}
	_result.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestReinitializeDevice(structType any) BACnetConfirmedServiceRequestReinitializeDevice {
	if casted, ok := structType.(BACnetConfirmedServiceRequestReinitializeDevice); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestReinitializeDevice); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestReinitializeDevice) GetTypeName() string {
	return "BACnetConfirmedServiceRequestReinitializeDevice"
}

func (m *_BACnetConfirmedServiceRequestReinitializeDevice) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).getLengthInBits(ctx))

	// Simple field (reinitializedStateOfDevice)
	lengthInBits += m.ReinitializedStateOfDevice.GetLengthInBits(ctx)

	// Optional Field (password)
	if m.Password != nil {
		lengthInBits += m.Password.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestReinitializeDevice) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConfirmedServiceRequestReinitializeDevice) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConfirmedServiceRequest, serviceRequestLength uint32) (__bACnetConfirmedServiceRequestReinitializeDevice BACnetConfirmedServiceRequestReinitializeDevice, err error) {
	m.BACnetConfirmedServiceRequestContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestReinitializeDevice"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestReinitializeDevice")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	reinitializedStateOfDevice, err := ReadSimpleField[BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged](ctx, "reinitializedStateOfDevice", ReadComplex[BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged](BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'reinitializedStateOfDevice' field"))
	}
	m.ReinitializedStateOfDevice = reinitializedStateOfDevice

	var password BACnetContextTagCharacterString
	_password, err := ReadOptionalField[BACnetContextTagCharacterString](ctx, "password", ReadComplex[BACnetContextTagCharacterString](BACnetContextTagParseWithBufferProducer[BACnetContextTagCharacterString]((uint8)(uint8(1)), (BACnetDataType)(BACnetDataType_CHARACTER_STRING)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'password' field"))
	}
	if _password != nil {
		password = *_password
		m.Password = password
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestReinitializeDevice"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestReinitializeDevice")
	}

	return m, nil
}

func (m *_BACnetConfirmedServiceRequestReinitializeDevice) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConfirmedServiceRequestReinitializeDevice) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestReinitializeDevice"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestReinitializeDevice")
		}

		if err := WriteSimpleField[BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged](ctx, "reinitializedStateOfDevice", m.GetReinitializedStateOfDevice(), WriteComplex[BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'reinitializedStateOfDevice' field")
		}

		if err := WriteOptionalField[BACnetContextTagCharacterString](ctx, "password", GetRef(m.GetPassword()), WriteComplex[BACnetContextTagCharacterString](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'password' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestReinitializeDevice"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestReinitializeDevice")
		}
		return nil
	}
	return m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConfirmedServiceRequestReinitializeDevice) IsBACnetConfirmedServiceRequestReinitializeDevice() {
}

func (m *_BACnetConfirmedServiceRequestReinitializeDevice) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
