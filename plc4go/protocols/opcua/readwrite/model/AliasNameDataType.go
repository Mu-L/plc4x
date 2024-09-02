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

// AliasNameDataType is the corresponding interface of AliasNameDataType
type AliasNameDataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetAliasName returns AliasName (property field)
	GetAliasName() QualifiedName
	// GetNoOfReferencedNodes returns NoOfReferencedNodes (property field)
	GetNoOfReferencedNodes() int32
	// GetReferencedNodes returns ReferencedNodes (property field)
	GetReferencedNodes() []ExpandedNodeId
	// IsAliasNameDataType is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsAliasNameDataType()
}

// _AliasNameDataType is the data-structure of this message
type _AliasNameDataType struct {
	ExtensionObjectDefinitionContract
	AliasName           QualifiedName
	NoOfReferencedNodes int32
	ReferencedNodes     []ExpandedNodeId
}

var _ AliasNameDataType = (*_AliasNameDataType)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_AliasNameDataType)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_AliasNameDataType) GetIdentifier() string {
	return "23470"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AliasNameDataType) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AliasNameDataType) GetAliasName() QualifiedName {
	return m.AliasName
}

func (m *_AliasNameDataType) GetNoOfReferencedNodes() int32 {
	return m.NoOfReferencedNodes
}

func (m *_AliasNameDataType) GetReferencedNodes() []ExpandedNodeId {
	return m.ReferencedNodes
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAliasNameDataType factory function for _AliasNameDataType
func NewAliasNameDataType(aliasName QualifiedName, noOfReferencedNodes int32, referencedNodes []ExpandedNodeId) *_AliasNameDataType {
	_result := &_AliasNameDataType{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		AliasName:                         aliasName,
		NoOfReferencedNodes:               noOfReferencedNodes,
		ReferencedNodes:                   referencedNodes,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastAliasNameDataType(structType any) AliasNameDataType {
	if casted, ok := structType.(AliasNameDataType); ok {
		return casted
	}
	if casted, ok := structType.(*AliasNameDataType); ok {
		return *casted
	}
	return nil
}

func (m *_AliasNameDataType) GetTypeName() string {
	return "AliasNameDataType"
}

func (m *_AliasNameDataType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (aliasName)
	lengthInBits += m.AliasName.GetLengthInBits(ctx)

	// Simple field (noOfReferencedNodes)
	lengthInBits += 32

	// Array field
	if len(m.ReferencedNodes) > 0 {
		for _curItem, element := range m.ReferencedNodes {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.ReferencedNodes), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_AliasNameDataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_AliasNameDataType) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__aliasNameDataType AliasNameDataType, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AliasNameDataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AliasNameDataType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	aliasName, err := ReadSimpleField[QualifiedName](ctx, "aliasName", ReadComplex[QualifiedName](QualifiedNameParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'aliasName' field"))
	}
	m.AliasName = aliasName

	noOfReferencedNodes, err := ReadSimpleField(ctx, "noOfReferencedNodes", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfReferencedNodes' field"))
	}
	m.NoOfReferencedNodes = noOfReferencedNodes

	referencedNodes, err := ReadCountArrayField[ExpandedNodeId](ctx, "referencedNodes", ReadComplex[ExpandedNodeId](ExpandedNodeIdParseWithBuffer, readBuffer), uint64(noOfReferencedNodes))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'referencedNodes' field"))
	}
	m.ReferencedNodes = referencedNodes

	if closeErr := readBuffer.CloseContext("AliasNameDataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AliasNameDataType")
	}

	return m, nil
}

func (m *_AliasNameDataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AliasNameDataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AliasNameDataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AliasNameDataType")
		}

		if err := WriteSimpleField[QualifiedName](ctx, "aliasName", m.GetAliasName(), WriteComplex[QualifiedName](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'aliasName' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfReferencedNodes", m.GetNoOfReferencedNodes(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfReferencedNodes' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "referencedNodes", m.GetReferencedNodes(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'referencedNodes' field")
		}

		if popErr := writeBuffer.PopContext("AliasNameDataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AliasNameDataType")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AliasNameDataType) IsAliasNameDataType() {}

func (m *_AliasNameDataType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
