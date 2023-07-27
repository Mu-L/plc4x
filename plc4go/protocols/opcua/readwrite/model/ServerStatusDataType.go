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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// ServerStatusDataType is the corresponding interface of ServerStatusDataType
type ServerStatusDataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetStartTime returns StartTime (property field)
	GetStartTime() int64
	// GetCurrentTime returns CurrentTime (property field)
	GetCurrentTime() int64
	// GetState returns State (property field)
	GetState() ServerState
	// GetBuildInfo returns BuildInfo (property field)
	GetBuildInfo() ExtensionObjectDefinition
	// GetSecondsTillShutdown returns SecondsTillShutdown (property field)
	GetSecondsTillShutdown() uint32
	// GetShutdownReason returns ShutdownReason (property field)
	GetShutdownReason() LocalizedText
}

// ServerStatusDataTypeExactly can be used when we want exactly this type and not a type which fulfills ServerStatusDataType.
// This is useful for switch cases.
type ServerStatusDataTypeExactly interface {
	ServerStatusDataType
	isServerStatusDataType() bool
}

// _ServerStatusDataType is the data-structure of this message
type _ServerStatusDataType struct {
	*_ExtensionObjectDefinition
	StartTime           int64
	CurrentTime         int64
	State               ServerState
	BuildInfo           ExtensionObjectDefinition
	SecondsTillShutdown uint32
	ShutdownReason      LocalizedText
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ServerStatusDataType) GetIdentifier() string {
	return "864"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ServerStatusDataType) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_ServerStatusDataType) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ServerStatusDataType) GetStartTime() int64 {
	return m.StartTime
}

func (m *_ServerStatusDataType) GetCurrentTime() int64 {
	return m.CurrentTime
}

func (m *_ServerStatusDataType) GetState() ServerState {
	return m.State
}

func (m *_ServerStatusDataType) GetBuildInfo() ExtensionObjectDefinition {
	return m.BuildInfo
}

func (m *_ServerStatusDataType) GetSecondsTillShutdown() uint32 {
	return m.SecondsTillShutdown
}

func (m *_ServerStatusDataType) GetShutdownReason() LocalizedText {
	return m.ShutdownReason
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewServerStatusDataType factory function for _ServerStatusDataType
func NewServerStatusDataType(startTime int64, currentTime int64, state ServerState, buildInfo ExtensionObjectDefinition, secondsTillShutdown uint32, shutdownReason LocalizedText) *_ServerStatusDataType {
	_result := &_ServerStatusDataType{
		StartTime:                  startTime,
		CurrentTime:                currentTime,
		State:                      state,
		BuildInfo:                  buildInfo,
		SecondsTillShutdown:        secondsTillShutdown,
		ShutdownReason:             shutdownReason,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastServerStatusDataType(structType any) ServerStatusDataType {
	if casted, ok := structType.(ServerStatusDataType); ok {
		return casted
	}
	if casted, ok := structType.(*ServerStatusDataType); ok {
		return *casted
	}
	return nil
}

func (m *_ServerStatusDataType) GetTypeName() string {
	return "ServerStatusDataType"
}

func (m *_ServerStatusDataType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (startTime)
	lengthInBits += 64

	// Simple field (currentTime)
	lengthInBits += 64

	// Simple field (state)
	lengthInBits += 32

	// Simple field (buildInfo)
	lengthInBits += m.BuildInfo.GetLengthInBits(ctx)

	// Simple field (secondsTillShutdown)
	lengthInBits += 32

	// Simple field (shutdownReason)
	lengthInBits += m.ShutdownReason.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_ServerStatusDataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ServerStatusDataTypeParse(ctx context.Context, theBytes []byte, identifier string) (ServerStatusDataType, error) {
	return ServerStatusDataTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func ServerStatusDataTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (ServerStatusDataType, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("ServerStatusDataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ServerStatusDataType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (startTime)
	_startTime, _startTimeErr := readBuffer.ReadInt64("startTime", 64)
	if _startTimeErr != nil {
		return nil, errors.Wrap(_startTimeErr, "Error parsing 'startTime' field of ServerStatusDataType")
	}
	startTime := _startTime

	// Simple Field (currentTime)
	_currentTime, _currentTimeErr := readBuffer.ReadInt64("currentTime", 64)
	if _currentTimeErr != nil {
		return nil, errors.Wrap(_currentTimeErr, "Error parsing 'currentTime' field of ServerStatusDataType")
	}
	currentTime := _currentTime

	// Simple Field (state)
	if pullErr := readBuffer.PullContext("state"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for state")
	}
	_state, _stateErr := ServerStateParseWithBuffer(ctx, readBuffer)
	if _stateErr != nil {
		return nil, errors.Wrap(_stateErr, "Error parsing 'state' field of ServerStatusDataType")
	}
	state := _state
	if closeErr := readBuffer.CloseContext("state"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for state")
	}

	// Simple Field (buildInfo)
	if pullErr := readBuffer.PullContext("buildInfo"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for buildInfo")
	}
	_buildInfo, _buildInfoErr := ExtensionObjectDefinitionParseWithBuffer(ctx, readBuffer, string("340"))
	if _buildInfoErr != nil {
		return nil, errors.Wrap(_buildInfoErr, "Error parsing 'buildInfo' field of ServerStatusDataType")
	}
	buildInfo := _buildInfo.(ExtensionObjectDefinition)
	if closeErr := readBuffer.CloseContext("buildInfo"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for buildInfo")
	}

	// Simple Field (secondsTillShutdown)
	_secondsTillShutdown, _secondsTillShutdownErr := readBuffer.ReadUint32("secondsTillShutdown", 32)
	if _secondsTillShutdownErr != nil {
		return nil, errors.Wrap(_secondsTillShutdownErr, "Error parsing 'secondsTillShutdown' field of ServerStatusDataType")
	}
	secondsTillShutdown := _secondsTillShutdown

	// Simple Field (shutdownReason)
	if pullErr := readBuffer.PullContext("shutdownReason"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for shutdownReason")
	}
	_shutdownReason, _shutdownReasonErr := LocalizedTextParseWithBuffer(ctx, readBuffer)
	if _shutdownReasonErr != nil {
		return nil, errors.Wrap(_shutdownReasonErr, "Error parsing 'shutdownReason' field of ServerStatusDataType")
	}
	shutdownReason := _shutdownReason.(LocalizedText)
	if closeErr := readBuffer.CloseContext("shutdownReason"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for shutdownReason")
	}

	if closeErr := readBuffer.CloseContext("ServerStatusDataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ServerStatusDataType")
	}

	// Create a partially initialized instance
	_child := &_ServerStatusDataType{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		StartTime:                  startTime,
		CurrentTime:                currentTime,
		State:                      state,
		BuildInfo:                  buildInfo,
		SecondsTillShutdown:        secondsTillShutdown,
		ShutdownReason:             shutdownReason,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_ServerStatusDataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ServerStatusDataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ServerStatusDataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ServerStatusDataType")
		}

		// Simple Field (startTime)
		startTime := int64(m.GetStartTime())
		_startTimeErr := writeBuffer.WriteInt64("startTime", 64, (startTime))
		if _startTimeErr != nil {
			return errors.Wrap(_startTimeErr, "Error serializing 'startTime' field")
		}

		// Simple Field (currentTime)
		currentTime := int64(m.GetCurrentTime())
		_currentTimeErr := writeBuffer.WriteInt64("currentTime", 64, (currentTime))
		if _currentTimeErr != nil {
			return errors.Wrap(_currentTimeErr, "Error serializing 'currentTime' field")
		}

		// Simple Field (state)
		if pushErr := writeBuffer.PushContext("state"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for state")
		}
		_stateErr := writeBuffer.WriteSerializable(ctx, m.GetState())
		if popErr := writeBuffer.PopContext("state"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for state")
		}
		if _stateErr != nil {
			return errors.Wrap(_stateErr, "Error serializing 'state' field")
		}

		// Simple Field (buildInfo)
		if pushErr := writeBuffer.PushContext("buildInfo"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for buildInfo")
		}
		_buildInfoErr := writeBuffer.WriteSerializable(ctx, m.GetBuildInfo())
		if popErr := writeBuffer.PopContext("buildInfo"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for buildInfo")
		}
		if _buildInfoErr != nil {
			return errors.Wrap(_buildInfoErr, "Error serializing 'buildInfo' field")
		}

		// Simple Field (secondsTillShutdown)
		secondsTillShutdown := uint32(m.GetSecondsTillShutdown())
		_secondsTillShutdownErr := writeBuffer.WriteUint32("secondsTillShutdown", 32, (secondsTillShutdown))
		if _secondsTillShutdownErr != nil {
			return errors.Wrap(_secondsTillShutdownErr, "Error serializing 'secondsTillShutdown' field")
		}

		// Simple Field (shutdownReason)
		if pushErr := writeBuffer.PushContext("shutdownReason"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for shutdownReason")
		}
		_shutdownReasonErr := writeBuffer.WriteSerializable(ctx, m.GetShutdownReason())
		if popErr := writeBuffer.PopContext("shutdownReason"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for shutdownReason")
		}
		if _shutdownReasonErr != nil {
			return errors.Wrap(_shutdownReasonErr, "Error serializing 'shutdownReason' field")
		}

		if popErr := writeBuffer.PopContext("ServerStatusDataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ServerStatusDataType")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ServerStatusDataType) isServerStatusDataType() bool {
	return true
}

func (m *_ServerStatusDataType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
