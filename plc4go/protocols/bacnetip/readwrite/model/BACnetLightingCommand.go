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

// BACnetLightingCommand is the corresponding interface of BACnetLightingCommand
type BACnetLightingCommand interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetLightningOperation returns LightningOperation (property field)
	GetLightningOperation() BACnetLightingOperationTagged
	// GetTargetLevel returns TargetLevel (property field)
	GetTargetLevel() BACnetContextTagReal
	// GetRampRate returns RampRate (property field)
	GetRampRate() BACnetContextTagReal
	// GetStepIncrement returns StepIncrement (property field)
	GetStepIncrement() BACnetContextTagReal
	// GetFadeTime returns FadeTime (property field)
	GetFadeTime() BACnetContextTagUnsignedInteger
	// GetPriority returns Priority (property field)
	GetPriority() BACnetContextTagUnsignedInteger
}

// BACnetLightingCommandExactly can be used when we want exactly this type and not a type which fulfills BACnetLightingCommand.
// This is useful for switch cases.
type BACnetLightingCommandExactly interface {
	BACnetLightingCommand
	isBACnetLightingCommand() bool
}

// _BACnetLightingCommand is the data-structure of this message
type _BACnetLightingCommand struct {
	LightningOperation BACnetLightingOperationTagged
	TargetLevel        BACnetContextTagReal
	RampRate           BACnetContextTagReal
	StepIncrement      BACnetContextTagReal
	FadeTime           BACnetContextTagUnsignedInteger
	Priority           BACnetContextTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetLightingCommand) GetLightningOperation() BACnetLightingOperationTagged {
	return m.LightningOperation
}

func (m *_BACnetLightingCommand) GetTargetLevel() BACnetContextTagReal {
	return m.TargetLevel
}

func (m *_BACnetLightingCommand) GetRampRate() BACnetContextTagReal {
	return m.RampRate
}

func (m *_BACnetLightingCommand) GetStepIncrement() BACnetContextTagReal {
	return m.StepIncrement
}

func (m *_BACnetLightingCommand) GetFadeTime() BACnetContextTagUnsignedInteger {
	return m.FadeTime
}

func (m *_BACnetLightingCommand) GetPriority() BACnetContextTagUnsignedInteger {
	return m.Priority
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetLightingCommand factory function for _BACnetLightingCommand
func NewBACnetLightingCommand(lightningOperation BACnetLightingOperationTagged, targetLevel BACnetContextTagReal, rampRate BACnetContextTagReal, stepIncrement BACnetContextTagReal, fadeTime BACnetContextTagUnsignedInteger, priority BACnetContextTagUnsignedInteger) *_BACnetLightingCommand {
	return &_BACnetLightingCommand{LightningOperation: lightningOperation, TargetLevel: targetLevel, RampRate: rampRate, StepIncrement: stepIncrement, FadeTime: fadeTime, Priority: priority}
}

// Deprecated: use the interface for direct cast
func CastBACnetLightingCommand(structType any) BACnetLightingCommand {
	if casted, ok := structType.(BACnetLightingCommand); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetLightingCommand); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetLightingCommand) GetTypeName() string {
	return "BACnetLightingCommand"
}

func (m *_BACnetLightingCommand) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (lightningOperation)
	lengthInBits += m.LightningOperation.GetLengthInBits(ctx)

	// Optional Field (targetLevel)
	if m.TargetLevel != nil {
		lengthInBits += m.TargetLevel.GetLengthInBits(ctx)
	}

	// Optional Field (rampRate)
	if m.RampRate != nil {
		lengthInBits += m.RampRate.GetLengthInBits(ctx)
	}

	// Optional Field (stepIncrement)
	if m.StepIncrement != nil {
		lengthInBits += m.StepIncrement.GetLengthInBits(ctx)
	}

	// Optional Field (fadeTime)
	if m.FadeTime != nil {
		lengthInBits += m.FadeTime.GetLengthInBits(ctx)
	}

	// Optional Field (priority)
	if m.Priority != nil {
		lengthInBits += m.Priority.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_BACnetLightingCommand) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetLightingCommandParse(ctx context.Context, theBytes []byte) (BACnetLightingCommand, error) {
	return BACnetLightingCommandParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetLightingCommandParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetLightingCommand, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetLightingCommand, error) {
		return BACnetLightingCommandParseWithBuffer(ctx, readBuffer)
	}
}

func BACnetLightingCommandParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetLightingCommand, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetLightingCommand"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetLightingCommand")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	lightningOperation, err := ReadSimpleField[BACnetLightingOperationTagged](ctx, "lightningOperation", ReadComplex[BACnetLightingOperationTagged](BACnetLightingOperationTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lightningOperation' field"))
	}

	_targetLevel, err := ReadOptionalField[BACnetContextTagReal](ctx, "targetLevel", ReadComplex[BACnetContextTagReal](BACnetContextTagParseWithBufferProducer[BACnetContextTagReal]((uint8)(uint8(1)), (BACnetDataType)(BACnetDataType_REAL)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'targetLevel' field"))
	}
	var targetLevel BACnetContextTagReal
	if _targetLevel != nil {
		targetLevel = *_targetLevel
	}

	_rampRate, err := ReadOptionalField[BACnetContextTagReal](ctx, "rampRate", ReadComplex[BACnetContextTagReal](BACnetContextTagParseWithBufferProducer[BACnetContextTagReal]((uint8)(uint8(2)), (BACnetDataType)(BACnetDataType_REAL)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'rampRate' field"))
	}
	var rampRate BACnetContextTagReal
	if _rampRate != nil {
		rampRate = *_rampRate
	}

	_stepIncrement, err := ReadOptionalField[BACnetContextTagReal](ctx, "stepIncrement", ReadComplex[BACnetContextTagReal](BACnetContextTagParseWithBufferProducer[BACnetContextTagReal]((uint8)(uint8(3)), (BACnetDataType)(BACnetDataType_REAL)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'stepIncrement' field"))
	}
	var stepIncrement BACnetContextTagReal
	if _stepIncrement != nil {
		stepIncrement = *_stepIncrement
	}

	_fadeTime, err := ReadOptionalField[BACnetContextTagUnsignedInteger](ctx, "fadeTime", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(4)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'fadeTime' field"))
	}
	var fadeTime BACnetContextTagUnsignedInteger
	if _fadeTime != nil {
		fadeTime = *_fadeTime
	}

	_priority, err := ReadOptionalField[BACnetContextTagUnsignedInteger](ctx, "priority", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(5)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'priority' field"))
	}
	var priority BACnetContextTagUnsignedInteger
	if _priority != nil {
		priority = *_priority
	}

	if closeErr := readBuffer.CloseContext("BACnetLightingCommand"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetLightingCommand")
	}

	// Create the instance
	return &_BACnetLightingCommand{
		LightningOperation: lightningOperation,
		TargetLevel:        targetLevel,
		RampRate:           rampRate,
		StepIncrement:      stepIncrement,
		FadeTime:           fadeTime,
		Priority:           priority,
	}, nil
}

func (m *_BACnetLightingCommand) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetLightingCommand) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetLightingCommand"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetLightingCommand")
	}

	// Simple Field (lightningOperation)
	if pushErr := writeBuffer.PushContext("lightningOperation"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for lightningOperation")
	}
	_lightningOperationErr := writeBuffer.WriteSerializable(ctx, m.GetLightningOperation())
	if popErr := writeBuffer.PopContext("lightningOperation"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for lightningOperation")
	}
	if _lightningOperationErr != nil {
		return errors.Wrap(_lightningOperationErr, "Error serializing 'lightningOperation' field")
	}

	// Optional Field (targetLevel) (Can be skipped, if the value is null)
	var targetLevel BACnetContextTagReal = nil
	if m.GetTargetLevel() != nil {
		if pushErr := writeBuffer.PushContext("targetLevel"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for targetLevel")
		}
		targetLevel = m.GetTargetLevel()
		_targetLevelErr := writeBuffer.WriteSerializable(ctx, targetLevel)
		if popErr := writeBuffer.PopContext("targetLevel"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for targetLevel")
		}
		if _targetLevelErr != nil {
			return errors.Wrap(_targetLevelErr, "Error serializing 'targetLevel' field")
		}
	}

	// Optional Field (rampRate) (Can be skipped, if the value is null)
	var rampRate BACnetContextTagReal = nil
	if m.GetRampRate() != nil {
		if pushErr := writeBuffer.PushContext("rampRate"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for rampRate")
		}
		rampRate = m.GetRampRate()
		_rampRateErr := writeBuffer.WriteSerializable(ctx, rampRate)
		if popErr := writeBuffer.PopContext("rampRate"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for rampRate")
		}
		if _rampRateErr != nil {
			return errors.Wrap(_rampRateErr, "Error serializing 'rampRate' field")
		}
	}

	// Optional Field (stepIncrement) (Can be skipped, if the value is null)
	var stepIncrement BACnetContextTagReal = nil
	if m.GetStepIncrement() != nil {
		if pushErr := writeBuffer.PushContext("stepIncrement"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for stepIncrement")
		}
		stepIncrement = m.GetStepIncrement()
		_stepIncrementErr := writeBuffer.WriteSerializable(ctx, stepIncrement)
		if popErr := writeBuffer.PopContext("stepIncrement"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for stepIncrement")
		}
		if _stepIncrementErr != nil {
			return errors.Wrap(_stepIncrementErr, "Error serializing 'stepIncrement' field")
		}
	}

	// Optional Field (fadeTime) (Can be skipped, if the value is null)
	var fadeTime BACnetContextTagUnsignedInteger = nil
	if m.GetFadeTime() != nil {
		if pushErr := writeBuffer.PushContext("fadeTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for fadeTime")
		}
		fadeTime = m.GetFadeTime()
		_fadeTimeErr := writeBuffer.WriteSerializable(ctx, fadeTime)
		if popErr := writeBuffer.PopContext("fadeTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for fadeTime")
		}
		if _fadeTimeErr != nil {
			return errors.Wrap(_fadeTimeErr, "Error serializing 'fadeTime' field")
		}
	}

	// Optional Field (priority) (Can be skipped, if the value is null)
	var priority BACnetContextTagUnsignedInteger = nil
	if m.GetPriority() != nil {
		if pushErr := writeBuffer.PushContext("priority"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for priority")
		}
		priority = m.GetPriority()
		_priorityErr := writeBuffer.WriteSerializable(ctx, priority)
		if popErr := writeBuffer.PopContext("priority"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for priority")
		}
		if _priorityErr != nil {
			return errors.Wrap(_priorityErr, "Error serializing 'priority' field")
		}
	}

	if popErr := writeBuffer.PopContext("BACnetLightingCommand"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetLightingCommand")
	}
	return nil
}

func (m *_BACnetLightingCommand) isBACnetLightingCommand() bool {
	return true
}

func (m *_BACnetLightingCommand) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
