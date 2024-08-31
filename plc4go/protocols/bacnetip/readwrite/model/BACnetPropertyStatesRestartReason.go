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

// BACnetPropertyStatesRestartReason is the corresponding interface of BACnetPropertyStatesRestartReason
type BACnetPropertyStatesRestartReason interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetPropertyStates
	// GetRestartReason returns RestartReason (property field)
	GetRestartReason() BACnetRestartReasonTagged
}

// BACnetPropertyStatesRestartReasonExactly can be used when we want exactly this type and not a type which fulfills BACnetPropertyStatesRestartReason.
// This is useful for switch cases.
type BACnetPropertyStatesRestartReasonExactly interface {
	BACnetPropertyStatesRestartReason
	isBACnetPropertyStatesRestartReason() bool
}

// _BACnetPropertyStatesRestartReason is the data-structure of this message
type _BACnetPropertyStatesRestartReason struct {
	*_BACnetPropertyStates
	RestartReason BACnetRestartReasonTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPropertyStatesRestartReason) InitializeParent(parent BACnetPropertyStates, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetPropertyStatesRestartReason) GetParent() BACnetPropertyStates {
	return m._BACnetPropertyStates
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesRestartReason) GetRestartReason() BACnetRestartReasonTagged {
	return m.RestartReason
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStatesRestartReason factory function for _BACnetPropertyStatesRestartReason
func NewBACnetPropertyStatesRestartReason(restartReason BACnetRestartReasonTagged, peekedTagHeader BACnetTagHeader) *_BACnetPropertyStatesRestartReason {
	_result := &_BACnetPropertyStatesRestartReason{
		RestartReason:         restartReason,
		_BACnetPropertyStates: NewBACnetPropertyStates(peekedTagHeader),
	}
	_result._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesRestartReason(structType any) BACnetPropertyStatesRestartReason {
	if casted, ok := structType.(BACnetPropertyStatesRestartReason); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesRestartReason); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesRestartReason) GetTypeName() string {
	return "BACnetPropertyStatesRestartReason"
}

func (m *_BACnetPropertyStatesRestartReason) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (restartReason)
	lengthInBits += m.RestartReason.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPropertyStatesRestartReason) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetPropertyStatesRestartReasonParse(ctx context.Context, theBytes []byte, peekedTagNumber uint8) (BACnetPropertyStatesRestartReason, error) {
	return BACnetPropertyStatesRestartReasonParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), peekedTagNumber)
}

func BACnetPropertyStatesRestartReasonParseWithBufferProducer(peekedTagNumber uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetPropertyStatesRestartReason, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetPropertyStatesRestartReason, error) {
		return BACnetPropertyStatesRestartReasonParseWithBuffer(ctx, readBuffer, peekedTagNumber)
	}
}

func BACnetPropertyStatesRestartReasonParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, peekedTagNumber uint8) (BACnetPropertyStatesRestartReason, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesRestartReason"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesRestartReason")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	restartReason, err := ReadSimpleField[BACnetRestartReasonTagged](ctx, "restartReason", ReadComplex[BACnetRestartReasonTagged](BACnetRestartReasonTaggedParseWithBufferProducer((uint8)(peekedTagNumber), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'restartReason' field"))
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesRestartReason"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesRestartReason")
	}

	// Create a partially initialized instance
	_child := &_BACnetPropertyStatesRestartReason{
		_BACnetPropertyStates: &_BACnetPropertyStates{},
		RestartReason:         restartReason,
	}
	_child._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _child
	return _child, nil
}

func (m *_BACnetPropertyStatesRestartReason) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesRestartReason) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesRestartReason"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesRestartReason")
		}

		// Simple Field (restartReason)
		if pushErr := writeBuffer.PushContext("restartReason"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for restartReason")
		}
		_restartReasonErr := writeBuffer.WriteSerializable(ctx, m.GetRestartReason())
		if popErr := writeBuffer.PopContext("restartReason"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for restartReason")
		}
		if _restartReasonErr != nil {
			return errors.Wrap(_restartReasonErr, "Error serializing 'restartReason' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesRestartReason"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesRestartReason")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesRestartReason) isBACnetPropertyStatesRestartReason() bool {
	return true
}

func (m *_BACnetPropertyStatesRestartReason) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
