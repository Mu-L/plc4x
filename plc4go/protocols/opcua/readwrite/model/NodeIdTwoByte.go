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

// NodeIdTwoByte is the corresponding interface of NodeIdTwoByte
type NodeIdTwoByte interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	NodeIdTypeDefinition
	// GetId returns Id (property field)
	GetId() uint8
	// GetIdentifier returns Identifier (virtual field)
	GetIdentifier() string
}

// NodeIdTwoByteExactly can be used when we want exactly this type and not a type which fulfills NodeIdTwoByte.
// This is useful for switch cases.
type NodeIdTwoByteExactly interface {
	NodeIdTwoByte
	isNodeIdTwoByte() bool
}

// _NodeIdTwoByte is the data-structure of this message
type _NodeIdTwoByte struct {
	NodeIdTypeDefinitionContract
	Id uint8
}

var _ NodeIdTwoByte = (*_NodeIdTwoByte)(nil)
var _ NodeIdTypeDefinitionRequirements = (*_NodeIdTwoByte)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_NodeIdTwoByte) GetNodeType() NodeIdType {
	return NodeIdType_nodeIdTypeTwoByte
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_NodeIdTwoByte) GetParent() NodeIdTypeDefinitionContract {
	return m.NodeIdTypeDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_NodeIdTwoByte) GetId() uint8 {
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

func (m *_NodeIdTwoByte) GetIdentifier() string {
	ctx := context.Background()
	_ = ctx
	return fmt.Sprintf("%v", m.GetId())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewNodeIdTwoByte factory function for _NodeIdTwoByte
func NewNodeIdTwoByte(id uint8) *_NodeIdTwoByte {
	_result := &_NodeIdTwoByte{
		NodeIdTypeDefinitionContract: NewNodeIdTypeDefinition(),
		Id:                           id,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastNodeIdTwoByte(structType any) NodeIdTwoByte {
	if casted, ok := structType.(NodeIdTwoByte); ok {
		return casted
	}
	if casted, ok := structType.(*NodeIdTwoByte); ok {
		return *casted
	}
	return nil
}

func (m *_NodeIdTwoByte) GetTypeName() string {
	return "NodeIdTwoByte"
}

func (m *_NodeIdTwoByte) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.NodeIdTypeDefinitionContract.(*_NodeIdTypeDefinition).getLengthInBits(ctx))

	// Simple field (id)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_NodeIdTwoByte) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_NodeIdTwoByte) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_NodeIdTypeDefinition) (__nodeIdTwoByte NodeIdTwoByte, err error) {
	m.NodeIdTypeDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("NodeIdTwoByte"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NodeIdTwoByte")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	id, err := ReadSimpleField(ctx, "id", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'id' field"))
	}
	m.Id = id

	identifier, err := ReadVirtualField[string](ctx, "identifier", (*string)(nil), id)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'identifier' field"))
	}
	_ = identifier

	if closeErr := readBuffer.CloseContext("NodeIdTwoByte"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NodeIdTwoByte")
	}

	return m, nil
}

func (m *_NodeIdTwoByte) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_NodeIdTwoByte) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("NodeIdTwoByte"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for NodeIdTwoByte")
		}

		if err := WriteSimpleField[uint8](ctx, "id", m.GetId(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'id' field")
		}
		// Virtual field
		identifier := m.GetIdentifier()
		_ = identifier
		if _identifierErr := writeBuffer.WriteVirtual(ctx, "identifier", m.GetIdentifier()); _identifierErr != nil {
			return errors.Wrap(_identifierErr, "Error serializing 'identifier' field")
		}

		if popErr := writeBuffer.PopContext("NodeIdTwoByte"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for NodeIdTwoByte")
		}
		return nil
	}
	return m.NodeIdTypeDefinitionContract.(*_NodeIdTypeDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_NodeIdTwoByte) isNodeIdTwoByte() bool {
	return true
}

func (m *_NodeIdTwoByte) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
