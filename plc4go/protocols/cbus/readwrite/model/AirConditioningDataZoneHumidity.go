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

// AirConditioningDataZoneHumidity is the corresponding interface of AirConditioningDataZoneHumidity
type AirConditioningDataZoneHumidity interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	AirConditioningData
	// GetZoneGroup returns ZoneGroup (property field)
	GetZoneGroup() byte
	// GetZoneList returns ZoneList (property field)
	GetZoneList() HVACZoneList
	// GetHumidity returns Humidity (property field)
	GetHumidity() HVACHumidity
	// GetSensorStatus returns SensorStatus (property field)
	GetSensorStatus() HVACSensorStatus
	// IsAirConditioningDataZoneHumidity is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsAirConditioningDataZoneHumidity()
}

// _AirConditioningDataZoneHumidity is the data-structure of this message
type _AirConditioningDataZoneHumidity struct {
	AirConditioningDataContract
	ZoneGroup    byte
	ZoneList     HVACZoneList
	Humidity     HVACHumidity
	SensorStatus HVACSensorStatus
}

var _ AirConditioningDataZoneHumidity = (*_AirConditioningDataZoneHumidity)(nil)
var _ AirConditioningDataRequirements = (*_AirConditioningDataZoneHumidity)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AirConditioningDataZoneHumidity) GetParent() AirConditioningDataContract {
	return m.AirConditioningDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AirConditioningDataZoneHumidity) GetZoneGroup() byte {
	return m.ZoneGroup
}

func (m *_AirConditioningDataZoneHumidity) GetZoneList() HVACZoneList {
	return m.ZoneList
}

func (m *_AirConditioningDataZoneHumidity) GetHumidity() HVACHumidity {
	return m.Humidity
}

func (m *_AirConditioningDataZoneHumidity) GetSensorStatus() HVACSensorStatus {
	return m.SensorStatus
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAirConditioningDataZoneHumidity factory function for _AirConditioningDataZoneHumidity
func NewAirConditioningDataZoneHumidity(zoneGroup byte, zoneList HVACZoneList, humidity HVACHumidity, sensorStatus HVACSensorStatus, commandTypeContainer AirConditioningCommandTypeContainer) *_AirConditioningDataZoneHumidity {
	_result := &_AirConditioningDataZoneHumidity{
		AirConditioningDataContract: NewAirConditioningData(commandTypeContainer),
		ZoneGroup:                   zoneGroup,
		ZoneList:                    zoneList,
		Humidity:                    humidity,
		SensorStatus:                sensorStatus,
	}
	_result.AirConditioningDataContract.(*_AirConditioningData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastAirConditioningDataZoneHumidity(structType any) AirConditioningDataZoneHumidity {
	if casted, ok := structType.(AirConditioningDataZoneHumidity); ok {
		return casted
	}
	if casted, ok := structType.(*AirConditioningDataZoneHumidity); ok {
		return *casted
	}
	return nil
}

func (m *_AirConditioningDataZoneHumidity) GetTypeName() string {
	return "AirConditioningDataZoneHumidity"
}

func (m *_AirConditioningDataZoneHumidity) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.AirConditioningDataContract.(*_AirConditioningData).getLengthInBits(ctx))

	// Simple field (zoneGroup)
	lengthInBits += 8

	// Simple field (zoneList)
	lengthInBits += m.ZoneList.GetLengthInBits(ctx)

	// Simple field (humidity)
	lengthInBits += m.Humidity.GetLengthInBits(ctx)

	// Simple field (sensorStatus)
	lengthInBits += 8

	return lengthInBits
}

func (m *_AirConditioningDataZoneHumidity) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_AirConditioningDataZoneHumidity) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_AirConditioningData) (__airConditioningDataZoneHumidity AirConditioningDataZoneHumidity, err error) {
	m.AirConditioningDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AirConditioningDataZoneHumidity"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AirConditioningDataZoneHumidity")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	zoneGroup, err := ReadSimpleField(ctx, "zoneGroup", ReadByte(readBuffer, 8))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'zoneGroup' field"))
	}
	m.ZoneGroup = zoneGroup

	zoneList, err := ReadSimpleField[HVACZoneList](ctx, "zoneList", ReadComplex[HVACZoneList](HVACZoneListParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'zoneList' field"))
	}
	m.ZoneList = zoneList

	humidity, err := ReadSimpleField[HVACHumidity](ctx, "humidity", ReadComplex[HVACHumidity](HVACHumidityParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'humidity' field"))
	}
	m.Humidity = humidity

	sensorStatus, err := ReadEnumField[HVACSensorStatus](ctx, "sensorStatus", "HVACSensorStatus", ReadEnum(HVACSensorStatusByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'sensorStatus' field"))
	}
	m.SensorStatus = sensorStatus

	if closeErr := readBuffer.CloseContext("AirConditioningDataZoneHumidity"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AirConditioningDataZoneHumidity")
	}

	return m, nil
}

func (m *_AirConditioningDataZoneHumidity) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AirConditioningDataZoneHumidity) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AirConditioningDataZoneHumidity"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AirConditioningDataZoneHumidity")
		}

		if err := WriteSimpleField[byte](ctx, "zoneGroup", m.GetZoneGroup(), WriteByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'zoneGroup' field")
		}

		if err := WriteSimpleField[HVACZoneList](ctx, "zoneList", m.GetZoneList(), WriteComplex[HVACZoneList](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'zoneList' field")
		}

		if err := WriteSimpleField[HVACHumidity](ctx, "humidity", m.GetHumidity(), WriteComplex[HVACHumidity](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'humidity' field")
		}

		if err := WriteSimpleEnumField[HVACSensorStatus](ctx, "sensorStatus", "HVACSensorStatus", m.GetSensorStatus(), WriteEnum[HVACSensorStatus, uint8](HVACSensorStatus.GetValue, HVACSensorStatus.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
			return errors.Wrap(err, "Error serializing 'sensorStatus' field")
		}

		if popErr := writeBuffer.PopContext("AirConditioningDataZoneHumidity"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AirConditioningDataZoneHumidity")
		}
		return nil
	}
	return m.AirConditioningDataContract.(*_AirConditioningData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AirConditioningDataZoneHumidity) IsAirConditioningDataZoneHumidity() {}

func (m *_AirConditioningDataZoneHumidity) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
