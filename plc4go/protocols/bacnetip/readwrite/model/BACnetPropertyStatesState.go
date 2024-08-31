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

// BACnetPropertyStatesState is the corresponding interface of BACnetPropertyStatesState
type BACnetPropertyStatesState interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetPropertyStates
	// GetState returns State (property field)
	GetState() BACnetEventStateTagged
}

// BACnetPropertyStatesStateExactly can be used when we want exactly this type and not a type which fulfills BACnetPropertyStatesState.
// This is useful for switch cases.
type BACnetPropertyStatesStateExactly interface {
	BACnetPropertyStatesState
	isBACnetPropertyStatesState() bool
}

// _BACnetPropertyStatesState is the data-structure of this message
type _BACnetPropertyStatesState struct {
	*_BACnetPropertyStates
	State BACnetEventStateTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPropertyStatesState) InitializeParent(parent BACnetPropertyStates, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetPropertyStatesState) GetParent() BACnetPropertyStates {
	return m._BACnetPropertyStates
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesState) GetState() BACnetEventStateTagged {
	return m.State
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStatesState factory function for _BACnetPropertyStatesState
func NewBACnetPropertyStatesState(state BACnetEventStateTagged, peekedTagHeader BACnetTagHeader) *_BACnetPropertyStatesState {
	_result := &_BACnetPropertyStatesState{
		State:                 state,
		_BACnetPropertyStates: NewBACnetPropertyStates(peekedTagHeader),
	}
	_result._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesState(structType any) BACnetPropertyStatesState {
	if casted, ok := structType.(BACnetPropertyStatesState); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesState); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesState) GetTypeName() string {
	return "BACnetPropertyStatesState"
}

func (m *_BACnetPropertyStatesState) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (state)
	lengthInBits += m.State.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPropertyStatesState) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetPropertyStatesStateParse(ctx context.Context, theBytes []byte, peekedTagNumber uint8) (BACnetPropertyStatesState, error) {
	return BACnetPropertyStatesStateParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), peekedTagNumber)
}

func BACnetPropertyStatesStateParseWithBufferProducer(peekedTagNumber uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetPropertyStatesState, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetPropertyStatesState, error) {
		return BACnetPropertyStatesStateParseWithBuffer(ctx, readBuffer, peekedTagNumber)
	}
}

func BACnetPropertyStatesStateParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, peekedTagNumber uint8) (BACnetPropertyStatesState, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesState"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesState")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	state, err := ReadSimpleField[BACnetEventStateTagged](ctx, "state", ReadComplex[BACnetEventStateTagged](BACnetEventStateTaggedParseWithBufferProducer((uint8)(peekedTagNumber), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'state' field"))
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesState"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesState")
	}

	// Create a partially initialized instance
	_child := &_BACnetPropertyStatesState{
		_BACnetPropertyStates: &_BACnetPropertyStates{},
		State:                 state,
	}
	_child._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _child
	return _child, nil
}

func (m *_BACnetPropertyStatesState) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesState) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesState"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesState")
		}

		// Simple Field (state)
		if pushErr := writeBuffer.PushContext("state"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for state")
		}
		_stateErr := writeBuffer.WriteSerializable(ctx, m.GetState())
		if popErr := writeBuffer.PopContext("state"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for state")
		}
		if _stateErr != nil {
			return errors.Wrap(_stateErr, "Error serializing 'state' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesState"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesState")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesState) isBACnetPropertyStatesState() bool {
	return true
}

func (m *_BACnetPropertyStatesState) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
