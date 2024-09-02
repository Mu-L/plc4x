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

// FindServersOnNetworkRequest is the corresponding interface of FindServersOnNetworkRequest
type FindServersOnNetworkRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetRequestHeader returns RequestHeader (property field)
	GetRequestHeader() ExtensionObjectDefinition
	// GetStartingRecordId returns StartingRecordId (property field)
	GetStartingRecordId() uint32
	// GetMaxRecordsToReturn returns MaxRecordsToReturn (property field)
	GetMaxRecordsToReturn() uint32
	// GetNoOfServerCapabilityFilter returns NoOfServerCapabilityFilter (property field)
	GetNoOfServerCapabilityFilter() int32
	// GetServerCapabilityFilter returns ServerCapabilityFilter (property field)
	GetServerCapabilityFilter() []PascalString
	// IsFindServersOnNetworkRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsFindServersOnNetworkRequest()
}

// _FindServersOnNetworkRequest is the data-structure of this message
type _FindServersOnNetworkRequest struct {
	ExtensionObjectDefinitionContract
	RequestHeader              ExtensionObjectDefinition
	StartingRecordId           uint32
	MaxRecordsToReturn         uint32
	NoOfServerCapabilityFilter int32
	ServerCapabilityFilter     []PascalString
}

var _ FindServersOnNetworkRequest = (*_FindServersOnNetworkRequest)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_FindServersOnNetworkRequest)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_FindServersOnNetworkRequest) GetIdentifier() string {
	return "12192"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_FindServersOnNetworkRequest) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_FindServersOnNetworkRequest) GetRequestHeader() ExtensionObjectDefinition {
	return m.RequestHeader
}

func (m *_FindServersOnNetworkRequest) GetStartingRecordId() uint32 {
	return m.StartingRecordId
}

func (m *_FindServersOnNetworkRequest) GetMaxRecordsToReturn() uint32 {
	return m.MaxRecordsToReturn
}

func (m *_FindServersOnNetworkRequest) GetNoOfServerCapabilityFilter() int32 {
	return m.NoOfServerCapabilityFilter
}

func (m *_FindServersOnNetworkRequest) GetServerCapabilityFilter() []PascalString {
	return m.ServerCapabilityFilter
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewFindServersOnNetworkRequest factory function for _FindServersOnNetworkRequest
func NewFindServersOnNetworkRequest(requestHeader ExtensionObjectDefinition, startingRecordId uint32, maxRecordsToReturn uint32, noOfServerCapabilityFilter int32, serverCapabilityFilter []PascalString) *_FindServersOnNetworkRequest {
	_result := &_FindServersOnNetworkRequest{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		RequestHeader:                     requestHeader,
		StartingRecordId:                  startingRecordId,
		MaxRecordsToReturn:                maxRecordsToReturn,
		NoOfServerCapabilityFilter:        noOfServerCapabilityFilter,
		ServerCapabilityFilter:            serverCapabilityFilter,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastFindServersOnNetworkRequest(structType any) FindServersOnNetworkRequest {
	if casted, ok := structType.(FindServersOnNetworkRequest); ok {
		return casted
	}
	if casted, ok := structType.(*FindServersOnNetworkRequest); ok {
		return *casted
	}
	return nil
}

func (m *_FindServersOnNetworkRequest) GetTypeName() string {
	return "FindServersOnNetworkRequest"
}

func (m *_FindServersOnNetworkRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (requestHeader)
	lengthInBits += m.RequestHeader.GetLengthInBits(ctx)

	// Simple field (startingRecordId)
	lengthInBits += 32

	// Simple field (maxRecordsToReturn)
	lengthInBits += 32

	// Simple field (noOfServerCapabilityFilter)
	lengthInBits += 32

	// Array field
	if len(m.ServerCapabilityFilter) > 0 {
		for _curItem, element := range m.ServerCapabilityFilter {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.ServerCapabilityFilter), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_FindServersOnNetworkRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_FindServersOnNetworkRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__findServersOnNetworkRequest FindServersOnNetworkRequest, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("FindServersOnNetworkRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for FindServersOnNetworkRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	requestHeader, err := ReadSimpleField[ExtensionObjectDefinition](ctx, "requestHeader", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("391")), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestHeader' field"))
	}
	m.RequestHeader = requestHeader

	startingRecordId, err := ReadSimpleField(ctx, "startingRecordId", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'startingRecordId' field"))
	}
	m.StartingRecordId = startingRecordId

	maxRecordsToReturn, err := ReadSimpleField(ctx, "maxRecordsToReturn", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'maxRecordsToReturn' field"))
	}
	m.MaxRecordsToReturn = maxRecordsToReturn

	noOfServerCapabilityFilter, err := ReadSimpleField(ctx, "noOfServerCapabilityFilter", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfServerCapabilityFilter' field"))
	}
	m.NoOfServerCapabilityFilter = noOfServerCapabilityFilter

	serverCapabilityFilter, err := ReadCountArrayField[PascalString](ctx, "serverCapabilityFilter", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer), uint64(noOfServerCapabilityFilter))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'serverCapabilityFilter' field"))
	}
	m.ServerCapabilityFilter = serverCapabilityFilter

	if closeErr := readBuffer.CloseContext("FindServersOnNetworkRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for FindServersOnNetworkRequest")
	}

	return m, nil
}

func (m *_FindServersOnNetworkRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_FindServersOnNetworkRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("FindServersOnNetworkRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for FindServersOnNetworkRequest")
		}

		if err := WriteSimpleField[ExtensionObjectDefinition](ctx, "requestHeader", m.GetRequestHeader(), WriteComplex[ExtensionObjectDefinition](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'requestHeader' field")
		}

		if err := WriteSimpleField[uint32](ctx, "startingRecordId", m.GetStartingRecordId(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'startingRecordId' field")
		}

		if err := WriteSimpleField[uint32](ctx, "maxRecordsToReturn", m.GetMaxRecordsToReturn(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'maxRecordsToReturn' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfServerCapabilityFilter", m.GetNoOfServerCapabilityFilter(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfServerCapabilityFilter' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "serverCapabilityFilter", m.GetServerCapabilityFilter(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'serverCapabilityFilter' field")
		}

		if popErr := writeBuffer.PopContext("FindServersOnNetworkRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for FindServersOnNetworkRequest")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_FindServersOnNetworkRequest) IsFindServersOnNetworkRequest() {}

func (m *_FindServersOnNetworkRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
