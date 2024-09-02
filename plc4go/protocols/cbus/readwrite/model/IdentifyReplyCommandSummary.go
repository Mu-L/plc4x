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

// IdentifyReplyCommandSummary is the corresponding interface of IdentifyReplyCommandSummary
type IdentifyReplyCommandSummary interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	IdentifyReplyCommand
	// GetPartName returns PartName (property field)
	GetPartName() string
	// GetUnitServiceType returns UnitServiceType (property field)
	GetUnitServiceType() byte
	// GetVersion returns Version (property field)
	GetVersion() string
}

// IdentifyReplyCommandSummaryExactly can be used when we want exactly this type and not a type which fulfills IdentifyReplyCommandSummary.
// This is useful for switch cases.
type IdentifyReplyCommandSummaryExactly interface {
	IdentifyReplyCommandSummary
	isIdentifyReplyCommandSummary() bool
}

// _IdentifyReplyCommandSummary is the data-structure of this message
type _IdentifyReplyCommandSummary struct {
	IdentifyReplyCommandContract
	PartName        string
	UnitServiceType byte
	Version         string
}

var _ IdentifyReplyCommandSummary = (*_IdentifyReplyCommandSummary)(nil)
var _ IdentifyReplyCommandRequirements = (*_IdentifyReplyCommandSummary)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_IdentifyReplyCommandSummary) GetAttribute() Attribute {
	return Attribute_Summary
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_IdentifyReplyCommandSummary) GetParent() IdentifyReplyCommandContract {
	return m.IdentifyReplyCommandContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_IdentifyReplyCommandSummary) GetPartName() string {
	return m.PartName
}

func (m *_IdentifyReplyCommandSummary) GetUnitServiceType() byte {
	return m.UnitServiceType
}

func (m *_IdentifyReplyCommandSummary) GetVersion() string {
	return m.Version
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewIdentifyReplyCommandSummary factory function for _IdentifyReplyCommandSummary
func NewIdentifyReplyCommandSummary(partName string, unitServiceType byte, version string, numBytes uint8) *_IdentifyReplyCommandSummary {
	_result := &_IdentifyReplyCommandSummary{
		IdentifyReplyCommandContract: NewIdentifyReplyCommand(numBytes),
		PartName:                     partName,
		UnitServiceType:              unitServiceType,
		Version:                      version,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastIdentifyReplyCommandSummary(structType any) IdentifyReplyCommandSummary {
	if casted, ok := structType.(IdentifyReplyCommandSummary); ok {
		return casted
	}
	if casted, ok := structType.(*IdentifyReplyCommandSummary); ok {
		return *casted
	}
	return nil
}

func (m *_IdentifyReplyCommandSummary) GetTypeName() string {
	return "IdentifyReplyCommandSummary"
}

func (m *_IdentifyReplyCommandSummary) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.IdentifyReplyCommandContract.(*_IdentifyReplyCommand).getLengthInBits(ctx))

	// Simple field (partName)
	lengthInBits += 48

	// Simple field (unitServiceType)
	lengthInBits += 8

	// Simple field (version)
	lengthInBits += 32

	return lengthInBits
}

func (m *_IdentifyReplyCommandSummary) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_IdentifyReplyCommandSummary) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_IdentifyReplyCommand, attribute Attribute, numBytes uint8) (__identifyReplyCommandSummary IdentifyReplyCommandSummary, err error) {
	m.IdentifyReplyCommandContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("IdentifyReplyCommandSummary"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for IdentifyReplyCommandSummary")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	partName, err := ReadSimpleField(ctx, "partName", ReadString(readBuffer, uint32(48)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'partName' field"))
	}
	m.PartName = partName

	unitServiceType, err := ReadSimpleField(ctx, "unitServiceType", ReadByte(readBuffer, 8))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'unitServiceType' field"))
	}
	m.UnitServiceType = unitServiceType

	version, err := ReadSimpleField(ctx, "version", ReadString(readBuffer, uint32(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'version' field"))
	}
	m.Version = version

	if closeErr := readBuffer.CloseContext("IdentifyReplyCommandSummary"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for IdentifyReplyCommandSummary")
	}

	return m, nil
}

func (m *_IdentifyReplyCommandSummary) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_IdentifyReplyCommandSummary) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("IdentifyReplyCommandSummary"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for IdentifyReplyCommandSummary")
		}

		if err := WriteSimpleField[string](ctx, "partName", m.GetPartName(), WriteString(writeBuffer, 48)); err != nil {
			return errors.Wrap(err, "Error serializing 'partName' field")
		}

		if err := WriteSimpleField[byte](ctx, "unitServiceType", m.GetUnitServiceType(), WriteByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'unitServiceType' field")
		}

		if err := WriteSimpleField[string](ctx, "version", m.GetVersion(), WriteString(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'version' field")
		}

		if popErr := writeBuffer.PopContext("IdentifyReplyCommandSummary"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for IdentifyReplyCommandSummary")
		}
		return nil
	}
	return m.IdentifyReplyCommandContract.(*_IdentifyReplyCommand).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_IdentifyReplyCommandSummary) isIdentifyReplyCommandSummary() bool {
	return true
}

func (m *_IdentifyReplyCommandSummary) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
