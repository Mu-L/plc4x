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

// BACnetConstructedDataLightingCommand is the corresponding interface of BACnetConstructedDataLightingCommand
type BACnetConstructedDataLightingCommand interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetLightingCommand returns LightingCommand (property field)
	GetLightingCommand() BACnetLightingCommand
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetLightingCommand
}

// BACnetConstructedDataLightingCommandExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataLightingCommand.
// This is useful for switch cases.
type BACnetConstructedDataLightingCommandExactly interface {
	BACnetConstructedDataLightingCommand
	isBACnetConstructedDataLightingCommand() bool
}

// _BACnetConstructedDataLightingCommand is the data-structure of this message
type _BACnetConstructedDataLightingCommand struct {
	BACnetConstructedDataContract
	LightingCommand BACnetLightingCommand
}

var _ BACnetConstructedDataLightingCommand = (*_BACnetConstructedDataLightingCommand)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataLightingCommand)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLightingCommand) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataLightingCommand) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LIGHTING_COMMAND
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLightingCommand) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLightingCommand) GetLightingCommand() BACnetLightingCommand {
	return m.LightingCommand
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataLightingCommand) GetActualValue() BACnetLightingCommand {
	ctx := context.Background()
	_ = ctx
	return CastBACnetLightingCommand(m.GetLightingCommand())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataLightingCommand factory function for _BACnetConstructedDataLightingCommand
func NewBACnetConstructedDataLightingCommand(lightingCommand BACnetLightingCommand, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLightingCommand {
	_result := &_BACnetConstructedDataLightingCommand{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		LightingCommand:               lightingCommand,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLightingCommand(structType any) BACnetConstructedDataLightingCommand {
	if casted, ok := structType.(BACnetConstructedDataLightingCommand); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLightingCommand); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLightingCommand) GetTypeName() string {
	return "BACnetConstructedDataLightingCommand"
}

func (m *_BACnetConstructedDataLightingCommand) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (lightingCommand)
	lengthInBits += m.LightingCommand.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataLightingCommand) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataLightingCommand) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataLightingCommand BACnetConstructedDataLightingCommand, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLightingCommand"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLightingCommand")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	lightingCommand, err := ReadSimpleField[BACnetLightingCommand](ctx, "lightingCommand", ReadComplex[BACnetLightingCommand](BACnetLightingCommandParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lightingCommand' field"))
	}
	m.LightingCommand = lightingCommand

	actualValue, err := ReadVirtualField[BACnetLightingCommand](ctx, "actualValue", (*BACnetLightingCommand)(nil), lightingCommand)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLightingCommand"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLightingCommand")
	}

	return m, nil
}

func (m *_BACnetConstructedDataLightingCommand) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataLightingCommand) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLightingCommand"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLightingCommand")
		}

		if err := WriteSimpleField[BACnetLightingCommand](ctx, "lightingCommand", m.GetLightingCommand(), WriteComplex[BACnetLightingCommand](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'lightingCommand' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLightingCommand"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLightingCommand")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLightingCommand) isBACnetConstructedDataLightingCommand() bool {
	return true
}

func (m *_BACnetConstructedDataLightingCommand) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
