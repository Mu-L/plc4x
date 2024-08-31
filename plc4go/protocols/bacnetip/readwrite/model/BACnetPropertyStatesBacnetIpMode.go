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

// BACnetPropertyStatesBacnetIpMode is the corresponding interface of BACnetPropertyStatesBacnetIpMode
type BACnetPropertyStatesBacnetIpMode interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetPropertyStates
	// GetBacnetIpMode returns BacnetIpMode (property field)
	GetBacnetIpMode() BACnetIPModeTagged
}

// BACnetPropertyStatesBacnetIpModeExactly can be used when we want exactly this type and not a type which fulfills BACnetPropertyStatesBacnetIpMode.
// This is useful for switch cases.
type BACnetPropertyStatesBacnetIpModeExactly interface {
	BACnetPropertyStatesBacnetIpMode
	isBACnetPropertyStatesBacnetIpMode() bool
}

// _BACnetPropertyStatesBacnetIpMode is the data-structure of this message
type _BACnetPropertyStatesBacnetIpMode struct {
	*_BACnetPropertyStates
	BacnetIpMode BACnetIPModeTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPropertyStatesBacnetIpMode) InitializeParent(parent BACnetPropertyStates, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetPropertyStatesBacnetIpMode) GetParent() BACnetPropertyStates {
	return m._BACnetPropertyStates
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesBacnetIpMode) GetBacnetIpMode() BACnetIPModeTagged {
	return m.BacnetIpMode
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStatesBacnetIpMode factory function for _BACnetPropertyStatesBacnetIpMode
func NewBACnetPropertyStatesBacnetIpMode(bacnetIpMode BACnetIPModeTagged, peekedTagHeader BACnetTagHeader) *_BACnetPropertyStatesBacnetIpMode {
	_result := &_BACnetPropertyStatesBacnetIpMode{
		BacnetIpMode:          bacnetIpMode,
		_BACnetPropertyStates: NewBACnetPropertyStates(peekedTagHeader),
	}
	_result._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesBacnetIpMode(structType any) BACnetPropertyStatesBacnetIpMode {
	if casted, ok := structType.(BACnetPropertyStatesBacnetIpMode); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesBacnetIpMode); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesBacnetIpMode) GetTypeName() string {
	return "BACnetPropertyStatesBacnetIpMode"
}

func (m *_BACnetPropertyStatesBacnetIpMode) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (bacnetIpMode)
	lengthInBits += m.BacnetIpMode.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPropertyStatesBacnetIpMode) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetPropertyStatesBacnetIpModeParse(ctx context.Context, theBytes []byte, peekedTagNumber uint8) (BACnetPropertyStatesBacnetIpMode, error) {
	return BACnetPropertyStatesBacnetIpModeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), peekedTagNumber)
}

func BACnetPropertyStatesBacnetIpModeParseWithBufferProducer(peekedTagNumber uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetPropertyStatesBacnetIpMode, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetPropertyStatesBacnetIpMode, error) {
		return BACnetPropertyStatesBacnetIpModeParseWithBuffer(ctx, readBuffer, peekedTagNumber)
	}
}

func BACnetPropertyStatesBacnetIpModeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, peekedTagNumber uint8) (BACnetPropertyStatesBacnetIpMode, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesBacnetIpMode"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesBacnetIpMode")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	bacnetIpMode, err := ReadSimpleField[BACnetIPModeTagged](ctx, "bacnetIpMode", ReadComplex[BACnetIPModeTagged](BACnetIPModeTaggedParseWithBufferProducer((uint8)(peekedTagNumber), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'bacnetIpMode' field"))
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesBacnetIpMode"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesBacnetIpMode")
	}

	// Create a partially initialized instance
	_child := &_BACnetPropertyStatesBacnetIpMode{
		_BACnetPropertyStates: &_BACnetPropertyStates{},
		BacnetIpMode:          bacnetIpMode,
	}
	_child._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _child
	return _child, nil
}

func (m *_BACnetPropertyStatesBacnetIpMode) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesBacnetIpMode) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesBacnetIpMode"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesBacnetIpMode")
		}

		// Simple Field (bacnetIpMode)
		if pushErr := writeBuffer.PushContext("bacnetIpMode"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for bacnetIpMode")
		}
		_bacnetIpModeErr := writeBuffer.WriteSerializable(ctx, m.GetBacnetIpMode())
		if popErr := writeBuffer.PopContext("bacnetIpMode"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for bacnetIpMode")
		}
		if _bacnetIpModeErr != nil {
			return errors.Wrap(_bacnetIpModeErr, "Error serializing 'bacnetIpMode' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesBacnetIpMode"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesBacnetIpMode")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesBacnetIpMode) isBACnetPropertyStatesBacnetIpMode() bool {
	return true
}

func (m *_BACnetPropertyStatesBacnetIpMode) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
