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

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// SubscribedDataSetDataType is the corresponding interface of SubscribedDataSetDataType
type SubscribedDataSetDataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
}

// SubscribedDataSetDataTypeExactly can be used when we want exactly this type and not a type which fulfills SubscribedDataSetDataType.
// This is useful for switch cases.
type SubscribedDataSetDataTypeExactly interface {
	SubscribedDataSetDataType
	isSubscribedDataSetDataType() bool
}

// _SubscribedDataSetDataType is the data-structure of this message
type _SubscribedDataSetDataType struct {
	*_ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SubscribedDataSetDataType) GetIdentifier() string {
	return "15632"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SubscribedDataSetDataType) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_SubscribedDataSetDataType) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

// NewSubscribedDataSetDataType factory function for _SubscribedDataSetDataType
func NewSubscribedDataSetDataType() *_SubscribedDataSetDataType {
	_result := &_SubscribedDataSetDataType{
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSubscribedDataSetDataType(structType any) SubscribedDataSetDataType {
	if casted, ok := structType.(SubscribedDataSetDataType); ok {
		return casted
	}
	if casted, ok := structType.(*SubscribedDataSetDataType); ok {
		return *casted
	}
	return nil
}

func (m *_SubscribedDataSetDataType) GetTypeName() string {
	return "SubscribedDataSetDataType"
}

func (m *_SubscribedDataSetDataType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	return lengthInBits
}

func (m *_SubscribedDataSetDataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func SubscribedDataSetDataTypeParse(ctx context.Context, theBytes []byte, identifier string) (SubscribedDataSetDataType, error) {
	return SubscribedDataSetDataTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func SubscribedDataSetDataTypeParseWithBufferProducer(identifier string) func(ctx context.Context, readBuffer utils.ReadBuffer) (SubscribedDataSetDataType, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (SubscribedDataSetDataType, error) {
		return SubscribedDataSetDataTypeParseWithBuffer(ctx, readBuffer, identifier)
	}
}

func SubscribedDataSetDataTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (SubscribedDataSetDataType, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("SubscribedDataSetDataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SubscribedDataSetDataType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("SubscribedDataSetDataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SubscribedDataSetDataType")
	}

	// Create a partially initialized instance
	_child := &_SubscribedDataSetDataType{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_SubscribedDataSetDataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SubscribedDataSetDataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SubscribedDataSetDataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SubscribedDataSetDataType")
		}

		if popErr := writeBuffer.PopContext("SubscribedDataSetDataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SubscribedDataSetDataType")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SubscribedDataSetDataType) isSubscribedDataSetDataType() bool {
	return true
}

func (m *_SubscribedDataSetDataType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
