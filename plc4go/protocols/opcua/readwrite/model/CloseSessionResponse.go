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

// CloseSessionResponse is the corresponding interface of CloseSessionResponse
type CloseSessionResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetResponseHeader returns ResponseHeader (property field)
	GetResponseHeader() ExtensionObjectDefinition
}

// CloseSessionResponseExactly can be used when we want exactly this type and not a type which fulfills CloseSessionResponse.
// This is useful for switch cases.
type CloseSessionResponseExactly interface {
	CloseSessionResponse
	isCloseSessionResponse() bool
}

// _CloseSessionResponse is the data-structure of this message
type _CloseSessionResponse struct {
	ExtensionObjectDefinitionContract
	ResponseHeader ExtensionObjectDefinition
}

var _ CloseSessionResponse = (*_CloseSessionResponse)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_CloseSessionResponse)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_CloseSessionResponse) GetIdentifier() string {
	return "476"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CloseSessionResponse) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CloseSessionResponse) GetResponseHeader() ExtensionObjectDefinition {
	return m.ResponseHeader
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCloseSessionResponse factory function for _CloseSessionResponse
func NewCloseSessionResponse(responseHeader ExtensionObjectDefinition) *_CloseSessionResponse {
	_result := &_CloseSessionResponse{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		ResponseHeader:                    responseHeader,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastCloseSessionResponse(structType any) CloseSessionResponse {
	if casted, ok := structType.(CloseSessionResponse); ok {
		return casted
	}
	if casted, ok := structType.(*CloseSessionResponse); ok {
		return *casted
	}
	return nil
}

func (m *_CloseSessionResponse) GetTypeName() string {
	return "CloseSessionResponse"
}

func (m *_CloseSessionResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (responseHeader)
	lengthInBits += m.ResponseHeader.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_CloseSessionResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_CloseSessionResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__closeSessionResponse CloseSessionResponse, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CloseSessionResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CloseSessionResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	responseHeader, err := ReadSimpleField[ExtensionObjectDefinition](ctx, "responseHeader", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("394")), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'responseHeader' field"))
	}
	m.ResponseHeader = responseHeader

	if closeErr := readBuffer.CloseContext("CloseSessionResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CloseSessionResponse")
	}

	return m, nil
}

func (m *_CloseSessionResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CloseSessionResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CloseSessionResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CloseSessionResponse")
		}

		if err := WriteSimpleField[ExtensionObjectDefinition](ctx, "responseHeader", m.GetResponseHeader(), WriteComplex[ExtensionObjectDefinition](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'responseHeader' field")
		}

		if popErr := writeBuffer.PopContext("CloseSessionResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CloseSessionResponse")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CloseSessionResponse) isCloseSessionResponse() bool {
	return true
}

func (m *_CloseSessionResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
