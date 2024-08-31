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

// BrowsePath is the corresponding interface of BrowsePath
type BrowsePath interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetStartingNode returns StartingNode (property field)
	GetStartingNode() NodeId
	// GetRelativePath returns RelativePath (property field)
	GetRelativePath() ExtensionObjectDefinition
}

// BrowsePathExactly can be used when we want exactly this type and not a type which fulfills BrowsePath.
// This is useful for switch cases.
type BrowsePathExactly interface {
	BrowsePath
	isBrowsePath() bool
}

// _BrowsePath is the data-structure of this message
type _BrowsePath struct {
	*_ExtensionObjectDefinition
	StartingNode NodeId
	RelativePath ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BrowsePath) GetIdentifier() string {
	return "545"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BrowsePath) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_BrowsePath) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BrowsePath) GetStartingNode() NodeId {
	return m.StartingNode
}

func (m *_BrowsePath) GetRelativePath() ExtensionObjectDefinition {
	return m.RelativePath
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBrowsePath factory function for _BrowsePath
func NewBrowsePath(startingNode NodeId, relativePath ExtensionObjectDefinition) *_BrowsePath {
	_result := &_BrowsePath{
		StartingNode:               startingNode,
		RelativePath:               relativePath,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBrowsePath(structType any) BrowsePath {
	if casted, ok := structType.(BrowsePath); ok {
		return casted
	}
	if casted, ok := structType.(*BrowsePath); ok {
		return *casted
	}
	return nil
}

func (m *_BrowsePath) GetTypeName() string {
	return "BrowsePath"
}

func (m *_BrowsePath) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (startingNode)
	lengthInBits += m.StartingNode.GetLengthInBits(ctx)

	// Simple field (relativePath)
	lengthInBits += m.RelativePath.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BrowsePath) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BrowsePathParse(ctx context.Context, theBytes []byte, identifier string) (BrowsePath, error) {
	return BrowsePathParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func BrowsePathParseWithBufferProducer(identifier string) func(ctx context.Context, readBuffer utils.ReadBuffer) (BrowsePath, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BrowsePath, error) {
		return BrowsePathParseWithBuffer(ctx, readBuffer, identifier)
	}
}

func BrowsePathParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (BrowsePath, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BrowsePath"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BrowsePath")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	startingNode, err := ReadSimpleField[NodeId](ctx, "startingNode", ReadComplex[NodeId](NodeIdParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'startingNode' field"))
	}

	relativePath, err := ReadSimpleField[ExtensionObjectDefinition](ctx, "relativePath", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("542")), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'relativePath' field"))
	}

	if closeErr := readBuffer.CloseContext("BrowsePath"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BrowsePath")
	}

	// Create a partially initialized instance
	_child := &_BrowsePath{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		StartingNode:               startingNode,
		RelativePath:               relativePath,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_BrowsePath) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BrowsePath) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BrowsePath"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BrowsePath")
		}

		// Simple Field (startingNode)
		if pushErr := writeBuffer.PushContext("startingNode"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for startingNode")
		}
		_startingNodeErr := writeBuffer.WriteSerializable(ctx, m.GetStartingNode())
		if popErr := writeBuffer.PopContext("startingNode"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for startingNode")
		}
		if _startingNodeErr != nil {
			return errors.Wrap(_startingNodeErr, "Error serializing 'startingNode' field")
		}

		// Simple Field (relativePath)
		if pushErr := writeBuffer.PushContext("relativePath"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for relativePath")
		}
		_relativePathErr := writeBuffer.WriteSerializable(ctx, m.GetRelativePath())
		if popErr := writeBuffer.PopContext("relativePath"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for relativePath")
		}
		if _relativePathErr != nil {
			return errors.Wrap(_relativePathErr, "Error serializing 'relativePath' field")
		}

		if popErr := writeBuffer.PopContext("BrowsePath"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BrowsePath")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BrowsePath) isBrowsePath() bool {
	return true
}

func (m *_BrowsePath) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
