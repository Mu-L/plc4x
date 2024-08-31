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

// BACnetUnconfirmedServiceRequestWhoIs is the corresponding interface of BACnetUnconfirmedServiceRequestWhoIs
type BACnetUnconfirmedServiceRequestWhoIs interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetUnconfirmedServiceRequest
	// GetDeviceInstanceRangeLowLimit returns DeviceInstanceRangeLowLimit (property field)
	GetDeviceInstanceRangeLowLimit() BACnetContextTagUnsignedInteger
	// GetDeviceInstanceRangeHighLimit returns DeviceInstanceRangeHighLimit (property field)
	GetDeviceInstanceRangeHighLimit() BACnetContextTagUnsignedInteger
}

// BACnetUnconfirmedServiceRequestWhoIsExactly can be used when we want exactly this type and not a type which fulfills BACnetUnconfirmedServiceRequestWhoIs.
// This is useful for switch cases.
type BACnetUnconfirmedServiceRequestWhoIsExactly interface {
	BACnetUnconfirmedServiceRequestWhoIs
	isBACnetUnconfirmedServiceRequestWhoIs() bool
}

// _BACnetUnconfirmedServiceRequestWhoIs is the data-structure of this message
type _BACnetUnconfirmedServiceRequestWhoIs struct {
	*_BACnetUnconfirmedServiceRequest
	DeviceInstanceRangeLowLimit  BACnetContextTagUnsignedInteger
	DeviceInstanceRangeHighLimit BACnetContextTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetUnconfirmedServiceRequestWhoIs) GetServiceChoice() BACnetUnconfirmedServiceChoice {
	return BACnetUnconfirmedServiceChoice_WHO_IS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetUnconfirmedServiceRequestWhoIs) InitializeParent(parent BACnetUnconfirmedServiceRequest) {
}

func (m *_BACnetUnconfirmedServiceRequestWhoIs) GetParent() BACnetUnconfirmedServiceRequest {
	return m._BACnetUnconfirmedServiceRequest
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetUnconfirmedServiceRequestWhoIs) GetDeviceInstanceRangeLowLimit() BACnetContextTagUnsignedInteger {
	return m.DeviceInstanceRangeLowLimit
}

func (m *_BACnetUnconfirmedServiceRequestWhoIs) GetDeviceInstanceRangeHighLimit() BACnetContextTagUnsignedInteger {
	return m.DeviceInstanceRangeHighLimit
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetUnconfirmedServiceRequestWhoIs factory function for _BACnetUnconfirmedServiceRequestWhoIs
func NewBACnetUnconfirmedServiceRequestWhoIs(deviceInstanceRangeLowLimit BACnetContextTagUnsignedInteger, deviceInstanceRangeHighLimit BACnetContextTagUnsignedInteger, serviceRequestLength uint16) *_BACnetUnconfirmedServiceRequestWhoIs {
	_result := &_BACnetUnconfirmedServiceRequestWhoIs{
		DeviceInstanceRangeLowLimit:      deviceInstanceRangeLowLimit,
		DeviceInstanceRangeHighLimit:     deviceInstanceRangeHighLimit,
		_BACnetUnconfirmedServiceRequest: NewBACnetUnconfirmedServiceRequest(serviceRequestLength),
	}
	_result._BACnetUnconfirmedServiceRequest._BACnetUnconfirmedServiceRequestChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetUnconfirmedServiceRequestWhoIs(structType any) BACnetUnconfirmedServiceRequestWhoIs {
	if casted, ok := structType.(BACnetUnconfirmedServiceRequestWhoIs); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetUnconfirmedServiceRequestWhoIs); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetUnconfirmedServiceRequestWhoIs) GetTypeName() string {
	return "BACnetUnconfirmedServiceRequestWhoIs"
}

func (m *_BACnetUnconfirmedServiceRequestWhoIs) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Optional Field (deviceInstanceRangeLowLimit)
	if m.DeviceInstanceRangeLowLimit != nil {
		lengthInBits += m.DeviceInstanceRangeLowLimit.GetLengthInBits(ctx)
	}

	// Optional Field (deviceInstanceRangeHighLimit)
	if m.DeviceInstanceRangeHighLimit != nil {
		lengthInBits += m.DeviceInstanceRangeHighLimit.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_BACnetUnconfirmedServiceRequestWhoIs) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetUnconfirmedServiceRequestWhoIsParse(ctx context.Context, theBytes []byte, serviceRequestLength uint16) (BACnetUnconfirmedServiceRequestWhoIs, error) {
	return BACnetUnconfirmedServiceRequestWhoIsParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), serviceRequestLength)
}

func BACnetUnconfirmedServiceRequestWhoIsParseWithBufferProducer(serviceRequestLength uint16) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetUnconfirmedServiceRequestWhoIs, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetUnconfirmedServiceRequestWhoIs, error) {
		return BACnetUnconfirmedServiceRequestWhoIsParseWithBuffer(ctx, readBuffer, serviceRequestLength)
	}
}

func BACnetUnconfirmedServiceRequestWhoIsParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, serviceRequestLength uint16) (BACnetUnconfirmedServiceRequestWhoIs, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetUnconfirmedServiceRequestWhoIs"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetUnconfirmedServiceRequestWhoIs")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	_deviceInstanceRangeLowLimit, err := ReadOptionalField[BACnetContextTagUnsignedInteger](ctx, "deviceInstanceRangeLowLimit", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'deviceInstanceRangeLowLimit' field"))
	}
	var deviceInstanceRangeLowLimit BACnetContextTagUnsignedInteger
	if _deviceInstanceRangeLowLimit != nil {
		deviceInstanceRangeLowLimit = *_deviceInstanceRangeLowLimit
	}

	_deviceInstanceRangeHighLimit, err := ReadOptionalField[BACnetContextTagUnsignedInteger](ctx, "deviceInstanceRangeHighLimit", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(1)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer), bool((deviceInstanceRangeLowLimit) != (nil)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'deviceInstanceRangeHighLimit' field"))
	}
	var deviceInstanceRangeHighLimit BACnetContextTagUnsignedInteger
	if _deviceInstanceRangeHighLimit != nil {
		deviceInstanceRangeHighLimit = *_deviceInstanceRangeHighLimit
	}

	if closeErr := readBuffer.CloseContext("BACnetUnconfirmedServiceRequestWhoIs"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetUnconfirmedServiceRequestWhoIs")
	}

	// Create a partially initialized instance
	_child := &_BACnetUnconfirmedServiceRequestWhoIs{
		_BACnetUnconfirmedServiceRequest: &_BACnetUnconfirmedServiceRequest{
			ServiceRequestLength: serviceRequestLength,
		},
		DeviceInstanceRangeLowLimit:  deviceInstanceRangeLowLimit,
		DeviceInstanceRangeHighLimit: deviceInstanceRangeHighLimit,
	}
	_child._BACnetUnconfirmedServiceRequest._BACnetUnconfirmedServiceRequestChildRequirements = _child
	return _child, nil
}

func (m *_BACnetUnconfirmedServiceRequestWhoIs) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetUnconfirmedServiceRequestWhoIs) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetUnconfirmedServiceRequestWhoIs"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetUnconfirmedServiceRequestWhoIs")
		}

		// Optional Field (deviceInstanceRangeLowLimit) (Can be skipped, if the value is null)
		var deviceInstanceRangeLowLimit BACnetContextTagUnsignedInteger = nil
		if m.GetDeviceInstanceRangeLowLimit() != nil {
			if pushErr := writeBuffer.PushContext("deviceInstanceRangeLowLimit"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for deviceInstanceRangeLowLimit")
			}
			deviceInstanceRangeLowLimit = m.GetDeviceInstanceRangeLowLimit()
			_deviceInstanceRangeLowLimitErr := writeBuffer.WriteSerializable(ctx, deviceInstanceRangeLowLimit)
			if popErr := writeBuffer.PopContext("deviceInstanceRangeLowLimit"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for deviceInstanceRangeLowLimit")
			}
			if _deviceInstanceRangeLowLimitErr != nil {
				return errors.Wrap(_deviceInstanceRangeLowLimitErr, "Error serializing 'deviceInstanceRangeLowLimit' field")
			}
		}

		// Optional Field (deviceInstanceRangeHighLimit) (Can be skipped, if the value is null)
		var deviceInstanceRangeHighLimit BACnetContextTagUnsignedInteger = nil
		if m.GetDeviceInstanceRangeHighLimit() != nil {
			if pushErr := writeBuffer.PushContext("deviceInstanceRangeHighLimit"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for deviceInstanceRangeHighLimit")
			}
			deviceInstanceRangeHighLimit = m.GetDeviceInstanceRangeHighLimit()
			_deviceInstanceRangeHighLimitErr := writeBuffer.WriteSerializable(ctx, deviceInstanceRangeHighLimit)
			if popErr := writeBuffer.PopContext("deviceInstanceRangeHighLimit"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for deviceInstanceRangeHighLimit")
			}
			if _deviceInstanceRangeHighLimitErr != nil {
				return errors.Wrap(_deviceInstanceRangeHighLimitErr, "Error serializing 'deviceInstanceRangeHighLimit' field")
			}
		}

		if popErr := writeBuffer.PopContext("BACnetUnconfirmedServiceRequestWhoIs"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetUnconfirmedServiceRequestWhoIs")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetUnconfirmedServiceRequestWhoIs) isBACnetUnconfirmedServiceRequestWhoIs() bool {
	return true
}

func (m *_BACnetUnconfirmedServiceRequestWhoIs) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
