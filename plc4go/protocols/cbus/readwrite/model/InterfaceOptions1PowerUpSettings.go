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

// InterfaceOptions1PowerUpSettings is the corresponding interface of InterfaceOptions1PowerUpSettings
type InterfaceOptions1PowerUpSettings interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetInterfaceOptions1 returns InterfaceOptions1 (property field)
	GetInterfaceOptions1() InterfaceOptions1
}

// InterfaceOptions1PowerUpSettingsExactly can be used when we want exactly this type and not a type which fulfills InterfaceOptions1PowerUpSettings.
// This is useful for switch cases.
type InterfaceOptions1PowerUpSettingsExactly interface {
	InterfaceOptions1PowerUpSettings
	isInterfaceOptions1PowerUpSettings() bool
}

// _InterfaceOptions1PowerUpSettings is the data-structure of this message
type _InterfaceOptions1PowerUpSettings struct {
	InterfaceOptions1 InterfaceOptions1
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_InterfaceOptions1PowerUpSettings) GetInterfaceOptions1() InterfaceOptions1 {
	return m.InterfaceOptions1
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewInterfaceOptions1PowerUpSettings factory function for _InterfaceOptions1PowerUpSettings
func NewInterfaceOptions1PowerUpSettings(interfaceOptions1 InterfaceOptions1) *_InterfaceOptions1PowerUpSettings {
	return &_InterfaceOptions1PowerUpSettings{InterfaceOptions1: interfaceOptions1}
}

// Deprecated: use the interface for direct cast
func CastInterfaceOptions1PowerUpSettings(structType any) InterfaceOptions1PowerUpSettings {
	if casted, ok := structType.(InterfaceOptions1PowerUpSettings); ok {
		return casted
	}
	if casted, ok := structType.(*InterfaceOptions1PowerUpSettings); ok {
		return *casted
	}
	return nil
}

func (m *_InterfaceOptions1PowerUpSettings) GetTypeName() string {
	return "InterfaceOptions1PowerUpSettings"
}

func (m *_InterfaceOptions1PowerUpSettings) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (interfaceOptions1)
	lengthInBits += m.InterfaceOptions1.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_InterfaceOptions1PowerUpSettings) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func InterfaceOptions1PowerUpSettingsParse(ctx context.Context, theBytes []byte) (InterfaceOptions1PowerUpSettings, error) {
	return InterfaceOptions1PowerUpSettingsParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func InterfaceOptions1PowerUpSettingsParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (InterfaceOptions1PowerUpSettings, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (InterfaceOptions1PowerUpSettings, error) {
		return InterfaceOptions1PowerUpSettingsParseWithBuffer(ctx, readBuffer)
	}
}

func InterfaceOptions1PowerUpSettingsParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (InterfaceOptions1PowerUpSettings, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("InterfaceOptions1PowerUpSettings"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for InterfaceOptions1PowerUpSettings")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	interfaceOptions1, err := ReadSimpleField[InterfaceOptions1](ctx, "interfaceOptions1", ReadComplex[InterfaceOptions1](InterfaceOptions1ParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'interfaceOptions1' field"))
	}

	if closeErr := readBuffer.CloseContext("InterfaceOptions1PowerUpSettings"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for InterfaceOptions1PowerUpSettings")
	}

	// Create the instance
	return &_InterfaceOptions1PowerUpSettings{
		InterfaceOptions1: interfaceOptions1,
	}, nil
}

func (m *_InterfaceOptions1PowerUpSettings) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_InterfaceOptions1PowerUpSettings) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("InterfaceOptions1PowerUpSettings"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for InterfaceOptions1PowerUpSettings")
	}

	// Simple Field (interfaceOptions1)
	if pushErr := writeBuffer.PushContext("interfaceOptions1"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for interfaceOptions1")
	}
	_interfaceOptions1Err := writeBuffer.WriteSerializable(ctx, m.GetInterfaceOptions1())
	if popErr := writeBuffer.PopContext("interfaceOptions1"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for interfaceOptions1")
	}
	if _interfaceOptions1Err != nil {
		return errors.Wrap(_interfaceOptions1Err, "Error serializing 'interfaceOptions1' field")
	}

	if popErr := writeBuffer.PopContext("InterfaceOptions1PowerUpSettings"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for InterfaceOptions1PowerUpSettings")
	}
	return nil
}

func (m *_InterfaceOptions1PowerUpSettings) isInterfaceOptions1PowerUpSettings() bool {
	return true
}

func (m *_InterfaceOptions1PowerUpSettings) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
