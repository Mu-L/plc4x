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

// DataTypeDescription is the corresponding interface of DataTypeDescription
type DataTypeDescription interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetDataTypeId returns DataTypeId (property field)
	GetDataTypeId() NodeId
	// GetName returns Name (property field)
	GetName() QualifiedName
	// IsDataTypeDescription is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsDataTypeDescription()
}

// _DataTypeDescription is the data-structure of this message
type _DataTypeDescription struct {
	ExtensionObjectDefinitionContract
	DataTypeId NodeId
	Name       QualifiedName
}

var _ DataTypeDescription = (*_DataTypeDescription)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_DataTypeDescription)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_DataTypeDescription) GetIdentifier() string {
	return "14527"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_DataTypeDescription) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_DataTypeDescription) GetDataTypeId() NodeId {
	return m.DataTypeId
}

func (m *_DataTypeDescription) GetName() QualifiedName {
	return m.Name
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewDataTypeDescription factory function for _DataTypeDescription
func NewDataTypeDescription(dataTypeId NodeId, name QualifiedName) *_DataTypeDescription {
	_result := &_DataTypeDescription{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		DataTypeId:                        dataTypeId,
		Name:                              name,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastDataTypeDescription(structType any) DataTypeDescription {
	if casted, ok := structType.(DataTypeDescription); ok {
		return casted
	}
	if casted, ok := structType.(*DataTypeDescription); ok {
		return *casted
	}
	return nil
}

func (m *_DataTypeDescription) GetTypeName() string {
	return "DataTypeDescription"
}

func (m *_DataTypeDescription) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (dataTypeId)
	lengthInBits += m.DataTypeId.GetLengthInBits(ctx)

	// Simple field (name)
	lengthInBits += m.Name.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_DataTypeDescription) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_DataTypeDescription) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__dataTypeDescription DataTypeDescription, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("DataTypeDescription"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DataTypeDescription")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	dataTypeId, err := ReadSimpleField[NodeId](ctx, "dataTypeId", ReadComplex[NodeId](NodeIdParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dataTypeId' field"))
	}
	m.DataTypeId = dataTypeId

	name, err := ReadSimpleField[QualifiedName](ctx, "name", ReadComplex[QualifiedName](QualifiedNameParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'name' field"))
	}
	m.Name = name

	if closeErr := readBuffer.CloseContext("DataTypeDescription"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DataTypeDescription")
	}

	return m, nil
}

func (m *_DataTypeDescription) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_DataTypeDescription) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("DataTypeDescription"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for DataTypeDescription")
		}

		if err := WriteSimpleField[NodeId](ctx, "dataTypeId", m.GetDataTypeId(), WriteComplex[NodeId](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'dataTypeId' field")
		}

		if err := WriteSimpleField[QualifiedName](ctx, "name", m.GetName(), WriteComplex[QualifiedName](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'name' field")
		}

		if popErr := writeBuffer.PopContext("DataTypeDescription"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for DataTypeDescription")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_DataTypeDescription) IsDataTypeDescription() {}

func (m *_DataTypeDescription) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
