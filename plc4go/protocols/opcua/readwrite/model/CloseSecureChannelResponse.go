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

// CloseSecureChannelResponse is the corresponding interface of CloseSecureChannelResponse
type CloseSecureChannelResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetResponseHeader returns ResponseHeader (property field)
	GetResponseHeader() ExtensionObjectDefinition
}

// CloseSecureChannelResponseExactly can be used when we want exactly this type and not a type which fulfills CloseSecureChannelResponse.
// This is useful for switch cases.
type CloseSecureChannelResponseExactly interface {
	CloseSecureChannelResponse
	isCloseSecureChannelResponse() bool
}

// _CloseSecureChannelResponse is the data-structure of this message
type _CloseSecureChannelResponse struct {
	ExtensionObjectDefinitionContract
	ResponseHeader ExtensionObjectDefinition
}

var _ CloseSecureChannelResponse = (*_CloseSecureChannelResponse)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_CloseSecureChannelResponse)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_CloseSecureChannelResponse) GetIdentifier() string {
	return "455"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CloseSecureChannelResponse) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CloseSecureChannelResponse) GetResponseHeader() ExtensionObjectDefinition {
	return m.ResponseHeader
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCloseSecureChannelResponse factory function for _CloseSecureChannelResponse
func NewCloseSecureChannelResponse(responseHeader ExtensionObjectDefinition) *_CloseSecureChannelResponse {
	_result := &_CloseSecureChannelResponse{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		ResponseHeader:                    responseHeader,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastCloseSecureChannelResponse(structType any) CloseSecureChannelResponse {
	if casted, ok := structType.(CloseSecureChannelResponse); ok {
		return casted
	}
	if casted, ok := structType.(*CloseSecureChannelResponse); ok {
		return *casted
	}
	return nil
}

func (m *_CloseSecureChannelResponse) GetTypeName() string {
	return "CloseSecureChannelResponse"
}

func (m *_CloseSecureChannelResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (responseHeader)
	lengthInBits += m.ResponseHeader.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_CloseSecureChannelResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_CloseSecureChannelResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__closeSecureChannelResponse CloseSecureChannelResponse, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CloseSecureChannelResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CloseSecureChannelResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	responseHeader, err := ReadSimpleField[ExtensionObjectDefinition](ctx, "responseHeader", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("394")), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'responseHeader' field"))
	}
	m.ResponseHeader = responseHeader

	if closeErr := readBuffer.CloseContext("CloseSecureChannelResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CloseSecureChannelResponse")
	}

	return m, nil
}

func (m *_CloseSecureChannelResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CloseSecureChannelResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CloseSecureChannelResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CloseSecureChannelResponse")
		}

		if err := WriteSimpleField[ExtensionObjectDefinition](ctx, "responseHeader", m.GetResponseHeader(), WriteComplex[ExtensionObjectDefinition](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'responseHeader' field")
		}

		if popErr := writeBuffer.PopContext("CloseSecureChannelResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CloseSecureChannelResponse")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CloseSecureChannelResponse) isCloseSecureChannelResponse() bool {
	return true
}

func (m *_CloseSecureChannelResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
