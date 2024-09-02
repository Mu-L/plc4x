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

// BACnetPropertyStatesProgramChange is the corresponding interface of BACnetPropertyStatesProgramChange
type BACnetPropertyStatesProgramChange interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetPropertyStates
	// GetProgramState returns ProgramState (property field)
	GetProgramState() BACnetProgramStateTagged
}

// BACnetPropertyStatesProgramChangeExactly can be used when we want exactly this type and not a type which fulfills BACnetPropertyStatesProgramChange.
// This is useful for switch cases.
type BACnetPropertyStatesProgramChangeExactly interface {
	BACnetPropertyStatesProgramChange
	isBACnetPropertyStatesProgramChange() bool
}

// _BACnetPropertyStatesProgramChange is the data-structure of this message
type _BACnetPropertyStatesProgramChange struct {
	BACnetPropertyStatesContract
	ProgramState BACnetProgramStateTagged
}

var _ BACnetPropertyStatesProgramChange = (*_BACnetPropertyStatesProgramChange)(nil)
var _ BACnetPropertyStatesRequirements = (*_BACnetPropertyStatesProgramChange)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPropertyStatesProgramChange) GetParent() BACnetPropertyStatesContract {
	return m.BACnetPropertyStatesContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesProgramChange) GetProgramState() BACnetProgramStateTagged {
	return m.ProgramState
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStatesProgramChange factory function for _BACnetPropertyStatesProgramChange
func NewBACnetPropertyStatesProgramChange(programState BACnetProgramStateTagged, peekedTagHeader BACnetTagHeader) *_BACnetPropertyStatesProgramChange {
	_result := &_BACnetPropertyStatesProgramChange{
		BACnetPropertyStatesContract: NewBACnetPropertyStates(peekedTagHeader),
		ProgramState:                 programState,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesProgramChange(structType any) BACnetPropertyStatesProgramChange {
	if casted, ok := structType.(BACnetPropertyStatesProgramChange); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesProgramChange); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesProgramChange) GetTypeName() string {
	return "BACnetPropertyStatesProgramChange"
}

func (m *_BACnetPropertyStatesProgramChange) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).getLengthInBits(ctx))

	// Simple field (programState)
	lengthInBits += m.ProgramState.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPropertyStatesProgramChange) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetPropertyStatesProgramChange) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetPropertyStates, peekedTagNumber uint8) (__bACnetPropertyStatesProgramChange BACnetPropertyStatesProgramChange, err error) {
	m.BACnetPropertyStatesContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesProgramChange"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesProgramChange")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	programState, err := ReadSimpleField[BACnetProgramStateTagged](ctx, "programState", ReadComplex[BACnetProgramStateTagged](BACnetProgramStateTaggedParseWithBufferProducer((uint8)(peekedTagNumber), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'programState' field"))
	}
	m.ProgramState = programState

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesProgramChange"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesProgramChange")
	}

	return m, nil
}

func (m *_BACnetPropertyStatesProgramChange) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesProgramChange) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesProgramChange"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesProgramChange")
		}

		if err := WriteSimpleField[BACnetProgramStateTagged](ctx, "programState", m.GetProgramState(), WriteComplex[BACnetProgramStateTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'programState' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesProgramChange"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesProgramChange")
		}
		return nil
	}
	return m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesProgramChange) isBACnetPropertyStatesProgramChange() bool {
	return true
}

func (m *_BACnetPropertyStatesProgramChange) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
