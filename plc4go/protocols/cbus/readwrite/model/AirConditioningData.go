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

// AirConditioningData is the corresponding interface of AirConditioningData
type AirConditioningData interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetCommandTypeContainer returns CommandTypeContainer (property field)
	GetCommandTypeContainer() AirConditioningCommandTypeContainer
	// GetCommandType returns CommandType (virtual field)
	GetCommandType() AirConditioningCommandType
}

// AirConditioningDataExactly can be used when we want exactly this type and not a type which fulfills AirConditioningData.
// This is useful for switch cases.
type AirConditioningDataExactly interface {
	AirConditioningData
	isAirConditioningData() bool
}

// _AirConditioningData is the data-structure of this message
type _AirConditioningData struct {
	_AirConditioningDataChildRequirements
	CommandTypeContainer AirConditioningCommandTypeContainer
}

type _AirConditioningDataChildRequirements interface {
	utils.Serializable
	GetLengthInBits(ctx context.Context) uint16
	GetCommandType() AirConditioningCommandType
}

type AirConditioningDataParent interface {
	SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child AirConditioningData, serializeChildFunction func() error) error
	GetTypeName() string
}

type AirConditioningDataChild interface {
	utils.Serializable
	InitializeParent(parent AirConditioningData, commandTypeContainer AirConditioningCommandTypeContainer)
	GetParent() *AirConditioningData

	GetTypeName() string
	AirConditioningData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AirConditioningData) GetCommandTypeContainer() AirConditioningCommandTypeContainer {
	return m.CommandTypeContainer
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_AirConditioningData) GetCommandType() AirConditioningCommandType {
	ctx := context.Background()
	_ = ctx
	return CastAirConditioningCommandType(m.GetCommandTypeContainer().CommandType())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAirConditioningData factory function for _AirConditioningData
func NewAirConditioningData(commandTypeContainer AirConditioningCommandTypeContainer) *_AirConditioningData {
	return &_AirConditioningData{CommandTypeContainer: commandTypeContainer}
}

// Deprecated: use the interface for direct cast
func CastAirConditioningData(structType any) AirConditioningData {
	if casted, ok := structType.(AirConditioningData); ok {
		return casted
	}
	if casted, ok := structType.(*AirConditioningData); ok {
		return *casted
	}
	return nil
}

func (m *_AirConditioningData) GetTypeName() string {
	return "AirConditioningData"
}

func (m *_AirConditioningData) GetParentLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (commandTypeContainer)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_AirConditioningData) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func AirConditioningDataParse(ctx context.Context, theBytes []byte) (AirConditioningData, error) {
	return AirConditioningDataParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func AirConditioningDataParseWithBufferProducer[T AirConditioningData]() func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		buffer, err := AirConditioningDataParseWithBuffer(ctx, readBuffer)
		if err != nil {
			var zero T
			return zero, err
		}
		return buffer.(T), err
	}
}

func AirConditioningDataParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (AirConditioningData, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("AirConditioningData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AirConditioningData")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(KnowsAirConditioningCommandTypeContainer(ctx, readBuffer)) {
		return nil, errors.WithStack(utils.ParseAssertError{Message: "no command type could be found"})
	}

	commandTypeContainer, err := ReadEnumField[AirConditioningCommandTypeContainer](ctx, "commandTypeContainer", "AirConditioningCommandTypeContainer", ReadEnum(AirConditioningCommandTypeContainerByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'commandTypeContainer' field"))
	}

	// Virtual field
	_commandType := commandTypeContainer.CommandType()
	commandType := AirConditioningCommandType(_commandType)
	_ = commandType

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type AirConditioningDataChildSerializeRequirement interface {
		AirConditioningData
		InitializeParent(AirConditioningData, AirConditioningCommandTypeContainer)
		GetParent() AirConditioningData
	}
	var _childTemp any
	var _child AirConditioningDataChildSerializeRequirement
	var typeSwitchError error
	switch {
	case commandType == AirConditioningCommandType_HVAC_SCHEDULE_ENTRY: // AirConditioningDataHvacScheduleEntry
		_childTemp, typeSwitchError = AirConditioningDataHvacScheduleEntryParseWithBuffer(ctx, readBuffer)
	case commandType == AirConditioningCommandType_HUMIDITY_SCHEDULE_ENTRY: // AirConditioningDataHumidityScheduleEntry
		_childTemp, typeSwitchError = AirConditioningDataHumidityScheduleEntryParseWithBuffer(ctx, readBuffer)
	case commandType == AirConditioningCommandType_REFRESH: // AirConditioningDataRefresh
		_childTemp, typeSwitchError = AirConditioningDataRefreshParseWithBuffer(ctx, readBuffer)
	case commandType == AirConditioningCommandType_ZONE_HVAC_PLANT_STATUS: // AirConditioningDataZoneHvacPlantStatus
		_childTemp, typeSwitchError = AirConditioningDataZoneHvacPlantStatusParseWithBuffer(ctx, readBuffer)
	case commandType == AirConditioningCommandType_ZONE_HUMIDITY_PLANT_STATUS: // AirConditioningDataZoneHumidityPlantStatus
		_childTemp, typeSwitchError = AirConditioningDataZoneHumidityPlantStatusParseWithBuffer(ctx, readBuffer)
	case commandType == AirConditioningCommandType_ZONE_TEMPERATURE: // AirConditioningDataZoneTemperature
		_childTemp, typeSwitchError = AirConditioningDataZoneTemperatureParseWithBuffer(ctx, readBuffer)
	case commandType == AirConditioningCommandType_ZONE_HUMIDITY: // AirConditioningDataZoneHumidity
		_childTemp, typeSwitchError = AirConditioningDataZoneHumidityParseWithBuffer(ctx, readBuffer)
	case commandType == AirConditioningCommandType_SET_ZONE_GROUP_OFF: // AirConditioningDataSetZoneGroupOff
		_childTemp, typeSwitchError = AirConditioningDataSetZoneGroupOffParseWithBuffer(ctx, readBuffer)
	case commandType == AirConditioningCommandType_SET_ZONE_GROUP_ON: // AirConditioningDataSetZoneGroupOn
		_childTemp, typeSwitchError = AirConditioningDataSetZoneGroupOnParseWithBuffer(ctx, readBuffer)
	case commandType == AirConditioningCommandType_SET_ZONE_HVAC_MODE: // AirConditioningDataSetZoneHvacMode
		_childTemp, typeSwitchError = AirConditioningDataSetZoneHvacModeParseWithBuffer(ctx, readBuffer)
	case commandType == AirConditioningCommandType_SET_PLANT_HVAC_LEVEL: // AirConditioningDataSetPlantHvacLevel
		_childTemp, typeSwitchError = AirConditioningDataSetPlantHvacLevelParseWithBuffer(ctx, readBuffer)
	case commandType == AirConditioningCommandType_SET_ZONE_HUMIDITY_MODE: // AirConditioningDataSetZoneHumidityMode
		_childTemp, typeSwitchError = AirConditioningDataSetZoneHumidityModeParseWithBuffer(ctx, readBuffer)
	case commandType == AirConditioningCommandType_SET_PLANT_HUMIDITY_LEVEL: // AirConditioningDataSetPlantHumidityLevel
		_childTemp, typeSwitchError = AirConditioningDataSetPlantHumidityLevelParseWithBuffer(ctx, readBuffer)
	case commandType == AirConditioningCommandType_SET_HVAC_UPPER_GUARD_LIMIT: // AirConditioningDataSetHvacUpperGuardLimit
		_childTemp, typeSwitchError = AirConditioningDataSetHvacUpperGuardLimitParseWithBuffer(ctx, readBuffer)
	case commandType == AirConditioningCommandType_SET_HVAC_LOWER_GUARD_LIMIT: // AirConditioningDataSetHvacLowerGuardLimit
		_childTemp, typeSwitchError = AirConditioningDataSetHvacLowerGuardLimitParseWithBuffer(ctx, readBuffer)
	case commandType == AirConditioningCommandType_SET_HVAC_SETBACK_LIMIT: // AirConditioningDataSetHvacSetbackLimit
		_childTemp, typeSwitchError = AirConditioningDataSetHvacSetbackLimitParseWithBuffer(ctx, readBuffer)
	case commandType == AirConditioningCommandType_SET_HUMIDITY_UPPER_GUARD_LIMIT: // AirConditioningDataSetHumidityUpperGuardLimit
		_childTemp, typeSwitchError = AirConditioningDataSetHumidityUpperGuardLimitParseWithBuffer(ctx, readBuffer)
	case commandType == AirConditioningCommandType_SET_HUMIDITY_LOWER_GUARD_LIMIT: // AirConditioningDataSetHumidityLowerGuardLimit
		_childTemp, typeSwitchError = AirConditioningDataSetHumidityLowerGuardLimitParseWithBuffer(ctx, readBuffer)
	case commandType == AirConditioningCommandType_SET_HUMIDITY_SETBACK_LIMIT: // AirConditioningDataSetHumiditySetbackLimit
		_childTemp, typeSwitchError = AirConditioningDataSetHumiditySetbackLimitParseWithBuffer(ctx, readBuffer)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [commandType=%v]", commandType)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of AirConditioningData")
	}
	_child = _childTemp.(AirConditioningDataChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("AirConditioningData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AirConditioningData")
	}

	// Finish initializing
	_child.InitializeParent(_child, commandTypeContainer)
	return _child, nil
}

func (pm *_AirConditioningData) SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child AirConditioningData, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("AirConditioningData"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for AirConditioningData")
	}

	// Simple Field (commandTypeContainer)
	if pushErr := writeBuffer.PushContext("commandTypeContainer"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for commandTypeContainer")
	}
	_commandTypeContainerErr := writeBuffer.WriteSerializable(ctx, m.GetCommandTypeContainer())
	if popErr := writeBuffer.PopContext("commandTypeContainer"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for commandTypeContainer")
	}
	if _commandTypeContainerErr != nil {
		return errors.Wrap(_commandTypeContainerErr, "Error serializing 'commandTypeContainer' field")
	}
	// Virtual field
	commandType := m.GetCommandType()
	_ = commandType
	if _commandTypeErr := writeBuffer.WriteVirtual(ctx, "commandType", m.GetCommandType()); _commandTypeErr != nil {
		return errors.Wrap(_commandTypeErr, "Error serializing 'commandType' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("AirConditioningData"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for AirConditioningData")
	}
	return nil
}

func (m *_AirConditioningData) isAirConditioningData() bool {
	return true
}

func (m *_AirConditioningData) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
