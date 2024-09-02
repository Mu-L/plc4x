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

// NodeIdByteString is the corresponding interface of NodeIdByteString
type NodeIdByteString interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	NodeIdTypeDefinition
	// GetNamespaceIndex returns NamespaceIndex (property field)
	GetNamespaceIndex() uint16
	// GetId returns Id (property field)
	GetId() PascalByteString
	// GetIdentifier returns Identifier (virtual field)
	GetIdentifier() string
}

// NodeIdByteStringExactly can be used when we want exactly this type and not a type which fulfills NodeIdByteString.
// This is useful for switch cases.
type NodeIdByteStringExactly interface {
	NodeIdByteString
	isNodeIdByteString() bool
}

// _NodeIdByteString is the data-structure of this message
type _NodeIdByteString struct {
	NodeIdTypeDefinitionContract
	NamespaceIndex uint16
	Id             PascalByteString
}

var _ NodeIdByteString = (*_NodeIdByteString)(nil)
var _ NodeIdTypeDefinitionRequirements = (*_NodeIdByteString)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_NodeIdByteString) GetNodeType() NodeIdType {
	return NodeIdType_nodeIdTypeByteString
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_NodeIdByteString) GetParent() NodeIdTypeDefinitionContract {
	return m.NodeIdTypeDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_NodeIdByteString) GetNamespaceIndex() uint16 {
	return m.NamespaceIndex
}

func (m *_NodeIdByteString) GetId() PascalByteString {
	return m.Id
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_NodeIdByteString) GetIdentifier() string {
	ctx := context.Background()
	_ = ctx
	return fmt.Sprintf("%v", m.GetId().GetStringValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewNodeIdByteString factory function for _NodeIdByteString
func NewNodeIdByteString(namespaceIndex uint16, id PascalByteString) *_NodeIdByteString {
	_result := &_NodeIdByteString{
		NodeIdTypeDefinitionContract: NewNodeIdTypeDefinition(),
		NamespaceIndex:               namespaceIndex,
		Id:                           id,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastNodeIdByteString(structType any) NodeIdByteString {
	if casted, ok := structType.(NodeIdByteString); ok {
		return casted
	}
	if casted, ok := structType.(*NodeIdByteString); ok {
		return *casted
	}
	return nil
}

func (m *_NodeIdByteString) GetTypeName() string {
	return "NodeIdByteString"
}

func (m *_NodeIdByteString) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.NodeIdTypeDefinitionContract.(*_NodeIdTypeDefinition).getLengthInBits(ctx))

	// Simple field (namespaceIndex)
	lengthInBits += 16

	// Simple field (id)
	lengthInBits += m.Id.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_NodeIdByteString) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_NodeIdByteString) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_NodeIdTypeDefinition) (__nodeIdByteString NodeIdByteString, err error) {
	m.NodeIdTypeDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("NodeIdByteString"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NodeIdByteString")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	namespaceIndex, err := ReadSimpleField(ctx, "namespaceIndex", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'namespaceIndex' field"))
	}
	m.NamespaceIndex = namespaceIndex

	id, err := ReadSimpleField[PascalByteString](ctx, "id", ReadComplex[PascalByteString](PascalByteStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'id' field"))
	}
	m.Id = id

	identifier, err := ReadVirtualField[string](ctx, "identifier", (*string)(nil), id.GetStringValue())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'identifier' field"))
	}
	_ = identifier

	if closeErr := readBuffer.CloseContext("NodeIdByteString"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NodeIdByteString")
	}

	return m, nil
}

func (m *_NodeIdByteString) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_NodeIdByteString) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("NodeIdByteString"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for NodeIdByteString")
		}

		if err := WriteSimpleField[uint16](ctx, "namespaceIndex", m.GetNamespaceIndex(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'namespaceIndex' field")
		}

		if err := WriteSimpleField[PascalByteString](ctx, "id", m.GetId(), WriteComplex[PascalByteString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'id' field")
		}
		// Virtual field
		identifier := m.GetIdentifier()
		_ = identifier
		if _identifierErr := writeBuffer.WriteVirtual(ctx, "identifier", m.GetIdentifier()); _identifierErr != nil {
			return errors.Wrap(_identifierErr, "Error serializing 'identifier' field")
		}

		if popErr := writeBuffer.PopContext("NodeIdByteString"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for NodeIdByteString")
		}
		return nil
	}
	return m.NodeIdTypeDefinitionContract.(*_NodeIdTypeDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_NodeIdByteString) isNodeIdByteString() bool {
	return true
}

func (m *_NodeIdByteString) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
