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

// SecurityDataCurrentAlarmType is the corresponding interface of SecurityDataCurrentAlarmType
type SecurityDataCurrentAlarmType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	SecurityData
}

// SecurityDataCurrentAlarmTypeExactly can be used when we want exactly this type and not a type which fulfills SecurityDataCurrentAlarmType.
// This is useful for switch cases.
type SecurityDataCurrentAlarmTypeExactly interface {
	SecurityDataCurrentAlarmType
	isSecurityDataCurrentAlarmType() bool
}

// _SecurityDataCurrentAlarmType is the data-structure of this message
type _SecurityDataCurrentAlarmType struct {
	*_SecurityData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SecurityDataCurrentAlarmType) InitializeParent(parent SecurityData, commandTypeContainer SecurityCommandTypeContainer, argument byte) {
	m.CommandTypeContainer = commandTypeContainer
	m.Argument = argument
}

func (m *_SecurityDataCurrentAlarmType) GetParent() SecurityData {
	return m._SecurityData
}

// NewSecurityDataCurrentAlarmType factory function for _SecurityDataCurrentAlarmType
func NewSecurityDataCurrentAlarmType(commandTypeContainer SecurityCommandTypeContainer, argument byte) *_SecurityDataCurrentAlarmType {
	_result := &_SecurityDataCurrentAlarmType{
		_SecurityData: NewSecurityData(commandTypeContainer, argument),
	}
	_result._SecurityData._SecurityDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSecurityDataCurrentAlarmType(structType any) SecurityDataCurrentAlarmType {
	if casted, ok := structType.(SecurityDataCurrentAlarmType); ok {
		return casted
	}
	if casted, ok := structType.(*SecurityDataCurrentAlarmType); ok {
		return *casted
	}
	return nil
}

func (m *_SecurityDataCurrentAlarmType) GetTypeName() string {
	return "SecurityDataCurrentAlarmType"
}

func (m *_SecurityDataCurrentAlarmType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	return lengthInBits
}

func (m *_SecurityDataCurrentAlarmType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func SecurityDataCurrentAlarmTypeParse(ctx context.Context, theBytes []byte) (SecurityDataCurrentAlarmType, error) {
	return SecurityDataCurrentAlarmTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func SecurityDataCurrentAlarmTypeParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (SecurityDataCurrentAlarmType, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (SecurityDataCurrentAlarmType, error) {
		return SecurityDataCurrentAlarmTypeParseWithBuffer(ctx, readBuffer)
	}
}

func SecurityDataCurrentAlarmTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (SecurityDataCurrentAlarmType, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("SecurityDataCurrentAlarmType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SecurityDataCurrentAlarmType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("SecurityDataCurrentAlarmType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SecurityDataCurrentAlarmType")
	}

	// Create a partially initialized instance
	_child := &_SecurityDataCurrentAlarmType{
		_SecurityData: &_SecurityData{},
	}
	_child._SecurityData._SecurityDataChildRequirements = _child
	return _child, nil
}

func (m *_SecurityDataCurrentAlarmType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SecurityDataCurrentAlarmType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SecurityDataCurrentAlarmType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SecurityDataCurrentAlarmType")
		}

		if popErr := writeBuffer.PopContext("SecurityDataCurrentAlarmType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SecurityDataCurrentAlarmType")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SecurityDataCurrentAlarmType) isSecurityDataCurrentAlarmType() bool {
	return true
}

func (m *_SecurityDataCurrentAlarmType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
