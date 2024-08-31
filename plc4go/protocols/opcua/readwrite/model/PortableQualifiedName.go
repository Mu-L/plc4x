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

// PortableQualifiedName is the corresponding interface of PortableQualifiedName
type PortableQualifiedName interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetNamespaceUri returns NamespaceUri (property field)
	GetNamespaceUri() PascalString
	// GetName returns Name (property field)
	GetName() PascalString
}

// PortableQualifiedNameExactly can be used when we want exactly this type and not a type which fulfills PortableQualifiedName.
// This is useful for switch cases.
type PortableQualifiedNameExactly interface {
	PortableQualifiedName
	isPortableQualifiedName() bool
}

// _PortableQualifiedName is the data-structure of this message
type _PortableQualifiedName struct {
	*_ExtensionObjectDefinition
	NamespaceUri PascalString
	Name         PascalString
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_PortableQualifiedName) GetIdentifier() string {
	return "24107"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_PortableQualifiedName) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_PortableQualifiedName) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_PortableQualifiedName) GetNamespaceUri() PascalString {
	return m.NamespaceUri
}

func (m *_PortableQualifiedName) GetName() PascalString {
	return m.Name
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewPortableQualifiedName factory function for _PortableQualifiedName
func NewPortableQualifiedName(namespaceUri PascalString, name PascalString) *_PortableQualifiedName {
	_result := &_PortableQualifiedName{
		NamespaceUri:               namespaceUri,
		Name:                       name,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastPortableQualifiedName(structType any) PortableQualifiedName {
	if casted, ok := structType.(PortableQualifiedName); ok {
		return casted
	}
	if casted, ok := structType.(*PortableQualifiedName); ok {
		return *casted
	}
	return nil
}

func (m *_PortableQualifiedName) GetTypeName() string {
	return "PortableQualifiedName"
}

func (m *_PortableQualifiedName) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (namespaceUri)
	lengthInBits += m.NamespaceUri.GetLengthInBits(ctx)

	// Simple field (name)
	lengthInBits += m.Name.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_PortableQualifiedName) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func PortableQualifiedNameParse(ctx context.Context, theBytes []byte, identifier string) (PortableQualifiedName, error) {
	return PortableQualifiedNameParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func PortableQualifiedNameParseWithBufferProducer(identifier string) func(ctx context.Context, readBuffer utils.ReadBuffer) (PortableQualifiedName, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (PortableQualifiedName, error) {
		return PortableQualifiedNameParseWithBuffer(ctx, readBuffer, identifier)
	}
}

func PortableQualifiedNameParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (PortableQualifiedName, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("PortableQualifiedName"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for PortableQualifiedName")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	namespaceUri, err := ReadSimpleField[PascalString](ctx, "namespaceUri", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'namespaceUri' field"))
	}

	name, err := ReadSimpleField[PascalString](ctx, "name", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'name' field"))
	}

	if closeErr := readBuffer.CloseContext("PortableQualifiedName"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for PortableQualifiedName")
	}

	// Create a partially initialized instance
	_child := &_PortableQualifiedName{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		NamespaceUri:               namespaceUri,
		Name:                       name,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_PortableQualifiedName) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_PortableQualifiedName) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("PortableQualifiedName"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for PortableQualifiedName")
		}

		// Simple Field (namespaceUri)
		if pushErr := writeBuffer.PushContext("namespaceUri"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for namespaceUri")
		}
		_namespaceUriErr := writeBuffer.WriteSerializable(ctx, m.GetNamespaceUri())
		if popErr := writeBuffer.PopContext("namespaceUri"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for namespaceUri")
		}
		if _namespaceUriErr != nil {
			return errors.Wrap(_namespaceUriErr, "Error serializing 'namespaceUri' field")
		}

		// Simple Field (name)
		if pushErr := writeBuffer.PushContext("name"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for name")
		}
		_nameErr := writeBuffer.WriteSerializable(ctx, m.GetName())
		if popErr := writeBuffer.PopContext("name"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for name")
		}
		if _nameErr != nil {
			return errors.Wrap(_nameErr, "Error serializing 'name' field")
		}

		if popErr := writeBuffer.PopContext("PortableQualifiedName"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for PortableQualifiedName")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_PortableQualifiedName) isPortableQualifiedName() bool {
	return true
}

func (m *_PortableQualifiedName) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
