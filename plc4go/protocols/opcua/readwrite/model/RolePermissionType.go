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

// RolePermissionType is the corresponding interface of RolePermissionType
type RolePermissionType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetRoleId returns RoleId (property field)
	GetRoleId() NodeId
	// GetPermissions returns Permissions (property field)
	GetPermissions() PermissionType
}

// RolePermissionTypeExactly can be used when we want exactly this type and not a type which fulfills RolePermissionType.
// This is useful for switch cases.
type RolePermissionTypeExactly interface {
	RolePermissionType
	isRolePermissionType() bool
}

// _RolePermissionType is the data-structure of this message
type _RolePermissionType struct {
	ExtensionObjectDefinitionContract
	RoleId      NodeId
	Permissions PermissionType
}

var _ RolePermissionType = (*_RolePermissionType)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_RolePermissionType)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_RolePermissionType) GetIdentifier() string {
	return "98"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_RolePermissionType) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_RolePermissionType) GetRoleId() NodeId {
	return m.RoleId
}

func (m *_RolePermissionType) GetPermissions() PermissionType {
	return m.Permissions
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewRolePermissionType factory function for _RolePermissionType
func NewRolePermissionType(roleId NodeId, permissions PermissionType) *_RolePermissionType {
	_result := &_RolePermissionType{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		RoleId:                            roleId,
		Permissions:                       permissions,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastRolePermissionType(structType any) RolePermissionType {
	if casted, ok := structType.(RolePermissionType); ok {
		return casted
	}
	if casted, ok := structType.(*RolePermissionType); ok {
		return *casted
	}
	return nil
}

func (m *_RolePermissionType) GetTypeName() string {
	return "RolePermissionType"
}

func (m *_RolePermissionType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (roleId)
	lengthInBits += m.RoleId.GetLengthInBits(ctx)

	// Simple field (permissions)
	lengthInBits += 32

	return lengthInBits
}

func (m *_RolePermissionType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_RolePermissionType) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__rolePermissionType RolePermissionType, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("RolePermissionType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for RolePermissionType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	roleId, err := ReadSimpleField[NodeId](ctx, "roleId", ReadComplex[NodeId](NodeIdParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'roleId' field"))
	}
	m.RoleId = roleId

	permissions, err := ReadEnumField[PermissionType](ctx, "permissions", "PermissionType", ReadEnum(PermissionTypeByValue, ReadUnsignedInt(readBuffer, uint8(32))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'permissions' field"))
	}
	m.Permissions = permissions

	if closeErr := readBuffer.CloseContext("RolePermissionType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for RolePermissionType")
	}

	return m, nil
}

func (m *_RolePermissionType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_RolePermissionType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("RolePermissionType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for RolePermissionType")
		}

		if err := WriteSimpleField[NodeId](ctx, "roleId", m.GetRoleId(), WriteComplex[NodeId](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'roleId' field")
		}

		if err := WriteSimpleEnumField[PermissionType](ctx, "permissions", "PermissionType", m.GetPermissions(), WriteEnum[PermissionType, uint32](PermissionType.GetValue, PermissionType.PLC4XEnumName, WriteUnsignedInt(writeBuffer, 32))); err != nil {
			return errors.Wrap(err, "Error serializing 'permissions' field")
		}

		if popErr := writeBuffer.PopContext("RolePermissionType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for RolePermissionType")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_RolePermissionType) isRolePermissionType() bool {
	return true
}

func (m *_RolePermissionType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
