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

// MonitoredItemModifyRequest is the corresponding interface of MonitoredItemModifyRequest
type MonitoredItemModifyRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetMonitoredItemId returns MonitoredItemId (property field)
	GetMonitoredItemId() uint32
	// GetRequestedParameters returns RequestedParameters (property field)
	GetRequestedParameters() ExtensionObjectDefinition
}

// MonitoredItemModifyRequestExactly can be used when we want exactly this type and not a type which fulfills MonitoredItemModifyRequest.
// This is useful for switch cases.
type MonitoredItemModifyRequestExactly interface {
	MonitoredItemModifyRequest
	isMonitoredItemModifyRequest() bool
}

// _MonitoredItemModifyRequest is the data-structure of this message
type _MonitoredItemModifyRequest struct {
	ExtensionObjectDefinitionContract
	MonitoredItemId     uint32
	RequestedParameters ExtensionObjectDefinition
}

var _ MonitoredItemModifyRequest = (*_MonitoredItemModifyRequest)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_MonitoredItemModifyRequest)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_MonitoredItemModifyRequest) GetIdentifier() string {
	return "757"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_MonitoredItemModifyRequest) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_MonitoredItemModifyRequest) GetMonitoredItemId() uint32 {
	return m.MonitoredItemId
}

func (m *_MonitoredItemModifyRequest) GetRequestedParameters() ExtensionObjectDefinition {
	return m.RequestedParameters
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewMonitoredItemModifyRequest factory function for _MonitoredItemModifyRequest
func NewMonitoredItemModifyRequest(monitoredItemId uint32, requestedParameters ExtensionObjectDefinition) *_MonitoredItemModifyRequest {
	_result := &_MonitoredItemModifyRequest{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		MonitoredItemId:                   monitoredItemId,
		RequestedParameters:               requestedParameters,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastMonitoredItemModifyRequest(structType any) MonitoredItemModifyRequest {
	if casted, ok := structType.(MonitoredItemModifyRequest); ok {
		return casted
	}
	if casted, ok := structType.(*MonitoredItemModifyRequest); ok {
		return *casted
	}
	return nil
}

func (m *_MonitoredItemModifyRequest) GetTypeName() string {
	return "MonitoredItemModifyRequest"
}

func (m *_MonitoredItemModifyRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (monitoredItemId)
	lengthInBits += 32

	// Simple field (requestedParameters)
	lengthInBits += m.RequestedParameters.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_MonitoredItemModifyRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_MonitoredItemModifyRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__monitoredItemModifyRequest MonitoredItemModifyRequest, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MonitoredItemModifyRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MonitoredItemModifyRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	monitoredItemId, err := ReadSimpleField(ctx, "monitoredItemId", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'monitoredItemId' field"))
	}
	m.MonitoredItemId = monitoredItemId

	requestedParameters, err := ReadSimpleField[ExtensionObjectDefinition](ctx, "requestedParameters", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("742")), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestedParameters' field"))
	}
	m.RequestedParameters = requestedParameters

	if closeErr := readBuffer.CloseContext("MonitoredItemModifyRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MonitoredItemModifyRequest")
	}

	return m, nil
}

func (m *_MonitoredItemModifyRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MonitoredItemModifyRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MonitoredItemModifyRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MonitoredItemModifyRequest")
		}

		if err := WriteSimpleField[uint32](ctx, "monitoredItemId", m.GetMonitoredItemId(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'monitoredItemId' field")
		}

		if err := WriteSimpleField[ExtensionObjectDefinition](ctx, "requestedParameters", m.GetRequestedParameters(), WriteComplex[ExtensionObjectDefinition](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'requestedParameters' field")
		}

		if popErr := writeBuffer.PopContext("MonitoredItemModifyRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MonitoredItemModifyRequest")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_MonitoredItemModifyRequest) isMonitoredItemModifyRequest() bool {
	return true
}

func (m *_MonitoredItemModifyRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
