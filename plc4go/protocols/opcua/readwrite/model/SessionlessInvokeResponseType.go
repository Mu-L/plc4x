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

// SessionlessInvokeResponseType is the corresponding interface of SessionlessInvokeResponseType
type SessionlessInvokeResponseType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetNoOfNamespaceUris returns NoOfNamespaceUris (property field)
	GetNoOfNamespaceUris() int32
	// GetNamespaceUris returns NamespaceUris (property field)
	GetNamespaceUris() []PascalString
	// GetNoOfServerUris returns NoOfServerUris (property field)
	GetNoOfServerUris() int32
	// GetServerUris returns ServerUris (property field)
	GetServerUris() []PascalString
	// GetServiceId returns ServiceId (property field)
	GetServiceId() uint32
}

// SessionlessInvokeResponseTypeExactly can be used when we want exactly this type and not a type which fulfills SessionlessInvokeResponseType.
// This is useful for switch cases.
type SessionlessInvokeResponseTypeExactly interface {
	SessionlessInvokeResponseType
	isSessionlessInvokeResponseType() bool
}

// _SessionlessInvokeResponseType is the data-structure of this message
type _SessionlessInvokeResponseType struct {
	*_ExtensionObjectDefinition
	NoOfNamespaceUris int32
	NamespaceUris     []PascalString
	NoOfServerUris    int32
	ServerUris        []PascalString
	ServiceId         uint32
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SessionlessInvokeResponseType) GetIdentifier() string {
	return "21001"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SessionlessInvokeResponseType) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_SessionlessInvokeResponseType) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SessionlessInvokeResponseType) GetNoOfNamespaceUris() int32 {
	return m.NoOfNamespaceUris
}

func (m *_SessionlessInvokeResponseType) GetNamespaceUris() []PascalString {
	return m.NamespaceUris
}

func (m *_SessionlessInvokeResponseType) GetNoOfServerUris() int32 {
	return m.NoOfServerUris
}

func (m *_SessionlessInvokeResponseType) GetServerUris() []PascalString {
	return m.ServerUris
}

func (m *_SessionlessInvokeResponseType) GetServiceId() uint32 {
	return m.ServiceId
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewSessionlessInvokeResponseType factory function for _SessionlessInvokeResponseType
func NewSessionlessInvokeResponseType(noOfNamespaceUris int32, namespaceUris []PascalString, noOfServerUris int32, serverUris []PascalString, serviceId uint32) *_SessionlessInvokeResponseType {
	_result := &_SessionlessInvokeResponseType{
		NoOfNamespaceUris:          noOfNamespaceUris,
		NamespaceUris:              namespaceUris,
		NoOfServerUris:             noOfServerUris,
		ServerUris:                 serverUris,
		ServiceId:                  serviceId,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSessionlessInvokeResponseType(structType any) SessionlessInvokeResponseType {
	if casted, ok := structType.(SessionlessInvokeResponseType); ok {
		return casted
	}
	if casted, ok := structType.(*SessionlessInvokeResponseType); ok {
		return *casted
	}
	return nil
}

func (m *_SessionlessInvokeResponseType) GetTypeName() string {
	return "SessionlessInvokeResponseType"
}

func (m *_SessionlessInvokeResponseType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (noOfNamespaceUris)
	lengthInBits += 32

	// Array field
	if len(m.NamespaceUris) > 0 {
		for _curItem, element := range m.NamespaceUris {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.NamespaceUris), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (noOfServerUris)
	lengthInBits += 32

	// Array field
	if len(m.ServerUris) > 0 {
		for _curItem, element := range m.ServerUris {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.ServerUris), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (serviceId)
	lengthInBits += 32

	return lengthInBits
}

func (m *_SessionlessInvokeResponseType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func SessionlessInvokeResponseTypeParse(ctx context.Context, theBytes []byte, identifier string) (SessionlessInvokeResponseType, error) {
	return SessionlessInvokeResponseTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func SessionlessInvokeResponseTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (SessionlessInvokeResponseType, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("SessionlessInvokeResponseType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SessionlessInvokeResponseType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (noOfNamespaceUris)
	_noOfNamespaceUris, _noOfNamespaceUrisErr := readBuffer.ReadInt32("noOfNamespaceUris", 32)
	if _noOfNamespaceUrisErr != nil {
		return nil, errors.Wrap(_noOfNamespaceUrisErr, "Error parsing 'noOfNamespaceUris' field of SessionlessInvokeResponseType")
	}
	noOfNamespaceUris := _noOfNamespaceUris

	// Array field (namespaceUris)
	if pullErr := readBuffer.PullContext("namespaceUris", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for namespaceUris")
	}
	// Count array
	namespaceUris := make([]PascalString, noOfNamespaceUris)
	// This happens when the size is set conditional to 0
	if len(namespaceUris) == 0 {
		namespaceUris = nil
	}
	{
		_numItems := uint16(noOfNamespaceUris)
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
			_item, _err := PascalStringParseWithBuffer(arrayCtx, readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'namespaceUris' field of SessionlessInvokeResponseType")
			}
			namespaceUris[_curItem] = _item.(PascalString)
		}
	}
	if closeErr := readBuffer.CloseContext("namespaceUris", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for namespaceUris")
	}

	// Simple Field (noOfServerUris)
	_noOfServerUris, _noOfServerUrisErr := readBuffer.ReadInt32("noOfServerUris", 32)
	if _noOfServerUrisErr != nil {
		return nil, errors.Wrap(_noOfServerUrisErr, "Error parsing 'noOfServerUris' field of SessionlessInvokeResponseType")
	}
	noOfServerUris := _noOfServerUris

	// Array field (serverUris)
	if pullErr := readBuffer.PullContext("serverUris", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for serverUris")
	}
	// Count array
	serverUris := make([]PascalString, noOfServerUris)
	// This happens when the size is set conditional to 0
	if len(serverUris) == 0 {
		serverUris = nil
	}
	{
		_numItems := uint16(noOfServerUris)
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
			_item, _err := PascalStringParseWithBuffer(arrayCtx, readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'serverUris' field of SessionlessInvokeResponseType")
			}
			serverUris[_curItem] = _item.(PascalString)
		}
	}
	if closeErr := readBuffer.CloseContext("serverUris", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for serverUris")
	}

	// Simple Field (serviceId)
	_serviceId, _serviceIdErr := readBuffer.ReadUint32("serviceId", 32)
	if _serviceIdErr != nil {
		return nil, errors.Wrap(_serviceIdErr, "Error parsing 'serviceId' field of SessionlessInvokeResponseType")
	}
	serviceId := _serviceId

	if closeErr := readBuffer.CloseContext("SessionlessInvokeResponseType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SessionlessInvokeResponseType")
	}

	// Create a partially initialized instance
	_child := &_SessionlessInvokeResponseType{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		NoOfNamespaceUris:          noOfNamespaceUris,
		NamespaceUris:              namespaceUris,
		NoOfServerUris:             noOfServerUris,
		ServerUris:                 serverUris,
		ServiceId:                  serviceId,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_SessionlessInvokeResponseType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SessionlessInvokeResponseType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SessionlessInvokeResponseType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SessionlessInvokeResponseType")
		}

		// Simple Field (noOfNamespaceUris)
		noOfNamespaceUris := int32(m.GetNoOfNamespaceUris())
		_noOfNamespaceUrisErr := writeBuffer.WriteInt32("noOfNamespaceUris", 32, (noOfNamespaceUris))
		if _noOfNamespaceUrisErr != nil {
			return errors.Wrap(_noOfNamespaceUrisErr, "Error serializing 'noOfNamespaceUris' field")
		}

		// Array Field (namespaceUris)
		if pushErr := writeBuffer.PushContext("namespaceUris", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for namespaceUris")
		}
		for _curItem, _element := range m.GetNamespaceUris() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetNamespaceUris()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'namespaceUris' field")
			}
		}
		if popErr := writeBuffer.PopContext("namespaceUris", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for namespaceUris")
		}

		// Simple Field (noOfServerUris)
		noOfServerUris := int32(m.GetNoOfServerUris())
		_noOfServerUrisErr := writeBuffer.WriteInt32("noOfServerUris", 32, (noOfServerUris))
		if _noOfServerUrisErr != nil {
			return errors.Wrap(_noOfServerUrisErr, "Error serializing 'noOfServerUris' field")
		}

		// Array Field (serverUris)
		if pushErr := writeBuffer.PushContext("serverUris", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for serverUris")
		}
		for _curItem, _element := range m.GetServerUris() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetServerUris()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'serverUris' field")
			}
		}
		if popErr := writeBuffer.PopContext("serverUris", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for serverUris")
		}

		// Simple Field (serviceId)
		serviceId := uint32(m.GetServiceId())
		_serviceIdErr := writeBuffer.WriteUint32("serviceId", 32, (serviceId))
		if _serviceIdErr != nil {
			return errors.Wrap(_serviceIdErr, "Error serializing 'serviceId' field")
		}

		if popErr := writeBuffer.PopContext("SessionlessInvokeResponseType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SessionlessInvokeResponseType")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SessionlessInvokeResponseType) isSessionlessInvokeResponseType() bool {
	return true
}

func (m *_SessionlessInvokeResponseType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
