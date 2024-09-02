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

// FindServersOnNetworkResponse is the corresponding interface of FindServersOnNetworkResponse
type FindServersOnNetworkResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetResponseHeader returns ResponseHeader (property field)
	GetResponseHeader() ExtensionObjectDefinition
	// GetLastCounterResetTime returns LastCounterResetTime (property field)
	GetLastCounterResetTime() int64
	// GetNoOfServers returns NoOfServers (property field)
	GetNoOfServers() int32
	// GetServers returns Servers (property field)
	GetServers() []ExtensionObjectDefinition
	// IsFindServersOnNetworkResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsFindServersOnNetworkResponse()
}

// _FindServersOnNetworkResponse is the data-structure of this message
type _FindServersOnNetworkResponse struct {
	ExtensionObjectDefinitionContract
	ResponseHeader       ExtensionObjectDefinition
	LastCounterResetTime int64
	NoOfServers          int32
	Servers              []ExtensionObjectDefinition
}

var _ FindServersOnNetworkResponse = (*_FindServersOnNetworkResponse)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_FindServersOnNetworkResponse)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_FindServersOnNetworkResponse) GetIdentifier() string {
	return "12193"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_FindServersOnNetworkResponse) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_FindServersOnNetworkResponse) GetResponseHeader() ExtensionObjectDefinition {
	return m.ResponseHeader
}

func (m *_FindServersOnNetworkResponse) GetLastCounterResetTime() int64 {
	return m.LastCounterResetTime
}

func (m *_FindServersOnNetworkResponse) GetNoOfServers() int32 {
	return m.NoOfServers
}

func (m *_FindServersOnNetworkResponse) GetServers() []ExtensionObjectDefinition {
	return m.Servers
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewFindServersOnNetworkResponse factory function for _FindServersOnNetworkResponse
func NewFindServersOnNetworkResponse(responseHeader ExtensionObjectDefinition, lastCounterResetTime int64, noOfServers int32, servers []ExtensionObjectDefinition) *_FindServersOnNetworkResponse {
	_result := &_FindServersOnNetworkResponse{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		ResponseHeader:                    responseHeader,
		LastCounterResetTime:              lastCounterResetTime,
		NoOfServers:                       noOfServers,
		Servers:                           servers,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastFindServersOnNetworkResponse(structType any) FindServersOnNetworkResponse {
	if casted, ok := structType.(FindServersOnNetworkResponse); ok {
		return casted
	}
	if casted, ok := structType.(*FindServersOnNetworkResponse); ok {
		return *casted
	}
	return nil
}

func (m *_FindServersOnNetworkResponse) GetTypeName() string {
	return "FindServersOnNetworkResponse"
}

func (m *_FindServersOnNetworkResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (responseHeader)
	lengthInBits += m.ResponseHeader.GetLengthInBits(ctx)

	// Simple field (lastCounterResetTime)
	lengthInBits += 64

	// Simple field (noOfServers)
	lengthInBits += 32

	// Array field
	if len(m.Servers) > 0 {
		for _curItem, element := range m.Servers {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.Servers), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_FindServersOnNetworkResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_FindServersOnNetworkResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__findServersOnNetworkResponse FindServersOnNetworkResponse, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("FindServersOnNetworkResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for FindServersOnNetworkResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	responseHeader, err := ReadSimpleField[ExtensionObjectDefinition](ctx, "responseHeader", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("394")), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'responseHeader' field"))
	}
	m.ResponseHeader = responseHeader

	lastCounterResetTime, err := ReadSimpleField(ctx, "lastCounterResetTime", ReadSignedLong(readBuffer, uint8(64)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lastCounterResetTime' field"))
	}
	m.LastCounterResetTime = lastCounterResetTime

	noOfServers, err := ReadSimpleField(ctx, "noOfServers", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfServers' field"))
	}
	m.NoOfServers = noOfServers

	servers, err := ReadCountArrayField[ExtensionObjectDefinition](ctx, "servers", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("12191")), readBuffer), uint64(noOfServers))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'servers' field"))
	}
	m.Servers = servers

	if closeErr := readBuffer.CloseContext("FindServersOnNetworkResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for FindServersOnNetworkResponse")
	}

	return m, nil
}

func (m *_FindServersOnNetworkResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_FindServersOnNetworkResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("FindServersOnNetworkResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for FindServersOnNetworkResponse")
		}

		if err := WriteSimpleField[ExtensionObjectDefinition](ctx, "responseHeader", m.GetResponseHeader(), WriteComplex[ExtensionObjectDefinition](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'responseHeader' field")
		}

		if err := WriteSimpleField[int64](ctx, "lastCounterResetTime", m.GetLastCounterResetTime(), WriteSignedLong(writeBuffer, 64)); err != nil {
			return errors.Wrap(err, "Error serializing 'lastCounterResetTime' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfServers", m.GetNoOfServers(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfServers' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "servers", m.GetServers(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'servers' field")
		}

		if popErr := writeBuffer.PopContext("FindServersOnNetworkResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for FindServersOnNetworkResponse")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_FindServersOnNetworkResponse) IsFindServersOnNetworkResponse() {}

func (m *_FindServersOnNetworkResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
